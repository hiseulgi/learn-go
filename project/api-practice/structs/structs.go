package structs

import "github.com/jinzhu/gorm"

type Mahasiswa struct {
	gorm.Model
	Nama    string
	NIM     string
	Kelas   string
	Jurusan string
	Prodi   string
	IPK     float64 `gorm:"column:ipk"`
}
