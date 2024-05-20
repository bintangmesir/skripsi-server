package models

import "time"

type Pengurus struct {
	PengurusId  string    `gorm:"primaryKey;unique;size:10" json:"pengurus_id"`
	Nama        string    `gorm:"size:50;not null" json:"nama"`
	Email       string    `gorm:"size:50;unique;not null" json:"email"`
	NoHandphone string    `gorm:"size:20" json:"no_handphone"`
	Password    string    `gorm:"size:72;not null" json:"-"`
	Jabatan     RoleEnum  `gorm:"default:ADMIN;size:15" json:"jabatan"`
	Foto        *string   `gorm:"size:255" json:"foto,omitempty"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"updated_at"`

	AdminId              *string                `json:"admin_id"`
	Pengurus             []Pengurus             `gorm:"foreignKey:AdminId" json:"pengurus"`
	Donatur              []Donatur              `gorm:"foreignKey:PengurusId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"donatur"`
	DanaSantunan         []DanaSantunan         `gorm:"foreignKey:PengurusId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"dana_santunan"`
	LaporanDanaSantunan  []LaporanDanaSantunan  `gorm:"foreignKey:PengurusId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"laporan_dana_santunan"`
	DanaSantunanAnakAsuh []DanaSantunanAnakAsuh `gorm:"foreignKey:PengurusId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"dana_santunan_anak_asuh"`
	AnakYatim            []AnakYatim            `gorm:"foreignKey:PengurusId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"anak_yatim"`
}
