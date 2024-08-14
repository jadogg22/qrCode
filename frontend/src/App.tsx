import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import HomePage from './HomePage';
import LoginPage from './LoginPage';
import MySitesPage from './MySitesPage';
import Register from './Register';
import Navigation from './Navigation';

const App: React.FC = () => {
    return (
        <Router>
            <Navigation />
            <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="/login" element={<LoginPage />} />
                <Route path="/my-sites" element={<MySitesPage />} />
                <Route path="/register" element={<Register />} />
                <Route path="*" element={<div>Not Found</div>} />
            </Routes>
        </Router>
    );
};

export default App;
