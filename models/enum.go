package models

type RoleEnum string

const (
	AdminRole        RoleEnum = "ADMIN"
	KetuaDkmRole     RoleEnum = "KETUA_DKM"
	BendaharaRole    RoleEnum = "BENDAHARA"
	HumasRole        RoleEnum = "HUMAS"
	DonaturRole      RoleEnum = "DONATUR"
	OrangTuaAsuhRole RoleEnum = "ORANG_TUA_ASUH"
)

type ValidationEnum string

const (
	Diverifikasi ValidationEnum = "DIVERIFIKASI"
	Pending      ValidationEnum = "PENDING"
	Diterima     ValidationEnum = "DITERIMA"
	Ditolak      ValidationEnum = "DITOLAK"
)

type PembayaranEnum string

const (
	Debit  StatusEnum = "DEBIT"
	Kredit StatusEnum = "KREDIT"
)

type StatusEnum string

const (
	BelumMemiliki StatusEnum = "BELUM_MEMILIKI"
	SudahMemiliki StatusEnum = "SUDAH_MEMILIKI"
)
