import React, {useState} from "react";
import { FormValues } from "../types/Types";
import * as yup from "yup";
import { Field, Form, Formik } from "formik";

const SignUpScheme = yup.object().shape({
    username: yup.string().trim().required("Username required"),
    email: yup.string().trim().email("Must be a valid email").max(255).required("Email is required"),
    password: yup.string().trim().min(6, "Password must be minimum 6 characters").matches(/[a-zA-z0-9!@#$%^&*]/).required("Password required"),
    conf_pass: yup.string().trim().min(6, "Password must be minimum 6 character").matches(/[a-zA-Z0-9!@#$%^&*]/).required("Password required")
});

const initial_values:FormValues = {
    username: "",
    email: "",
    password: "",
    conf_pass: ""
};

const SERVER_URI = "http://localhost:8080"

const SignUpForm: React.FC = () => {
    const [exists, setExists] = useState<boolean>(false);
    return (
        <div className="flex flex-col items-center">
            <Formik validationSchema={SignUpScheme} validateOnBlur={false} validateOnChange={false} initialValues={initial_values} onSubmit={async (values, actions) => {
                console.log(values.username);
                
                await fetch(SERVER_URI+`/users/exists?username=${values.username}`, {
                    method: 'GET'
                })
                .then(response => {
                    if(response.status === 200)
                        return response.json();
                })
                .then(data => {
                    setExists(data.exists);
                    console.log(exists);
                    console.log(data.message);
                })
                .catch(err => { console.log(err); });

                if(exists) {
                    alert(`Username ${values.username} already exists.`);
                }
                else if(values.password !== values.conf_pass) {
                    alert('Passwords do not match');
                }
                else {
                    await fetch(SERVER_URI+'/users', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify({
                            username: values.username,
                            email: values.email,
                            password: values.password
                        })
                    })
                    .catch(error => { console.log(error); });
                    actions.setSubmitting(false);
                }
            }}>
                {({errors, isSubmitting}) => (
                    <Form className="flex flex-col justify-center items-center">
                        <label htmlFor="username">Username</label>
                        <Field id="username" name="username" type="text" placeholder="Username" autoComplete="off"/>
                        <label htmlFor="email">Email</label>
                        <Field id="email" name="email" type="text" placeholder="Email" autoComplete="off"/>
                        <label htmlFor="password">Password</label>
                        <Field id="password" name="password" type="password" placeholder="Password" autoComplete="off"/>
                        <label htmlFor="conf_pass">Confirm Password</label>
                        <Field id="conf_pass" name="conf_pass" type="password" placeholder="Confirm Password" autoComplete="off"/>
                        <button className="border-2 w-auto p-1 rounded-md m-1 font-semibold" type="submit">Sign Up</button>
                        {errors.username ? <div>{errors.username}</div> : null}
                        {errors.email ? <div>{errors.email}</div> : null}
                        {errors.password ? <div>{errors.password}</div> : null}
                        {errors.conf_pass ? <div>{errors.conf_pass}</div> : null}
                        {isSubmitting ? (<div className="animate-pulse font-semibold text-lg">Loading...</div>) : null}
                    </Form>
                )}
            </Formik>
        </div>
    );
};

export default SignUpForm;