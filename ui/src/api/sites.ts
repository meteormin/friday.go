import axios from "axios";
import {WithToken} from "./common.ts";

export interface Site {
    id: number
    name: string
    url: string
    createdAt: string
    updatedAt: string
}

