import React, { Component } from 'react'

export default class InputFolder extends Component {
    render() {
        return (
            <div>
                <input type="text" id="folderName" name="folderName"></input>
                <button>Submit</button>
            </div>
        )
    }
}
