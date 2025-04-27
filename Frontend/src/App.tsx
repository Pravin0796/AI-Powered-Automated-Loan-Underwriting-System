import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "./components/Login";
import Register from "./components/Register";
import DashboardLayout from "./components/DashboardLayout";
//import Dashboard from "./components/Dashboard";
import Profile from "./components/Profile";
import Header from "./components/Header";
import Footer from "./components/Footer";
//import ProtectedRoute from "./components/ProtectedRoute";
import Home from "./components/Home";
import LoanStatus from "./components/LoanStatus";
import ApplyLoan from "./components/ApplyLoan";
import LoanDetailsPage from "./components/ViewLoan";
import ViewAllLoan from "./components/ViewAllLoan";
import Dashboard from "./components/Dashboard";

import LoanManagement from "./components/LoanManagement";
import UserManagement from "./components/UserManagement";
import Notifications from "./components/Notification";
import Settings from "./components/Settings";

export default function App() {

  const sampleNotifications = [
    "Your loan application has been approved!",
    "Reminder: Complete your KYC verification.",
    "New loan offer available: Check it out now!",
  ];

  return (
    <Router>
      <div className="flex flex-col min-h-screen">
        <Header />
        <main className="flex-grow mt-15">

          {/* <Routes>
          <Route path="/admin">
          <Header1/>
            <Route path="" element={<Sidebar1 />}/>
            <Route path="/admin/dashboard" element={<Dashboard1 />} />
            <Route path="/admin/loans" element={<LoanManagement />} />
            <Route path="/admin/users" element={<UserManagement />} />
            <Route path="/admin/notifications" element={<Notifications1 />} />
            <Route path="/admin/settings" element={<Settings1 />} />
          </Route>


            <Route path="/dashboard" element={<Dashboard />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/" element={
              // <ProtectedRoute>
              <Home />
              // </ProtectedRoute>
            } />
            <Route path="/viewloan" element={
              // <ProtectedRoute>
              <LoanStatus />
              // </ProtectedRoute>
            } />
            <Route path="/loan/:id" element={
              // <ProtectedRoute>
              <LoanDetailsPage />
              // </ProtectedRoute>
            } />
            <Route
              path="/applyloan"
              element={
                // <ProtectedRoute>
                <ApplyLoan />
                // </ProtectedRoute>
              }
            />

            <Route
              path="/loan"
              element={
                // <ProtectedRoute>
                <ViewAllLoan />
                // </ProtectedRoute>
              }
            />

            <Route
              path="/profile"
              element={
                // <ProtectedRoute>
                <DashboardLayout><Profile /></DashboardLayout>
                // </ProtectedRoute>
              }
            />
          </Routes> */}

<Routes>
  {/* Public user routes */}
  <Route path="/dashboard" element={<Dashboard />} />
  <Route path="/login" element={<Login />} />
  <Route path="/register" element={<Register />} />
  <Route path="/" element={<Home />} />
  <Route path="/viewloan" element={<LoanStatus />} />
  <Route path="/loan/:id" element={<LoanDetailsPage />} />
  <Route path="/applyloan" element={<ApplyLoan />} />
  <Route path="/loan" element={<ViewAllLoan />} />
  <Route path="/profile" element={<DashboardLayout><Profile /></DashboardLayout>} />

    <Route path="loans" element={<LoanManagement />} />
    <Route path="users" element={<UserManagement />} />
    <Route path="notifications" element={<Notifications notifications={sampleNotifications} />} />
    <Route path="settings" element={<Settings />} />
</Routes>

        </main>
        <Footer />
      </div>
    </Router>
  );
}