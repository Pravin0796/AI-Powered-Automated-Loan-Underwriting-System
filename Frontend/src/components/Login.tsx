import React, { useState } from "react";
import { loginUser } from "../services/UserService";
import { FaEye, FaEyeSlash } from "react-icons/fa";
import { useNavigate } from "react-router-dom";
import { setToken } from "../utils/Auth";

const Login = () => {
  const [form, setForm] = useState({ email: "", password: "" });
  const [showPassword, setShowPassword] = useState(false);
  const navigate = useNavigate();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleLogin = async () => {
    try {
      const res = await loginUser(form);
      if (!res.token) throw new Error("Invalid login credentials");
      setToken(res.token);
      navigate("/dashboard");
    } catch (err) {
      if (err instanceof Error) {
        alert(err.message || "Login failed. Please try again.");
      } else {
        alert("Login failed. Please try again.");
      }
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100 px-4">
      <div className="w-full max-w-md bg-white shadow-lg rounded-lg p-8">
        <h2 className="text-3xl font-bold text-center text-blue-800">Login</h2>
        <div className="mt-6">
          {/* Email Input */}
          <label className="block text-sm font-medium text-gray-700">
            Email
          </label>
          <input
            type="email"
            name="email"
            value={form.email}
            onChange={handleChange}
            className="w-full border border-gray-300 rounded-md px-4 py-2 mt-1"
            placeholder="Enter your email"
          />
        </div>
        <div className="mt-4">
          {/* Password Input */}
          <label className="block text-sm font-medium text-gray-700">
            Password
          </label>
          <div className="flex flex-row items-center relative">
            <input
              type={showPassword ? "text" : "password"}
              name="password"
              value={form.password}
              onChange={handleChange}
              className="w-full border border-gray-300 rounded-md px-4 py-2 mt-1"
              placeholder="Enter your password"
            />
            <div
              className="absolute right-3 top-3 cursor-pointer"
              onClick={() => setShowPassword(!showPassword)}
            >
              {showPassword ? <FaEyeSlash /> : <FaEye />}
            </div>
          </div>
        </div>
        <button
          onClick={handleLogin}
          className="w-full bg-blue-800 text-white font-semibold rounded-md py-2 mt-6 hover:bg-blue-700"
        >
          Login
        </button>
        <p className="text-center mt-4">
          Don't have an account?{" "}
          <span
            onClick={() => navigate("/register")}
            className="text-blue-600 cursor-pointer"
          >
            Register
          </span>
        </p>
      </div>
    </div>
  );
};

export default Login;