import React, { useEffect, useState } from "react";
import { loanClient } from "../services/Grpc";
import { LoanApplicationResponse } from "../proto/loan";
import { Card, CardContent } from "@/components/ui/card";
import { Loader2 } from "lucide-react";

export default function AllLoanApplicationsDashboard() {
  const [applications, setApplications] = useState<LoanApplicationResponse[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchApplications() {
      try {
        const response = await loanClient.GetAllLoanApplications({});
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
    <div className="p-4 md:p-8 grid gap-4 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
      {applications.map((app) => (
        <Card key={app.loanId} className="shadow-md rounded-2xl">
          <CardContent className="p-4">
            <h2 className="text-lg font-semibold">Loan ID: {app.loanId}</h2>
            <p><span className="font-medium">User ID:</span> {app.userId}</p>
            <p><span className="font-medium">Amount:</span> ${app.loanAmount.toLocaleString()}</p>
            <p><span className="font-medium">Status:</span> {app.applicationStatus}</p>
            <p><span className="font-medium">Purpose:</span> {app.loanPurpose}</p>
            <p><span className="font-medium">Employment:</span> {app.employmentStatus}</p>
            <p className="text-sm text-gray-500 mt-1">Created: {app.createdAt?.split("T")[0]}</p>
          </CardContent>
        </Card>
      ))}
    </div>
  );
}
