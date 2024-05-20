package midtranscontroller

import (
	"fmt"
	"os"
	danasantunananakasuh "server/controllers/dana-santunan-anak-asuh"
	"server/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func MidtransDanaSantunanAnakAsuhCreate(c *fiber.Ctx) error {

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
	danaSantunanAnakAsuhId, err := danasantunananakasuh.HandleDanaSantunanAnakAsuhId()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Mohon maaf terjadi kesalahan pada server.",
		})
	}

	fmt.Println(danaSantunanAnakAsuhId)

	midtrans.ServerKey = pkg.MIDTRANS_SECRET_KEY
	if os.Getenv("APP_ENV") == "PRODUCTION" {
		midtrans.Environment = midtrans.Production
	} else {
		midtrans.Environment = midtrans.Sandbox
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  danaSantunanAnakAsuhId,
			GrossAmt: int64(nominalConverted),
		},
	}

	snapResp, _ := snap.CreateTransaction(req)

	newMidtransJSON := map[string]interface{}{
		"message":                    "Sedang menunggu proses pembayaran, mohon untuk tidak menutup halaman ini",
		"snap":                       snapResp,
		"dana_santunan_anak_asuh_id": danaSantunanAnakAsuhId,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": newMidtransJSON,
	})
}
