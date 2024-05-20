package midtranscontroller

import (
	"os"
	"server/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func MidtransDanaSantunanAnakAsuhGetById(c *fiber.Ctx) error {

	orderId := c.Params("id")

	r := coreapi.Client{}
	if os.Getenv("APP_ENV") == "PRODUCTION" {
		r.New(pkg.MIDTRANS_SECRET_KEY, midtrans.Production)
	} else {
		r.New(pkg.MIDTRANS_SECRET_KEY, midtrans.Sandbox)
	}

	res, _ := r.CheckTransaction(orderId)
	if res != nil {
		newMidtransJSON := map[string]interface{}{
			"order_id":     res.OrderID,
			"gross_amount": res.GrossAmount,
			"status_code":  res.StatusCode,
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"data": newMidtransJSON,
		})
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Data transaksi tidak ditemukan.",
	})
}
