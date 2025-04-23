import { useState } from "react";
import { NavLink, useNavigate } from "react-router-dom";
import classNames from "classnames";
import { removeToken } from "../utils/Auth";
import { FaBars, FaTimes } from "react-icons/fa";

const navItemClass = (isActive: boolean) =>
  classNames("block px-4 py-2 rounded hover:bg-gray-200 transition-all", {
    "bg-blue-100 text-blue-600 font-semibold": isActive,
  });

export default function Sidebar() {
  const [isCollapsed, setIsCollapsed] = useState(false);
  const navigate = useNavigate();

  const handleLogout = () => {
    removeToken();
    navigate("/login");
  };

  return (
    <div className="md:flex">
      {/* Sidebar Toggle Button */}
      <button
        className="p-4 md:hidden text-2xl"
        onClick={() => setIsCollapsed(!isCollapsed)}
      >
        {isCollapsed ? <FaTimes /> : <FaBars />}
      </button>

      {/* Sidebar */}
      <div
        className={classNames(
          "bg-white shadow-md md:shadow-none md:block space-y-4 p-4 md:min-w-[200px] md:h-screen",
          {
            "block": isCollapsed,
            "hidden": !isCollapsed,
          },
          "md:block"
        )}
      >
        <NavLink to="/dashboard" className={({ isActive }) => navItemClass(isActive)}>
          Dashboard
        </NavLink>
        <NavLink to="/profile" className={({ isActive }) => navItemClass(isActive)}>
          Profile
        </NavLink>
        <button
          onClick={handleLogout}
          className="text-red-500 hover:underline px-4 py-2 w-full text-left"
        >
          Logout
        </button>
      </div>
    </div>
  );
}
