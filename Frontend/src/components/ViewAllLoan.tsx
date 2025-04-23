import { useEffect, useState } from "react";
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
        setApplications(response.applications || []);
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
    <div className="overflow-x-auto p-4">
      <table className="min-w-full bg-white border border-gray-200 shadow-md rounded-xl text-sm">
        <thead className="bg-gray-100">
          <tr>
            <th className="px-4 py-2">Loan ID</th>
            <th className="px-4 py-2">User Name</th>
            <th className="px-4 py-2">Loan Amount</th>
            <th className="px-4 py-2">Purpose</th>
            <th className="px-4 py-2">Employment</th>
            <th className="px-4 py-2">Income</th>
            <th className="px-4 py-2">DTI</th>
            <th className="px-4 py-2">Status</th>
            <th className="px-4 py-2">Score</th>
            <th className="px-4 py-2">Reasoning</th>
            <th className="px-4 py-2">Updated</th>
          </tr>
        </thead>
        <tbody>
          {applications.map((app) => (
            <tr key={app.loanId} className="border-t text-center hover:bg-gray-50">
              <td className="px-4 py-2">{app.loanId}</td>
              <td className="px-4 py-2">{app.userName}</td>
              <td className="px-4 py-2">${app.loanAmount?.toLocaleString()}</td>
              <td className="px-4 py-2">{app.loanPurpose}</td>
              <td className="px-4 py-2">{app.employmentStatus}</td>
              <td className="px-4 py-2">${app.grossMonthlyIncome}</td>
              <td className="px-4 py-2">{app.dtiRatio}</td>
              <td className="px-4 py-2">{app.applicationStatus}</td>
              <td className="px-4 py-2">{app.creditScore}</td>
              <td className="px-4 py-2">{app.reasoning}</td>
              <td className="px-4 py-2">{app.createdAt?.split("T")[0]}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
