import React, {Component, useContext, useState} from "react"
import {Link} from 'react-router-dom'
import './navbar.css'
import {AppContext} from "../../../stateManager";
import AuthService from "../../auth-service";
import Spinner from "../spinner";

export default function Navbar(props){
    const { appState } = useContext(AppContext)
    const {isAuthorized, userData} = appState

    let {activeLink} = props

    let authBlock = !isAuthorized ? (
        <>
            <li className="block text-lg flex items-center mr-4 text-blue-500 font-medium">
                <Link to="/signin">
                    Sign in
                </Link>
            </li>
            <li className="block px-6 ml-2 py-2 transition text-medium text-white bg-blue-500 hover:bg-blue-600 rounded items-center flex text-lg">
                <Link to="/signup">
                    Sign up
                </Link>
            </li>
        </>
    ) : (
        <>
            <li className="block px-6 ml-2 py-2 transition text-medium text-white bg-blue-500 hover:bg-blue-600 rounded items-center flex text-lg">
                <Link to="/signup">
                    {userData ? userData.name : ''}
                </Link>
            </li>
        </>
    )

    return (
        <>
            <nav className="navbar">
                <ul className="nav-links">
                    <Link to="/" className="logo">
                            <span>
                                ONLINE COURSES
                            </span>
                    </Link>
                    <li className={`${activeLink === 'main' ? `text-blue-500` : ``} ml-12 block text-lg flex items-center mr-4 font-medium`}><Link to="/">Main page</Link></li>
                    <li className={`${activeLink === 'contacts' ? `text-blue-500` : ``} ml-8 block text-lg flex items-center mr-4 font-medium`}><Link to="/contacts">Contacts</Link></li>
                    <li className={`${activeLink === 'help' ? `text-blue-500` : ``} ml-8 block text-lg flex items-center mr-4 font-medium`}><Link to="/help">Help</Link></li>
                </ul>
                <ul className="nav-links">
                    {authBlock}
                </ul>
            </nav>
        </>
    )
}