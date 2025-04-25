# Step 1: Import libraries
import pandas as pd
import numpy as np
import sqlalchemy as sa
import joblib
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import LabelEncoder
from xgboost import XGBClassifier
from sklearn.metrics import accuracy_score, classification_report
import matplotlib.pyplot as plt
from xgboost import plot_importance

# Step 2: Connect to PostgreSQL
DATABASE_URL = "postgresql://postgres:postgres@localhost:5432/mydata"
engine = sa.create_engine(DATABASE_URL)

# Step 3: Helper function to query SQL
def query(sql):
    with engine.connect() as connection:
        result = connection.execute(sa.text(sql))
        df = pd.DataFrame(result.fetchall(), columns=result.keys())
        return df

# Step 4: Fetch data
user_df = query("SELECT * FROM users")
loan_df = query("SELECT * FROM loan_applications")
credit_df = query("SELECT * FROM credit_reports")
decision_df = query("SELECT * FROM loan_decisions")
payment_df = query("SELECT * FROM loan_payments")

# Step 5: Merge data
user_df = user_df.rename(columns={'id': 'user_id'})
loan_df = loan_df.rename(columns={'id': 'id_loan'})
credit_df = credit_df.rename(columns={'id': 'id_credit'})
decision_df = decision_df.rename(columns={'id': 'id_decision'})
payment_df = payment_df.rename(columns={'id': 'id_payment'})

df = loan_df.merge(user_df, on='user_id', suffixes=('_loan', '_user'))
df = df.merge(credit_df, left_on='id_loan', right_on='loan_application_id')
df = df.merge(decision_df, left_on='id_loan', right_on='loan_application_id')
df = df.merge(payment_df, left_on='id_loan', right_on='loan_application_id', how='left', suffixes=('', '_payment'))

print("✅ Data merged successfully",df)

# Step 6: Feature engineering from payment history
df['annual_income'] = df['gross_monthly_income'] * 12

# Ensure datetime for safe comparison
df['payment_date'] = pd.to_datetime(df['payment_date'], errors='coerce')
df['due_date'] = pd.to_datetime(df['due_date'], errors='coerce')

# Late payment logic
df['late_payment'] = (df['payment_date'] > df['due_date']) & (df['status'] == 'failed')

# Aggregation on payment behavior
agg_features = df.groupby('id_loan').agg({
    'id_payment': 'count',
    'late_payment': 'sum',
    'amount_paid': 'sum',
    'status': lambda x: (x == 'Success').mean(),
}).rename(columns={
    'id_payment': 'num_payments_made',
    'late_payment': 'num_late_payments',
    'amount_paid': 'total_amount_paid',
    'status': 'payment_success_ratio'
}).reset_index()

# Step 7: Merge back aggregated features
df = df.drop_duplicates(subset=['id_loan'])
df = df.merge(agg_features, on='id_loan', how='left')


# Step 8: Feature selection
# Use correct credit score fields
df = df.rename(columns={
    'credit_score': 'report_credit_score',
    'credit_score_user': 'user_credit_score'
})


features = df[[
    'loan_amount',
    'loan_purpose',
    'employment_status',
    'annual_income',
    'dti_ratio',
    'report_credit_score',
    'user_credit_score',
    'delinquency_flag',
    'num_payments_made',
    'num_late_payments',
    'total_amount_paid',
    'payment_success_ratio'
]]


# Step 9: Target variable (ai_decision as binary)
target = df['ai_decision'].astype(int)

# Step 10: Encode categoricals
for col in ['loan_purpose', 'employment_status']:
    le = LabelEncoder()
    features[col] = le.fit_transform(features[col].astype(str))

# Step 11: Clean and align with target
features = features.apply(pd.to_numeric, errors='coerce')
features.dropna(inplace=True)
target = target.loc[features.index]

# Step 12: Train/test split and training
X_train, X_test, y_train, y_test = train_test_split(features, target, test_size=0.2, random_state=42)
model = XGBClassifier(use_label_encoder=False, eval_metric='logloss')
model.fit(X_train, y_train)

# Step 13: Evaluation
y_pred = model.predict(X_test)
y_prob = model.predict_proba(X_test)[:, 1]

# plot_importance(model)
# plt.show()

print("✅ Accuracy:", accuracy_score(y_test, y_pred))
print("✅ Classification Report:\n", classification_report(y_test, y_pred))

le_loan_purpose = LabelEncoder().fit(features['loan_purpose'])
le_employment_status = LabelEncoder().fit(features['employment_status'])

# joblib.dump(le_loan_purpose, "le_loan_purpose.pkl")
# joblib.dump(le_employment_status, "le_employment_status.pkl")

# # Step 14: Save model
# joblib.dump(model, "loan_model.pkl")
# print("✅ Model saved as loan_model.pkl")
