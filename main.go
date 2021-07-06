package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	// input := user.LoginInput{
	// 	Email:    "kunti@gmail.com",
	// 	Password: "dsdsd",
	// }

	// user, err := userService.Login(input)
	// if err != nil {
	// 	fmt.Println("Login gagal")
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(user.Email)
	// fmt.Println(user.Name)

	// userByEmail, err := userRepository.FindByEmail("kunti@gmail.com")

	// if userByEmail.ID == 0 {
	// 	fmt.Println("User tidak ditemukan")
	// } else {
	// 	fmt.Println(userByEmail.Name)
	// }

	userHandler := handler.NewUserHandler(userService)
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run()

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Tes simpan dari service"
	// userInput.Email = "contoh@gmaail.com"
	// userInput.Occupation = "anak band"
	// userInput.Password = "password"

	// userService.RegisterUser(userInput)

	// fmt.Println("Connected to Database successfully")

	// var users []user.User

	// db.Find(&users)

	// length := len(users)

	// fmt.Println(length)

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// }

	// router := gin.Default()
	// router.GET("/handler", handler)
	// router.Run()

}
