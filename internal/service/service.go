package service

import (
	"github.com/osamikoyo/test-task/internal/data"
	"github.com/osamikoyo/test-task/internal/data/models"
	"github.com/osamikoyo/test-task/pkg/loger"
)

type SongService struct {
    repo *data.SongRepository
	logger loger.Logger
}

func NewSongService(repo *data.SongRepository, logger loger.Logger) *SongService {
	return &SongService{repo: repo, logger: logger}
}

func (s *SongService) CreateSong(song *models.Song) error {
	s.logger.Info().Msg("Creating song in service layer")
	return s.repo.Create(song)
}

func (s *SongService) GetSongByID(id uint) (*models.Song, error) {
	s.logger.Info().Uint("id", id).Msg("Fetching song by ID in service layer")
	return s.repo.GetByID(id)
}

func (s *SongService) UpdateSong(song *models.Song) error {
	s.logger.Info().Uint("id", song.ID).Msg("Updating song in service layer")
	return s.repo.Update(song)
}

func (s *SongService) DeleteSong(id uint) error {
	s.logger.Info().Uint("id", id).Msg("Deleting song in service layer")
	return s.repo.Delete(id)
}

func (s *SongService) GetAllSongs(filter map[string]string, offset, limit int) ([]models.Song, error) {
	s.logger.Info().Interface("filter", filter).Int("offset", offset).Int("limit", limit).Msg("Fetching all songs in service layer")
	return s.repo.GetAll(filter, offset, limit)
}