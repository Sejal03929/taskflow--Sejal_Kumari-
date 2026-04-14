package handlers

import (
	"net/http"
	"os"
	"time"

	"taskflow/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ✅ Claims struct (IMPORTANT)
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// ---------------- REGISTER ----------------
func Register(c *gin.Context) {
	var user User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.ID = uuid.New().String()

	_, err := db.DB.Exec(
		"INSERT INTO users (id, name, email, password) VALUES ($1,$2,$3,$4)",
		user.ID, user.Name, user.Email, string(hashedPassword),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered"})
}

// ---------------- LOGIN ----------------
func Login(c *gin.Context) {
	var input User
	var dbUser User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err := db.DB.QueryRow(
		"SELECT id, password FROM users WHERE email=$1",
		input.Email,
	).Scan(&dbUser.ID, &dbUser.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid"})
		return
	}

	// ✅ CREATE TOKEN USING CLAIMS STRUCT
	claims := Claims{
		UserID: dbUser.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not login"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}