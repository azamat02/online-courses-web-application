import React,{Component} from "react"
import {Link} from 'react-router-dom'
import './navbar.css'

export default class Navbar extends Component{
    render(){
        let {activeLink} = this.props

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
                        <li className="block text-lg flex items-center mr-4 text-blue-500 font-medium"><Link to="/signin">Sign in</Link></li>
                        <li className="block px-6 ml-2 py-2 transition text-medium text-white bg-blue-500 hover:bg-blue-600 rounded items-center flex text-lg"><Link to="/signup">Sign up</Link></li>
                    </ul>
                </nav>
            </>
        )
    }
}