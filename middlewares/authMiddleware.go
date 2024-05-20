package middlewares

import (
	"server/config"
	"server/models"
	"server/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(allowedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// * Check if user already login
		cookie := c.Cookies("jwt")
		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(pkg.SECRET_KEY), nil
		})
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message": "Tolong login terlebih dahulu.",
			})
		}

		if len(allowedRoles) == 0 {
			return c.Next()
		}
		// * Check role user from token
		claims := token.Claims.(*jwt.StandardClaims)
		var userRole string
		roleAllowed := false
		if err := config.DB.Model(&models.Pengurus{}).Where("pengurus_id = ?", claims.Issuer).Pluck("jabatan", &userRole).Error; err == nil {
			for _, allowedRole := range allowedRoles {
				if userRole == allowedRole {
					roleAllowed = true
					break
				}
			}
		}
		if err := config.DB.Model(&models.Donatur{}).Where("donatur_id = ?", claims.Issuer).Pluck("validasi", &userRole).Error; err == nil {
			for _, allowedRole := range allowedRoles {
				if userRole == allowedRole {
					roleAllowed = true
					break
				}
			}
		}
		if !roleAllowed {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Anda tidak diizinkan untuk melakukan akses pada menu ini.",
			})
		}
		return c.Next()
	}
}
