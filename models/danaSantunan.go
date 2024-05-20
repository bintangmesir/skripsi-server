package models

import (
	"time"
)

type DanaSantunan struct {
	DanaSantunanId string         `gorm:"primaryKey;unique;size:50" json:"dana_santunan_id"`
	Index          *int           `json:"index"`
	Tanggal        time.Time      `gorm:"default:CURRENT_TIMESTAMP()" json:"tanggal"`
	Nama           string         `gorm:"size:50;default:'Hamba Allah'" json:"nama"`
	Nominal        int            `json:"nominal"`
	Tipe           PembayaranEnum `gorm:"size:6;default:DEBIT" json:"tipe"`
	File           *string        `gorm:"size:255" json:"file,omitempty"`
	Validasi       ValidationEnum `gorm:"default:PENDING;size:15" json:"validasi"`
	Keterangan     *string        `gorm:"size:255" json:"keterangan"`
	CreatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP()" json:"updated_at"`

	PengurusId *string  `gorm:"size:10" json:"pengurus_id"`
	Pengurus   Pengurus `json:"pengurus"`

	LaporanDanaSantunanId *string             `gorm:"size:10" json:"laporan_dana_santunan_id"`
	LaporanDanaSantunan   LaporanDanaSantunan `json:"laporan_dana_santunan"`
}
