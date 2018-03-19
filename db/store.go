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
	s.db.Create(site)
	return site.Model.ID, nil
}

func (s *databaseStore) DeleteSite(site *Site) error {
	s.db.Delete(&site)
	return nil
}

func (s *databaseStore) GetSite(id uint) (*Site, error) {
	var site Site
	s.db.First(&site, id)
	return &site, nil
}

func (s *databaseStore) PatchSite(site *Site) error {
	return nil
}

func (s *databaseStore) CheckSitenameExists(sitename string) (bool, error) {
	// exists, return true
	// not exists, return false
	var site Site
	s.db.Where("sitename = ?", sitename).First(&site)
	return site != Site{}, nil
}

func (s *databaseStore) GetSiteIDByUserID(userID string) (uint, error) {
	var site Site
	s.db.Where("user_id = ?", userID).First(&site)
	return site.Model.ID, nil
}
