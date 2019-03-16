package main

import (
        "github.com/alvaroenriqueds/dinamoPruebas/controllers"
        "github.com/labstack/echo"
)

func main()  {
        e := echo.New()

        e.POST("/register", controllers.CreateUser)
        e.Start(":80")
}
