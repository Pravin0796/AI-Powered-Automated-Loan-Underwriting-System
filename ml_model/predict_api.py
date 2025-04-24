from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import joblib
import numpy as np

# Load trained model
model = joblib.load("loan_model.pkl")

# Label encoding maps (match with training phase)
label_maps = {
    'loan_purpose': {'home': 0, 'car': 1, 'business': 2},
    'employment_status': {'employed': 0, 'self-employed': 1, 'unemployed': 2}
}

# Define expected input schema
class LoanInput(BaseModel):
    loan_amount: float
    loan_purpose: str
    employment_status: str
    annual_income: float
    dti_ratio: float
    report_credit_score: int
    user_credit_score: int
    delinquency_flag: bool
    num_payments_made: int
    num_late_payments: int
    total_amount_paid: float
    payment_success_ratio: float

app = FastAPI()

@app.post("/predict")
def predict(data: LoanInput):
    try:
        x_input = np.array([[
            data.loan_amount,
            label_maps['loan_purpose'].get(data.loan_purpose.lower(), 0),
            label_maps['employment_status'].get(data.employment_status.lower(), 0),
            data.annual_income,
            data.dti_ratio,
            data.report_credit_score,
            data.user_credit_score,
            int(data.delinquency_flag),
            data.num_payments_made,
            data.num_late_payments,
            data.total_amount_paid,
            data.payment_success_ratio
        ]])

        prediction = model.predict(x_input)[0]
        return {"decision": "approved" if prediction == 1 else "rejected"}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
