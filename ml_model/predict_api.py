from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import joblib
import numpy as np

# Load trained model
model = joblib.load("loan_model.pkl")

# Load Label Encoders
le_loan_purpose = joblib.load("le_loan_purpose.pkl")
le_employment_status = joblib.load("le_employment_status.pkl")

# Define expected input schema
class LoanInput(BaseModel):
    loan_amount: float
    loan_purpose: str
    employment_status: str
    annual_income: float
    dti_ratio: float
    report_credit_score: int
    delinquency_flag: bool
    # num_payments_made: int
    # num_late_payments: int
    # total_amount_paid: float
    # payment_success_ratio: float

app = FastAPI()

# Helper function to safely transform the category values
def safe_transform(encoder, value):
    try:
        # Try to transform the value using the encoder
        return encoder.transform([value])[0]
    except ValueError:
        # If the value is not in the encoder's classes (unknown category), return a default value (e.g., 0)
        return 0

@app.post("/predict")
def predict(data: LoanInput):
    try:
        # Safely transform categorical features with the custom safe_transform function
        loan_purpose_encoded = safe_transform(le_loan_purpose, data.loan_purpose.lower())
        employment_status_encoded = safe_transform(le_employment_status, data.employment_status.lower())

        # Prepare the input array for prediction
        x_input = np.array([[ 
            data.loan_amount,
            loan_purpose_encoded,
            employment_status_encoded,
            data.annual_income,
            data.dti_ratio,
            data.report_credit_score,
            int(data.delinquency_flag),
            # data.num_payments_made,
            # data.num_late_payments,
            # data.total_amount_paid,
            # data.payment_success_ratio
        ]])

        # Make the prediction using the trained model
        prediction = model.predict(x_input)[0]
        decision = "approved" if prediction == 1 else "rejected"

        # Simple reasoning based on the prediction
        reasoning = "High credit score and low DTI" if prediction == 1 else "Low creditworthiness"

        # Return the decision and reasoning
        return {"decision": decision, "reasoning": reasoning}

    except Exception as e:
        # Handle any exceptions that occur during prediction and return an HTTP error
        raise HTTPException(status_code=500, detail=str(e))



# commAN to run
# uvicorn predict_api:app --reload --port 8000