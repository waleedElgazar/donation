package Controllers

import (
	"Donation/DbFunctions"
	"Donation/Models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mail.v2"
	"os"
	"strconv"
)

func AddUser(ctx *fiber.Ctx) error {
	var user Models.User
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error in parsing user",
		})
	}
	otp := DbFunctions.CreateOtp()
	verification:=Models.VerificationCode{
		VerificationCode:otp,
		UserEmail:user.Email,
	}
	_=SendVerifyCode(verification)
	err = DbFunctions.InsertUser(user)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error in adding user",
		})
	}
	DbFunctions.AddVerificationCode(verification)
	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "user added successfully",
		"verificationCode":verification,
	})
}

func GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userId, _ := strconv.Atoi(id)
	user, err := DbFunctions.RetrieveUser(userId)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error in getting user",
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "user found successfully",
		"user":    user,
	})
}

func GetAllUsers(ctx *fiber.Ctx) error {
	users, err := DbFunctions.RetrieveAllUsers()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error in getting users",
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "users found successfully",
		"users":   users,
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userId, _ := strconv.Atoi(id)
	var user Models.User
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error in parsing user",
		})

	}
	err = DbFunctions.UpdateTheUser(userId, user)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error in updating user",
		})
	}
	return ctx.Status(200).JSON(fiber.Map{

		"status":  "success",
		"message": "user updated successfully",
	})
}

func DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userId, _ := strconv.Atoi(id)
	err := DbFunctions.DeleteTheUser(userId)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error in deleting user",
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "user deleted successfully",
	})
}

func SendVerifyCode(data Models.VerificationCode) error {
	messenger := mail.NewMessage()
	email:=os.Getenv("EMAIL")
	password:=os.Getenv("PASSWORD")
	messenger.SetHeader("From", email)
	messenger.SetHeader("To", data.UserEmail)
	messenger.SetHeader("Subject", "hi from Donation App")
	messenger.SetBody("text/plain", "your verification code is "+data.VerificationCode)
	a := mail.NewDialer("smtp.gmail.com", 587, email, password)
	if err := a.DialAndSend(messenger); err != nil {
		fmt.Println("error ", err)
		panic(err)
	}
	return nil
}

func Login(ctx *fiber.Ctx) error {
	var userRequest Models.User
	err := ctx.BodyParser(&userRequest)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error in parsing user",
		})

	}
	user, err := DbFunctions.RetrieveUserByEmail(userRequest.Email)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "error in finding user",
		})
	}
	if user.Password != userRequest.Password {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "password or email is incorrect",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "user logged in successfully",
		"user":    user,
	})
}
