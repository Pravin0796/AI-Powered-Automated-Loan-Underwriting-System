import { useState } from "react";
import { NavLink, useNavigate } from "react-router-dom";
import { FaBars, FaTimes, FaChartLine, FaUser, FaSignOutAlt, FaFileInvoiceDollar, FaHome } from "react-icons/fa";
import classNames from "classnames";
import { removeToken } from "../utils/Auth";

const navItemClass = (isActive: boolean) =>
  classNames(
    "flex items-center gap-3 px-4 py-3 rounded-md hover:bg-gray-100 transition-all",
    {
      "bg-blue-100 text-blue-600 font-semibold": isActive,
      "text-gray-700": !isActive,
    }
  );

export default function Sidebar() {
  const [isCollapsed, setIsCollapsed] = useState(false);
  const navigate = useNavigate();

  const handleLogout = () => {
    removeToken();
    navigate("/login");
  };

  return (
    <div className="flex">
      {/* Sidebar Toggle Button */}
      <button
        className="p-4 text-2xl text-gray-600 md:hidden focus:outline-none"
        onClick={() => setIsCollapsed(!isCollapsed)}
      >
        {isCollapsed ? <FaTimes /> : <FaBars />}
      </button>

      {/* Sidebar */}
      <div
        className={classNames(
          "fixed top-0 left-0 w-64 h-full bg-white shadow-lg space-y-6 p-6 transform transition-transform md:relative md:translate-x-0",
          {
            "translate-x-0": isCollapsed,
            "-translate-x-full": !isCollapsed,
          }
        )}
      >
        {/* Sidebar Header */}
        <div className="text-center text-xl font-bold text-blue-600">
          Loan Underwriting
        </div>

        {/* Navigation Links */}
        <nav className="space-y-2">
          <NavLink
            to="/dashboard"
            className={({ isActive }) => navItemClass(isActive)}
          >
            <FaHome />
            Dashboard
          </NavLink>
          <NavLink
            to="/loan-applications"
            className={({ isActive }) => navItemClass(isActive)}
          >
            <FaFileInvoiceDollar />
            Loan Applications
          </NavLink>
          <NavLink
            to="/analytics"
            className={({ isActive }) => navItemClass(isActive)}
          >
            <FaChartLine />
            Analytics
          </NavLink>
          <NavLink
            to="/profile"
            className={({ isActive }) => navItemClass(isActive)}
          >
            <FaUser />
            Profile
          </NavLink>
        </nav>

        {/* Logout Button */}
        <button
          onClick={handleLogout}
          className="flex items-center gap-3 px-4 py-3 text-red-500 hover:bg-red-100 rounded-md transition-all w-full"
        >
          <FaSignOutAlt />
          Logout
        </button>
      </div>

      {/* Main Content Placeholder */}
      <div className="flex-grow p-6">
        {/* Content for the selected page will be displayed here */}
      </div>
    </div>
  );
}


// import { NavLink } from "react-router-dom";

// const Sidebar = () => {
//   const links = [
//     { name: "Dashboard", path: "/admin/dashboard" },
//     { name: "Manage Loans", path: "/admin/loans" },
//     { name: "User Management", path: "/admin/users" },
//     { name: "Reports", path: "/admin/reports" },
//     { name: "Notifications", path: "/admin/notifications" },
//     { name: "Settings", path: "/admin/settings" },
//   ];

//   return (
//     <aside className="w-64 bg-blue-600 text-white h-screen p-4">
//       <h2 className="text-2xl font-bold mb-6">Admin Panel</h2>
//       <nav>
//         {links.map((link) => (
//           <NavLink
//             key={link.name}
//             to={link.path}
//             className={({ isActive }) =>
//               `block py-2 px-4 rounded hover:bg-blue-700 ${isActive ? "bg-blue-800" : ""}`
//             }
//           >
//             {link.name}
//           </NavLink>
//         ))}
//       </nav>
//     </aside>
//   );
// };

// export default Sidebar;