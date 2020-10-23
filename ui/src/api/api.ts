
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

    public getFolderInfo(folderName: string): Promise<any> {
        if (folderName === "") {
            folderName = 'xxxoooxxx'
        }
        return axios.get(this.url+this.port+this.path+folderName+"/");
    }
}
