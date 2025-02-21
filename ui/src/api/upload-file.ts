import axios from "axios"
import {WithToken} from "./common.ts";

export interface UploadFile {
    file: File
}

export interface UploadFileResponse {
    id: number
    url: string
}

const uploadFile = async (url: string, token: string, req: UploadFile): Promise<UploadFileResponse> => {
    const formData = new FormData();
    formData.append("file", req.file);

    const res = await axios.post(url + "/api/upload-file", formData, {
        headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "multipart/form-data"
        }
    });

    return res.data as UploadFileResponse;
}

export interface UploadFileClient extends WithToken {
    uploadFile: (req: UploadFile) => Promise<UploadFileResponse>
}

export default function UploadFile(url: string, withToken: WithToken): UploadFileClient {
    return {
        ...withToken,
        uploadFile: (req: UploadFile) => uploadFile(url, withToken.getToken().token, req),
    }
}