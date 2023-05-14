"use client";
import {useEffect, useState} from "react";

export default function Home() {
    const [message, setMessage] = useState<string>("");
    const getMessage = async () => {
        try {
            const response = await fetch(`/api`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            if (response.ok) {
                const data = await response.json();
                setMessage(data.message);
            } else {
                const errorResponse = await response.json();
                console.error('Error:', errorResponse);
            }
        } catch (error) {
            console.error(error);
        }
    };
    useEffect(() => {
        getMessage();
    }, []);

    return (
        <main>
            <div className={"text-9xl"}>
                test
                {message}
            </div>
        </main>
    )
}
