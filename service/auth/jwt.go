package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yann-fk-21/todo-platform/config"
	"github.com/yann-fk-21/todo-platform/types"
	"github.com/yann-fk-21/todo-platform/utils"
)

type contextKey string
const UserKey contextKey = "userID"

func CreateJWT(secret []byte, userID int) (string, error) {
	expirationToken := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
    
    token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
        "userID": strconv.Itoa(userID),
        "expiratedAt": time.Now().Add(expirationToken).Unix(),
    })

    tokenString, err := token.SignedString(secret)
    if err != nil {
        return " ", err
    }

    return tokenString, nil
}

func WithJWTAuth(handleFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
    //   get the token from the user request
    tokenString := getTokenFromRequest(r)
    // validate JWT 
    t, err := validateToken(tokenString)
    if err != nil {
        log.Fatal(err)
        utils.WriteError(w, http.StatusUnauthorized, err)
        return 
    }

    if !t.Valid {
        log.Println("invalid token")
        utils.WriteError(w, http.StatusUnauthorized, err)
        return
    }

    claims := t.Claims.(jwt.MapClaims)
    str := claims["userID"].(string)

    userID, _ := strconv.Atoi(str)

    // Get userId
    
    ctx := r.Context()
    ctx = context.WithValue(ctx, UserKey, userID)
    r = r.WithContext(ctx)

    handleFunc(w, r)
    }  
}

func getTokenFromRequest(r *http.Request) string {
    tokenAuth := r.Header.Get("Authorization")

    if tokenAuth != " " {
        return tokenAuth
    }

    return " "
}

func validateToken(t string) (*jwt.Token, error){
   return jwt.Parse(t, func (t *jwt.Token) (interface{}, error) {
    if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
       return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
    }
    return []byte(config.Envs.JWTSecret), nil
   })
}