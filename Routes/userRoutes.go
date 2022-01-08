package Routes

import (
	"Donation/Controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUpUSer(app *fiber.App){
	app.Post("/user/addUser",Controllers.AddUser)
	app.Get("/user/getUser/:id",Controllers.GetUser)
	app.Get("/user/getUsers",Controllers.GetAllUsers)
	app.Patch("/user/updateUser/:id",Controllers.UpdateUser)
	app.Delete("/user/deleteUser/:id",Controllers.DeleteUser)
	app.Post("/user/login",Controllers.Login)


}
