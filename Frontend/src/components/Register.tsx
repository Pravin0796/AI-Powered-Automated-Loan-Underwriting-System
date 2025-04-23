import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { registerUser } from "../services/UserService.ts";
import { FaEye, FaEyeSlash } from "react-icons/fa";

const Register = () => {
    const [form, setForm] = useState({
        email: "",
        password: "",
        confirmPassword: "",
        full_name: "",
        date_of_birth: "",
        phone: "",
        address: "",
    });
    const [showPassword, setShowPassword] = useState(false);
    const [showConfirmPassword, setShowConfirmPassword] = useState(false);
    const [loading, setLoading] = useState(false);
    const navigate = useNavigate();

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setForm({ ...form, [e.target.name]: e.target.value });
    };

    const handleRegister = async () => {
        if (form.password !== form.confirmPassword) {
            return alert("Passwords do not match!");
        }

        setLoading(true);
        try {
            const { confirmPassword, ...formData } = form;
            const res = await registerUser(formData);
            if (res.status !== 200) {
                throw new Error("Registration failed");
            }
            alert("Registration successful!");
            navigate("/login");
        } catch (err) {
            console.error("Registration error:", err);
            alert("Registration failed.");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="max-w-md w-full sm:max-w-xl lg:max-w-2xl mx-auto bg-white p-6 sm:p-10 shadow-lg rounded-lg my-8">
            <h2 className="text-3xl font-bold text-center text-blue-700 mb-6">Create your account</h2>
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                <div className="flex flex-col">
                    <label>Full Name <span className="text-red-600">*</span></label>
                    <input name="full_name" onChange={handleChange} placeholder="Full Name"
                        className="p-2 border rounded-md border-gray-500 focus:ring-2 focus:ring-blue-500 focus:outline-none" />
                </div>
                <div className="flex flex-col">
                    <label>Email <span className="text-red-600">*</span></label>
                    <input type="email" name="email" onChange={handleChange} placeholder="Email"
                        className="p-2 border rounded-md border-gray-500 focus:ring-2 focus:ring-blue-500 focus:outline-none" />
                </div>
                <div className="flex flex-col relative col-span-1 sm:col-span-2">
                    <label>Password <span className="text-red-600">*</span></label>
                    <div className="relative">
              <input
                type={showPassword ? "text" : "password"}
                name="password"
                placeholder="Enter your password"
                onChange={handleChange}
                className="w-full px-4 py-2 pr-10 border border-gray-500 rounded-md focus:ring-2 focus:ring-blue-500 focus:outline-none"
                required
              />
                        <div className="absolute inset-y-0 right-3 flex items-center cursor-pointer"
                            onClick={() => setShowPassword(!showPassword)}>
                            {showPassword ? <FaEyeSlash className="text-gray-600" /> : <FaEye className="text-gray-600" />}
                        </div>
                    </div>
                </div>
                <div className="flex flex-col col-span-1 sm:col-span-2">
                    <label>Confirm Password <span className="text-red-600">*</span></label>
                    <div className="relative">
              <input
                type={showConfirmPassword ? "text" : "password"}
                name="confirmPassword"
                placeholder="Enter your password"
                onChange={handleChange}
                className="w-full px-4 py-2 pr-10 border border-gray-500 rounded-md focus:ring-2 focus:ring-blue-500 focus:outline-none"
                required
              />
                        <div className="absolute inset-y-0 right-3 flex items-center cursor-pointer"
                            onClick={() => setShowConfirmPassword(!showConfirmPassword)}>
                            {showConfirmPassword ? <FaEyeSlash className="text-gray-600" /> : <FaEye className="text-gray-600" />}
                        </div>
                    </div>
                </div>
                    <div className="flex flex-col">
                        <label>Phone <span className="text-red-600">*</span></label>
                        <input name="phone" onChange={handleChange} placeholder="Phone"
                            className="p-2 border rounded-md focus:outline-none border-gray-500 focus:ring-2 focus:ring-blue-500" />
                    </div>
                    <div className="flex flex-col">
                        <label>Date of Birth <span className="text-red-600">*</span></label>
                        <input type="date" name="date_of_birth" onChange={handleChange}
                            className="p-2 border rounded-md focus:outline-none border-gray-500 focus:ring-2 focus:ring-blue-500" />
                    </div>
                    <div className="flex flex-col col-span-1 sm:col-span-2">
                        <label>Address <span className="text-red-600">*</span></label>
                        <input name="address" onChange={handleChange} placeholder="Address"
                            className="p-2 border rounded-md focus:outline-none border-gray-500 focus:ring-2 focus:ring-blue-500" />
                    </div>
                </div>

                <button
                    onClick={handleRegister}
                    disabled={loading}
                    className="mt-6 w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded transition-all"
                >
                    {loading ? "Registering..." : "Register"}
                </button>

                <p className="mt-4 text-center text-sm">
                    Already have an account?{" "}
                    <span
                        onClick={() => navigate("/login")}
                        className="text-blue-600 cursor-pointer hover:underline"
                    >
                        Login here
                    </span>
                </p>
            </div>
            );
};

            export default Register;
