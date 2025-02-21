import axios from "axios"
import {WithToken} from "./common.ts";

export interface User {
    id: number
    username: string
    name: string
    createdAt: string
    updatedAt: string
}

const getUsers = async (url: string, token: string): Promise<User[]> => {
    const res = await axios.get(url + "/api/users", {
        headers: {
            Authorization: `Bearer ${token}`
        }
    });

    return res.data as User[];
}

export interface UsersClient extends WithToken {
    getUsers: () => Promise<User[]>
}

export default function Users(url: string, withToken: WithToken): UsersClient {
    return {
        ...withToken,
        getUsers: () => getUsers(url, withToken.getToken().token),
    }
}
