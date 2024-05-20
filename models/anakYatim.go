package models

import "time"

type AnakYatim struct {
	AnakYatimId         string     `gorm:"primaryKey;unique;size:10" json:"anak_yatim_id"`
	Nama                string     `gorm:"size:50;not null" json:"nama"`
	Status              string     `gorm:"size:15;not null" json:"status"`
	TanggalLahir        time.Time  `gorm:"default:CURRENT_TIMESTAMP()" json:"tanggal_lahir"`
	JenisKelamin        string     `gorm:"size:2;not null" json:"jenis_kelamin"`
	Pendidikan          string     `gorm:"size:10;not null" json:"pendidikan"`
	PenghasilanOrangTua int        `gorm:"not null" json:"penghasilan_orang_tua"`
	TanggunganOrangTua  int        `gorm:"size:2;not null" json:"tanggungan_orang_tua"`
	PekerjaanOrangTua   string     `gorm:"size:50;not null" json:"pekerjaan_orang_tua"`
	Kebutuhan           string     `gorm:"not null" json:"kebutuhan"`
	Deskripsi           string     `gorm:"not null" json:"deskripsi"`
	StatusSantunan      StatusEnum `gorm:"default:BELUM_MEMILIKI;size:20;not null" json:"status_santunan"`
	NominalSantunan     int        `gorm:"not null" json:"nominal_santunan"`
	Foto                *string    `gorm:"size:255" json:"foto,omitempty"`
	CreatedAt           time.Time  `gorm:"default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt           time.Time  `gorm:"default:CURRENT_TIMESTAMP()" json:"updated_at"`

	PengurusId *string  `gorm:"size:10" json:"pengurus_id"`
	Pengurus   Pengurus `json:"pengurus"`

	DonaturId *string `gorm:"size:10" json:"donatur_id"`
	Donatur   Donatur `json:"donatur"`

	DanaSantunanAnakAsuh []DanaSantunanAnakAsuh `gorm:"foreignKey:AnakYatimId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"dana_santunan_anak_asuh"`
}
