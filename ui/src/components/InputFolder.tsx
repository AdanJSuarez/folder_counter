import React, { Component } from 'react'
import API from "../api/api"
import "./inputFolder.css"

interface IInputFolderProps {
    folderName: string;
}

interface IInputFolderState {
    folderName: string;
    folderInfo: any;
    notFound: boolean;
}
export default class InputFolder extends Component<IInputFolderProps, IInputFolderState> {
    constructor(props: IInputFolderProps) {
        super(props)
        this.state = {
            folderName: "xxxoooxxx",
            folderInfo: {filesStats:null, totalSize: 0, totalNumberOfFiles: 0},
            notFound: false
        };
        this.getFolderInfo = this.getFolderInfo.bind(this);
    }
    private getFolderInfo(): any {
        const api = new API();
        api.getFolderInfo(this.state.folderName)
        .then((res: any)=> {
            if (typeof(res.data)==="string") {
                this.setState({folderInfo:{filesStats: null, totalSize: 0, totalNumberOfFiles: 0}, notFound: true})
            } else if (res.data.filesStats) {
                this.setState({folderInfo: res.data, notFound: false});
            } else {
                this.setState({folderInfo:{filesStats: null, totalSize: 0, totalNumberOfFiles: 0}});
            }
        })
        .catch((err: any) => {
            this.setState({folderInfo:{filesStats: null}});
        })
    }
    render() {
        const filesInfo = this.state.folderInfo;
        return (
            <div>
                <input  onChange={evt => this.setState({folderName: evt.target.value})}></input>
                <button onClick={this.getFolderInfo}>Submit</button>
                <div>
                    {this.state.notFound ? <div>Folder not found!</div>: null}
                    <div>Total size: {filesInfo.totalSize}</div>
                    <div>Total number of files: {filesInfo.totalNumberOfFiles}</div>
                </div>
                {
                    filesInfo.filesStats ?
                        <table>
                            <tbody>
                                <tr>
                                    <th>File name</th>
                                    <th>File size</th>
                                    <th>Last modification</th>
                                    <th>File is folder</th>
                                </tr>
                                {
                                    filesInfo.filesStats.map((fileStat: any, index: number) =>
                                    <tr>
                                        <td  key={index+1}>{JSON.stringify(fileStat.fileName)}</td>
                                        <td  key={index+2}>{JSON.stringify(fileStat.fileSize)}</td>
                                        <td  key={index+3}>{JSON.stringify(fileStat.fileLastModification)}</td>
                                        <td  key={index+4}>{JSON.stringify(fileStat.fileIsDirectory)}</td>
                                    </tr>
                                    )
                                }
                            </tbody>
                        </table>:
                    null
                }
            </div>
        )
    }
}
