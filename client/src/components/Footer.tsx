import React from "react";
import { FooterProps } from "../types/Types";

const Footer:React.FC<FooterProps> = () => {
    return (
        <footer className="h-20 flex justify-center w-full">
            <div className="flex items-center font-bold">
                React/Go Application
            </div>
        </footer>
    );
}

export default Footer;