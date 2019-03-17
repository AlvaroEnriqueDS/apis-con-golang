package controllers

import (
        "encoding/json"
        "fmt"
        "github.com/alvaroenriqueds/dinamoPruebas/configuration"
        "github.com/alvaroenriqueds/dinamoPruebas/models"
        "github.com/labstack/echo"
        "net/http"
)

func CommentCreate(c echo.Context) error  {
        comment := models.Comment{}
        //user := models.User{}

        err := json.NewDecoder(c.Request().Body).Decode(&comment)
        if err != nil {
                fmt.Printf("Error al leer el comentario a registrar: %s", err)
                return c.NoContent(http.StatusBadRequest)
        }

        //se abre conexion con la base de datos
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //se inserta el usuario
        q := "insert into comments (user_id, parent_id, votes, content) values ($1, $2, $3, $4);"

        stmt, err := db.Prepare(q)
        if err != nil {
                fmt.Printf("Error al crear el registro: %s", err)
                return c.NoContent(http.StatusBadRequest)
        }

        stmt.QueryRow(comment.UserID, comment.ParentId, comment.Votes, comment.Content)
        //err = row.Scan(&comment.Id)
        //if err != nil {
        //        fmt.Printf("Error al scanear el registro: %s", err)
        //        return c.NoContent(http.StatusBadRequest)
        //}

        return c.NoContent(http.StatusCreated)
}
