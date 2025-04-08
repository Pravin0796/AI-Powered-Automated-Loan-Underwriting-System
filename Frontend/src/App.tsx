import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
// import Home from "./pages/Home";
import Auth from "./components/user.tsx";
// import Register from "./pages/Register";
// import Dashboard from "./pages/Dashboard";

const App: React.FC = () => {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<Auth />} />
            </Routes>
        </Router>
    );
};

export default App;