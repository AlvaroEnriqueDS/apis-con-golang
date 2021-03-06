package main

import (
        "github.com/alvaroenriqueds/apis-con-golang/controllers"
        "github.com/labstack/echo"
)

func main()  {
        e := echo.New()

        /*
        server, err := socketio.NewServer(nil)
        if err != nil {
                log.Fatal(err)
        }
        */

        e.Static("/", "public")

        e.POST("/register", controllers.CreateUser)
        e.POST("/login", controllers.LoginUser)
        e.POST("/comment", controllers.CommentCreate)
        e.GET("/comment", controllers.CommentGetAll)
        e.POST("/upload", controllers.Upload)
        e.GET("/validate", controllers.ValidateToken)
        e.Start(":80")
}
