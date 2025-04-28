import React from 'react';
import { Link } from 'react-router-dom'; // For navigation
import LoadingSpinner from './LoadingSpinner';

const Home = () => {
  const [isLoading, setIsLoading] = React.useState(true);

  React.useEffect(() => {
    setTimeout(() => setIsLoading(false), 2000); // Simulate loading
  }, []);

  if (isLoading) return <LoadingSpinner />;

  return (
    <div className="p-6 sm:p-12 bg-gray-50 min-h-screen">
      <h1 className="text-3xl sm:text-4xl font-bold text-blue-700 mb-6">Welcome to the Loan Underwriting System</h1>
      <p className="text-lg text-gray-700 mb-6">
        This platform allows users to apply for loans, track their applications, and manage their profiles.
      </p>

      {/* Main Action Cards */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <div className="bg-white rounded-2xl shadow-md hover:shadow-xl transition-shadow p-6 text-center">
          <h2 className="text-xl font-semibold text-gray-800 mb-4">Apply for Loan</h2>
          <p className="text-gray-600 mb-4">Start a new loan application quickly and easily.</p>
          <Link to="/loan-applications" className="inline-block bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition-colors">
            Start Application
          </Link>
        </div>

        <div className="bg-white rounded-2xl shadow-md hover:shadow-xl transition-shadow p-6 text-center">
          <h2 className="text-xl font-semibold text-gray-800 mb-4">Track Applications</h2>
          <p className="text-gray-600 mb-4">Monitor the progress and status of your loan applications.</p>
          <Link to="/viewloan" className="inline-block bg-green-600 text-white px-6 py-2 rounded-lg hover:bg-green-700 transition-colors">
            View Applications
          </Link>
        </div>

        <div className="bg-white rounded-2xl shadow-md hover:shadow-xl transition-shadow p-6 text-center">
          <h2 className="text-xl font-semibold text-gray-800 mb-4">View Credit Score</h2>
          <p className="text-gray-600 mb-4">Check your current credit score and improve your chances of approval.</p>
          <Link to="/viewcredit" className="inline-block bg-yellow-600 text-white px-6 py-2 rounded-lg hover:bg-yellow-700 transition-colors">
            Check Score
          </Link>
        </div>
      </div>

      {/* Additional Information Section */}
      <div className="mt-12">
        <h2 className="text-2xl sm:text-3xl font-semibold text-gray-800 mb-4">How it Works</h2>
        <div className="flex flex-col sm:flex-row gap-6">
          <div className="bg-white rounded-2xl shadow-md p-6 flex-1">
            <h3 className="text-lg font-semibold text-gray-800 mb-2">Step 1: Apply</h3>
            <p className="text-gray-600">Fill out a simple application form to begin your loan process.</p>
          </div>
          <div className="bg-white rounded-2xl shadow-md p-6 flex-1">
            <h3 className="text-lg font-semibold text-gray-800 mb-2">Step 2: Track</h3>
            <p className="text-gray-600">Stay updated with the status of your application through the dashboard.</p>
          </div>
          <div className="bg-white rounded-2xl shadow-md p-6 flex-1">
            <h3 className="text-lg font-semibold text-gray-800 mb-2">Step 3: Get Approved</h3>
            <p className="text-gray-600">Receive your loan approval decision and manage repayments online.</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Home;
