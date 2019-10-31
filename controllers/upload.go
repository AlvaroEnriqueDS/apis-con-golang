package controllers

import (
        "encoding/json"
        "fmt"
        "github.com/labstack/echo"
        "golang.org/x/net/websocket"
        "io/ioutil"
        "log"
        "net/http"
)

func Upload(c echo.Context) error  {
        file, handle, err := c.Request().FormFile("myFile")
        if err != nil {
                log.Printf("Error al cargar el archivo %v", err)
                return nil
        }
        defer file.Close()

        data, err := ioutil.ReadAll(file)
        if err != nil {
                log.Printf("Error al leer el archivo %v", err)
                fmt.Fprintf(c.Response(), "Error al leer el archivo %v", err)
                return nil
        }

        err = ioutil.WriteFile("./public/files/"+handle.Filename, data, 0666)
        if err != nil {
                log.Printf("Error al escribir el archivo %v", err)
                fmt.Fprintf(c.Response(), "Error al escribir el archivo %v", err)
                return nil
        }

        port := 2020
        origin := fmt.Sprintf("http://localhost:%d/", port)
        url := fmt.Sprintf("ws://localhost:%d/wsi", port)
        wsi, err := websocket.Dial(url, "", origin)
        if err != nil {
                log.Fatal(err)
        }


        nombre := "./files/" + handle.Filename
        j, err := json.Marshal(&nombre)
        if _, err := wsi.Write(j); err != nil {
                log.Fatal(err)
        }

        fmt.Println(nombre)
        return c.NoContent(http.StatusOK)
}
