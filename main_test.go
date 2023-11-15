package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_gin/controllers"
	"go_gin/database"
	"go_gin/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID uint
var CPF string

type StudentJson struct {
	Student struct {
		models.Student
	} `json:"student"`
}

func SetupRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{
		Name: "Student Test", CPF: "12345678901", RG: "123456789",
	}
	database.DB.Create(&student)
	ID = student.ID
	CPF = student.CPF
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}
func TestIndexStudents(t *testing.T) {
	database.DBConnect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutes()
	r.GET("/students", controllers.IndexStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

}

func TestShowStudentByCPF(t *testing.T) {
	database.DBConnect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutes()
	r.GET("/students/cpf/:cpf", controllers.ShowStudentByCPF)
	path := fmt.Sprintf("/students/cpf/%s", CPF)
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

}

func TestShowStudent(t *testing.T) {

	database.DBConnect()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutes()
	r.GET("/students/:id", controllers.ShowStudent)
	path := fmt.Sprintf("/students/%d", ID)
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var studentTest StudentJson
	json.Unmarshal(res.Body.Bytes(), &studentTest)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Student Test", studentTest.Student.Name)

}

func TestDeleteStudent(t *testing.T) {
	database.DBConnect()
	CreateStudentMock()
	r := SetupRoutes()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	path := "/students/" + strconv.Itoa(int(ID))
	req, _ := http.NewRequest("DELETE", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestUpdateStudent(t *testing.T) {
	database.DBConnect()
	CreateStudentMock()
	defer DeleteStudentMock()
	studentParam := models.Student{
		Name: "Student Test", CPF: "12345678989", RG: "343456789",
	}
	studentParamJson, _ := json.Marshal(studentParam)
	r := SetupRoutes()
	r.PATCH("/students/:id", controllers.UpdateStudent)
	path := "/students/" + strconv.Itoa(int(ID))
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(studentParamJson))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	responseMock := `{"data":"student updated"}`
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, responseMock, res.Body.String())
}
