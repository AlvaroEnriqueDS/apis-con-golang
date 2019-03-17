package main

import (
        "github.com/alvaroenriqueds/dinamoPruebas/controllers"
        "github.com/labstack/echo"
)

func main()  {
        e := echo.New()

        e.POST("/register", controllers.CreateUser)
        e.POST("/login", controllers.LoginUser)
        e.POST("/comment", controllers.CommentCreate)
        e.GET("/comment", controllers.CommentGetAll)
        e.Start(":80")
}
