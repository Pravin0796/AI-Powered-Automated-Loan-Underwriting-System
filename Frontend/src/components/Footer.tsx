// src/components/Footer.tsx

const Footer = () => {
  return (
    <footer className="bg-blue-600 text-white py-4">
      <div className="container mx-auto text-center">
        <p>&copy; {new Date().getFullYear()} MyLoanApp. All rights reserved.</p>
        <div className="space-x-4 mt-2">
          <a href="#" className="hover:text-gray-300">Privacy Policy</a>
          <a href="#" className="hover:text-gray-300">Terms of Service</a>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
