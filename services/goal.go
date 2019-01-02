package services

import (
	"github.com/jackinf/golang-goals/app"
	"github.com/jackinf/golang-goals/models"
)

// goalDAO specifies the interface of the goal DAO needed by GoalService.
type goalDAO interface {
	// Get returns the goal with the specified goal ID.
	Get(rs app.RequestScope, id int) (*models.Goal, error)
	// Count returns the number of goals.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of goals with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.Goal, error)
	// Create saves a new goal in the storage.
	Create(rs app.RequestScope, goal *models.Goal) error
	// Update updates the goal with given ID in the storage.
	Update(rs app.RequestScope, id int, goal *models.Goal) error
	// Delete removes the goal with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

// GoalService provides services related with goals.
type GoalService struct {
	dao goalDAO
}

// NewGoalService creates a new GoalService with the given goal DAO.
func NewGoalService(dao goalDAO) *GoalService {
	return &GoalService{dao}
}

// Get returns the goal with the specified the goal ID.
func (s *GoalService) Get(rs app.RequestScope, id int) (*models.Goal, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new goal.
func (s *GoalService) Create(rs app.RequestScope, model *models.Goal) (*models.Goal, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Update updates the goal with the specified ID.
func (s *GoalService) Update(rs app.RequestScope, id int, model *models.Goal) (*models.Goal, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

// Delete deletes the goal with the specified ID.
func (s *GoalService) Delete(rs app.RequestScope, id int) (*models.Goal, error) {
	goal, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return goal, err
}

// Count returns the number of goals.
func (s *GoalService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

// Query returns the goals with the specified offset and limit.
func (s *GoalService) Query(rs app.RequestScope, offset, limit int) ([]models.Goal, error) {
	return s.dao.Query(rs, offset, limit)
}
