# Step 1: Import libraries
import pandas as pd
import numpy as np
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import LabelEncoder
from xgboost import XGBClassifier
from sklearn.metrics import accuracy_score, classification_report, mean_squared_error, r2_score
import joblib
import matplotlib.pyplot as plt
from xgboost import plot_importance

# Step 2: Load preprocessed CSV
df = pd.read_csv("loan_data_preprocessed.csv")
print("✅ Loaded CSV:", df.shape)

# Step 3: Select features and target
feature_columns = [
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
]

target_column = 'ai_decision'

# Step 4: Encode categorical features
for col in ['loan_purpose', 'employment_status']:
    le = LabelEncoder()
    df[col] = le.fit_transform(df[col].astype(str))

# Step 5: Prepare data
features = df[feature_columns]
features = features.apply(pd.to_numeric, errors='coerce')
features.dropna(inplace=True)

target = df.loc[features.index, target_column].astype(int)

# Step 6: Train/test split
X_train, X_test, y_train, y_test = train_test_split(features, target, test_size=0.2, random_state=42)

# Step 7: Train the model
model = XGBClassifier(use_label_encoder=False, eval_metric='logloss')
model.fit(X_train, y_train)

# Step 8: Evaluate the model
y_pred = model.predict(X_test)
y_prob = model.predict_proba(X_test)[:, 1]

# plot_importance(model)
# plt.show()

print("✅ Accuracy:", accuracy_score(y_test, y_pred))
print("✅ Classification Report:\n", classification_report(y_test, y_pred))
print("RMSE:", mean_squared_error(y_test, y_pred, squared=False))
print("R² Score:", r2_score(y_test, y_pred))

le_loan_purpose = LabelEncoder().fit(features['loan_purpose'])
le_employment_status = LabelEncoder().fit(features['employment_status'])

joblib.dump(le_loan_purpose, "le_loan_purpose.pkl")
joblib.dump(le_employment_status, "le_employment_status.pkl")

# Step 14: Save model
joblib.dump(model, "loan_model.pkl")
print("✅ Model saved as loan_model.pkl")