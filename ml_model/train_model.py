# Step 1: Import libraries
import pandas as pd
import numpy as np
import sqlalchemy as sa
import joblib  # For .pkl
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import LabelEncoder
from xgboost import XGBClassifier
from sklearn.metrics import accuracy_score, classification_report

# Step 2: Connect to PostgreSQL
DATABASE_URL = "postgresql://username:password@localhost:5432/yourdb"
engine = sa.create_engine(DATABASE_URL)

# Step 3: Query data
def query(sql):
    result = engine.execute(sql)
    df = pd.DataFrame(result.fetchall())
    df.columns = result.keys()
    return df

user_df = query("SELECT * FROM users")
loan_df = query("SELECT * FROM loan_applications")
credit_df = query("SELECT * FROM credit_reports")
decision_df = query("SELECT * FROM loan_decisions")

# Step 4: Merge data
df = loan_df.merge(user_df, left_on='user_id', right_on='id', suffixes=('_loan', '_user'))
df = df.merge(credit_df, on='loan_application_id')
df = df.merge(decision_df, on='loan_application_id')

# Step 5: Feature Selection
features = df[[
    'loan_amount',
    'loan_purpose',
    'employment_status',
    'annual_income',
    'dti_ratio',
    'credit_score',          # From credit_reports
    'credit_score_user',     # From users
    'delinquency_flag'
]].rename(columns={'credit_score_user': 'user_credit_score'})

target = df['ai_decision'].apply(lambda x: 1 if x.lower() == 'approved' else 0)

# Step 6: Encode categoricals
for col in ['loan_purpose', 'employment_status']:
    features[col] = LabelEncoder().fit_transform(features[col])

# Step 7: Train
X_train, X_test, y_train, y_test = train_test_split(features, target, test_size=0.2, random_state=42)
model = XGBClassifier(use_label_encoder=False, eval_metric='logloss')
model.fit(X_train, y_train)

# Step 8: Evaluation (Optional)
y_pred = model.predict(X_test)
print("Accuracy:", accuracy_score(y_test, y_pred))
print("Report:\n", classification_report(y_test, y_pred))

# Step 9: Save model to pkl
joblib.dump(model, "loan_model.pkl")
