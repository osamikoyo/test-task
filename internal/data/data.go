package data

import (
	"fmt"

	"github.com/osamikoyo/test-task/internal/config"
	"github.com/osamikoyo/test-task/internal/data/models"
	"github.com/osamikoyo/test-task/pkg/loger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SongRepository struct{
	db *gorm.DB
	logger loger.Logger
}

func New(cfg *config.Config) (*SongRepository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
	cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil{
		return nil, fmt.Errorf("cant connect to db: %v", err)
	}

	err = db.AutoMigrate(&models.Song{})

	return &SongRepository{db: db, logger: loger.New()}, err
}

func (r *SongRepository) Create(song *models.Song) error {
	r.logger.Info().Msg("Creating a new song")
	if err := r.db.Create(song).Error; err != nil {
		r.logger.Error().Err(err).Msg("Failed to create song")
		return err
	}
	r.logger.Info().Msg("Song created successfully")
	return nil
}

func (r *SongRepository) GetByID(id uint) (*models.Song, error) {
	r.logger.Info().Uint("id", id).Msg("Fetching song by ID")
	var song models.Song
	if err := r.db.First(&song, id).Error; err != nil {
		r.logger.Error().Err(err).Uint("id", id).Msg("Failed to fetch song by ID")
		return nil, err
	}
	r.logger.Info().Uint("id", id).Msg("Song fetched successfully")
	return &song, nil
}

func (r *SongRepository) Update(song *models.Song) error {
	r.logger.Info().Uint("id", song.ID).Msg("Updating song")
	if err := r.db.Save(song).Error; err != nil {
		r.logger.Error().Err(err).Uint("id", song.ID).Msg("Failed to update song")
		return err
	}
	r.logger.Info().Uint("id", song.ID).Msg("Song updated successfully")
	return nil
}

func (r *SongRepository) Delete(id uint) error {
	r.logger.Info().Uint("id", id).Msg("Deleting song")
	if err := r.db.Delete(&models.Song{}, id).Error; err != nil {
		r.logger.Error().Err(err).Uint("id", id).Msg("Failed to delete song")
		return err
	}
	r.logger.Info().Uint("id", id).Msg("Song deleted successfully")
	return nil
}

func (r *SongRepository) GetAll(filter map[string]string, offset, limit int) ([]models.Song, error) {
	r.logger.Info().Interface("filter", filter).Int("offset", offset).Int("limit", limit).Msg("Fetching all songs")
	var songs []models.Song
	query := r.db
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}
	if err := query.Offset(offset).Limit(limit).Find(&songs).Error; err != nil {
		r.logger.Error().Err(err).Msg("Failed to fetch songs")
		return nil, err
	}
	r.logger.Info().Int("count", len(songs)).Msg("Songs fetched successfully")
	return songs, nil
}