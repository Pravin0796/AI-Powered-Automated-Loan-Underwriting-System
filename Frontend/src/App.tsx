import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "./components/Login";
import Register from "./components/Register";
import DashboardLayout from "./components/DashboardLayout";
import Profile from "./components/Profile";
import Header from "./components/Header";
import Footer from "./components/Footer";
import Home from "./components/Home";
import LoanStatus from "./components/LoanStatus";
import ApplyLoan from "./components/ApplyLoan";
import LoanDetailsPage from "./components/ViewLoan";
import ViewAllLoan from "./components/ViewAllLoan";
import Dashboard from "./components/Dashboard";
import ErrorBoundary from "./components/ErrorBoundary";
import LoanManagement from "./components/LoanManagement";
import UserManagement from "./components/UserManagement";
import Notifications from "./components/Notification";
import Settings from "./components/Settings";
// import ProtectedRoute from "./components/ProtectedRoute";
import ViewCreditPage from "./components/ViewCreditPage";
import { getUserRole } from "./utils/Auth"; // Assuming this function gets the token


export default function App() {
  const sampleNotifications = [
    "Your loan application has been approved!",
    "Reminder: Complete your KYC verification.",
    "New loan offer available: Check it out now!",
  ];

  const userRole = getUserRole(); // Get user role from token

  return (
    <Router>
      <ErrorBoundary>
        <div className="flex flex-col min-h-screen">
          <Header />
          <main className="flex-grow mt-15">
            <Routes>
              {/* Public Routes */}
              <Route path="/" element={<Home />} />
              <Route path="/login" element={<Login />} />
              <Route path="/register" element={<Register />} />
              <Route path="/viewloan" element={<LoanStatus />} />
              <Route path="/loan/:id" element={<LoanDetailsPage />} />
              <Route path="/loan" element={<ViewAllLoan />} />
              <Route path="/viewcredit" element={<ViewCreditPage />} />

              {/* Protected Routes - User only */}
              {userRole === "user" && (
                <>
                  <Route
                    path="/profile"
                    element={<DashboardLayout><Profile /></DashboardLayout>}
                  />
                  <Route
                    path="/loan-applications"
                    element={<DashboardLayout><ApplyLoan /></DashboardLayout>}
                  />
                  <Route
                    path="/notifications"
                    element={<Notifications notifications={sampleNotifications} />}
                  />
                </>
              )}

              {/* Protected Routes - Admin only */}
              {userRole === "admin" && (
                <>
                  <Route path="/dashboard" element={<Dashboard />} />
                  <Route
                    path="/loans"
                    element={<LoanManagement />}
                  />
                  <Route
                    path="/users"
                    element={<UserManagement />}
                  />
                  <Route path="/settings" element={<Settings />} />
                </>
              )}
            </Routes>
          </main>
          <Footer />
        </div>
      </ErrorBoundary>
    </Router>
  );
}
