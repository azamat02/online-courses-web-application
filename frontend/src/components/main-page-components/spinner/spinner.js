import React, {Component} from "react";
import './spinner.css'

export default class Spinner extends Component{
    render() {
        return (
            <>
                <div className="loader">Loading...</div>
                <div className="text-center">Loading...</div>
            </>
        )
    }
}