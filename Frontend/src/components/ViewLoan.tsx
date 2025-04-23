import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { loanClient } from "../services/Grpc";
import { LoanApplicationResponse } from "../proto/loan";

const LoanDetailsPage = () => {
const { loanId } = useParams();
  const [loanDetails, setLoanDetails] = useState<LoanApplicationResponse | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (loanId) {
        console.log("Fetching loan details for ID:", loanId);
      loanClient.GetLoanApplicationDetails({ loanId: Number(loanId) })
        .then((res) => {
          setLoanDetails(res);
          setLoading(false);
        })
        .catch((err) => {
          console.error("Error fetching loan details:", err);
          setLoading(false);
        });
    }
  }, [loanId]);

  if (loading) {
    return <div className="text-center mt-8">Loading...</div>;
  }

  if (!loanDetails) {
    return <div className="text-center mt-8 text-red-600">Loan details not found.</div>;
  }

  return (
    <div className="max-w-4xl mx-auto px-4 py-6">
      <h1 className="text-2xl font-semibold mb-4 text-center">Loan Application Details</h1>
      <div className="bg-white shadow-lg rounded-lg p-6 space-y-4">
        <div><strong>Loan ID:</strong> {loanDetails.loanId}</div>
        <div><strong>User ID:</strong> {loanDetails.userName}</div>
        <div><strong>SSN:</strong> {loanDetails.ssn}</div>
        <div><strong>Address Area:</strong> {loanDetails.addressArea}</div>
        <div><strong>Loan Amount:</strong> ${loanDetails.loanAmount}</div>
        <div><strong>Loan Purpose:</strong> {loanDetails.loanPurpose}</div>
        <div><strong>Employment Status:</strong> {loanDetails.employmentStatus}</div>
        <div><strong>Gross Monthly Income:</strong> ${loanDetails.grossMonthlyIncome}</div>
        <div><strong>Total Monthly Debt Payment:</strong> ${loanDetails.totalMonthlyDebtPayment}</div>
        <div><strong>DTI Ratio:</strong> {loanDetails.dtiRatio}</div>
        <div><strong>Status:</strong> <span className="font-medium text-blue-600">{loanDetails.applicationStatus}</span></div>
        <div><strong>Credit Score:</strong> {loanDetails.creditScore}</div>
        <div><strong>Reasoning:</strong> {loanDetails.reasoning}</div>
        <div><strong>Created At:</strong> {loanDetails.createdAt}</div>
        <div><strong>Updated At:</strong> {loanDetails.updatedAt}</div>
      </div>
    </div>
  );
};

export default LoanDetailsPage;
