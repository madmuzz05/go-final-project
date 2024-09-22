package postgres

import (
	"fmt"

	"github.com/madmuzz05/go-final-project/internal/config"
	entityComment "github.com/madmuzz05/go-final-project/service/comment/entity"
	orderEntity "github.com/madmuzz05/go-final-project/service/order/entity"
	entityPhoto "github.com/madmuzz05/go-final-project/service/photo/entity"
	entitySosmed "github.com/madmuzz05/go-final-project/service/sosial_media/entity"
	userEntity "github.com/madmuzz05/go-final-project/service/user/entity"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormDB struct {
	DB  *gorm.DB
	Trx *gorm.DB
}

func LoadGorm(cfg *config.Config) (*GormDB, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseUsername, cfg.DatabasePassword, cfg.DatabaseName, cfg.DatabaseSSL)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}
	db.Debug().AutoMigrate(orderEntity.Order{}, orderEntity.Item{}, userEntity.User{}, entitySosmed.SosialMedia{}, entityPhoto.Photo{}, entityComment.Comment{})
	log.Info().Msg("connected successfully to the database with gorm!")

	return &GormDB{
		DB: db,
	}, nil
}
func (g *GormDB) BeginTransaction() {
	g.Trx = g.DB.Begin()
}

func (g *GormDB) CommitTransaction() {
	g.Trx.Commit()
	g.Trx = nil
}

func (g *GormDB) RollbackTransaction() {
	g.Trx.Rollback()
	g.Trx = nil
}

func (g *GormDB) GetDB() *gorm.DB {
	return g.DB
}
