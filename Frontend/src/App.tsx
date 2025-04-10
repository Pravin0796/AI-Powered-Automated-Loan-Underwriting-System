import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
// import Home from "./pages/Home";
import Register from "./components/user.tsx";
import Login from "./components/login.tsx";
import ApplyLoan from "./services/loanService.tsx";
import GetLoanStatus from "./services/GetLoanStatus.tsx";
// import Register from "./pages/Register";
// import Dashboard from "./pages/Dashboard";

const App: React.FC = () => {
    return (
        <Router>
            <Routes>
                <Route path="/register" element={<Register />} />
                <Route path="/login" element={<Login/>} />
                <Route path="/loan" element={<ApplyLoan />} />
                <Route path="/getloan" element={<GetLoanStatus />} />
             </Routes>
        </Router>
    );
};

export default App;