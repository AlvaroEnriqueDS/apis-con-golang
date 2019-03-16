package controllers

import (
        "crypto/sha256"
        "encoding/json"
        "fmt"
        "github.com/alvaroenriqueds/dinamoPruebas/configuration"
        "github.com/alvaroenriqueds/dinamoPruebas/models"
        "github.com/labstack/echo"
        "net/http"
)

func CreateUser( c echo.Context) error  {
        user := models.User{}

        //se lee el json entrante y vuelca en el modelo user
        err := json.NewDecoder(c.Request().Body).Decode(&user)
        if err != nil {
                fmt.Sprintf("Error al leer el usuario a registrar: %s", err)
                return c.NoContent(http.StatusBadRequest)
        }

        //se confirma que las contraseñas seas iguales
        if user.Password != user.ConfirmPassword {
                fmt.Sprintf("Las contraseñas no coinciden: %s1 | %s2", user.Password, models.User{}.ConfirmPassword)
                return c.NoContent(http.StatusBadRequest)
        }

        //se codifica la contraseña en sha256
        pass := sha256.Sum256([]byte(user.Password))
        pwd := fmt.Sprintf("%x", pass)
        user.Password = pwd

        //agregar validacion para la imagen del usuario

        //se abre conexion con la base de datos
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //se inserta el usuario
        q := "insert into cliente (correoelectronico, nombre, apellido, contrasena) values ($1,$2,$3,$4) RETURNING idcliente;"

        stmt, err := db.Prepare(q)
        if err != nil {
                fmt.Printf("Error al crear el registro: %s", err)
                return c.NoContent(http.StatusBadRequest)
        }

        stmt.QueryRow(user.Password)

        return c.NoContent(http.StatusCreated)
}
