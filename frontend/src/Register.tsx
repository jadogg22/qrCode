import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';


// Define the types for state and response
interface RegisterResponse {
  message: string;
}

const Register: React.FC = () => {
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [email, setEmail] = useState<string>('');
  const [message, setMessage] = useState<string>('');
  const [error, setError] = useState<string>('');
    const navigate = useNavigate();

  const handleRegister = async () => {
    try {
      const response = await axios.post<RegisterResponse>(
        'http://localhost:8080/api/register',
        new URLSearchParams({
          username,
          password,
          email
        }).toString(),
        {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
          }
        }
      );
      setMessage(response.data.message);
      setError('');
      // wait 2 secounds then redirect to my-sites
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
    <div>
      <h2>Register</h2>
      <input
        type="text"
        placeholder="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <input
        type="email"
        placeholder="Email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />
      <button onClick={handleRegister}>Register</button>
      {message && <p style={{ color: 'green' }}>{message}</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </div>
  );
};

export default Register;
