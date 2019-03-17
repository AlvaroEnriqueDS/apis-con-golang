package controllers

import (
        "encoding/json"
        "fmt"
        "github.com/alvaroenriqueds/dinamoPruebas/configuration"
        "github.com/alvaroenriqueds/dinamoPruebas/models"
        "github.com/labstack/echo"
        "net/http"
)

//CommentCreate funcion para crear comentarios
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

        //se inserta el comentario
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

//CommentGetAll llama a todos los comentarios PADRES
func CommentGetAll(c echo.Context) error  {
        //comments := []models.Comment{}
        comments := make([]models.Comment, 0)

        row := models.Comment{}
        //user := models.User{}

        //c.Request().Context().Value(&user)
        //vars := c.Request().URL.Query() // lee la URL que llega ->
        // /api/comments/?order=votes&idlimit=10

        //se abre conexion con la base de datos
        db := configuration.GetConnectionPsql()
        defer db.Close()

        //esto solo nos traera los comentarios padres
        q := "select id, user_id, parent_id, votes, content from comments where parent_id = 0;"

        stmt, err := db.Prepare(q)
        if err != nil {
                fmt.Printf("Error al crear el registro: %s", err)
                return c.NoContent(http.StatusBadRequest)
        }
        rows, err := stmt.Query()
        if err != nil {
                fmt.Printf("Error ejecutar el registro: %s", err)
                return c.NoContent(http.StatusBadRequest)
        }

        for rows.Next() {
                err := rows.Scan(
                        &row.Id,
                        &row.UserID,
                        &row.ParentId,
                        &row.Votes,
                        &row.Content,
                )
                if err != nil {
                        fmt.Println("3")
                        return err
                }
                /*
                ctm := models.Comment{
                        Id: row.Id,
                        UserID: row.UserID,
                        ParentId: row.ParentId,
                        Votes: row.Votes,
                        Content: row.Content,
                }
                */
                comments = append(comments, row)
        }

        return c.JSON(http.StatusOK, comments)
}
