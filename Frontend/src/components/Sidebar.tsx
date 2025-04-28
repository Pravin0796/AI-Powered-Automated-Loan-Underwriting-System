import React, { useState } from "react";
import { FaHome, FaUser, FaChartLine } from "react-icons/fa";

interface SidebarProps {
  isOpen: boolean;
}

const Sidebar: React.FC<SidebarProps> = ({ isOpen }) => {
  return (
    <div
      className={`fixed my-10 left-0 h-full bg-blue-800 text-white shadow-lg transform transition-transform duration-300 ${
        isOpen ? "translate-x-0" : "-translate-x-full"
      } w-64 z-50`}
    >
      <div className="p-6">
        <h2 className="text-2xl font-bold mb-6">My App</h2>
        <ul className="space-y-4">
          <li className="flex items-center space-x-3 cursor-pointer hover:bg-blue-700 p-2 rounded">
            <FaHome />
            <span>Home</span>
          </li>
          <li className="flex items-center space-x-3 cursor-pointer hover:bg-blue-700 p-2 rounded">
            <FaUser />
            <span>Profile</span>
          </li>
          <li className="flex items-center space-x-3 cursor-pointer hover:bg-blue-700 p-2 rounded">
            <FaChartLine />
            <span>Analytics</span>
          </li>
        </ul>
      </div>
    </div>
  );
};

export default Sidebar;
