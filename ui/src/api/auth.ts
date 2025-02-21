import axios from "axios"
import {WithToken} from "./common.ts";

export interface SignUpRequest {
    username: string
    password: string
    name: string
}

export interface SignUpResponse {
    id: number
    username: string
    name: string
    createdAt: string
    updatedAt: string
}

const signUP = async (url: string, req: SignUpRequest): Promise<SignUpResponse> => {
    const res = await axios.post(url + "/api/auth/sign-up", req);

    return res.data as SignUpResponse
}

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

export interface User {
    id: number
    username: string
    name: string
    createdAt: string
    updatedAt: string
}


const me = async (url: string, token: string): Promise<User> => {
    const res = await axios.get(url + "/api/auth/me", {
        headers: {
            Authorization: `Bearer ${token}`
        }
    });

    return res.data as User;
}

export interface AuthClient extends WithToken {
    signUp: (req: SignUpRequest) => Promise<SignUpResponse>
    signIn: (req: SignInRequest) => Promise<SignInResponse>
    me: () => Promise<User>
}

export default function Auth(apiUrl: string, withToken: WithToken): AuthClient {
    return {
        ...withToken,
        signUp: (req: SignUpRequest): Promise<SignUpResponse> => signUP(apiUrl, req),
        signIn: (req: SignInRequest): Promise<SignInResponse> => signIn(apiUrl, req),
        me: (): Promise<User> => me(apiUrl, withToken.getToken().token)
    };
};