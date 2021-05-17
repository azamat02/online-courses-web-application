import React,{Component, useState} from "react"
import {BrowserRouter as Router, Route} from 'react-router-dom'
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

export default class App extends Component{
  constructor(props){
    super(props)

      this.state = {
        coursesList: [],
        coursesSlider: []
      }
  }

  api = new CoursesApi();

  componentDidMount() {
      this.api.getAllCourses().then(data=>{
          this.setState({courses:data})
      })
  }

    render() {
        const {courses} = this.state

        if(courses){
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
                    <Route exact path='/signup/' component={SignUp}>
                        <Navbar/>
                        <SignUp/>
                    </Route>
                    <Route exact path='/course/:id/' render={({match})=>{
                        const {id} = match.params
                        console.log(id);
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
}