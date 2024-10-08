package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/helltale/EffectiveMobile2024/test-api/internal/db"
	"github.com/helltale/EffectiveMobile2024/test-api/internal/models"
)

// @Summary Get songs
// @Description Get songs with pagination
// @Tags songs
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Limit per page"
// @Success 200 {array} models.Song
// @Router /songs [get]
func GetSongs(c *gin.Context) {
	var songs []models.Song
	page := c.Query("page")
	limit := c.Query("limit")

	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}

	// Преобразуем параметры в целые числа
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
		return
	}

	offset := (pageInt - 1) * limitInt

	db.DB.Offset(offset).Limit(limitInt).Find(&songs)
	c.JSON(http.StatusOK, songs)
}

// @Summary Add song
// @Description Add a new song
// @Tags songs
// @Accept json
// @Produce json
// @Param song body models.Song true "Song data"
// @Success 201 {object} models.Song
// @Router /songs [post]
func AddSong(c *gin.Context) {
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	db.DB.Create(&song)
	c.JSON(http.StatusCreated, song)
}

// @Summary Delete song
// @Description Delete a song by ID
// @Tags songs
// @Param id path int true "Song ID"
// @Success 204
// @Router /songs/{id} [delete]
func DeleteSong(c *gin.Context) {
	id := c.Param("id")
	var song models.Song
	if err := db.DB.Where("id = ?", id).Delete(&song).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Song not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary Update song
// @Description Update a song by ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param song body models.Song true "Song data"
// @Success 200 {object} models.Song
// @Router /songs/{id} [put]
func UpdateSong(c *gin.Context) {
	id := c.Param("id")
	var song models.Song

	if err := db.DB.First(&song, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Song not found"})
		return
	}

	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	db.DB.Save(&song)
	c.JSON(http.StatusOK, song)
}

// @Router /songs/{id} [get]
func GetSongByID(c *gin.Context) {
	id := c.Param("id")
	var song models.Song

	if err := db.DB.First(&song, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Song not found"})
		return
	}

	c.JSON(http.StatusOK, song)
}

func RegisterRoutes(router *gin.Engine) {
	router.GET("/songs", GetSongs)
	router.POST("/songs", AddSong)
	router.DELETE("/songs/:id", DeleteSong)
	router.PUT("/songs/:id", UpdateSong)
	router.GET("/songs/:id", GetSongByID)
}
