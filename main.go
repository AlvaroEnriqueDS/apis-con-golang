package main

import (
        "github.com/alvaroenriqueds/apis-con-golang/controllers"
        "github.com/labstack/echo"
        "github.com/labstack/echo/middleware"
        //"gopkg.in/olahol/melody.v1"
)

func main()  {
        e := echo.New()
        e.Use(middleware.Logger())
        e.Use(middleware.Recover())
        e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
                AllowOrigins:     []string{"*"},
                AllowMethods:     []string{echo.GET, echo.POST, echo.DELETE, echo.PUT},
        }))
        //mel = melody.New()
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
        socketRoute(e)
        /*
        e.GET("/ws", func(c echo.Context) error {
                m.HandleRequest(c.Response().Writer, c.Request())
                return nil
        })
        m.HandleMessage(func(s *melody.Session, msg []byte) {
                m.Broadcast(msg)
        })
        */
        e.Start(":2020")

}

func socketRoute(e *echo.Echo) {
        e.GET("/ws", controllers.WebSockets)
        e.GET("/wsi", controllers.WebSockets)

}


