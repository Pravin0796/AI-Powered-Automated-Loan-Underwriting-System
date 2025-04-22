// src/pages/Home.tsx
import React from "react";

const Home = () => {
  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-4">Welcome to the Loan Underwriting System</h1>
      <p className="text-gray-700 mb-6">
        This platform allows users to apply for loans, track their applications, and manage their profiles.
      </p>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div className="bg-white rounded-2xl shadow-md p-4">
          <h2 className="text-xl font-semibold mb-2">Apply for Loan</h2>
          <p>Start a new loan application quickly and easily.</p>
        </div>
        <div className="bg-white rounded-2xl shadow-md p-4">
          <h2 className="text-xl font-semibold mb-2">Track Applications</h2>
          <p>Monitor the progress and status of your loan applications.</p>
        </div>
        <div className="bg-white rounded-2xl shadow-md p-4">
          <h2 className="text-xl font-semibold mb-2">View Credit Score</h2>
          <p>Check your current credit score and improve your chances of approval.</p>
        </div>
      </div>
    </div>
  );
};

export default Home;
