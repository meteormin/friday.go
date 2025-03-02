import {AxiosError} from 'axios';
import {Token, WithToken} from './common';
import Auth, {AuthClient} from './auth';
import User, {UsersClient} from "./users.ts";
import UploadFile, {UploadFileClient} from "./upload-file.ts";
import dayjs from "dayjs";

const getToken = (): Token | null => {
    const token = JSON.parse(localStorage.getItem('token') ?? "{}") as Token;
    const exp = dayjs(token.expiresAt, 'YYYY-MM-DD HH:mm:ss');
    const now = dayjs();
    if (now.isBefore(exp)) {
        return token;
    }

    return null;
}

const setToken = (token: Token) => {
    localStorage.setItem('token', JSON.stringify(token));
}

const handleError = (error: AxiosError) => {
    if (error.status === undefined) {
        alert(error.message);
    } else if (error.status === 401) {
        alert("로그인이 필요합니다.");
        window.location.href = "/sign-in";
    } else if (error.status >= 500) {
        alert("서버에러가 발생했습니다.");
    } else {
        alert(error.message);
    }
}

function ErrorProxy<T extends WithToken>(client: T): T {
    return new Proxy(client, {
        get(target, prop, receiver) {
            const original = Reflect.get(target, prop, receiver);

            // 만약 속성이 함수라면, 프록시를 씌워 실행 시 에러 핸들링 추가
            if (typeof original === "function") {
                return (...args: any[]) => {
                    try {
                        console.log(`Calling method: ${String(prop)} with args:`, args);
                        const result = original.apply(target, args);
                        console.log(`Method ${String(prop)} executed successfully.`);
                        return result;
                    } catch (error) {
                        console.error(`Error in method: ${String(prop)}`, error);
                        if (error instanceof AxiosError) {
                            handleError(error);
                        } else {
                            throw error;
                        }
                    }
                };
            }

            // 함수가 아니면 원래 속성 그대로 반환
            return original;
        },
    });
}

export interface ApiClient {
    getToken: () => Token | null,
    setToken: (token: Token) => void,
    auth: AuthClient
    users: UsersClient
    uploadFile: UploadFileClient
}

export function newApiClient(apiUrl: string): ApiClient {
    const withToken = {getToken, setToken};
    return {
        ...withToken,
        auth: ErrorProxy(Auth(apiUrl, withToken)),
        users: ErrorProxy(User(apiUrl, withToken)),
        uploadFile: ErrorProxy(UploadFile(apiUrl, withToken)),
    }
}