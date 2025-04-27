import React, { useState } from "react";
import { FaHome, FaUser, FaChartLine, FaBars, FaTimes } from "react-icons/fa";

const Sidebar: React.FC = () => {
  const [isOpen, setIsOpen] = useState(false);

  const toggleSidebar = () => {
    setIsOpen(!isOpen);
  };

  return (
    <div className="flex">
      {/* Sidebar */}
      <div
        className={`fixed top-0 left-0 h-full bg-blue-800 text-white shadow-lg transform ${
          isOpen ? "translate-x-0" : "-translate-x-full"
        } transition-transform duration-300 w-64 z-50`}
      >
        <button
          onClick={toggleSidebar}
          className="absolute top-4 right-4 text-white hover:text-gray-300"
        >
          <FaTimes size={20} />
        </button>
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

      {/* Main Content */}
      <div
        className={`flex-1 transition-all duration-300 ${
          isOpen ? "ml-64" : "ml-0"
        }`}
      >
        {/* Toggle Button */}
        <button
          onClick={toggleSidebar}
          className="fixed top-4 left-4 text-blue-800 z-50"
        >
          <FaBars size={24} />
        </button>

        {/* Main Content Area */}
        <div className="p-6">
          <h1 className="text-2xl font-bold">Main Content</h1>
          <p>This is where your main content will go.</p>
        </div>
      </div>
    </div>
  );
};

export default Sidebar;