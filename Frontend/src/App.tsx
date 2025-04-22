import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "./components/Login";
import Register from "./components/Register";
import DashboardLayout from "./components/DashboardLayout";
import Dashboard from "./components/Dashboard";
import Profile from "./components/Profile";
import Header from "./components/Header";
import Footer from "./components/Footer";
//import ProtectedRoute from "./components/ProtectedRoute";
import Home from "./components/Home";
import Loan from "./components/Loan";

export default function App() {
  return (
    <Router>
        <Header/>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/" element={
            // <ProtectedRoute>
                <Home />
                // </ProtectedRoute>
            } />
        <Route path="/loan" element={
            // <ProtectedRoute>
                <Loan />
            // </ProtectedRoute>
        } />
        <Route
          path="/dashboard"
          element={
            // <ProtectedRoute>
              <DashboardLayout><Dashboard /></DashboardLayout>
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
      <Footer/>
    </Router>
  );
}