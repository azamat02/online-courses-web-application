import React, {Component} from 'react'
import {ProgressBar} from 'react-bootstrap'
import '../../course-page-components/course-rating/course-rating.css'
import {
    AcademicCapIcon, BookOpenIcon,
    LightningBoltIcon, LockClosedIcon,
    PresentationChartBarIcon,
    UserIcon
} from "@heroicons/react/outline";
import Modal from "../../tools/modal";

export default class ProfilePage extends Component {
    constructor(props) {
        super(props);

        this.state = {
            activeButton: 0
        }
    }

    changeButton = (num) => {
        this.setState({activeButton: num})
    }

    render() {
        let {activeButton} = this.state

        return (
            <>
                <div className="grid grid-rows-10 grid-flow-col mt-10 gap-4">
                    {/*Sidebar*/}
                    <div className="col-span-1 row-span-3 ">
                        <div className="sticky top-0 rounded shadow pt-8 px-10 pb-8 h-50 mt">
                            <h1 className="text-3xl font-bold">Profile</h1>
                            <a href="#information">
                                <button onClick={()=> this.changeButton(0)} className={`${activeButton === 0 ? `bg-gray-300` : `bg-gray-50`} hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 text-gray-700 transition px-7 py-3 mt-8 shadow font-bold rounded-lg block w-full text-left`}>
                                    <UserIcon className="w-5 h-5 inline mb-1 mr-2"/>
                                    Profile info
                                </button>
                            </a>
                            <a href="#courses">
                                <button onClick={()=> this.changeButton(1)} className={`${activeButton === 1 ? `bg-gray-300` : `bg-gray-50`} hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 transition text-gray-700 shadow px-7 py-3 mt-4 font-bold rounded-lg block w-full text-left`}>
                                    <BookOpenIcon className="w-5 h-5 inline mb-1 mr-2"/>
                                    My courses
                                </button>
                            </a>
                            <a href="#change_password">
                                <button onClick={()=> this.changeButton(2)} className={`${activeButton === 2 ? `bg-gray-300` : `bg-gray-50`} hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 transition text-gray-700 shadow px-7 py-3 mt-4 font-bold rounded-lg block w-full text-left`}>
                                    <LockClosedIcon className="w-5 h-5 inline mb-1 mr-2"/>
                                    Change password
                                </button>

                            </a>
                        </div>
                    </div>

                    {/*Personal info*/}
                    <div className="col-span-1" id="information">
                        <div className="col-span-3 rounded container shadow pt-2">
                            <div className="px-6 py-2">
                                <h1 className="text-2xl font-medium ">Profile information</h1>
                                <h1 className="text-lg text-gray-500 ">Personal data</h1>
                            </div>
                            <hr className="mt-2 "/>
                            <div className="grid grid-cols-6 bg-gray-100 p-6">
                                <div className="text-gray-600 font-medium col-span-2">Full name</div>
                                <div className="font-medium col-span-4">Azamat Saiduly</div>
                            </div>
                            <div className="grid grid-cols-6 p-6">
                                <div className="text-gray-600 font-medium col-span-2">Email adress</div>
                                <div className="font-medium col-span-4">azamattolegenov1@gmail.com</div>
                            </div>
                            <div className="grid grid-cols-6 bg-gray-100 p-6">
                                <div className="text-gray-600 font-medium col-span-2">Login</div>
                                <div className="font-medium col-span-4">azamat02</div>
                            </div>
                        </div>
                    </div>

                    {/*Courses*/}
                    <div className="col-span-1" id="courses">
                        <div className="container rounded mt-2 shadow pt-2">
                            <div className="px-6 py-2">
                                <h1 className="text-2xl font-medium ">Purchased courses</h1>
                                <h1 className="text-lg text-gray-500 ">Your courses</h1>
                            </div>
                            <hr className="mt-2 mb-2"/>
                            <div>
                                <div className="px-6 mb-2 py-4 cursor-pointer">
                                    <p className="mb-3 text-2xl font-bold align-middle">
                                        <AcademicCapIcon className="w-10 h-10 inline mb-1"/>
                                        Python course
                                        <p className="text-right inline ml-10 text-gray-600 font-medium text-sm">Purchased date: 20-05-2021</p>
                                        <p className="text-right inline ml-10 text-gray-600 font-medium text-sm">Last study date: 20-05-2021</p>
                                    </p>
                                    <p className="text-gray-700 font-bold">Progress 60%</p>
                                    <ProgressBar className="mb-6" now={60} />
                                    <button className="bg-green-500 focus:outline-none focus:ring-2 focus:ring-green-600 focus:ring-opacity-50 py-3 px-10 transition font-bold text-white rounded-lg hover:bg-green-600">
                                        <LightningBoltIcon className="mb-1 w-5 h-5 inline mr-2"/>
                                        Continue study
                                    </button>
                                    <button className="ml-10 bg-green-100 focus:outline-none focus:ring-2 focus:ring-green-600 focus:ring-opacity-50 py-3 px-10 transition font-bold text-green-700 rounded-lg hover:bg-green-200">
                                        <PresentationChartBarIcon className="mb-1 w-5 h-5 inline mr-2"/>
                                        View analytics
                                    </button>
                                </div>
                                <hr/>
                                <div className="px-6 mb-2 py-4 cursor-pointer">
                                    <p className="mb-3 text-2xl font-bold align-middle">
                                        <AcademicCapIcon className="w-10 h-10 inline mb-1"/>
                                        Golang course
                                        <p className="text-right inline ml-10 text-gray-600 font-medium text-sm">Purchased date: 20-05-2021</p>
                                        <p className="text-right inline ml-10 text-gray-600 font-medium text-sm">Last study date: 20-05-2021</p>
                                    </p>
                                    <p className="text-gray-700 font-bold">Progress 80%</p>
                                    <ProgressBar className="mb-6" now={80} />
                                    <button className="bg-green-500 focus:outline-none focus:ring-2 focus:ring-green-600 focus:ring-opacity-50 py-3 px-10 transition font-bold text-white rounded-lg hover:bg-green-600">
                                        <LightningBoltIcon className="mb-1 w-5 h-5 inline mr-2"/>
                                        Continue study
                                    </button>
                                    <button className="ml-10 bg-green-100 focus:outline-none focus:ring-2 focus:ring-green-600 focus:ring-opacity-50 py-3 px-10 transition font-bold text-green-700 rounded-lg hover:bg-green-200">
                                        <PresentationChartBarIcon className="mb-1 w-5 h-5 inline mr-2"/>
                                        View analytics
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>

                    {/*Change password*/}
                    <div className="col-span-1 mb-36" id="change_password">
                        <div className="container rounded mt-2 shadow pt-2 pb-10">
                            <div className="px-6 py-2">
                                <h1 className="text-2xl font-medium ">Change password</h1>
                                <h1 className="text-lg text-gray-500 ">Password details</h1>
                            </div>
                            <hr className="mt-2 mb-2"/>
                            <div className="px-6 py-2">
                                <form>
                                    <div className="grid grid-cols-2 gap-5">
                                        <div>
                                            <label htmlFor="price" className="block text-sm font-medium text-gray-700">Password</label>
                                            <input className="border-gray-400 border-2 focus:outline-none rounded-md py-2 block mb-4 mt-2 focus:ring-indigo-500 focus:border-indigo-500 block w-full px-4 sm:text-sm" type="text" placeholder="Enter password"/>
                                        </div>
                                        <div>
                                            <label htmlFor="price" className="block text-sm font-medium text-gray-700">Repeat password</label>
                                            <input className="border-gray-400 border-2 focus:outline-none rounded-md py-2 block mb-4 mt-2 focus:ring-indigo-500 focus:border-indigo-500 block w-full px-4 sm:text-sm" type="text" placeholder="Enter password"/>
                                        </div>
                                    </div>
                                    <button className="bg-green-500 focus:outline-none focus:ring-2 focus:ring-green-600 focus:ring-opacity-50 py-3 px-10 transition font-bold text-white rounded-lg hover:bg-green-600">
                                        Change password
                                    </button>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            </>
        )
    }
}