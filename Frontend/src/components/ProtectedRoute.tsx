import { ReactNode } from "react";
import { Navigate } from "react-router-dom";
import { getToken } from "../utils/Auth";

interface ProtectedRouteProps {
  children: ReactNode;
}

export default function ProtectedRoute({ children }: ProtectedRouteProps) {
  const token = getToken(); // Replace with sessionStorage if needed

  if (!token) {
    return <Navigate to="/login" replace />;
  }

  return <>{children}</>; // Wrapping children in a fragment
}
