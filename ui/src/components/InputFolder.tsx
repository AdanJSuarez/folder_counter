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
    private api: API;
    constructor(props: IInputFolderProps) {
        super(props)
        this.api = new API();
        this.state = {
            folderName: "",
            folderInfo: {listOfComponent:null, totalSize: 0, totalNumberOfFiles: 0},
            notFound: false
        };
        this.getFolderInfo = this.getFolderInfo.bind(this);
    }
    /**
     * getFolderInfo set component state with the info received from the API.
     *
     * @private
     * @memberof InputFolder
     */
    public getFolderInfo() {
        this.api.getFolderInfo(this.state.folderName)
        .then((res: any)=> {
            if (res.status === 204) {
                this.setState({folderInfo:{listOfComponent: null, totalSize: 0, totalNumberOfFiles: 0}, notFound: true})
            } else if (res.data.listOfComponent) {
                this.setState({folderInfo: res.data, notFound: false});
            } else {
                this.setState({folderInfo:{listOfComponent: null, totalSize: 0, totalNumberOfFiles: 0}});
            }
        })
        .catch((err: any) => {
            this.setState({folderInfo:{listOfComponent: null}});
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
                    filesInfo.listOfComponent ?
                        <table>
                            <tbody>
                                <tr>
                                    <th>File name</th>
                                    <th>File size</th>
                                    <th>Last modification</th>
                                    <th>File is folder</th>
                                </tr>
                                {
                                    filesInfo.listOfComponent.map((component: any, index: number) =>
                                    <tr key={index}>
                                        <td  key={index+1}>{JSON.stringify(component.name)}</td>
                                        <td  key={index+2}>{JSON.stringify(component.size)}</td>
                                        <td  key={index+3}>{JSON.stringify(component.lastModification)}</td>
                                        <td  key={index+4}>{JSON.stringify(component.isFolder)}</td>
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
