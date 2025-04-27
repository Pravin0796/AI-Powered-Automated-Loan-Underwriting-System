// import React from "react";

// interface PaginationProps {
//   currentPage: number;
//   totalPages: number;
//   onPageChange: (page: number) => void;
// }

// const Pagination: React.FC<PaginationProps> = ({
//   currentPage,
//   totalPages,
//   onPageChange,
// }) => {
//   const pages = Array.from({ length: totalPages }, (_, i) => i + 1);

//   return (
//     <div className="flex justify-center space-x-2 p-4">
//       <button
//         disabled={currentPage === 1}
//         onClick={() => onPageChange(currentPage - 1)}
//         className="px-3 py-1 bg-gray-200 hover:bg-gray-300 rounded disabled:opacity-50"
//       >
//         Previous
//       </button>
//       {pages.map((page) => (
//         <button
//           key={page}
//           onClick={() => onPageChange(page)}
//           className={`px-3 py-1 rounded ${
//             page === currentPage
//               ? "bg-blue-600 text-white"
//               : "bg-gray-200 hover:bg-gray-300"
//           }`}
//         >
//           {page}
//         </button>
//       ))}
//       <button
//         disabled={currentPage === totalPages}
//         onClick={() => onPageChange(currentPage + 1)}
//         className="px-3 py-1 bg-gray-200 hover:bg-gray-300 rounded disabled:opacity-50"
//       >
//         Next
//       </button>
//     </div>
//   );
// };

// export default Pagination;


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
  const pages = Array.from({ length: totalPages }, (_, i) => i + 1);

  return (
    <div className="flex justify-center mt-4">
      <button
        disabled={currentPage === 1}
        onClick={() => onPageChange(currentPage - 1)}
        className="px-3 py-2 bg-gray-200 hover:bg-gray-300 rounded disabled:opacity-50"
      >
        Previous
      </button>
      {pages.map((page) => (
        <button
          key={page}
          onClick={() => onPageChange(page)}
          className={`px-3 py-2 mx-1 rounded ${
            currentPage === page
              ? "bg-blue-600 text-white"
              : "bg-gray-200 hover:bg-gray-300"
          }`}
        >
          {page}
        </button>
      ))}
      <button
        disabled={currentPage === totalPages}
        onClick={() => onPageChange(currentPage + 1)}
        className="px-3 py-2 bg-gray-200 hover:bg-gray-300 rounded disabled:opacity-50"
      >
        Next
      </button>
    </div>
  );
};

export default Pagination;