from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import joblib
import numpy as np

model = joblib.load("loan_model.pkl")

class LoanInput(BaseModel):
    loan_amount: float
    loan_purpose: str
    employment_status: str
    annual_income: float
    dti_ratio: float
    credit_score: int
    user_credit_score: int
    delinquency_flag: bool

app = FastAPI()

# Dummy label encoders (match with training)
label_maps = {
    'loan_purpose': {'home': 0, 'car': 1, 'business': 2},
    'employment_status': {'employed': 0, 'self-employed': 1, 'unemployed': 2}
}

@app.post("/predict")
def predict(data: LoanInput):
    try:
        x_input = np.array([[
            data.loan_amount,
            label_maps['loan_purpose'].get(data.loan_purpose.lower(), 0),
            label_maps['employment_status'].get(data.employment_status.lower(), 0),
            data.annual_income,
            data.dti_ratio,
            data.credit_score,
            data.user_credit_score,
            int(data.delinquency_flag)
        ]])
        prediction = model.predict(x_input)[0]
        return {"decision": "approved" if prediction == 1 else "rejected"}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
