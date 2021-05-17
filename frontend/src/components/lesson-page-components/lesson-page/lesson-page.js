import React, {Component} from 'react'
import CoursesApi from "../../api";
import Spinner from "../../main-page-components/spinner";
import {Link} from "react-router-dom";
import {Menu, Transition} from '@headlessui/react'
import {AdjustmentsIcon, LogoutIcon, UserCircleIcon, UserIcon} from "@heroicons/react/outline";
import ReactPlayer from "react-player";

import { Fragment, useEffect, useRef, useState } from 'react'
import { ChevronDownIcon } from '@heroicons/react/solid'
import ModuleItem from "../../course-page-components/module-item";
import Modal from "../../tools/modal";

export default class LessonPage extends Component {
    constructor(props) {
        super(props);

        this.isAutorized = false

        this.state = {
            id: props.id,
            lesson: null,
            modules: null
        }
    }

    api = new CoursesApi()

    componentDidMount() {
        // Getting lesson and modules
        this.api.getLesson(this.state.id).then((data)=>{
            this.api.getModule(data.lessonModuleId).then(data=>{
                this.api.getModulesByCourseId(data.moduleCourseId).then(data=>{
                    this.setState({
                        modules: data
                    })
                })
            })
            this.setState({
                lesson: data
            })
        })
    }

    componentDidUpdate() {
        if (this.props.id != this.state.id){
            this.setState({id: this.props.id, lesson: null})
            // Getting lesson and modules
            this.api.getLesson(this.props.id).then((data)=>{
                this.setState({
                    lesson: data
                })
            })
        }
    }

    render() {
        if (!this.isAuthorized) {
            return <Modal info="You are not authorized. Please sign in/up first" buttonLink={`/`} title="Error" open={true}/>
        }

        let {lesson, modules} = this.state
        if (!lesson || !modules) {
            return <Spinner/>
        }

        let {lessonId,
            lessonType,
            lessonModuleId,
            lessonTitle,
            lessonLink,
            lessonContent} = lesson

        // Set video or lecture container
        let videoContainer = lessonLink !== "no data" ? (<div className="video-container w-full h-full bg-gray-700">
                                                            <ReactPlayer url={lessonLink} controls width="100%" height="100%"/>
                                                        </div>) : (<h1 className="text-2xl font-bold text-red-500 border-l-4 pl-4 border-red-500 bg-red-200 py-3" >No video link added yet</h1>)

        let lectureContainer = (<div className="content mt-5 mb-32">
                                    <h1 className="text-3xl font-bold border-l-4 border-gray-900 p-2 bg-indigo-100 text-gray-900">Lecture: {lessonTitle}</h1>
                                    <p className="text-lg mt-2 italic p-2">{lessonContent}</p>
                                </div>)

        let lessonContainer = lessonType === "video" ? videoContainer : lectureContainer
        let rightBlock = lessonType === "video" ? (<div className="col-span-2">
                                                        <hr/>
                                                        <div className="modules my-3">
                                                            <h1 className="text-gray-900 text-2xl font-bold my-4">{lessonTitle}</h1>
                                                            <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Eos exercitationem iusto, magnam numquam officia provident rem tempora temporibus tenetur ut.</p>
                                                        </div>
                                                        <hr/>
                                                    </div>) :null

        // Rendering modules
        let renderedModules = "No modules added yet."
        if(modules.length != 0) {
            renderedModules = modules.map((elem)=>{
                return <ModuleItem id={elem.moduleId}/>
            })
        }

        return (
            <>
                {/*Navbar*/}
                <section className="h-16 bg-gray-800 pl-44 pr-44 items-center relative flex justify-between">
                    <div className="flex relative items-center">
                        <Link to="/" className="text-2xl font-bold text-white mr-10">ONLINE COURSES</Link>
                        <Link to="/" className="bg-gray-900 text-white px-3 py-2 rounded-md font-medium mr-2">Courses</Link>
                        <Link to="/" className="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md font-medium mr-2">Contacts</Link>
                        <Link to="/" className="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md font-medium mr-2">Help</Link>
                    </div>
                    <div className="flex relative">
                        <div className="flex">
                            <div className="relative inline-block text-left">
                                <Menu>
                                    {({ open }) => (
                                        <>
                                            <Menu.Button className="focus:outline-none inline-flex justify-center w-full">
                                                <button className="hover:bg-indigo-600 focus:ring-2 transition flex items-center text-gray-50 relative focus:outline-none px-5 py-2 bg-indigo-500 rounded-lg">
                                                        <span className="text-md mr-1 font-bold">
                                                            Azamat
                                                        </span>
                                                    <UserCircleIcon className="inline w-7 h-7"/>
                                                </button>
                                            </Menu.Button>

                                            <Transition
                                                show={open}
                                                enter="transition ease-out duration-100"
                                                enterFrom="transform opacity-0 scale-95"
                                                enterTo="transform opacity-100 scale-100"
                                                leave="transition ease-in duration-75"
                                                leaveFrom="transform opacity-100 scale-100"
                                                leaveTo="transform opacity-0 scale-95"
                                            >
                                                <Menu.Items static className="absolute right-0 w-56 mt-2 origin-top-right bg-gray-800 text-white mt-4 divide-y divide-gray-100 rounded-md shadow-xl outline-none">
                                                    <div className="py-1">
                                                        <Menu.Item>
                                                            {({ active }) => (
                                                                <Link
                                                                    to="/profile/"
                                                                    className={`${
                                                                        active
                                                                            ? "bg-gray-100 text-gray-900"
                                                                            : "text-white"
                                                                    } flex mx-4 rounded my-2 px-4 py-2 font-medium text-sm leading-5`}
                                                                >
                                                                    <UserIcon className="w-5 h-5 inline mr-2"/>
                                                                    Your profile
                                                                </Link>
                                                            )}
                                                        </Menu.Item>
                                                        <Menu.Item>
                                                            {({ active }) => (
                                                                <Link
                                                                    to="/settings/"
                                                                    className={`${
                                                                        active
                                                                            ? "bg-gray-100 text-gray-900"
                                                                            : "text-white"
                                                                    } flex mx-4 rounded my-2 px-4 py-2 font-medium text-sm leading-5`}
                                                                >
                                                                    <AdjustmentsIcon className="w-5 h-5 inline mr-2"/>
                                                                    Settings
                                                                </Link>
                                                            )}
                                                        </Menu.Item>
                                                        <hr className="mx-7"/>
                                                        <Menu.Item>
                                                            {({ active }) => (
                                                                <Link
                                                                    to="/signout/"
                                                                    className={`${
                                                                        active
                                                                            ? "bg-gray-100 text-gray-900"
                                                                            : "text-white"
                                                                    } flex mx-4 rounded my-2 px-4 py-2 font-medium text-sm leading-5`}
                                                                >
                                                                    <LogoutIcon className="w-5 h-5 inline mr-2"/>
                                                                    Sign out
                                                                </Link>
                                                            )}
                                                        </Menu.Item>
                                                    </div>
                                                </Menu.Items>
                                            </Transition>
                                        </>
                                    )}
                                </Menu>
                            </div>
                        </div>
                    </div>
                </section>
                <main>
                    <div className="container mt-5 grid grid-cols-6 gap-4 mb-32">
                        <div className="col-span-4">
                            {lessonContainer}
                            <br/><br/>
                            <hr/>
                            <h1 className="text-gray-900 text-2xl font-bold my-4">Course materials</h1>
                            {renderedModules}
                        </div>
                        {rightBlock}
                    </div>
                </main>
            </>
        )
    }
}



