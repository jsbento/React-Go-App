import React from "react";
import NavBar from "./NavBar";
import Footer from "./Footer";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import SignIn from "../pages/SignUp";
import { LayoutProps } from "../types/Types";

const Layout:React.FC<LayoutProps> = () => {
    return (
        <div className="flex flex-col min-h-screen">
            <BrowserRouter>
                <NavBar/>
                <Routes>
                    <Route path="signup" element={<SignIn/>}/>
                </Routes>
            </BrowserRouter>
            <div className="absolute bottom-0 flex justify-center w-full">
                <Footer/>
            </div>
        </div>
    );
}

export default Layout;