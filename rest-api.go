package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	email     string  `json:"email"`
	GPA       float32 `json:"gpa"`
}

var students []Student = []Student{
	{1, "Amina", "Rakhimova", "a_rakhimova@kbtu.kz", 2.6},
	{2, "Mariyam", "Seydalieva", "m_seydalieva@kbtu.kz", 3.0},
	{3, "Alibek", "Ikramuly", "al_ikramuly@kbtu.kz", 2.2},
	{4, "Assyl", "Erzhanuly", "a_erzhanuly@kbtu.kz", 2.7},
	{5, "Ramses", "Massalim", "r_massalim@kbtu.kz", 3.8},
}

func getStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}

func postStudents(c *gin.Context) {
	var newStudent Student
	if err := c.BindJSON(&newStudent); err != nil {
		return
	}
	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}

func getStudentByID(c *gin.Context) {
	var foundId string = c.Param("id")
	idInt, err := strconv.ParseInt(foundId, 10, 32)
	if err != nil {
		return
	}

	for _, stud := range students {
		if stud.ID == int(idInt) {
			c.IndentedJSON(http.StatusOK, stud)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func deleteStudentByID(c *gin.Context) {
	var foundId string = c.Param("id")
	idInt, err := strconv.ParseInt(foundId, 10, 32)
	if err != nil {
		return
	}
	for i, stud := range students {
		if stud.ID == int(idInt) {
			students = append(students[:i], students[i+1:]...)
			c.IndentedJSON(http.StatusOK, students)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func main() {
	router := gin.Default()
	router.GET("/students", getStudents)
	router.POST("/postStudent", postStudents)
	router.GET("/student/:id", getStudentByID)
	router.DELETE("/deleteStudent/:id", deleteStudentByID)

	router.Run("localhost:8081")