import React, { useState } from "react";
import { AppProps, Validation, FormValues, SignUpProps } from "../types/Types";
import * as yup from "yup";
import { Field, FieldInputProps, Form, Formik } from "formik";

const SignUpScheme = yup.object().shape({
    username: yup.string().trim().required("Username required"),
    email: yup.string().email("Must be a valid email").max(255).required("Email is required"),
    password: yup.string().min(8, "Password must be minimum 8 characters").matches(/[a-zA-z0-9!@#$%^&*]/).required("Password required"),
    confirm_password: yup.string().min(8, "Password must be minimum 8 character").matches(/[a-zA-Z0-9!@#$%^&*]/).required("Password required")
});

const initial_values:FormValues = {
    username: "",
    email: "",
    password: "",
    confirm_password: ""
}

const SignUpForm: React.FC<SignUpProps> = () => {
    const [username, setUsername] = useState<string | null>(null);
    const [email, setEmail] = useState<string | null>(null);
    const [password, setPassword] = useState<string | null>(null);
    const [conf_pass, setConfPass] = useState<string | null>(null);

    return (
        <div>
            <Formik validationSchema={SignUpScheme} initialValues={initial_values} onSubmit={async (values, actions) => {
                //Check username doesnt already exist
                //Confirm passwords match
                //Post user
            }}>
                {({errors, isSubmitting}) => (
                    <Form>
                        <label htmlFor="username">Username</label>
                        <Field id="username" name="username" type="text" placeholder="Username" autoComplete="off"/>
                        <label htmlFor="email">Email</label>
                        <Field id="email" name="email" type="text" placeholder="Email" autoComplete="off"/>
                        <label htmlFor="password">Password</label>
                        <Field id="password" name="password" type="text" placeholder="Password" autoComplete="off"/>
                        <label htmlFor="conf_pass">Confirm Password</label>
                        <Field id="conf_pass" name="conf_pass" type="text" placeholder="Confirm Password" autoComplete="off"/>
                        <button className="border-2 w-auto p-1 rounded-md m-1 font-semibold" type="submit">Sign Up</button>
                        {isSubmitting ? (<div className="animate-pulse font-semibold text-lg">Loading...</div>) : null}
                    </Form>
                )}
            </Formik>
        </div>
    );
};

export default SignUpForm;