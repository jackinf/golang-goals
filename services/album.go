package services

import (
	"golang-restful-starter-kit/app"
	"golang-restful-starter-kit/models"
)

// AlbumDAO specifies the interface of the Album DAO needed by AlbumService.
type albumDAO interface {
	// Get returns the Album with the specified Album ID.
	Get(rs app.RequestScope, id int) (*models.Album, error)
	// Count returns the number of Albums.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of Albums with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Album, error)
	// Create saves a new Album in the storage.
	Create(rs app.RequestScope, Album *models.Album) error
	// Update updates the Album with given ID in the storage.
	Update(rs app.RequestScope, id int, Album *models.Album) error
	// Delete removes the Album with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// AlbumService provides services related with Albums.
type AlbumService struct {
	dao albumDAO
}

// NewAlbumService creates a new AlbumService with the given Album DAO.
func NewAlbumService(dao albumDAO) *AlbumService {
	return &AlbumService{dao}
}

// Get returns the Album with the specified the Album ID.
func (s *AlbumService) Get(rs app.RequestScope, id int) (*models.Album, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new Album.
func (s *AlbumService) Create(rs app.RequestScope, model *models.Album) (*models.Album, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Update updates the Album with the specified ID.
func (s *AlbumService) Update(rs app.RequestScope, id int, model *models.Album) (*models.Album, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the Album with the specified ID.
func (s *AlbumService) Delete(rs app.RequestScope, id int) (*models.Album, error) {
	Album, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return Album, err
}

// Count returns the number of Albums.
func (s *AlbumService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the Albums with the specified offset and limit.
func (s *AlbumService) Query(rs app.RequestScope, offset, limit int) ([]models.Album, error) {
	return s.dao.Query(rs, offset, limit)
}
