package midtranscontroller

import (
	"os"
	danasantunan "server/controllers/dana-santunan"
	"server/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func MidtransDanaSantunanCreate(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	nominal := form.Value["nominal"][0]
	nominalConverted, err := strconv.Atoi(nominal)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	//* Handle order id
	danaSantunanId, err := danasantunan.HandleDanaSantunanId()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	midtrans.ServerKey = pkg.MIDTRANS_SECRET_KEY
	if os.Getenv("APP_ENV") == "PRODUCTION" {
		midtrans.Environment = midtrans.Production
	} else {
		midtrans.Environment = midtrans.Sandbox
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  danaSantunanId,
			GrossAmt: int64(nominalConverted),
		},
	}

	snapResp, _ := snap.CreateTransaction(req)

	newMidtransJSON := map[string]interface{}{
		"message":          "Sedang menunggu proses pembayaran, mohon untuk tidak menutup halaman ini",
		"snap":             snapResp,
		"dana_santunan_id": danaSantunanId,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": newMidtransJSON,
	})
}
