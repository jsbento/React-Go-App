import React from "react"

export interface FooterProps {}

export interface FormValues {
    username: string,
    email: string,
    password: string,
    conf_pass: string,
};

export interface LayoutProps {
    children: React.ReactNode
};