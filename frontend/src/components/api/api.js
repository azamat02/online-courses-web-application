export default class CoursesApi{
    constructor() {
        this._apiBase = 'http://localhost:4200';
    }

    getResource = async (url) => {
        const res = await fetch(`${this._apiBase}${url}`);
        if (!res.ok) {
            throw new Error(`Could not fetch ${url}` +
                `, received ${res.status}`);
        }
        return await res.json();
    }

    getAllModules = async () => {
        const res = await this.getResource(`/modules/`);
        return res.map(this._transformModule);
    }

    getAllLessons = async () => {
        const res = await this.getResource(`/lessons/`);
        return res.map(this._transformModule);
    }

    getModulesByCourseId = async (id) =>{
        const res = await this.getResource(`/modules/`);
        return res.filter((elem)=>{
            if (elem.c_id == id){
                return elem
            }
        }).map(this._transformModule);
    }

    getLessonsByModuleId = async (id) =>{
        const res = await this.getResource(`/lessons/`);
        return res.filter((elem)=>{
            if (elem.m_id == id){
                return elem
            }
        }).map(this._transformLesson);
    }

    getAllCourses = async () => {
        const res = await this.getResource(`/courses/`);
        const modules = await  this.getAllModules()
        return res.map(course=>{
            let newElem = {...course, lessonsCount: 0}
            let lessonsCount = 0

            modules.forEach(elem=>{
                if(elem.moduleCourseId == course.id){
                    lessonsCount += + elem.moduleNumberOfLessons
                }
            })

            newElem.lessonsCount = lessonsCount
            return newElem
        }).map(this._transformCourse);
    }

    getCourse = async (id) => {
        const res = await this.getResource(`/courses/${id}`);
        return this._transformCourse(res);
    }

    getModule = async (id) => {
        const res = await this.getResource(`/modules/${id}`);
        return this._transformModule(res);
    }

    getLesson = async (id) => {
        const res = await this.getResource(`/lessons/${id}`);
        return this._transformLesson(res);
    }

    isSet(data) {
        if (data) {
            return data
        } else {
            return 'no data'
        }
    }

    _transformCourse = (course) => {
        return {
            courseId: this.isSet(course.id),
            courseLessons: this.isSet(course.lessonsCount),
            courseImage: this.isSet(course.img),
            courseTitle: this.isSet(course.title),
            courseDesc: this.isSet(course.desc),
            courseReqs: this.isSet(course.req),
            courseAdv: this.isSet(course.what_you_will_learn)
        };
    }

    _transformModule = (module) => {
        return {
            moduleId: this.isSet(module.id),
            moduleTitle: this.isSet(module.title),
            moduleCourseId: this.isSet(module.c_id),
            moduleNumberOfLessons: this.isSet(module.number_of_lessons)
        };
    }

    _transformLesson = (lesson) => {
        return {
            lessonId: this.isSet(lesson.id),
            lessonType: this.isSet(lesson.type),
            lessonModuleId: this.isSet(lesson.m_id),
            lessonTitle: this.isSet(lesson.title),
            lessonLink: this.isSet(lesson.link),
            lessonContent: this.isSet(lesson.content)
        };
    }
}

