package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/osamikoyo/test-task/internal/data/models"
	"github.com/osamikoyo/test-task/internal/service"
	"github.com/osamikoyo/test-task/pkg/loger"
)

// SongHandler структура для обработчиков песен.
type SongHandler struct {
	service *service.SongService
	logger  loger.Logger
}

// NewSongHandler создает новый экземпляр SongHandler.
func NewSongHandler(service *service.SongService, logger loger.Logger) *SongHandler {
	return &SongHandler{service: service, logger: logger}
}

// CreateSong создает новую песню.
// @Summary Создать новую песню
// @Description Добавляет новую песню в библиотеку
// @Tags songs
// @Accept json
// @Produce json
// @Param song body models.Song true "Данные песни"
// @Success 201 {object} models.Song
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /songs [post]
func (h *SongHandler) CreateSong(c *gin.Context) {
	h.logger.Info().Msg("Handling create song request")
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		h.logger.Error().Err(err).Msg("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateSong(&song); err != nil {
		h.logger.Error().Err(err).Msg("Failed to create song")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.logger.Info().Msg("Song created successfully")
	c.JSON(http.StatusCreated, song)
}

// GetSong возвращает песню по ID.
// @Summary Получить песню по ID
// @Description Возвращает информацию о песне по её ID
// @Tags songs
// @Produce json
// @Param id path int true "ID песни"
// @Success 200 {object} models.Song
// @Failure 404 {object} map[string]string
// @Router /songs/{id} [get]
func (h *SongHandler) GetSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.logger.Info().Uint("id", uint(id)).Msg("Handling get song request")
	song, err := h.service.GetSongByID(uint(id))
	if err != nil {
		h.logger.Error().Err(err).Uint("id", uint(id)).Msg("Failed to fetch song")
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}
	h.logger.Info().Uint("id", uint(id)).Msg("Song fetched successfully")
	c.JSON(http.StatusOK, song)
}

// UpdateSong обновляет данные песни.
// @Summary Обновить песню
// @Description Обновляет информацию о песне по её ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param song body models.Song true "Обновленные данные песни"
// @Success 200 {object} models.Song
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /songs/{id} [put]
func (h *SongHandler) UpdateSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.logger.Info().Uint("id", uint(id)).Msg("Handling update song request")
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		h.logger.Error().Err(err).Msg("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	song.ID = uint(id)
	if err := h.service.UpdateSong(&song); err != nil {
		h.logger.Error().Err(err).Uint("id", uint(id)).Msg("Failed to update song")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.logger.Info().Uint("id", uint(id)).Msg("Song updated successfully")
	c.JSON(http.StatusOK, song)
}

// DeleteSong удаляет песню по ID.
// @Summary Удалить песню
// @Description Удаляет песню по её ID
// @Tags songs
// @Param id path int true "ID песни"
// @Success 204
// @Failure 500 {object} map[string]string
// @Router /songs/{id} [delete]
func (h *SongHandler) DeleteSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.logger.Info().Uint("id", uint(id)).Msg("Handling delete song request")
	if err := h.service.DeleteSong(uint(id)); err != nil {
		h.logger.Error().Err(err).Uint("id", uint(id)).Msg("Failed to delete song")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.logger.Info().Uint("id", uint(id)).Msg("Song deleted successfully")
	c.JSON(http.StatusNoContent, nil)
}

// GetAllSongs возвращает список песен с фильтрацией и пагинацией.
// @Summary Получить все песни
// @Description Возвращает список песен с возможностью фильтрации и пагинации
// @Tags songs
// @Produce json
// @Param group query string false "Группа"
// @Param song query string false "Название песни"
// @Param offset query int false "Смещение" default(0)
// @Param limit query int false "Лимит" default(10)
// @Success 200 {array} models.Song
// @Failure 500 {object} map[string]string
// @Router /songs [get]
func (h *SongHandler) GetAllSongs(c *gin.Context) {
	h.logger.Info().Msg("Handling get all songs request")
	filter := make(map[string]string)
	for key, value := range c.Request.URL.Query() {
		filter[key] = value[0]
	}
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	songs, err := h.service.GetAllSongs(filter, offset, limit)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to fetch songs")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.logger.Info().Int("count", len(songs)).Msg("Songs fetched successfully")
	c.JSON(http.StatusOK, songs)
}