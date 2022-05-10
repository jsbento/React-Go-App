import React from "react";

const NavBar:React.FC = () => {
    return (
        <div>
            <ul className="flex justify-end p-8 gap-5 items-center bg-gray-800 text-white mb-5">
                <li className="mr-auto">Logo</li>
                <li className="hover:scale-105 font-bold">Page 1</li>
                <li className="hover:scale-105 font-bold">Page 2</li>
                <li className="hover:scale-105 font-bold">
                    <a href="/signup">Sign Up</a>
                </li>
            </ul>
        </div>
    );
}

export default NavBar;