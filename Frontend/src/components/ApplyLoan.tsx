import React, { useState } from "react";
import { loanClient } from "../services/Grpc";
import { LoanRequest } from "../proto/loan";

export default function LoanApplicationForm() {
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

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: name.includes("Amount") || name.includes("Income") || name.includes("Payment") ? parseFloat(value) : value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await loanClient.ApplyForLoan(formData);
      setResponse({ loanId: res.loanId.toString(), status: res.status });
    } catch (err) {
      console.error("Error applying for loan:", err);
    }
  };

  return (
    <div className="max-w-2xl mx-auto p-4">
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
        <input
          type="text"
          name="ssn"
          placeholder="SSN"
          className="w-full p-2 border rounded"
          value={formData.ssn}
          onChange={handleChange}
          required
        />
        <select
          name="addressArea"
          className="w-full p-2 border rounded"
          value={formData.addressArea}
          onChange={handleChange}
        >
          <option value="urban">Urban</option>
          <option value="rural">Rural</option>
        </select>
        <input
          type="number"
          name="loanAmount"
          placeholder="Loan Amount"
          className="w-full p-2 border rounded"
          value={formData.loanAmount}
          onChange={handleChange}
          required
        />
        <input
          type="text"
          name="loanPurpose"
          placeholder="Loan Purpose"
          className="w-full p-2 border rounded"
          value={formData.loanPurpose}
          onChange={handleChange}
          required
        />
        <input
          type="text"
          name="employmentStatus"
          placeholder="Employment Status"
          className="w-full p-2 border rounded"
          value={formData.employmentStatus}
          onChange={handleChange}
          required
        />
        <input
          type="number"
          name="grossMonthlyIncome"
          placeholder="Gross Monthly Income"
          className="w-full p-2 border rounded"
          value={formData.grossMonthlyIncome}
          onChange={handleChange}
          required
        />
        <input
          type="number"
          name="totalMonthlyDebtPayment"
          placeholder="Total Monthly Debt Payment"
          className="w-full p-2 border rounded"
          value={formData.totalMonthlyDebtPayment}
          onChange={handleChange}
          required
        />
        <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">
          Submit Application
        </button>
      </form>

      {response && (
        <div className="mt-6 p-4 bg-green-100 text-green-700 rounded">
          <p><strong>Loan ID:</strong> {response.loanId}</p>
          <p><strong>Status:</strong> {response.status}</p>
        </div>
      )}
    </div>
  );
}
