import React from "react";
import ViewAllLoan from "./ViewAllLoan";

const LoanManagement = () => {

  function LoanList() {
    const [currentPage, setCurrentPage] = useState(1);
    const totalPages = 5; // Example total pages
  
    const handlePageChange = (page: number) => {
      console.log(`Fetching data for page ${page}`);
      setCurrentPage(page);
    };


  return (
    <div className="p-6">
      <h2 className="text-2xl font-bold mb-4">Loan Management</h2>
      <ViewAllLoan />
      <Pagination
        currentPage={currentPage}
        totalPages={totalPages}
        onPageChange={handlePageChange}
      />
    </div>
  );
};

export default LoanManagement;