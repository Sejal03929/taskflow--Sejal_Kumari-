package handlers

import (
	//ne/http"

	"taskflow/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	ProjectID   string `json:"project_id"`
}

// ✅ Create Task
func CreateTask(c *gin.Context) {
	var task Task
	userID := c.GetString("user_id")
	projectID := c.Param("id") // ✅ from URL

	if err := c.BindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "invalid"})
		return
	}

	task.ID = uuid.New().String()

	if task.Status == "" {
		task.Status = "todo"
	}

	_, err := db.DB.Exec(
		`INSERT INTO tasks (id,title,description,status,priority,project_id,assignee_id)
		 VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		task.ID,
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
		projectID, // ✅ fixed
		userID,
	)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "task created"})
}

// ✅ Get Tasks
func GetTasks(c *gin.Context) {
	projectID := c.Query("project_id")

	rows, err := db.DB.Query(
		"SELECT id,title,description,status,priority FROM tasks WHERE project_id=$1",
		projectID,
	)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var t Task
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.Priority)
		tasks = append(tasks, t)
	}

	c.JSON(200, gin.H{"tasks": tasks})
}

// ✅ Update Task (mark done)
func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	_, err := db.DB.Exec(
		"UPDATE tasks SET status='done' WHERE id=$1",
		id,
	)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "task updated"})
}

// ✅ Delete Task
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	_, err := db.DB.Exec(
		"DELETE FROM tasks WHERE id=$1",
		id,
	)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "task deleted"})
}