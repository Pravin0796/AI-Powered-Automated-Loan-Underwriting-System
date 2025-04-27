import { useState } from "react";
import { Link } from "react-router-dom";
import { FaBars, FaTimes } from "react-icons/fa";

const Header = () => {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  const toggleMenu = () => setIsMobileMenuOpen(!isMobileMenuOpen);

  return (
    <header className="fixed top-0 left-0 w-full z-50 bg-blue-800 text-white p-4 shadow-md">
      <div className="container mx-auto flex justify-between items-center">
        {/* Logo */}
        <div className="text-xl font-bold">
          <Link to="/">LoanUnderwriter</Link>
        </div>

        {/* Desktop Links */}
        <nav className="hidden md:flex space-x-6">
          <Link to="/" className="hover:text-gray-300">
            Home
          </Link>
          <Link to="/profile" className="hover:text-gray-300">
            Profile
          </Link>
          <Link to="/loan-applications" className="hover:text-gray-300">
            Loans
          </Link>
          <Link to="/analytics" className="hover:text-gray-300">
            Analytics
          </Link>
          <Link to="/login" className="hover:text-gray-300">
            Login
          </Link>
        </nav>

        {/* Mobile Menu Toggle */}
        <button onClick={toggleMenu} className="md:hidden text-2xl">
          {isMobileMenuOpen ? <FaTimes /> : <FaBars />}
        </button>
      </div>

      {/* Mobile Menu */}
      {isMobileMenuOpen && (
        <div className="md:hidden mt-4 space-y-2 bg-blue-700 p-4">
          <Link to="/" className="block px-4 py-2 hover:bg-blue-600 rounded">
            Home
          </Link>
          <Link to="/profile" className="block px-4 py-2 hover:bg-blue-600 rounded">
            Profile
          </Link>
          <Link to="/loan-applications" className="block px-4 py-2 hover:bg-blue-600 rounded">
            Loans
          </Link>
          <Link to="/analytics" className="block px-4 py-2 hover:bg-blue-600 rounded">
            Analytics
          </Link>
          <Link to="/login" className="block px-4 py-2 hover:bg-blue-600 rounded">
            Login
          </Link>
        </div>
      )}
    </header>
  );
};

export default Header;