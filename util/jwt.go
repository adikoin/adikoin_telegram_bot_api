package util

import (
	"telegram_bot_api/config"
	model "telegram_bot_api/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GenerateJwtToken(user *model.User) (string, error) {

	// expTimeMs, _ := strconv.Atoi(config.JWTExpirationMs)
	// exp := time.Now().Add(time.Millisecond * time.Duration(expTimeMs)).Unix()
	name := user.FirstName
	surName := user.LastName
	// role := user.Role
	// companyID := user.CompanyID
	// companyCategory := user.CompanyCategory

	// // Set custom claims
	// claims := &model.JwtCustomClaims{
	// 	user.ID.Hex(),
	// 	name,
	// 	surname,
	// 	jwt.StandardClaims{
	// 		ExpiresAt: exp,
	// 	},
	// }

	claims := &model.JwtCustomClaims{
		ID:      user.ID.Hex(),
		Name:    name,
		SurName: surName,
		// Role:            role,
		// CompanyID:       companyID.Hex(),
		// CompanyCategory: companyCategory,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return "", err
	}

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"id":              user.ID.Hex(),
	// 	"name":            name,
	// 	"surname":         surname,
	// 	"role":            role,
	// 	"companyID":       companyID,
	// 	"companyCategory": companyCategory,
	// 	"ExpiresAt":       time.Now().Add(time.Millisecond * time.Duration(expTimeMs)).Unix(),
	// })

	// Create token with claims
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	// jwt, err := token.SignedString([]byte(config.JWTSecret))
	// log.Println("jwt is:", jwt)
	return jwt, err
}

func GetUserIdFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	return claims.ID
}

// func GetUserRoleFromToken(c echo.Context) string {

// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(*model.JwtCustomClaims)
// 	print(claims)
// 	return claims.Role
// }

// func GetUserRoleFromToken(c echo.Context) string {
// 	bearerToken := c.Request().Header.Get("Authorization")
// 	// log.Println("jwt is:", c.Request().Header.Get("Access-Control-Request-Headers"))

// 	jwtString := strings.Split(bearerToken, "Bearer ")[1]

// 	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(config.JWTSecret), nil
// 	})

// 	if err != nil {
// 		return err.Error()
// 	}

// 	role := token.Claims.(jwt.MapClaims)["role"].(string)

// 	return role
// }
