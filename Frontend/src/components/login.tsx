import React, { useState } from "react";
import { loginUser } from "../services/userService.ts";
import { FaEye } from "react-icons/fa";

const Login = () => {
    const [form, setForm] = useState({ email: "", password: "" });
    const [showPassword, setShowPassword] = useState(false);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setForm({ ...form, [e.target.name]: e.target.value });
    };

    const handleLogin = async () => {
        try {
            const res = await loginUser({ email: form.email, password: form.password });
            console.log("Login successful:", res);
        } catch (err) {
            console.error("Login error:", err);
        }
    };

    return (
        <div className="flex flex-col justify-center bg-gray-200 border m-auto my-10 w-sm rounded-lg p-10 gap-6">
            <h2 className="text-center text-4xl text-blue-800">Login</h2>
            <div className="flex flex-col justify-center gap-2">
                <label htmlFor="email">Email <span className="text-red-700">*</span></label>
                <input className="outline-none border p-2 bg-white" type="email" name="email" placeholder="Email"
                    onChange={handleChange} />
            </div>
            <div className="flex flex-col gap-2">
                <label htmlFor="password">Password <span className="text-red-700">*</span></label>
                <div className="flex flex-row justify-between bg-white items-center p-2 border">
                    <input className="outline-none w-full" type={showPassword ? 'text' : 'password'} name="password" placeholder="Password"
                        onChange={handleChange} />
                    <FaEye
                        className="cursor-pointer"
                        type="button"
                        onClick={() => setShowPassword(!showPassword)}
                    >
                        {showPassword ? 'Hide' : 'Show'}
                    </FaEye>
                </div>
            </div>
            <button className="border w-25 bg-blue-400 m-auto p-1" onClick={handleLogin}>Login</button>
        </div>

    );
};

export default Login;