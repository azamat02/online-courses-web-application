import React from 'react'
import {StarIcon as StarSolid} from "@heroicons/react/solid";
import {StarIcon as StarOutline} from "@heroicons/react/outline";

export default function CommentsSection(){
    return (
        <div className="mb-36">
            <h1 className="text-3xl font-bold mb-10">Comments</h1>
            <div className="comment-block rounded-2xl bg-gray-100 transition shadow w-4/5 p-6 mb-5">
                <p id="name" className="font-medium text-lg text-blue-600">Azamat Saiduly</p>
                <p className="stars mb-2 text-yellow-500">
                    <StarSolid className="inline w-5 h-5"/>
                    <StarSolid className="inline w-5 h-5"/>
                    <StarSolid className="inline w-5 h-5"/>
                    <StarSolid className="inline w-5 h-5"/>
                    <StarSolid className="inline w-5 h-5"/>
                </p>
                <p id="comment" className="italic mb-2">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Aspernatur ipsam, magnam minima minus nostrum rem rerum? Cupiditate eum illum voluptas.</p>
                <p id="date" className="text-gray-600">5 may 2021</p>
            </div>
            <div className="comment-block rounded-2xl bg-gray-100 transition shadow w-4/5 p-6 mb-5">
                <p id="name" className="font-medium text-lg text-blue-600">Azamat Saiduly</p>
                <p className="stars mb-2 text-yellow-500">
                    <StarSolid className="inline w-5 h-5"/>
                    <StarSolid className="inline w-5 h-5"/>
                    <StarSolid className="inline w-5 h-5"/>
                    <StarOutline className="inline w-5 h-5"/>
                    <StarOutline className="inline w-5 h-5"/>
                </p>
                <p id="comment" className="italic mb-2">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Aspernatur ipsam, magnam minima minus nostrum rem rerum? Cupiditate eum illum voluptas.</p>
                <p id="date" className="text-gray-600">5 may 2021</p>
            </div>

        </div>
    )
}
