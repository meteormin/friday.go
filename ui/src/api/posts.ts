import axios from "axios"
import {WithToken} from "./common.ts";

export interface Post {
    id: number
    title: string
    content: string
    url: string
    tags: string[]
    createdAt: string
    updatedAt: string
}

export interface CreatePostRequest {
    title: string
    content: string
    tags: string[]
    fileId: number
    siteId: number
}

export interface UpdatePostRequest {
    title: string
    content: string
    tags: string[]
    fileId: number
}

export interface PostsClient extends WithToken {
    getPosts: () => Promise<Post[]>
    getPost: (id: number) => Promise<Post>
    createPost: (req: CreatePostRequest) => Promise<Post>
    updatePost: (id: number, req: UpdatePostRequest) => Promise<Post>
    deletePost: (id: number, ) => Promise<void>
}

const getPosts = async (url: string, token: string): Promise<Post[]> => {
    const res = await axios.get(url + "/api/posts", {
        headers: {
            Authorization: `Bearer ${token}`
        }
    });

    return res.data as Post[];
}

const getPost = async (url: string, token: string, id: number): Promise<Post> => {
    const res = await axios.get(url + `/api/posts/${id}`, {
        headers: {
            Authorization: `Bearer ${token}`
        }
    });

    return res.data as Post;
}

const createPost = async (url: string, token: string, req: CreatePostRequest): Promise<Post> => {
    const res = await axios.post(url + "/api/posts", req, {
        headers: {
            Authorization: `Bearer ${token}`
        }
    });

    return res.data as Post;
}

const updatePost = async (url: string, token: string, id: number, req: UpdatePostRequest): Promise<Post> => {
    const res = await axios.put(url + `/api/posts/${id}`, req, {
        headers: {
            Authorization: `Bearer ${token}`
        }
    });

    return res.data as Post;
}

const deletePost = async (url: string, token: string, id: number): Promise<void> => {
    await axios.delete(url + `/api/posts/${id}`, {
        headers: {
            Authorization: `Bearer ${token}`
        }
    });
};

export default function Posts(url: string, withToken: WithToken): PostsClient {
    return {
        ...withToken,
        getPosts: () => getPosts(url, withToken.getToken().token),
        getPost: (id: number) => getPost(url, withToken.getToken().token, id),
        createPost: (req: CreatePostRequest) => createPost(url, withToken.getToken().token, req),
        updatePost: (id: number, req: UpdatePostRequest) => updatePost(url, withToken.getToken().token, id, req),
        deletePost: (id: number) => deletePost(url, withToken.getToken().token, id)
    }
}