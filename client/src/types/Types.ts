export type AppProps = {
    username: string,
    token: string
};

export type Validation = {
    name: string,
    is_valid: boolean
};

export interface FormValues {
    username: string,
    email: string,
    password: string,
    conf_pass: string,
};

export interface SignUpProps {};

export interface IState {
    username: string | null,
    email: string | null, 
    password: string | null,
    confirm_password: string | null,
    pending: boolean,
    username_valid: Validation[],
    email_valid: Validation[],
    password_valid: Validation[]
};