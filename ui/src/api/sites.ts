import {WithToken} from "./common.ts";
import axios from "axios";

export interface Site {
    id: number
    host: string
    name: string
    createdAt: string
    updatedAt: string
}

export interface CreateSite {
    host: string
    name: string
}

export interface UpdateSite {
    name: string
}

export interface SiteClient extends WithToken {
    getSite: (id: number) => Promise<Site>
    getSites: () => Promise<Site[]>
    createSite: (req: CreateSite) => Promise<Site>
    updateSite: (id: number, req: UpdateSite) => Promise<Site>
    deleteSite: (id: number) => Promise<void>
}

const getSite = async (url: string, token: string, id: number): Promise<Site> => {
    const res = await axios.get(url + `/api/sites/${id}`, {
        headers: {
            Authorization: `Bearer ${token}`
        }
    });

    return res.data as Site;
}

const getSites = async (url: string, token: string): Promise<Site[]> => {
    const res = await axios.get(url + "/api/sites", {
        headers: {
            Authorization: `Bearer ${token}`
        }
    });

    return res.data.content as Site[];
}

const createSite = async (url: string, token: string, req: CreateSite): Promise<Site> => {
    const res = await axios.post(url + "/api/sites", req, {
        headers: {
            Authorization: `Bearer ${token}`
        }
    });

    return res.data as Site;
}

const updateSite = async (url: string, token: string, id: number, req: UpdateSite): Promise<Site> => {
    const res = await axios.put(url + `/api/sites/${id}`, req, {
        headers: {
            Authorization: `Bearer ${token}`
        }
    });

    return res.data as Site;
}

const deleteSite = async (url: string, token: string, id: number): Promise<void> => {
    await axios.delete(url + `/api/sites/${id}`, {
        headers: {
            Authorization: `Bearer ${token}`
        }
    });
}

export function Sites(apiUrl: string, withToken: WithToken): SiteClient {
    return {
        ...withToken,
        getSite: (id: number) => getSite(apiUrl, withToken.getToken().token, id),
        getSites: () => getSites(apiUrl, withToken.getToken().token),
        createSite: (req: CreateSite) => createSite(apiUrl, withToken.getToken().token, req),
        updateSite: (id: number, req: UpdateSite) => updateSite(apiUrl, withToken.getToken().token, id, req),
        deleteSite: (id: number) => deleteSite(apiUrl, withToken.getToken().token, id),
    };
}