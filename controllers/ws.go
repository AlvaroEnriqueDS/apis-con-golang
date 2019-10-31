package controllers

import (
        "encoding/json"
        "github.com/labstack/echo"
        "github.com/olahol/melody"
        "log"
)

func init()  {
        mel = melody.New()
}


var mel *melody.Melody
func WebSockets(c echo.Context) error {

        mel.HandleRequest(c.Response().Writer, c.Request())
        //mel.HandleConnect(hConnect)
        //mel.HandleDisconnect(hDisconnect)
        mel.HandleMessage(hMessage)
        return nil
}
func hConnect(s *melody.Session) {

        sendMessage("conectado")
}

func hDisconnect(s *melody.Session) {

        sendMessage("desconectado")
}
func hMessage(s *melody.Session, msg []byte) {

        mel.Broadcast(msg)
}

func sendMessage(m string) {
        j, err := json.Marshal(m)
        if err != nil {
                log.Printf("no se pudo convertir el mensaje a json: %v", err)
                return
        }

        mel.Broadcast(j)
}

