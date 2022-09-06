package main

// Imports of the modules needed for the application
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Structures needed --------------------------------------------------------------
// This structure represents the data of the student
type student struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	LastName string  `json:"lastName"`
	Semester float64 `json:"semester"`
}

// Variables needed --------------------------------------------------------------
// students to use as a start
var students = []student{
	{ID: "1", Name: "Cristian", LastName: "Sanchez", Semester: 9},
	{ID: "2", Name: "Alejandro", LastName: "Gonzales", Semester: 8},
	{ID: "3", Name: "isabella", LastName: "Correa", Semester: 7},
}

// Main function needed --------------------------------------------------------------
func main() {
	router := gin.Default()
	router.GET("/students", getStudents)
	router.GET("/students/:id", getStudent)
	router.POST("/students", postStudent)
	router.DELETE("/students/:id", deleteStudent)
	router.PUT("/students/:id", updateStudent)

	router.Run("localhost:8080")
}

// Auxiliar functions needed --------------------------------------------------------------
// Gets all current students
func getStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}

// Gets the student with the specified Id
func getStudent(c *gin.Context) {
	id := c.Param("id")

	for _, a := range students {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

// Creates a new Student
func postStudent(c *gin.Context) {
	var newStudent student

	if err := c.BindJSON(&newStudent); err != nil {
		return
	}

	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}

// Deletes the student with the specified Id
func deleteStudent(c *gin.Context) {
	id := c.Param("id")

	for i, a := range students {
		if a.ID == id {
			students = append(students[:i], students[i+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func updateStudent(c *gin.Context) {
	id := c.Param("id")
	var updatedStudent student

	if err := c.BindJSON(&updatedStudent); err != nil {
		return
	}

	for i, a := range students {
		if a.ID == id {
			students[i] = updatedStudent
			c.IndentedJSON(http.StatusOK, updatedStudent)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}
