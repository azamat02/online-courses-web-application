import React, {useEffect, useState} from 'react'
import { Link } from 'react-router-dom'
import './auth.css'

export default function SignIn(){
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')

    let submit = (e) => {
        e.preventDefault()
        console.log(email, password)
    }

    return (
        <>
            <div className="auth-body">
                <div className="auth-container">
                    <p>Sign in</p>
                    <form className="auth-form" onSubmit={(e)=>submit(e)}>
                        <label for="email" className="block">
                            <svg xmlns="http://www.w3.org/2000/svg" className="icon icon-tabler icon-tabler-mail" width="28" height="28" viewBox="0 0 24 24" stroke-width="1.5" stroke="#3A7CA5" fill="none" stroke-linecap="round" stroke-linejoin="round">
                                <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
                                <rect x="3" y="5" width="18" height="14" rx="2" />
                                <polyline points="3 7 12 13 21 7" />
                            </svg>
                            <input type="email" value={email} onChange={(e)=>{setEmail(e.target.value)}} placeholder="Email"/>
                        </label>
                        <label for="pass" className="block">
                            <svg xmlns="http://www.w3.org/2000/svg" className="icon icon-tabler icon-tabler-lock" width="28" height="28" viewBox="0 0 24 24" stroke-width="1.5" stroke="#3A7CA5" fill="none" stroke-linecap="round" stroke-linejoin="round">
                                <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
                                <rect x="5" y="11" width="14" height="10" rx="2" />
                                <circle cx="12" cy="16" r="1" />
                                <path d="M8 11v-4a4 4 0 0 1 8 0v4" />
                            </svg>
                            <input type="password" value={password} onChange={(e)=>{setPassword(e.target.value)}} placeholder="Password"/>
                        </label>
                        <a href="/" className="right">Forget password?</a>
                        <button>Sign in</button>
                        <Link to="/signup" className="center">New user? Sign up</Link>
                    </form>
                </div>
            </div>
        </>
    )

}