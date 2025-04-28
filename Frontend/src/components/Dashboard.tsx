import React, { useEffect, useState } from "react";
import LoanStatusCharts from "../components/LoanStatusCharts";
import { loanClient } from "../services/Grpc"; // Assuming loanClient is exported from loanService

const Dashboard = () => {
  const [loanStats, setLoanStats] = useState({
    totalLoans: 0,
    approvedLoans: 0,
    rejectedLoans: 0,
    pendingLoans: 0,
  });
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchLoanStats = async () => {
      try {
        const response = await loanClient.GetLoanStatusCount({});
        setLoanStats({
          totalLoans: response.totalApplications,
          approvedLoans: response.approved,
          rejectedLoans: response.rejected,
          pendingLoans: response.pending,
        });
        console.log("Fetched loan stats:", response);
      } catch (err) {
        setError("Failed to fetch loan statistics. Please try again later.");
      } finally {
        setLoading(false);
      }
    };

    fetchLoanStats();
  }, []);

  if (loading) {
    return <p>Loading...</p>;
  }

  if (error) {
    return <p className="text-red-600">{error}</p>;
  }

  return (
    <div className="p-6 bg-gray-50 min-h-screen">
      <h2 className="text-2xl font-bold mb-6 text-blue-700">Dashboard</h2>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="bg-white shadow-md rounded-lg p-6 text-center">
          <h3 className="text-lg font-semibold text-gray-700">Total Loans</h3>
          <p className="text-3xl font-bold text-blue-600">{loanStats.totalLoans}</p>
        </div>
        <div className="bg-white shadow-md rounded-lg p-6 text-center">
          <h3 className="text-lg font-semibold text-gray-700">Approved Loans</h3>
          <p className="text-3xl font-bold text-green-600">{loanStats.approvedLoans}</p>
        </div>
        <div className="bg-white shadow-md rounded-lg p-6 text-center">
          <h3 className="text-lg font-semibold text-gray-700">Rejected Loans</h3>
          <p className="text-3xl font-bold text-red-600">{loanStats.rejectedLoans}</p>
        </div>
        <div className="bg-white shadow-md rounded-lg p-6 text-center">
          <h3 className="text-lg font-semibold text-gray-700">Pending Loans</h3>
          <p className="text-3xl font-bold text-yellow-600">{loanStats.pendingLoans}</p>
        </div>
      </div>
      <div className="mt-8">
        <LoanStatusCharts />
      </div>
    </div>
  );
};

export default Dashboard;