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

export default function App() {
  return (
    <Router>
      <div className="flex flex-col min-h-screen">
        <Header/>
        <main className="flex-grow mt-15">
      <Routes>
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
          path="/profile"
          element={
            // <ProtectedRoute>
              <DashboardLayout><Profile /></DashboardLayout>
            // </ProtectedRoute>
          }
        />
      </Routes>
      </main>
      <Footer/>
      </div>
    </Router>
  );
}