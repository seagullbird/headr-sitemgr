package db

import (
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	// used for database connection
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Store deals with database operations with table site.
type Store interface {
	InsertSite(site *Site) (id uint, err error)
	DeleteSite(site *Site) error
	GetSite(id uint) (*Site, error)
	PatchSite(site *Site) error
	CheckSitenameExists(sitename string) (bool, error)
	GetSiteIDByUserID(userID string) (uint, error)
}

type databaseStore struct {
	db *gorm.DB
}

// New creates a databaseStore instance
func New(logger log.Logger) Store {
	db, err := gorm.Open("postgres", "host=postgresql-postgresql port=5432 user=postgres dbname=postgres password=qBDXNlz276 sslmode=disable")
	if err != nil {
		logger.Log("error_desc", "Failed to connected to PostgreSQL", "error", err)
	}
	db.AutoMigrate(&Site{})
	return &databaseStore{
		db: db,
	}
}

func (s *databaseStore) InsertSite(site *Site) (id uint, err error) {
	err = s.db.Create(site).Error
	return site.Model.ID, err
}

func (s *databaseStore) DeleteSite(site *Site) error {
	return s.db.Delete(&site).Error
}

func (s *databaseStore) GetSite(id uint) (*Site, error) {
	var site Site
	err := s.db.First(&site, id).Error
	return &site, err
}

func (s *databaseStore) PatchSite(site *Site) error {
	return nil
}

func (s *databaseStore) CheckSitenameExists(sitename string) (bool, error) {
	// exists, return true
	// not exists, return false
	var site Site
	err := s.db.Where("sitename = ?", sitename).First(&site).Error
	return site != Site{}, err
}

func (s *databaseStore) GetSiteIDByUserID(userID string) (uint, error) {
	var site Site
	err := s.db.Where("user_id = ?", userID).First(&site).Error
	return site.Model.ID, err
}
