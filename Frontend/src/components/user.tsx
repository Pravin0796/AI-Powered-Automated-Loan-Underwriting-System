import React, { useState } from "react";
import { registerUser, loginUser } from "../services/user";

const Auth = () => {
    const [form, setForm] = useState({ email: "", password: "", full_name: "", date_of_birth: "", phone: "", address: "" });

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

    const handleLogin = async () => {
        try {
            const res = await loginUser({ email: form.email, password: form.password });
            console.log("Login successful:", res);
        } catch (err) {
            console.error("Login error:", err);
        }
    };

    return (
        <div>
            <h2>Register</h2>
            <input name="full_name" placeholder="Full Name" onChange={handleChange} />
            <input name="email" placeholder="Email" onChange={handleChange} />
            <input name="password" placeholder="Password" type="password" onChange={handleChange} />
            <input name="phone" placeholder="Phone" onChange={handleChange} />
            <input name="date_of_birth" placeholder="Date of Birth" onChange={handleChange} />
            <input name="address" placeholder="Address" onChange={handleChange} />
            <button onClick={handleRegister}>Register</button>

            <h2>Login</h2>
            <input name="email" placeholder="Email" onChange={handleChange} />
            <input name="password" placeholder="Password" type="password" onChange={handleChange} />
            <button onClick={handleLogin}>Login</button>
        </div>
    );
};

export default Auth;