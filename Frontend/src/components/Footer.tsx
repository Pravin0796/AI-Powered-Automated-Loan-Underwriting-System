const Footer = () => {
  return (
    <footer className="bg-blue-800 text-white py-6">
      <div className="max-w-6xl mx-auto px-4 text-center">
        <p className="text-sm md:text-base">
          &copy; {new Date().getFullYear()} MyLoanApp. All rights reserved.
        </p>
        <div className="space-x-4 mt-3">
          <a
            href="#"
            className="text-sm md:text-base hover:text-gray-300 transition duration-200"
          >
            Privacy Policy
          </a>
          <a
            href="#"
            className="text-sm md:text-base hover:text-gray-300 transition duration-200"
          >
            Terms of Service
          </a>
        </div>
      </div>
    </footer>
  );
};

export default Footer;