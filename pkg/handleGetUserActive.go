package pkg

import (
	"errors"
	"server/config"
	"server/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func GetUserActive(c *fiber.Ctx) (idACtiveUser string, err error) {

	var idUser string
	// * Check if cookies valid
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return "", errors.New("unknown id")
	}

	// * Get values of Cookies
	claims := token.Claims.(*jwt.StandardClaims)
	var pengurus models.Pengurus
	var donatur models.Donatur

	if err := config.DB.Where("pengurus_id = ?", claims.Issuer).First(&pengurus).Error; err == nil {
		// * Check if data exist
		pengurus := models.Pengurus{PengurusId: pengurus.PengurusId}
		if err := config.DB.First(&pengurus).Error; err == nil {
			idUser = pengurus.PengurusId
			return idUser, nil
		}
	}

	if err := config.DB.Where("donatur_id = ?", claims.Issuer).First(&donatur).Error; err == nil {
		// * Check if data exist
		donatur := models.Donatur{DonaturId: donatur.DonaturId}
		if err := config.DB.First(&donatur).Error; err == nil {
			idUser = donatur.DonaturId
			return idUser, nil
		}
	}
	return "PA-1", nil
}
