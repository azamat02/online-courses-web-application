import React, {Component, useEffect, useState} from "react"
import {BrowserRouter as Router, Link, Route} from 'react-router-dom'
import './App.css'
import CoursesSlider from "./components/main-page-components/courses-slider"
import Navbar from './components/main-page-components/navbar'
import CoursesList from "./components/main-page-components/courses-list"
import Footer from "./components/main-page-components/footer"
import CoursePage from './components/course-page-components/course-page'
import {SignIn, SignUp} from './components/main-page-components/authorization'
import CoursesApi from "./components/api";
import ContactsPage from './components/contacts-page-components'
import Spinner from "./components/main-page-components/spinner";
import ProfilePage from "./components/profile-page-components/profile-page";
import LessonPage from "./components/lesson-page-components/lesson-page";
import PaymentsPage from "./components/payment-page-components/payments-page";
import AuthService from "./components/auth-service";


export default function App(){
    const [courses, setCourses] = useState([])

    let api = new CoursesApi();
    let authService = new AuthService()

    useEffect(()=>{
        if (courses.length == 0){
            api.getAllCourses().then(data=>{
                setCourses(data)
            })
        }
    },[courses])

    let checkAuth = () =>{
      (
          async () =>{
              let isAuthorized = await authService.isAuthorized()
              if (isAuthorized) {
                  console.log(`User name: ${isAuthorized.name}`)
              }
          }
      )()
    }

    if(courses.length != 0){
        return (
            <Router>
                <Route exact path="/">
                    <Navbar activeLink="main"/>
                    <div className="app">
                        <main>
                            <CoursesSlider courses = {courses}/>
                            <CoursesList courses = {courses}/>
                        </main>
                        <Footer/>
                    </div>
                </Route>
                <Route exact path='/signin/'>
                    <Navbar/>
                    <SignIn/>
                </Route>
                <Route exact path='/signup/'>
                    <Navbar/>
                    <SignUp/>
                </Route>
                <Route exact path='/course/:id/' render={({match})=>{
                    const {id} = match.params
                    return (
                        <>
                            <Navbar/>
                            <main>
                                <CoursePage courseId={id}/>
                            </main>
                            <Footer/>
                        </>
                    )
                }}/>
                <Route exact path='/contacts/'>
                    <Navbar activeLink='contacts'/>
                    <ContactsPage/>
                </Route>
                <Route exact path='/profile/' render={()=>{
                    return (
                        <>
                            <Navbar/>
                            <main>
                                <ProfilePage/>
                            </main>
                        </>
                    )
                }}/>
                <Route exact path='/lesson/:id' render={({match})=>{
                    const {id} = match.params
                    return (
                        <LessonPage id={id}/>
                    )
                }}/>
                <Route exact path='/payments/:id' render={({match})=>{
                    const {id} = match.params
                    return (
                        <>
                            <Navbar/>
                            <main>
                                <PaymentsPage courseId={id}/>
                            </main>
                        </>
                    )
                }
                } />
            </Router>
        )
    } else {
        return <Spinner/>
    }
}