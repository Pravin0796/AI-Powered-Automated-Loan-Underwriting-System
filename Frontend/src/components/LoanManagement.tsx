import { useState } from "react";
import Pagination from "./Pagination";
import ViewAllLoans from "./ViewAllLoan";

const LoanManagement = () => {
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1); // Dynamically set this

  const handlePageChange = (page: number) => {
    console.log(`Fetching data for page ${page}`);
    setCurrentPage(page);
  };

  return (
    <div className="p-6 bg-gray-50 min-h-screen">
      <h2 className="text-2xl font-bold mb-4 text-blue-600">Loan Management</h2>
      <div className="bg-white p-4 rounded shadow-md">
        <ViewAllLoans
          currentPage={currentPage}
          setTotalPages={setTotalPages}
        />
      </div>
      <Pagination
        currentPage={currentPage}
        totalPages={totalPages}
        onPageChange={handlePageChange}
      />
    </div>
  );
};

export default LoanManagement;
