package models

import (
	"time"
)

type LaporanDanaSantunan struct {
	LaporanDanaSantunanId string         `gorm:"primaryKey;unique;size:10" json:"laporan_dana_santunan_id"`
	Judul                 string         `gorm:"size:50" json:"judul"`
	Keterangan            *string        `gorm:"size:255;not null" json:"keterangan"`
	SaldoAwal             int            `json:"saldo_awal"`
	SaldoSisa             int            `json:"saldo_sisa"`
	TandaTangan           *string        `gorm:"size:255" json:"tanda_tangan,omitempty"`
	NamaTandaTangan       *string        `gorm:"size:50" json:"nama_tanda_tangan"`
	TanggalTandaTangan    time.Time      `gorm:"default:CURRENT_TIMESTAMP()" json:"tanggal_tanda_tangan"`
	File                  *string        `gorm:"size:255" json:"file,omitempty"`
	Validasi              ValidationEnum `gorm:"default:PENDING;size:15" json:"validasi"`
	CreatedAt             time.Time      `gorm:"default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt             time.Time      `gorm:"default:CURRENT_TIMESTAMP()" json:"updated_at"`

	PengurusId   *string        `gorm:"size:10" json:"pengurus_id"`
	Pengurus     Pengurus       `json:"pengurus"`
	DanaSantunan []DanaSantunan `gorm:"foreignKey:LaporanDanaSantunanId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"dana_santunan"`
}
