package controllers

import (
        "fmt"
        "github.com/alvaroenriqueds/dinamoPruebas/commons"
        "github.com/alvaroenriqueds/dinamoPruebas/models"
        "github.com/dgrijalva/jwt-go"
        "github.com/dgrijalva/jwt-go/request"
        "github.com/labstack/echo"
        "net/http"
)

func ValidateToken(c echo.Context) error  {
        token, err := request.ParseFromRequestWithClaims(
                c.Request(),
                request.OAuth2Extractor,
                &models.Claim{},
                func(token *jwt.Token) (interface{}, error) {
                        return commons.PublicKey, nil
                },
        )
        if err != nil {
                switch err.(type) {
                case *jwt.ValidationError:
                        vError := err.(*jwt.ValidationError)
                        switch vError.Errors {
                        case jwt.ValidationErrorExpired:
                                fmt.Fprintln(c.Response(), "Su token ha expirado")
                                return c.NoContent(http.StatusNoContent)
                        case jwt.ValidationErrorSignatureInvalid:
                                fmt.Fprintln(c.Response(), "La firma del token no coincide")
                                return c.NoContent(http.StatusNoContent)
                        default:
                                fmt.Fprintln(c.Response(), "Su token no es valido")
                                return c.NoContent(http.StatusNoContent)
                        }

                }
        }

        if token.Valid {
                fmt.Fprintln(c.Response(), "Bienvenido al sistema")
                return c.String(http.StatusOK, "Si ingreso")
        }else {
                return c.String(http.StatusUnauthorized, "Su token no es validp")
        }

        /*
        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(*models.Claim)
        name := claims.User.Fullname

        return c.String(http.StatusOK, "Bienvenido "+name)
        */
}
