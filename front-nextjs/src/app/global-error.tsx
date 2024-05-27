'use client'

import Navbar from "@/components/Navbar/Navbar";
import { Inter } from "next/font/google";
import "./globals.scss";
import styles from './not-found.module.scss'

const inter = Inter({ subsets: ["latin"] });

export default function GlobalError({
    error,
    reset,
}: {
    error: Error & { digest?: string }
    reset: () => void
}) {
    return (
        <html lang="en">
            <body className={inter.className}>
                <Navbar />
                <div className="main">
                    <div style={{ margin: 'auto auto' }}>
                        <h1 className={styles.message}>500 - Internal Server Error</h1>
                        <p className={styles.message}>{error.message}</p>
                        <p className={styles.message}>{error.stack}</p>
                        <button onClick={reset}>Reload</button>
                    </div>
                </div>
            </body>
        </html>
    );
}
