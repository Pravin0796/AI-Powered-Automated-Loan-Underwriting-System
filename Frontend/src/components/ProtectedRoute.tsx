import React from "react";
import { Navigate, Outlet } from "react-router-dom";
import { getUserRole, unAuthorized } from "../utils/Auth"; // Import role utility functions
import { toast } from 'react-toastify';

type ProtectedRouteProps = {
  allowedRoles?: string[]; 
};

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ allowedRoles }) => {
  const tokenValid = unAuthorized(); // Check if token is present
  const role = getUserRole(); // Get user role from token

  if (!tokenValid || (allowedRoles && !allowedRoles.includes(role))) {
    toast.error("You are not authorized to access this page.");
    return <Navigate to="/login" replace />;
  }

  return <Outlet />;
};

export default ProtectedRoute;



// import { ReactNode } from "react";
// import { Navigate } from "react-router-dom";
// import { getToken } from "../utils/Auth";

// interface ProtectedRouteProps {
//   children: ReactNode;
// }

// export default function ProtectedRoute({ children }: ProtectedRouteProps) {
//   const token = getToken(); // Replace with sessionStorage if needed

//   if (!token) {
//     return <Navigate to="/login" replace />;
//   }

//   return <>{children}</>; // Wrapping children in a fragment
// }


