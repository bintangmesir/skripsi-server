package models

import "time"

type Donatur struct {
	DonaturId    string    `gorm:"primaryKey;unique;size:10" json:"donatur_id"`
	Nama         string    `gorm:"size:50;not null" json:"nama"`
	Email        string    `gorm:"size:50;unique;not null" json:"email"`
	Password     string    `gorm:"size:72;not null" json:"-"`
	JenisKelamin string    `gorm:"size:2;not null" json:"jenis_kelamin"`
	NoHandphone  string    `gorm:"size:20;not null" json:"no_handphone"`
	Foto         *string   `gorm:"size:255" json:"foto,omitempty"`
	Validasi     RoleEnum  `gorm:"default:DONATUR;size:15" json:"validasi"`
	NamaJalan    string    `gorm:"size:50" json:"nama_jalan"`
	Rt           string    `gorm:"size:5" json:"rt"`
	Rw           string    `gorm:"size:5" json:"rw"`
	Kelurahan    string    `gorm:"size:50" json:"kelurahan"`
	Kecamatan    string    `gorm:"size:50" json:"kecamatan"`
	Kota         string    `gorm:"size:50" json:"kota"`
	Provinsi     string    `gorm:"size:50" json:"provinsi"`
	KodePos      string    `gorm:"size:50" json:"kode_pos"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"updated_at"`

	PengurusId *string `gorm:"size:10" json:"pengurus_id"`

	AnakYatim            []AnakYatim            `gorm:"foreignKey:DonaturId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"anak_yatim"`
	DanaSantunanAnakAsuh []DanaSantunanAnakAsuh `gorm:"foreignKey:DonaturId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"dana_santunan_anak_asuh"`
}
