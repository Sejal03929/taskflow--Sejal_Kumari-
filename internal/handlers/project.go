package handlers


import (
	"taskflow/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateProject(c *gin.Context) {
	var project Project

	// Bind JSON
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	// Get user_id from middleware
	userID, _ := c.Get("user_id")

	// Generate UUID
	id := uuid.New()

	// Insert into DB
	_, err := db.DB.Exec(
		"INSERT INTO projects (id, name, description, owner_id) VALUES ($1, $2, $3, $4)",
		id, project.Name, project.Description, userID,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// ✅ RETURN ID (IMPORTANT)
	c.JSON(200, gin.H{
		"id":          id,
		"name":        project.Name,
		"description": project.Description,
	})
}

func GetProjects(c *gin.Context) {
	userID, _ := c.Get("user_id")

	rows, err := db.DB.Query(
		"SELECT id,name,description FROM projects WHERE owner_id=$1",
		userID,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var projects []map[string]interface{}

	for rows.Next() {
		var id, name, description string

		rows.Scan(&id, &name, &description)

		projects = append(projects, map[string]interface{}{
			"id":          id,
			"name":        name,
			"description": description,
		})
	}

	c.JSON(200, gin.H{"projects": projects})
}