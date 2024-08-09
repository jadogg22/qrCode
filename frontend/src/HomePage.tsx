import React, { useState } from 'react';

const HomePage: React.FC = () => {
    const [url, setUrl] = useState<string>('');
    const [qrCode, setQrCode] = useState<Blob | null>(null);

    const handleGenerate = async () => {
        try {
            // Ensure the URL is not empty
            if (!url) {
                throw new Error('URL is required');
            }

            // Create a URL-encoded string
            console.log("url: " + url);
            const encodedBody = new URLSearchParams();
            encodedBody.append('name', url);
            console.log("body: " + encodedBody.toString());

            const response = await fetch('http://localhost:8080/Generate', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: encodedBody.toString(), // Set the body to the URL-encoded string
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            // Assuming the backend returns the QR code image as a Blob
            const resultBlob = await response.blob();
            setQrCode(resultBlob);
        } catch (error) {
            console.error('Error:', error);
        }
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-sky-500">
            <div className="text-center">
                <h1 className="text-white text-4xl font-bold mb-8">QR Code Generator</h1>
                <div className="flex justify-center">
                    <input
                        type="text"
                        placeholder="Enter your URL here"
                        value={url}
                        onChange={(e) => setUrl(e.target.value)}
                        className="w-80 p-4 text-lg rounded-l-full outline-none shadow-md"
                    />
                    <button
                        onClick={handleGenerate}
                        className="bg-blue-800 text-white px-6 py-4 text-lg rounded-r-full shadow-md hover:bg-blue-700"
                    >
                        Generate
                    </button>
                </div>
                {qrCode && (
                    <div className="mt-8">
                        <img
                            src={URL.createObjectURL(qrCode)}
                            alt="Generated QR Code"
                            className="mx-auto"
                        />
                    </div>
                )}
            </div>
        </div>
    );
};

export default HomePage;
