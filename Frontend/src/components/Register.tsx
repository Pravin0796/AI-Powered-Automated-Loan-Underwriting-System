import React, { useState } from "react";
import { registerUser } from "../services/UserService.ts";
import {FaEye} from "react-icons/fa";

const Register = () => {
    const [form, setForm] = useState({ email: "", password: "", full_name: "", date_of_birth: "", phone: "", address: "" });
    const [showPassword, setShowPassword] = useState(false);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setForm({ ...form, [e.target.name]: e.target.value });
    };

    const handleRegister = async () => {
        try {
            const res = await registerUser(form);
            console.log("Registration successful:", res);
        } catch (err) {
            console.error("Registration error:", err);
        }
    };

    return (
        <div className="flex flex-col justify-center  w-lg shadow-[0_0_20px_3px_rgba(0,0,0,0.12)] bg-white rounded-lg p-10 m-auto my-10 gap-6">
            <h2 className="text-center text-4xl text-blue-800">Welcome! Register here:</h2>
            <div className="flex flex-col  gap-2 p-2">
                <label htmlFor="first_name">Full Name <span className="text-red-700">*</span></label>
                <input className="outline-none bg-white border p-2" name="full_name" placeholder="Full Name" onChange={handleChange} />
            </div>
            <div className="flex flex-col gap-2 p-2">
                <label htmlFor="email">Email <span className="text-red-700">*</span></label>
                <input className="outline-none bg-white border p-2" type="email" name="email" placeholder="Email" onChange={handleChange} />
            </div>
            <div className="flex flex-col gap-2 p-2">
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
            <div className="flex flex-col gap-2 p-2">
                <label htmlFor="phone">Phone <span className="text-red-700">*</span></label>
                <input className="outline-none border bg-white p-2" name="phone" placeholder="Phone" onChange={handleChange} />
            </div>
            <div className="flex flex-col gap-2 p-2">
                <label htmlFor="date">Date of Birth <span className="text-red-700">*</span></label>
                <input className="outline-none border bg-white p-2" type="date" name="date_of_birth" placeholder="Date of Birth" onChange={handleChange} />
            </div>
            <div className="flex flex-col gap-2 p-2">
                <label htmlFor="address">Address <span className="text-red-700">*</span></label>
                <input className="outline-none border bg-white p-2" name="address" placeholder="Address" onChange={handleChange} />
            </div>
            <button className="border w-25 bg-blue-400 m-auto p-1" onClick={handleRegister}>Register</button>
        </div>
    );
};

export default Register;