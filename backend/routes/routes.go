package routes

import (
	"github.com/gofiber/fiber/v2"
	"online-courses-app/controllers"
)

func Setup(app *fiber.App) {
	//Authorization routes
	app.Post("api/sign_up", controllers.SignUp)
	app.Post("api/sign_in", controllers.SignIn)
	app.Get("api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)

	//Courses routes
	app.Get("api/courses/:id", controllers.GetCourse)
	app.Get("api/courses", controllers.GetAllCourses)
	app.Post("api/courses", controllers.CreateCourse)

	//Module routes
	app.Get("api/modules/:id", controllers.GetModule)
	app.Get("api/modules/by_course/:id", controllers.GetModulesByCourseId)
	app.Get("api/modules", controllers.GetAllModules)
	app.Post("api/modules", controllers.CreateModule)
}
