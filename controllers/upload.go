package controllers

import (
        "fmt"
        "github.com/labstack/echo"
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

        err = ioutil.WriteFile("./files/"+handle.Filename, data, 0666)
        if err != nil {
                log.Printf("Error al escribir el archivo %v", err)
                fmt.Fprintf(c.Response(), "Error al escribir el archivo %v", err)
                return nil
        }

        return c.NoContent(http.StatusOK)
}
