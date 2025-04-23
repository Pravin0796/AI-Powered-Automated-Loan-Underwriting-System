import React, { useEffect, useState } from "react";
import { loanClient } from "../services/Grpc";
import { LoanApplicationResponse } from "../proto/loan";
import { Loader2 } from "lucide-react";

export default function AllLoanApplicationsDashboard() {
  const [applications, setApplications] = useState<LoanApplicationResponse[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchApplications() {
      try {
        const response = await loanClient.GetAllLoanApplications({});
        console.log("Fetched applications:", response);
        setApplications(response.applications);
      } catch (err) {
        console.error("Error fetching applications:", err);
      } finally {
        setLoading(false);
      }
    }
    fetchApplications();
  }, []);

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <Loader2 className="animate-spin h-8 w-8 text-gray-500" />
      </div>
    );
  }

  return (
    <div className="p-4 md:p-8 grid gap-6 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
      {applications.map((app) => (
        <div
          key={app.loanId}
          className="bg-white shadow-lg rounded-2xl p-6 hover:shadow-xl transition-shadow duration-300 ease-in-out"
        >
          <h2 className="text-xl font-bold text-gray-800 mb-2">Loan ID: {app.loanId}</h2>
          <div className="space-y-1 text-gray-700 text-sm">
            <p><span className="font-medium">User ID:</span> {app.userId}</p>
            <p><span className="font-medium">Amount:</span> ${app.loanAmount.toLocaleString()}</p>
            <p><span className="font-medium">Status:</span> {app.applicationStatus}</p>
            <p><span className="font-medium">Purpose:</span> {app.loanPurpose}</p>
            <p><span className="font-medium">Employment:</span> {app.employmentStatus}</p>
            <p className="text-gray-500 mt-2"><span className="font-medium">Created:</span> {app.createdAt?.split("T")[0]}</p>
          </div>
        </div>
      ))}
    </div>
  );
}
