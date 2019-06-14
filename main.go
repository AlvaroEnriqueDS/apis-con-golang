package main

import (
        "github.com/alvaroenriqueds/apis-con-golang/controllers"
        "github.com/labstack/echo"
        "github.com/labstack/echo/middleware"
        "github.com/olahol/melody"
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
        mel = melody.New()
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
        e.GET("/ws", webSockets)
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

var mel *melody.Melody


func webSockets(c echo.Context) error  {
        mel.HandleRequest(c.Response().Writer, c.Request())
        //mel.HandleConnect(hConnect)
        //mel.HandleDisconnect(hDisconnect)
        mel.HandleMessage(hMessage)
        return nil
}
func hMessage(s *melody.Session, msg []byte) {
        mel.Broadcast(msg)
}


