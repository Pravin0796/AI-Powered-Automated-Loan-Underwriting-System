// src/components/Header.tsx
import { useState } from "react";
import { Link } from "react-router-dom"; // Assuming you have routing set up
import '@fortawesome/fontawesome-free/css/all.min.css';

const Header = () => {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  const toggleMenu = () => setIsMobileMenuOpen(!isMobileMenuOpen);

  return (
    <header className="bg-blue-600 text-white p-4 shadow-md">
      <div className="container mx-auto flex justify-between items-center">
        <div className="text-xl font-semibold">
          <Link to="/">MyLoanApp</Link>
        </div>

        <div className="hidden md:flex space-x-4">
          <Link to="/" className="hover:text-gray-300">Home</Link>
          <Link to="/profile" className="hover:text-gray-300">Profile</Link>
          <Link to="/loans" className="hover:text-gray-300">Loans</Link>
          <Link to="/register" className="hover:text-gray-300">Register</Link>
          <Link to="/login" className="hover:text-gray-300">Login</Link>
        </div>

        <button onClick={toggleMenu} className="md:hidden text-white">
          <i className={`fas ${isMobileMenuOpen ? 'fa-times' : 'fa-bars'}`}></i>
        </button>
      </div>

      {/* Mobile Menu */}
      {isMobileMenuOpen && (
        <div className="md:hidden mt-4 space-y-2">
          <Link to="/" className="block px-4 py-2 hover:bg-blue-500 rounded">Home</Link>
          <Link to="/profile" className="block px-4 py-2 hover:bg-blue-500 rounded">Profile</Link>
          <Link to="/loans" className="block px-4 py-2 hover:bg-blue-500 rounded">Loans</Link>
          <Link to="/register" className="block px-4 py-2 hover:bg-blue-500 rounded">Register</Link>
          <Link to="/login" className="block px-4 py-2 hover:bg-blue-500 rounded">Login</Link>
        </div>
      )}
    </header>
  );
};

export default Header;
