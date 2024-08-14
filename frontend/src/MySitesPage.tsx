import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';

interface QRCode {
    id: string;
    url: string;
    imageUrl: string;
}

const MySitesPage: React.FC = () => {
    const [qrCodes, setQrCodes] = useState<QRCode[]>([]);

    useEffect(() => {
        fetchQRCodes();
    }, []);

    const fetchQRCodes = async () => {
        try {
            // In a real application, you'd fetch this data from your backend
            // For this example, we'll use mock data
            const mockData: QRCode[] = [
                { id: '1', url: 'https://example.com', imageUrl: 'path/to/qr1.png' },
                { id: '2', url: 'https://example.org', imageUrl: 'path/to/qr2.png' },
            ];
            setQrCodes(mockData);
        } catch (error) {
            console.error('Error fetching QR codes:', error);
        }
    };

    return (
        <div className="container mx-auto p-4">
            <h1 className="text-3xl font-bold mb-6">My Sites</h1>
            <Link to="/" className="bg-blue-500 text-white px-4 py-2 rounded mb-4 inline-block">
                Generate New QR Code
            </Link>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                {qrCodes.map((qr) => (
                    <div key={qr.id} className="border p-4 rounded">
                        <img src={qr.imageUrl} alt={`QR for ${qr.url}`} className="mx-auto mb-2" />
                        <p className="text-center">{qr.url}</p>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default MySitesPage;
