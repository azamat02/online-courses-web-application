import React, {Component} from 'react'
import CourseItem from '../course-item/course-item'
import './courses-list.css'
import Spinner from "../spinner";

export default class CoursesList extends Component{
    constructor(props) {
        super(props)

        this.state = {
            courses: props.courses
        }
    }

    renderItems = (arr)=>{
        return arr.map((elem)=>{
            let {courseLessons, courseImage, courseTitle, courseDesc, courseId} = elem
            return <CourseItem courseLessons={courseLessons} courseImage={courseImage} courseTitle={courseTitle} courseDesc={courseDesc} courseId={courseId}/>
        })
    }

    render() {

        let {courses} = this.state
        let data = null

        if(!courses) {
            data = <Spinner/>
        } else {
            data = this.renderItems(courses)
        }

        return (
            <>
                <p className="title">Courses</p>
                <div className="courses-list">
                    {data}
                </div>
            </>
        )
    }
}