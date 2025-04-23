import React, { useState } from "react";
import { loginUser } from "../services/UserService.ts";
import { FaEye, FaEyeSlash } from "react-icons/fa";
import { useNavigate } from "react-router-dom";
import { setToken } from "../utils/Auth.ts";

const Login = () => {
  const [form, setForm] = useState({ email: "", password: "" });
  const [showPassword, setShowPassword] = useState(false);
  const navigate = useNavigate();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleLogin = async () => {
    try {
      console.log("Sending login request:", form);
      const res = await loginUser({ email: form.email, password: form.password });
      console.log("gRPC Login response:", res);
  
      if (!res.token) {
        throw new Error("No token returned from backend");
      }
  
      setToken(res.token);
      navigate("/dashboard");
    } catch (err: any) {
      console.error("Login failed:", err);
      alert(err.message || "Login failed. Please check your credentials.");
    }
  };  

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100 px-4">
      <div className="w-full max-w-md bg-white shadow-lg rounded-xl p-8 space-y-6">
        <h2 className="text-3xl font-bold text-center text-blue-700">Welcome Back 👋</h2>

        <div className="space-y-4">
          <div>
            <label htmlFor="email" className="block mb-1 text-sm font-medium text-gray-700">
              Email <span className="text-red-600">*</span>
            </label>
            <input
              type="email"
              name="email"
              placeholder="you@example.com"
              onChange={handleChange}
              className="w-full px-4 py-2 border border-gray-500 rounded-md focus:ring-2 focus:ring-blue-500 focus:outline-none"
              required
            />
          </div>

          <div>
            <label htmlFor="password" className="block mb-1 text-sm font-medium text-gray-700">
              Password <span className="text-red-600">*</span>
            </label>
            <div className="relative">
              <input
                type={showPassword ? "text" : "password"}
                name="password"
                placeholder="Enter your password"
                onChange={handleChange}
                className="w-full px-4 py-2 pr-10 border border-gray-500 rounded-md focus:ring-2 focus:ring-blue-500 focus:outline-none"
                required
              />
              <div
                className="absolute inset-y-0 right-3 flex items-center cursor-pointer"
                onClick={() => setShowPassword(!showPassword)}
              >
                {showPassword ? <FaEyeSlash className="text-gray-600" /> : <FaEye className="text-gray-600" />}
              </div>
            </div>
          </div>

          <div className="text-right">
            <button className="text-sm text-blue-600 hover:underline" onClick={() => navigate("/forgot-password")}>
              Forgot Password?
            </button>
          </div>
        </div>

        <button
          onClick={handleLogin}
          className="w-full py-2 mt-2 bg-blue-600 text-white font-semibold rounded-md hover:bg-blue-700 transition duration-200"
        >
          Login
        </button>

        <p className="text-center text-sm text-gray-600">
          Don't have an account?{" "}
          <button className="text-blue-600 hover:underline" onClick={() => navigate("/register")}>
            Register
          </button>
        </p>
      </div>
    </div>
  );
};

export default Login;
