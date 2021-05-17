import React,{Component} from "react"
import { Link } from "react-router-dom";
import './courses-slider.css'
import Spinner from "../spinner";

export default class CoursesSlider extends Component{
    constructor(props) {
        super(props)

        this.state = {
            courses: props.courses,
            timerId: null,
            currentSlide: 0   
        }
    }

    componentDidMount() {
        this.timerId = setInterval(
            ()=>{this.changeData()}, 5000
        )
    }

    componentWillUnmount() {
        clearInterval(this.timerId)
    }

    changeData = ()=>{
        let dataLength = this.state.courses.length
        dataLength -= 1
        if (this.state.currentSlide === dataLength) {
            this.setState({currentSlide:0})
        } else {
            let num = this.state.currentSlide
            num++
            this.setState({currentSlide: num})
        }
    }

    nextSlide = ()=>{
        let {courses} = this.state
        let {currentSlide} = this.state
        let num = currentSlide;
        if(currentSlide === courses.length-1) {
            num = 0
        }
        else {
            num ++
        }
        this.setState({currentSlide:num})
        console.log("next");
    }

    prevSlide = ()=>{
        let {courses} = this.state
        let {currentSlide} = this.state
        let num = currentSlide;
        if(currentSlide === 0) {
            num = courses.length-1
        }
        else {
            num --
        }
        this.setState({currentSlide:num})
        console.log("prev");
    }

    render() {
        let {currentSlide, courses} = this.state

        if(!courses) {
            return <Spinner/>
        }

        let {courseId, courseImage, courseTitle, courseDesc} = courses[currentSlide]


        let dots = courses.map((elem,index)=>{
            if(index === currentSlide){
                return (<span className="active-dot"></span>)
            } else {
                return (<span></span>)
            }
        })
        return (
            <>
                <div className="slider-container shadow-md transition">
                    <div className="slider-image">
                        <img className="course-img" src={courseImage} alt=""/>
                    </div>
                    <div className="slider-content">
                        <div className="indicators">
                            <div className="indicator-dots">
                                  {dots}
                            </div>
                            <div className="indicator-controllers">
                                <span onClick={()=>this.prevSlide()}><i className="fas fa-arrow-left"></i></span>
                                <span onClick={()=>this.nextSlide()}><i className="fas fa-arrow-right"></i></span>
                            </div>
                        </div>
                        <p className="course-title">
                            {courseTitle}
                        </p>
                        <p className="course-desc">
                            {courseDesc}
                        </p>
                        <Link to={'/course/'+courseId+'/'} className="px-4 py-3 hover:bg-blue-600 transition shadow text-lg font-bold text-white mt-10 block bg-blue-500 w-2/4 text-center rounded">
                            Read more
                        </Link>
                    </div>
                </div>
            </>
        )
    }

}