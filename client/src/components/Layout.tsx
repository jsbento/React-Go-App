import NavBar from "./NavBar";
import Footer from "./Footer";
import { LayoutProps } from "../types/Types";
import React from "react";

const Layout:React.FC<LayoutProps> = ({children}:LayoutProps) => {
    return (
        <div className="flex flex-col min-h-screen">
            <NavBar></NavBar>
            <div className="flex flex-1 justify-center">
                {children}
            </div>
            <Footer></Footer>
        </div>
    );
}

export default Layout;