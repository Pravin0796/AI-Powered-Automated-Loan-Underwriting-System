import React, { useState, useEffect } from "react";
import { loanClient } from "../services/Grpc";
import { LoanRequest } from "../proto/loan";
import { getUserId } from "../utils/Auth";

export default function ApplyLoan() {
  const userIdStr = getUserId();
  console.log("User ID from token:", userIdStr);
  const [formData, setFormData] = useState<LoanRequest>({
    userId: userIdStr ? Number(userIdStr) : 0,
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
  const [loading, setLoading] = useState(false);

  // useEffect(() => {
  //   const userIdStr = getToken();
  //   const parsedId = parseInt(userIdStr || "0");
  //   if (!isNaN(parsedId)) {
  //     setFormData((prev) => ({ ...prev, userId: parsedId }));
  //   }
  // }, []);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: name.includes("Amount") || name.includes("Income") || name.includes("Payment")
        ? value === "" ? undefined : Number(value)
        : value,
    }));
  };

  const isFormValid = () => {
    return (
      formData.userId > 0 &&
      formData.ssn.trim() !== "" &&
      formData.loanPurpose.trim() !== "" &&
      formData.employmentStatus.trim() !== "" &&
      formData.loanAmount > 0 &&
      formData.grossMonthlyIncome > 0 &&
      formData.totalMonthlyDebtPayment >= 0
    );
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    if (!isFormValid()) {
      setError("Please fill in all required fields with valid values.");
      return;
    }

    try {
      setLoading(true);
      // Prepare the form data correctly if needed (for BigInt fields)
      const preparedFormData = {
        ...formData,
        userId: Number(formData.userId),
        loanAmount: Number(formData.loanAmount),
        grossMonthlyIncome: Number(formData.grossMonthlyIncome),
        totalMonthlyDebtPayment: Number(formData.totalMonthlyDebtPayment),
      };

      const res = await loanClient.ApplyForLoan(preparedFormData);
      setResponse({ loanId: res.loanId.toString(), status: res.status });
    } catch (err) {
      console.error("Error applying for loan:", err);
      setError("Failed to apply for loan. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="max-w-3xl mx-auto p-6 bg-white shadow-md rounded-lg my-10">
      <h2 className="text-3xl font-bold mb-6 text-blue-600 text-center">Apply for a Loan</h2>

      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block font-semibold text-gray-700 mb-1">SSN</label>
          <input
            type="text"
            name="ssn"
            placeholder="Enter your SSN"
            className="w-full p-3 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            value={formData.ssn}
            onChange={handleChange}
            required
          />
        </div>

        <div>
          <label className="block font-semibold text-gray-700 mb-1">Address Area</label>
          <select
            name="addressArea"
            className="w-full p-3 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            value={formData.addressArea}
            onChange={handleChange}
            required
          >
            <option value="urban">Urban</option>
            <option value="rural">Rural</option>
          </select>
        </div>

        <div>
          <label className="block font-semibold text-gray-700 mb-1">Loan Amount</label>
          <input
            type="number"
            name="loanAmount"
            placeholder="Enter Loan Amount"
            className="w-full p-3 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            value={formData.loanAmount || ""}
            onChange={handleChange}
            required
          />
        </div>

        <div>
          <label className="block font-semibold text-gray-700 mb-1">Loan Purpose</label>
          <input
            type="text"
            name="loanPurpose"
            placeholder="Enter Loan Purpose"
            className="w-full p-3 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            value={formData.loanPurpose}
            onChange={handleChange}
            required
          />
        </div>

        <div>
          <label className="block font-semibold text-gray-700 mb-1">Employment Status</label>
          <input
            type="text"
            name="employmentStatus"
            placeholder="Enter Employment Status"
            className="w-full p-3 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            value={formData.employmentStatus}
            onChange={handleChange}
            required
          />
        </div>

        <div>
          <label className="block font-semibold text-gray-700 mb-1">Gross Monthly Income</label>
          <input
            type="number"
            name="grossMonthlyIncome"
            placeholder="Enter Gross Monthly Income"
            className="w-full p-3 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            value={formData.grossMonthlyIncome || ""}
            onChange={handleChange}
            required
          />
        </div>

        <div>
          <label className="block font-semibold text-gray-700 mb-1">Total Monthly Debt Payment</label>
          <input
            type="number"
            name="totalMonthlyDebtPayment"
            placeholder="Enter Total Monthly Debt Payment"
            className="w-full p-3 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
            value={formData.totalMonthlyDebtPayment || ""}
            onChange={handleChange}
            required
          />
        </div>

        <button
          type="submit"
          disabled={!isFormValid() || loading}
          className={`w-full font-semibold py-3 rounded-lg transition ${
            loading || !isFormValid()
              ? "bg-gray-400 cursor-not-allowed"
              : "bg-blue-600 text-white hover:bg-blue-700"
          }`}
        >
          {loading ? "Submitting..." : "Submit Application"}
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
