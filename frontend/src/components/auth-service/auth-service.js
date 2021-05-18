import React, {Component} from "react";

export default class AuthService extends Component{
    isAuthorized = async () => {
        let res = await fetch("http://localhost:8000/api/user", {
            headers: {'Content-type': 'application/json'},
            credentials: 'include'
        })

        let user = await res.json()

        if (user.id == null) {
            return false
        } else {
            return user
        }
    }

    SignIn = async (email, password) => {
        let response = await fetch("http://localhost:8000/api/sign_in", {
            method: "POST",
            headers: {'Content-type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify({
                email,
                password
            })
        })

        let json = await response.json()

        return json.message
    }

    SignUp = async (email, name, surname, login, pass) => {
        let response = await fetch("http://localhost:8000/api/sign_up", {
            method: "POST",
            headers: {'Content-type': 'application/json'},
            body: JSON.stringify({
                name,
                surname,
                email,
                login,
                password: pass
            })
        })
        let json = await response.json()

        return json.id
    }

    LogOut = async () => {
        await fetch("http://localhost:8000/api/logout", {
            method: "POST",
            headers: {'Content-type': 'application/json'},
            credentials: 'include'
        })
    }
}