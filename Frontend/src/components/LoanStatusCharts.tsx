import React, { useEffect, useState } from "react";
import { loanClient } from "../services/Grpc";
import { Empty } from "../proto/loan";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

const LoanStatsChart = () => {
  const [stats, setStats] = useState({
    total_applications: 0,
    approved: 0,
    rejected: 0,
    pending: 0,
  });

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const response = await loanClient.GetLoanStats(Empty.create());
        setStats(response);
      } catch (error) {
        console.error("Error fetching stats:", error);
      }
    };

    fetchStats();
  }, []);

  const data = [
    { name: "Total", count: stats.total_applications },
    { name: "Approved", count: stats.approved },
    { name: "Rejected", count: stats.rejected },
    { name: "Pending", count: stats.pending },
  ];

  return (
    <div className="w-full max-w-4xl mx-auto mt-8">
      <h2 className="text-2xl font-bold mb-4 text-center">Loan Application Stats</h2>
      <ResponsiveContainer width="100%" height={300}>
        <BarChart data={data} margin={{ top: 20, right: 30, left: 0, bottom: 5 }}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="name" />
          <YAxis />
          <Tooltip />
          <Bar dataKey="count" fill="#3b82f6" radius={[6, 6, 0, 0]} />
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
};

export default LoanStatsChart;
