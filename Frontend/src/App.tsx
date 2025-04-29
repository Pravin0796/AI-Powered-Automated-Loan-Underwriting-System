
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "./components/Login";
import Register from "./components/Register";
import DashboardLayout from "./components/DashboardLayout";
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import Profile from "./components/Profile";
import Header from "./components/Header";
import Footer from "./components/Footer";
import Home from "./components/Home";
import LoanStatus from "./components/LoanStatus";
import ApplyLoan from "./components/ApplyLoan";
import LoanDetailsPage from "./components/LoanDetailsPage";
// import ViewAllLoan from "./components/ViewAllLoan";
import Dashboard from "./components/Dashboard";
import ErrorBoundary from "./components/ErrorBoundary";
import LoanManagement from "./components/LoanManagement";
import UserManagement from "./components/UserManagement";
import Notifications from "./components/Notification";
import Settings from "./components/Settings";
import ProtectedRoute from "./components/ProtectedRoute";
import ViewCreditPage from "./components/ViewCreditPage";
import FAQ from "./components/FAQ";
import ContactUs from "./components/ContactUs";

export default function App() {
  const sampleNotifications = [
    "Your loan application has been approved!",
    "Reminder: Complete your KYC verification.",
    "New loan offer available: Check it out now!",
  ];

  

  return (
    <Router>
      <ErrorBoundary>
        <div className="flex flex-col min-h-screen">
          <Header />
          <main className="flex-grow mt-15">
            <Routes>
              <Route path="/login" element={<Login />} />
              <Route path="/register" element={<Register />} />
              <Route path="/" element={<Home />} />
              <Route path="/contact" element={<ContactUs />} />
              <Route path="/faq" element={<FAQ />} />

              <Route element={<ProtectedRoute allowedRoles={["user", "admin"]} />}>
                <Route path="/loan-applications" element={<ApplyLoan />} />
                <Route path="/loan-status" element={<LoanStatus />} />
                <Route path="/loan/:id" element={<LoanDetailsPage />} />
                <Route path="/viewcredit" element={<ViewCreditPage />} />
                <Route path="/profile" element={<Profile />} />
                <Route path="/notifications" element={<Notifications notifications={sampleNotifications} />} />
                <Route path="/settings" element={<Settings />} />
              </Route>

              <Route element={<ProtectedRoute allowedRoles={["admin"]} />}>
                <Route path="/dashboard" element={<Dashboard />} />
                <Route path="/loans" element={<LoanManagement />} />
                <Route path="/users" element={<UserManagement />} />
              </Route>
            </Routes>
          </main>
          <Footer />
          <ToastContainer />
        </div>
      </ErrorBoundary>
    </Router>
  );
}


{/* <Route element={<ProtectedRoute allowedRoles={["customer", "manager", "admin"]} />}>
<Route>

  <Route element={<ProtectedRoute allowedRoles={["admin"]} />}>
<Route> */}