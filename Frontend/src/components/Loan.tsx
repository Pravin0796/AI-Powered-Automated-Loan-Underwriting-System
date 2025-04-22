// src/pages/Loan.tsx
import React from "react";

const Loan = () => {
  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-4">Your Loans</h1>
      <p className="text-gray-700 mb-6">Here you can view your loan applications and their statuses.</p>

      {/* Placeholder Table */}
      <div className="overflow-x-auto">
        <table className="min-w-full bg-white shadow-md rounded-2xl overflow-hidden">
          <thead className="bg-gray-100">
            <tr>
              <th className="py-2 px-4 text-left">Loan ID</th>
              <th className="py-2 px-4 text-left">Amount</th>
              <th className="py-2 px-4 text-left">Purpose</th>
              <th className="py-2 px-4 text-left">Status</th>
            </tr>
          </thead>
          <tbody>
            <tr className="border-t">
              <td className="py-2 px-4">#1001</td>
              <td className="py-2 px-4">$10,000</td>
              <td className="py-2 px-4">Education</td>
              <td className="py-2 px-4 text-green-600">Approved</td>
            </tr>
            <tr className="border-t">
              <td className="py-2 px-4">#1002</td>
              <td className="py-2 px-4">$5,000</td>
              <td className="py-2 px-4">Medical</td>
              <td className="py-2 px-4 text-yellow-600">Pending</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default Loan;
