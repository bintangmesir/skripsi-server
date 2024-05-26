package models

import (
	"time"
)

type DanaSantunanAnakAsuh struct {
	DanaSantunanAnakAsuhId string         `gorm:"primaryKey;unique;size:50" json:"dana_santunan_anak_asuh_id"`
	Index                  *int           `json:"index"`
	Tanggal                time.Time      `gorm:"default:CURRENT_TIMESTAMP()" json:"tanggal"`
	Nominal                int            `json:"nominal"`
	Tipe                   PembayaranEnum `gorm:"size:6;default:DEBIT" json:"tipe"`
	File                   *string        `gorm:"size:255" json:"file,omitempty"`
	Validasi               ValidationEnum `gorm:"default:PENDING;size:15" json:"validasi"`
	Keterangan             *string        `gorm:"size:255" json:"keterangan"`
	CreatedAt              time.Time      `gorm:"default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt              time.Time      `gorm:"default:CURRENT_TIMESTAMP()" json:"updated_at"`

	PengurusId *string  `gorm:"size:10" json:"pengurus_id"`
	Pengurus   Pengurus `json:"pengurus"`

	DonaturId *string `gorm:"size:10" json:"donatur_id"`
	Donatur   Donatur `json:"donatur"`

	AnakYatimId *string   `gorm:"size:10" json:"anak_yatim_id"`
	AnakYatim   AnakYatim `json:"anak_yatim"`

	LaporanDanaSantunanAnakAsuhId *string                     `gorm:"size:10" json:"laporan_dana_santunan_anak_asuh_id"`
	LaporanDanaSantunanAnakAsuh   LaporanDanaSantunanAnakAsuh `json:"laporan_dana_santunan_anak_asuh"`
}
