package Controllers

import (
	"Donation/DbFunctions"
	"Donation/Models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AddUser(ctx *fiber.Ctx)error{
	var user Models.User
	err:=ctx.BodyParser(&user)
	if  err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "error in parsing user",
		})

	}
	err=DbFunctions.InsertUser(user)
	if err!=nil{
		return ctx.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "error in adding user",
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "user added successfully",
	})
}

func GetUser(ctx *fiber.Ctx)error{
	id:=ctx.Params("id")
	userId,_:=strconv.Atoi(id)
	user,err:=DbFunctions.RetrieveUser(userId)
	if err!=nil{
		return ctx.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "error in getting user",
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "user found successfully",
		"user": user,
	})
}

func GetAllUsers(ctx *fiber.Ctx)error{
	users,err:=DbFunctions.RetrieveAllUsers()
	if err!=nil{
		return ctx.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "error in getting users",
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "users found successfully",
		"users": users,
	})
}

func UpdateUser(ctx *fiber.Ctx)error{
	id:=ctx.Params("id")
	userId,_:=strconv.Atoi(id)
	var user Models.User
	err:=ctx.BodyParser(&user)
	if  err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "error in parsing user",
		})

	}
	err=DbFunctions.UpdateTheUser(userId,user)
	if err!=nil{
		return ctx.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "error in updating user",
		})
	}
	return ctx.Status(200).JSON(fiber.Map{

		"status": "success",
		"message": "user updated successfully",
	})
}

func DeleteUser(ctx *fiber.Ctx)error{
	id:=ctx.Params("id")
	userId,_:=strconv.Atoi(id)
	err:=DbFunctions.DeleteTheUser(userId)
	if err!=nil{
		return ctx.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "error in deleting user",
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "user deleted successfully",
	})
}