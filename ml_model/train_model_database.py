import pandas as pd
import numpy as np
import joblib
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import LabelEncoder
from xgboost import XGBClassifier
from sklearn.metrics import accuracy_score, classification_report, root_mean_squared_error, r2_score
import matplotlib.pyplot as plt
from xgboost import plot_importance

# Step 2: Fetch data from CSV files
user_df = pd.read_csv("./csv_files/users.csv")
loan_df = pd.read_csv("./csv_files/loan_applications.csv")
credit_df = pd.read_csv("./csv_files/credit_reports.csv")
decision_df = pd.read_csv("./csv_files/loan_decisions.csv")
payment_df = pd.read_csv("./csv_files/loan_payments.csv")

# Step 3: Rename columns to align for merging
user_df = user_df.rename(columns={'ID': 'user_id', 'CreditScore': 'user_credit_score'})
loan_df = loan_df.rename(columns={'ID': 'id_loan', 'UserID': 'user_id', 'LoanAmount': 'loan_amount', 
                                   'LoanPurpose': 'loan_purpose', 'EmploymentStatus': 'employment_status',
                                   'GrossMonthlyIncome': 'gross_monthly_income', 'DTIRatio': 'dti_ratio'})
credit_df = credit_df.rename(columns={'ID': 'id_credit', 'LoanApplicationID': 'loan_application_id', 
                                      'CreditScore': 'report_credit_score'})
decision_df = decision_df.rename(columns={'ID': 'id_decision', 'LoanApplicationID': 'loan_application_id', 
                                          'AiDecision': 'ai_decision'})
payment_df = payment_df.rename(columns={'ID': 'id_payment', 'LoanApplicationID': 'loan_application_id',
                                        'AmountPaid': 'amount_paid', 'PaymentDate': 'payment_date', 
                                        'Status': 'status'})

# Step 4: Merge dataframes
df = loan_df.merge(user_df, on='user_id', suffixes=('_loan', '_user'))
df = df.merge(credit_df, left_on='id_loan', right_on='loan_application_id')
df = df.merge(decision_df, left_on='id_loan', right_on='loan_application_id')
df = df.merge(payment_df, left_on='id_loan', right_on='loan_application_id', how='left', suffixes=('', '_payment'))

# Step 5: Feature engineering from payment history
df['annual_income'] = df['gross_monthly_income'] * 12

df['payment_date'] = pd.to_datetime(df['payment_date'], errors='coerce')
# df['due_date'] = pd.to_datetime(df['due_date'], errors='coerce')

# df['late_payment'] = (df['payment_date'] > df['due_date']) & (df['status'] == 'failed')

agg_features = df.groupby('id_loan').agg({
    'id_payment': 'count',
    # 'late_payment': 'sum',
    'amount_paid': 'sum',
    'status': lambda x: (x == 'Success').mean(),
}).rename(columns={
    'id_payment': 'num_payments_made',
    # 'late_payment': 'num_late_payments',
    'amount_paid': 'total_amount_paid',
    'status': 'payment_success_ratio'
}).reset_index()

df = df.drop_duplicates(subset=['id_loan'])
df = df.merge(agg_features, on='id_loan', how='left')

# Step 6: Feature selection
features = df[[
    'loan_amount',
    'loan_purpose',
    'employment_status',
    'annual_income',
    'dti_ratio',
    'report_credit_score',
    'user_credit_score',
    'DelinquencyFlag',
    'num_payments_made',
    # 'num_late_payments',
    'total_amount_paid',
    'payment_success_ratio'
]]

# Step 7: Target variable
target = df['ai_decision'].astype(int)

# Step 8: Encode categoricals
for col in ['loan_purpose', 'employment_status']:
    le = LabelEncoder()
    features[col] = le.fit_transform(features[col].astype(str))

# Step 9: Clean and align with target
features = features.apply(pd.to_numeric, errors='coerce')
features.dropna(inplace=True)
target = target.loc[features.index]

# Step 10: Train/test split and training
X_train, X_test, y_train, y_test = train_test_split(features, target, test_size=0.2, random_state=42)
model = XGBClassifier(use_label_encoder=False, eval_metric='logloss')
model.fit(X_train, y_train)

# Step 11: Evaluation
y_pred = model.predict(X_test)
y_prob = model.predict_proba(X_test)[:, 1]

print("✅ Accuracy:", accuracy_score(y_test, y_pred))
print("✅ Classification Report:\n", classification_report(y_test, y_pred))
print("RMSE:", root_mean_squared_error(y_test, y_pred))
print("R² Score:", r2_score(y_test, y_pred))

# Save model
# joblib.dump(model, "loan_model.pkl")
