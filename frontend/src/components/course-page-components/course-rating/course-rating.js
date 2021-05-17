import React from 'react'
import {StarIcon as StarSolid} from "@heroicons/react/solid";
import {StarIcon as StarOutline} from "@heroicons/react/outline";
import {ProgressBar} from 'react-bootstrap'
import './course-rating.css'

export default function CourseRating(){


    return (
        <div className="mb-24">
            <h1 className="text-3xl font-bold">Students feedback</h1>
            <div className="grid grid-cols-4 gap-4 mt-10 w-4/5">
                <div className="h-52 leading-8 text-yellow-600">
                    <p className="rating font-bold text-8xl">
                        4.0
                    </p>
                    <p className="stars">
                        <StarSolid className="inline w-8 h-8"/>
                        <StarSolid className="inline w-8 h-8"/>
                        <StarSolid className="inline w-8 h-8"/>
                        <StarSolid className="inline w-8 h-8"/>
                        <StarOutline className="inline w-8 h-8"/>
                    </p>
                    <p className="text-2xl font-medium mt-2">
                        Course Rating
                    </p>
                </div>
                <div className="h-52 col-span-2 pt-3">
                    <ProgressBar className="mb-6" now={60} />
                    <ProgressBar className="mb-6" now={23} />
                    <ProgressBar className="mb-6" now={12} />
                    <ProgressBar className="mb-6" now={4} />
                    <ProgressBar className="mb-6" now={1} />
                </div>
                <div className="h-52 text-yellow-600 leading-10">
                    <p className="stars">
                        <StarSolid className="inline w-7 h-7"/>
                        <StarSolid className="inline w-7 h-7"/>
                        <StarSolid className="inline w-7 h-7"/>
                        <StarSolid className="inline w-7 h-7"/>
                        <StarSolid className="inline w-7 h-7"/>
                        <p className='text-blue-500 text-2xl inline ml-2 align-middle'>60%</p>
                    </p>
                    <p className="stars">
                        <StarSolid className="inline w-7 h-7"/>
                        <StarSolid className="inline w-7 h-7"/>
                        <StarSolid className="inline w-7 h-7"/>
                        <StarSolid className="inline w-7 h-7"/>
                        <StarOutline className="inline w-7 h-7"/>
                        <p className='text-blue-500 text-2xl inline ml-2 align-middle'>23%</p>
                    </p>
                    <p className="stars">
                        <StarSolid className="inline w-7 h-7"/>
                        <StarSolid className="inline w-7 h-7"/>
                        <StarSolid className="inline w-7 h-7"/>
                        <StarOutline className="inline w-7 h-7"/>
                        <StarOutline className="inline w-7 h-7"/>
                        <p className='text-blue-500 text-2xl inline ml-2 align-middle'>12%</p>
                    </p>
                    <p className="stars">
                        <StarSolid className="inline w-7 h-7"/>
                        <StarSolid className="inline w-7 h-7"/>
                        <StarOutline className="inline w-7 h-7"/>
                        <StarOutline className="inline w-7 h-7"/>
                        <StarOutline className="inline w-7 h-7"/>
                        <p className='text-blue-500 text-2xl inline ml-2 align-middle'>4%</p>
                    </p>
                    <p className="stars">
                        <StarSolid className="inline w-7 h-7"/>
                        <StarOutline className="inline w-7 h-7"/>
                        <StarOutline className="inline w-7 h-7"/>
                        <StarOutline className="inline w-7 h-7"/>
                        <StarOutline className="inline w-7 h-7"/>
                        <p className='text-blue-500 text-2xl inline ml-2 align-middle'>1%</p>
                    </p>

                </div>
            </div>
        </div>
    )
}
