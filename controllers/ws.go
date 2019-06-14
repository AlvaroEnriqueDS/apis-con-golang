package controllers

import (
        "github.com/labstack/echo"
        "github.com/olahol/melody"
)

var mel *melody.Melody
//mel = melody.New()

func WebSockets(c echo.Context) error  {
        mel = melody.New()
        mel.HandleRequest(c.Response().Writer, c.Request())
        //mel.HandleConnect(hConnect)
        //mel.HandleDisconnect(hDisconnect)
        mel.HandleMessage(hMessage)
        return nil
}
func hMessage(s *melody.Session, msg []byte) {
        mel.Broadcast(msg)
}
