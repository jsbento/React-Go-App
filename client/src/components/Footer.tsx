import React from "react";
import { FooterProps } from "../types/Types";

const Footer:React.FC<FooterProps> = () => {
    return (
        <footer className="h-20 flex justify-center">
            <div className="flex justify-center items-center font-bold">
                React/Go Application
            </div>
        </footer>
    );
}

export default Footer;