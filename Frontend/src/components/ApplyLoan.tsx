import React, { useState } from "react";
import { loanClient } from "../services/Grpc";
import { LoanRequest } from "../proto/loan";

export default function ApplyLoan() {
  const [formData, setFormData] = useState<LoanRequest>({
    userId: 0,
    ssn: "",
    addressArea: "urban",
    loanAmount: 0,
    loanPurpose: "",
    employmentStatus: "",
    grossMonthlyIncome: 0,
    totalMonthlyDebtPayment: 0,
  });

  const [response, setResponse] = useState<{ loanId: string; status: string } | null>(null);
  const [error, setError] = useState<string | null>(null);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: name.includes("Amount") || name.includes("Income") || name.includes("Payment")
        ? parseFloat(value)
        : value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    try {
      const res = await loanClient.ApplyForLoan(formData);
      setResponse({ loanId: res.loanId.toString(), status: res.status });
    } catch (err) {
      console.error("Error applying for loan:", err);
      setError("Failed to apply for loan. Please try again.");
    }
  };

  return (
    <div className="max-w-2xl mx-auto p-6 bg-white shadow-md rounded-lg">
      <h2 className="text-2xl font-bold mb-4">Apply for a Loan</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <input
          type="number"
          name="userId"
          placeholder="User ID"
          className="w-full p-2 border rounded"
          value={formData.userId}
          onChange={handleChange}
          required
        />
        {/* Other input fields go here */}
        <button
          type="submit"
          className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
        >
          Submit Application
        </button>
      </form>

      {response && (
        <div className="mt-6 p-4 bg-green-100 text-green-700 rounded">
          <p><strong>Loan ID:</strong> {response.loanId}</p>
          <p><strong>Status:</strong> {response.status}</p>
        </div>
      )}
      {error && <div className="mt-4 text-red-600">{error}</div>}
    </div>
  );
}