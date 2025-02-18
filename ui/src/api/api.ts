
import axios from "axios"

export default class API {

    private url: string;
    private port: string;
    private path: string;

    constructor() {
        this.url = "http://localhost";
        this.port = ":5000";
        this.path = "/api/v1/folder/";
    }
    /**
     * getFolderInfo return a promise of http get call for reading folder info
     * passed as parameter.
     *
     * @param {string} folderName
     * @returns {Promise<any>}
     * @memberof API
     */
    public getFolderInfo(folderName: string): Promise<any> {
        return axios.post(this.url+this.port+this.path, folderName);
    }
}
