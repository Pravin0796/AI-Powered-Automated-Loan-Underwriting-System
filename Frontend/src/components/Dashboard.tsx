// export default function Dashboard() {
//     return (
//       <div className="p-6 bg-gray-100 min-h-screen">
//         <h1 className="text-3xl font-bold text-blue-800 mb-4">Dashboard</h1>
//         <p className="text-gray-600">
//           Welcome to the Loan Underwriting System Dashboard. Use the navigation menu to explore loan applications, analytics, and your profile.
//         </p>
//       </div>
//     );
//   }

import React from "react";
import LoanStatusCharts from "../components/LoanStatusCharts";

const Dashboard = () => {
  return (
    <div className="p-6">
      <h2 className="text-2xl font-bold mb-4">Dashboard</h2>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div className="bg-white shadow-md rounded-lg p-4">
          <h3 className="text-lg font-semibold">Total Loans</h3>
          <p className="text-2xl font-bold">1,234</p>
        </div>
        <div className="bg-white shadow-md rounded-lg p-4">
          <h3 className="text-lg font-semibold">Approved Loans</h3>
          <p className="text-2xl font-bold">567</p>
        </div>
      </div>
      <LoanStatusCharts />
    </div>
  );
};

export default Dashboard;