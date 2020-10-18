package middleware

import (
	"github.com/gin-gonic/gin"
)

// Messages -
var Messages = map[string]string{
	"auth-not-found":    "Credenciais estão incorretas",
	"unexpected-verify": "Erro ao verificar jwt.",
	"id-not-found":      "Não foi possível identificar o solicitante",
}

// AuthRequired -
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var (
		// auth  = c.GetHeader("Authorization")
		// token string
		// )

		// if !strings.Contains(auth, "Bearer ") || auth == "" {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, router.Response{Status: "error", Message: Messages["auth-not-found"], Error: "auth-not-found"}) // TODO add router.Reponse
		// 	return
		// }

		// token = strings.Replace(auth, "Bearer ", "", 1)
		// if token == "" {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, router.Response{Status: "error", Message: Messages["auth-not-found"], Error: "auth-not-found"})
		// 	return
		// }

		// check, claims, err := jwt.Verify(token, true)
		// if !check {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, router.Response{Status: "error", Message: Messages["auth-not-found"], Error: "auth-not-found"})
		// 	return
		// }

		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, router.Response{Status: "error", Message: Messages["unexpected-verify"], Error: "unexpected-verify"})
		// 	return
		// }

		// operatorID := fmt.Sprintf("%v", claims["sub"])
		// if operatorID == "" {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, router.Response{Status: "error", Message: Messages["id-not-found"], Error: "id-not-found"})
		// 	return
		// }
		// c.Set("operator_id", operatorID)

		id := c.Param("id")
		c.Set("id", id)

		limit := c.Query("limit")
		if limit != "" {
			c.Set("limit", limit)
		}

		offset := c.Query("offset")
		if offset != "" {
			c.Set("offset", offset)
		}

		lat := c.Query("lat")
		if offset != "" {
			c.Set("lat", lat)
		}

		long := c.Query("long")
		if offset != "" {
			c.Set("long", long)
		}

		city := c.Query("city")
		if offset != "" {
			c.Set("city", city)
		}

		key := c.Query("key")
		if offset != "" {
			c.Set("key", key)
		}
		c.Next()
	}
}
