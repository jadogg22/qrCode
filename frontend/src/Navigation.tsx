import React from 'react';
import { Link, useNavigate } from 'react-router-dom';

const Navigation: React.FC = () => {
    const navigate = useNavigate();
    const token = localStorage.getItem('token'); // Check for token

    const handleLogout = () => {
        localStorage.removeItem('token'); // Remove token on logout
        navigate('/login'); // Redirect to login page
    };

    return (
        <nav className="bg-blue-800 text-white p-4">
            <ul className="flex justify-between items-center">
                <li>
                    <Link to="/" className="hover:underline">Home</Link>
                </li>
                {!token ? (
                    <>
                        <li>
                            <Link to="/register" className="hover:underline">Register</Link>
                        </li>
                        <li>
                            <Link to="/login" className="hover:underline">Login</Link>
                        </li>
                    </>
                ) : (
                    <li>
                        <button onClick={handleLogout} className="hover:underline">
                            Logout
                        </button>
                    </li>
                )}
            </ul>
        </nav>
    );
};

export default Navigation;
