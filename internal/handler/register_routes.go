package handler

import "github.com/gin-gonic/gin"

func (h *SongHandler) RegisterRoutes(g *gin.Engine){
	g.GET("/songs", h.GetAllSongs)
    g.GET("/songs/:id", h.GetSong)
    g.POST("/songs", h.GetAllSongs)
    g.PUT("/songs/:id", h.UpdateSong)
    g.DELETE("/songs/:id", h.DeleteSong)
}