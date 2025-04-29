import React from "react";

interface PaginationProps {
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
}

const Pagination: React.FC<PaginationProps> = ({
  currentPage,
  totalPages,
  onPageChange,
}) => {
  const getPageNumbers = (current: number, total: number) => {
    const delta = 2;
    const start = Math.max(1, current - delta);
    const end = Math.min(total, current + delta);
    const pages = [];

    for (let i = start; i <= end; i++) {
      pages.push(i);
    }

    return pages;
  };

  const pageNumbers = getPageNumbers(currentPage, totalPages);

  const handleFirst = () => onPageChange(1);
  const handleLast = () => onPageChange(totalPages);
  const handlePrevious = () => onPageChange(currentPage - 1);
  const handleNext = () => onPageChange(currentPage + 1);

  return (
    <div className="flex justify-center items-center space-x-1 mt-6 flex-wrap">
      <button
        onClick={handleFirst}
        disabled={currentPage === 1}
        className="px-3 py-2 bg-gray-200 hover:bg-gray-300 rounded disabled:opacity-50"
      >
        First
      </button>
      <button
        onClick={handlePrevious}
        disabled={currentPage === 1}
        className="px-3 py-2 bg-gray-200 hover:bg-gray-300 rounded disabled:opacity-50"
      >
        Prev
      </button>

      {pageNumbers.map((page) => (
        <button
          key={page}
          onClick={() => onPageChange(page)}
          className={`px-3 py-2 rounded ${
            currentPage === page
              ? "bg-blue-600 text-white"
              : "bg-gray-200 hover:bg-gray-300"
          }`}
        >
          {page}
        </button>
      ))}

      <button
        onClick={handleNext}
        disabled={currentPage === totalPages}
        className="px-3 py-2 bg-gray-200 hover:bg-gray-300 rounded disabled:opacity-50"
      >
        Next
      </button>
      <button
        onClick={handleLast}
        disabled={currentPage === totalPages}
        className="px-3 py-2 bg-gray-200 hover:bg-gray-300 rounded disabled:opacity-50"
      >
        Last
      </button>
    </div>
  );
};

export default Pagination;
