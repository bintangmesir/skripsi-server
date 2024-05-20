package pkg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var URI_HOST string
var URI_PORT string

var DB_USER string
var DB_PASSWORD string
var DB_URI string
var DB_NAME string

var DIR_PUBLIC string
var DIR_IMG_PENGURUS string
var DIR_IMG_DONATUR string
var DIR_IMG_ANAK_YATIM string
var DIR_IMG_TANDA_TANGAN string
var DIR_IMG_BUKTI_PENGGUNAAN string
var DIR_FILE_DANA_SANTUNAN string
var DIR_FILE_DANA_SANTUNAN_ANAK_ASUH string

var SECRET_KEY string
var MIDTRANS_SECRET_KEY string
var MIDTRANS_MERCHANT_ID string

func HandleEnv() {

	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error:", err)
	}

	if err := godotenv.Load(filepath.Join(cwd, ".env")); err != nil {
		panic("File .env tidak tersedia...")
	}

	if os.Getenv("APP_ENV") == "PRODUCTION" {
		URI_HOST = os.Getenv("URI_HOST_PRODUCTION")
		MIDTRANS_SECRET_KEY = os.Getenv("MIDTRANS_SERVER_KEY_PRODUCTION")
	} else {
		URI_HOST = os.Getenv("URI_HOST_DEVELOPMENT")
		MIDTRANS_SECRET_KEY = os.Getenv("MIDTRANS_SERVER_KEY_DEVELOPMENT")
	}

	URI_PORT = os.Getenv("URI_PORT")

	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_URI = os.Getenv("DB_URI")
	DB_NAME = os.Getenv("DB_NAME")

	DIR_PUBLIC = os.Getenv("DIR_PUBLIC")
	DIR_IMG_PENGURUS = os.Getenv("DIR_IMG_PENGURUS")
	DIR_IMG_DONATUR = os.Getenv("DIR_IMG_DONATUR")
	DIR_IMG_ANAK_YATIM = os.Getenv("DIR_IMG_ANAK_YATIM")
	DIR_IMG_TANDA_TANGAN = os.Getenv("DIR_IMG_TANDA_TANGAN")
	DIR_IMG_BUKTI_PENGGUNAAN = os.Getenv("DIR_IMG_BUKTI_PENGGUNAAN")
	DIR_FILE_DANA_SANTUNAN = os.Getenv("DIR_FILE_DANA_SANTUNAN")
	DIR_FILE_DANA_SANTUNAN_ANAK_ASUH = os.Getenv("DIR_FILE_DANA_SANTUNAN_ANAK_ASUH")
	SECRET_KEY = os.Getenv("SECRET_KEY")

	MIDTRANS_MERCHANT_ID = os.Getenv("MIDTRANS_MERCHANT_ID")
}
