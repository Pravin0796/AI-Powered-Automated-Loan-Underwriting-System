// import React from "react";
// import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
// import Sidebar from "../components/admin/Sidebar";
// import Header from "../components/admin/Header";
// import Dashboard from "../components/admin/Dashboard";
// import LoanManagement from "../components/admin/LoanManagement";
// import UserManagement from "../components/UserManagement";
// import Notifications from "../components/Notifications";
// import Settings from "../components/Settings";

// const AdminDashboard = () => {
//   return (
//     <Router>
//       <div className="flex">
//         <Sidebar />
//         <div className="flex-1">
//           <Header />
//           <Routes>
//             <Route path="/admin/dashboard" element={<Dashboard />} />
//             <Route path="/admin/loans" element={<LoanManagement />} />
//             <Route path="/admin/users" element={<UserManagement />} />
//             <Route path="/admin/notifications" element={<Notifications />} />
//             <Route path="/admin/settings" element={<Settings />} />
//           </Routes>
//         </div>
//       </div>
//     </Router>
//   );
// };

// export default AdminDashboard;