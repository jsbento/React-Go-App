import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

// Components
import NavBar from "./NavBar";
import Footer from "./Footer";

// Pages
import SignUp from "../pages/SignUp";
import Login from "../pages/Login";

// Types
import { LayoutProps } from "../types/Types";

const Layout:React.FC<LayoutProps> = () => {
    return (
        <div className="flex flex-col min-h-screen">
            <BrowserRouter>
                <NavBar/>
                <Routes>
                    <Route path="" element={<Login/>}/>
                    <Route path="signup" element={<SignUp/>}/>
                </Routes>
            </BrowserRouter>
            <div className="absolute bottom-0 flex justify-center w-full">
                <Footer/>
            </div>
        </div>
    );
};

export default Layout;