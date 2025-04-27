import React from "react";
import LoanStatusCharts from "../components/LoanStatusCharts";

const Dashboard = () => {
  return (
    <div className="p-6 bg-gray-50 min-h-screen">
      <h2 className="text-2xl font-bold mb-6 text-blue-700">Dashboard</h2>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="bg-white shadow-md rounded-lg p-6 text-center">
          <h3 className="text-lg font-semibold text-gray-700">Total Loans</h3>
          <p className="text-3xl font-bold text-blue-600">1,234</p>
        </div>
        <div className="bg-white shadow-md rounded-lg p-6 text-center">
          <h3 className="text-lg font-semibold text-gray-700">Approved Loans</h3>
          <p className="text-3xl font-bold text-green-600">567</p>
        </div>
      </div>
      <div className="mt-8">
        <LoanStatusCharts />
      </div>
    </div>
  );
};

export default Dashboard;