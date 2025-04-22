import { useState } from "react";
import { loanClient } from "../services/Grpc";
import { LoanStatusRequest } from "../proto/loan";

export default function LoanStatus() {
  const [loanId, setLoanId] = useState("");
  const [status, setStatus] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleCheckStatus = async () => {
    if (!loanId) return;
    setLoading(true);
    setError("");
    setStatus("");
    try {
      const request: LoanStatusRequest = { loanId: Number(loanId) };
      const response = await loanClient.GetLoanStatus(request);
      setStatus(response.status);
    } catch (err) {
      setError("Failed to fetch loan status. Try again later.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="max-w-xl mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4 text-center">Check Loan Status</h1>
      <div className="space-y-4">
        <input
          type="text"
          placeholder="Enter Loan ID"
          value={loanId}
          onChange={(e) => setLoanId(e.target.value)}
          className="w-full px-4 py-2 border rounded-md"
        />
        <button
          onClick={handleCheckStatus}
          disabled={loading || !loanId}
          className="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700"
        >
          {loading ? "Checking..." : "Check Status"}
        </button>

        {status && (
          <div className="text-green-600 font-semibold text-center">
            Loan Status: {status}
          </div>
        )}
        {error && (
          <div className="text-red-500 text-center">{error}</div>
        )}
      </div>
    </div>
  );
}
