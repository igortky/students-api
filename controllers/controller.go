package controllers

import (
	"go_gin/database"
	"go_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	data := gin.H{
		"students": students,
	}
	c.JSON(200, data)
}

func ShowStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Find(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student with id " + id + "not found"})
		return
	}
	data := gin.H{
		"student": student,
	}
	c.JSON(200, data)
}

func ShowStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")
	if err := database.DB.Where("cpf = ?", cpf).First(&student).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Student not found",
		})
		return
	}

	data := gin.H{
		"student": student,
	}
	c.JSON(200, data)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student with id " + id + "not found"})
		return
	}
	database.DB.Delete(&student, id)
	data := gin.H{
		"data": "student deleted",
	}
	c.JSON(200, data)
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	database.DB.Model(&student).UpdateColumns(student)
	data := gin.H{
		"data": "student updated",
	}
	c.JSON(200, data)
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "not_found.html", nil)
}
