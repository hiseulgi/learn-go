package controllers

import (
	"api-practice/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// get mahasiswa data by id
func (idb *InDB) GetMahasiswa(c *gin.Context) {
	var (
		mahasiswa structs.Mahasiswa
		result    gin.H
	)

	id := c.Param("id")
	if err := idb.DB.Where("id = ?", id).First(&mahasiswa).Error; err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
		c.JSON(http.StatusNotFound, result)
	} else {
		result = gin.H{
			"result": mahasiswa,
			"count":  1,
		}
		c.JSON(http.StatusOK, result)
	}

}

// get all mahasiswa data
func (idb *InDB) GetMahasiswas(c *gin.Context) {
	var (
		mahasiswas []structs.Mahasiswa
		result     gin.H
	)

	idb.DB.Find(&mahasiswas)

	if len(mahasiswas) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": mahasiswas,
			"count":  len(mahasiswas),
		}
	}

	c.JSON(http.StatusOK, result)
}

// create new mahasiswa data
func (idb *InDB) CreateMahasiswa(c *gin.Context) {
	var (
		mahasiswa structs.Mahasiswa
		result    gin.H
	)
	nama := c.PostForm("nama")
	nim := c.PostForm("nim")
	kelas := c.PostForm("kelas")
	jurusan := c.PostForm("jurusan")
	prodi := c.PostForm("prodi")
	ipk := c.PostForm("ipk")

	mahasiswa.Nama = nama
	mahasiswa.NIM = nim
	mahasiswa.Kelas = kelas
	mahasiswa.Jurusan = jurusan
	mahasiswa.Prodi = prodi

	if len(ipk) != 0 {
		if ipkFloat, err := strconv.ParseFloat(ipk, 64); err != nil {
			result = gin.H{
				"result": err.Error(),
			}
			c.JSON(http.StatusInternalServerError, result)
			return
		} else {
			mahasiswa.IPK = ipkFloat
		}
	}

	idb.DB.Create(&mahasiswa)
	result = gin.H{
		"result": mahasiswa,
	}
	c.JSON(http.StatusOK, result)
}

// update mahasiswa data by id as Param
func (idb *InDB) UpdateMahasiswa(c *gin.Context) {
	var (
		mahasiswa    structs.Mahasiswa
		newMahasiswa structs.Mahasiswa
		result       gin.H
	)

	id := c.Param("id")

	if err := idb.DB.First(&mahasiswa, id).Error; err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusNotFound, result)
	}

	nama := c.PostForm("nama")
	nim := c.PostForm("nim")
	kelas := c.PostForm("kelas")
	jurusan := c.PostForm("jurusan")
	prodi := c.PostForm("prodi")
	ipk := c.PostForm("ipk")

	newMahasiswa.Nama = nama
	newMahasiswa.NIM = nim
	newMahasiswa.Kelas = kelas
	newMahasiswa.Jurusan = jurusan
	newMahasiswa.Prodi = prodi

	if len(ipk) != 0 {
		if ipkFloat, err := strconv.ParseFloat(ipk, 64); err != nil {
			result = gin.H{
				"result": err.Error(),
			}
			c.JSON(http.StatusInternalServerError, result)
			return
		} else {
			newMahasiswa.IPK = ipkFloat
		}
	}

	if err := idb.DB.Model(&mahasiswa).Updates(newMahasiswa).Error; err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	result = gin.H{
		"result": "succesfully update data!",
	}
	c.JSON(http.StatusOK, result)
}

// delete mahasiswa data by id
func (idb *InDB) DeleteMahasiswa(c *gin.Context) {
	var (
		mahasiswa structs.Mahasiswa
		result    gin.H
	)

	id := c.Param("id")

	if err := idb.DB.First(&mahasiswa, id).Error; err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	if err := idb.DB.Delete(&mahasiswa).Error; err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	result = gin.H{
		"result": "successfully delete data!",
	}
	c.JSON(http.StatusOK, result)
}
