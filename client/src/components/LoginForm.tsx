import React from "react";
import { FormValues } from "../types/Types";
import * as yup from "yup";
import { Field, Form, Formik } from "formik";
import "../../src/styles/Card.css"

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
                // Generate token on backend, recieve token and set as a cookie
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