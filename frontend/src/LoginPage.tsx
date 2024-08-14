import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

// Define the types for state and response
interface LoginResponse {
    message: string;
}

const LoginPage: React.FC = () => {
    const [username, setUsername] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [message, setMessage] = useState<string>('');
    const [error, setError] = useState<string>('');
    const navigate = useNavigate();

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();

        try {
            const response = await axios.post<LoginResponse>(
                'http://localhost:8080/api/login',
                new URLSearchParams({
                    username,
                    password
                }).toString(),
                {
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded'
                    }
                }
            );

            // Handle the response from the server
            setMessage(response.data.message || 'Login successful');
            setError('');

            // Wait 2 seconds then redirect to /my-sites
            setTimeout(() => {
                navigate('/my-sites');
            }, 2000);
        } catch (err: any) {
            if (err.response && err.response.data) {
                setError(err.response.data.error || 'An error occurred');
            } else {
                setError('An error occurred');
            }
            setMessage('');
        }
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-sky-500">
            <form onSubmit={handleLogin} className="bg-white p-8 rounded-lg shadow-md">
                <h2 className="text-2xl font-bold mb-6 text-center">Login</h2>
                <input
                    type="text"
                    placeholder="Username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    className="w-full p-2 mb-4 border rounded"
                />
                <input
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    className="w-full p-2 mb-6 border rounded"
                />
                <button
                    type="submit"
                    className="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600"
                >
                    Login
                </button> 
                {error && <p className="text-red-500 text-center mt-4">{error}</p>}
                {message && <p className="text-green-500 text-center mt-4">{message}</p>}
            </form>
        </div>
    );
};

export default LoginPage;
