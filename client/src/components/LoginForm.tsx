import React from "react";
import { FormValues } from "../types/Types";
import * as yup from "yup";
import { Field, Form, Formik } from "formik";
import CryptoJS from "crypto-js";
import "../../src/styles/Card.css";

const SERVER_URI = "http://localhost:8080";

const LoginScheme = yup.object().shape({
    username: yup.string().trim().required("Username required"),
    password: yup.string().trim().matches(/[a-zA-z0-9!@#$%^&*]/).required("Password required")
});

const initial_values:FormValues = {
    username: "",
    password: ""
};

const LoginForm: React.FC = () => {
    return (
        <div id="login-card" className="card flex flex-col items-center">
            <Formik validationSchema={LoginScheme} validateOnBlur={false} validateOnChange={false} initialValues={initial_values} onSubmit={async (values, actions) => {
                document.cookie = "token=; Max-Age=0; path=/; domain=" + window.location.hostname;
                const token = await fetch(`${SERVER_URI}/users/login`, {
                    method: 'POST',
                    body: JSON.stringify({
                        "username": values.username,
                        "password": CryptoJS.AES.encrypt(values.password, process.env.REACT_APP_CRYPTO_KEY!).toString()
                    })
                })
                .then(response => { return response.json(); })
                .then(data => {
                    console.log(data.message);
                    return data.token;
                })
                .catch(err => { console.log(err); });
                document.cookie = `token=${token}; SameSite=None; Secure`;
                window.location.href = '/dashboard';
                actions.setSubmitting(false);
            }}>
                {({errors, isSubmitting}) => (
                    <>
                        <h2 className="font-bold p-2">Login</h2>
                        <Form className="flex flex-col justify-center items-center mx-auto my-auto w-fit">
                            <label id="form-elt" htmlFor="username">Username</label>
                            <Field id="username" name="username" type="text" placeholder="Username" autoComplete="off"/>
                            <label id="form-elt" htmlFor="password">Password</label>
                            <Field id="password" name="password" type="password" placeholder="Password" autoComplete="off"/>
                            <button className="border-2 w-auto p-1 rounded-md m-2 font-semibold" type="submit">Login</button>
                            <p>Don't have an account? <a href="/signup" className="text-sky-500">Sign up</a></p>
                            {errors.username ? <div className="text-red-600">{errors.username}</div> : null}
                            {errors.password ? <div className="text-red-600">{errors.password}</div> : null}
                            {isSubmitting ? (<div className="animate-pulse font-semibold text-lg">Loading...</div>) : null}
                        </Form>
                    </>
                )}
            </Formik>
        </div>
    );
};

export default LoginForm;