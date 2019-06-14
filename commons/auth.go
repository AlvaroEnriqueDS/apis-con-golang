package commons

import (
        "crypto/rsa"
        "github.com/alvaroenriqueds/apis-con-golang/models"
        "github.com/dgrijalva/jwt-go"
        "io/ioutil"
        "log"
        "time"
)

var (
        privateKey *rsa.PrivateKey
        PublicKey *rsa.PublicKey
)

func init()  {
        privateBytes, err := ioutil.ReadFile("./keys/private.rsa")
        if err != nil {
                log.Fatal("No se pudo leer el archivo privado")
        }

        publicBytes, err := ioutil.ReadFile("./keys/public.rsa")
        if err != nil {
                log.Fatal("No se pudo leer el archivo publico")
        }

        privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
        if err != nil {
                log.Fatal("No se pudo hacer el parse a privateKey")
        }

        PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
        if err != nil {
                log.Fatal("No se pudo hacer el parse a publicKey")
        }
}

func GenerateJWT(user models.User) string {
        claims := models.Claim{
                User: user,
                StandardClaims: jwt.StandardClaims{
                        ExpiresAt: time.Now().Add(time.Hour*1).Unix(),
                        //objetivo
                        Issuer: "Pruebas con JWT",
                },
        }

        //convirtiendo el claim a token y luego a string
        token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
        result, err := token.SignedString(privateKey)
        if err != nil {
                log.Fatal("No se pudo firmar el token")
        }

        return result
}


