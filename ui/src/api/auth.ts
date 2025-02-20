import axios from "axios"

export interface SignUpRequest {
    username: string
    password: string
    name: string
}

export interface SignUpResponse {
    username: string
    name: string
    createdAt: string
    updatedAt: string
}

const signUP = async (url: string, req: SignUpRequest): Promise<SignUpResponse> => {
    const res = await axios.post(url + "/api/auth/sign-up", req);

    return res.data as SignUpResponse
};

export interface SignInRequest {
    username: string
    password: string
}

export interface SignInResponse {
    exp: number
    token: string
}

const signIn = async (url: string, req: SignInRequest): Promise<SignInResponse> => {
    const res = await axios.post(url + "/api/auth/sign-in", req);
    return res.data as SignInResponse;
}

const Auth = (apiUrl: string) => {
    return {
        signUp: (req: SignUpRequest): Promise<SignUpResponse> => signUP(apiUrl, req),
        signIn: (req: SignInRequest): Promise<SignInResponse> => signIn(apiUrl, req),
    };
}

export default Auth;