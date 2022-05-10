import { Route } from "react-router-dom";
import SignUpForm from "./components/SignUpForm";

export default function Pages() {
    return (
        <Route path="/signin">
            <SignUpForm/>
        </Route>
    );
};