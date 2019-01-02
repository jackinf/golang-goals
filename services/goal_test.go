package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/jackinf/golang-goals/app"
	"github.com/jackinf/golang-goals/models"
)

func TestNewGoalService(t *testing.T) {
	dao := newMockGoalDAO()
	s := NewGoalService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestGoalService_Get(t *testing.T) {
	s := NewGoalService(newMockGoalDAO())
	goal, err := s.Get(nil, 1)
	if assert.Nil(t, err) && assert.NotNil(t, goal) {
		assert.Equal(t, "aaa", goal.Name)
	}

	goal, err = s.Get(nil, 100)
	assert.NotNil(t, err)
}

func TestGoalService_Create(t *testing.T) {
	s := NewGoalService(newMockGoalDAO())
	goal, err := s.Create(nil, &models.Goal{
		Name: "ddd",
	})
	if assert.Nil(t, err) && assert.NotNil(t, goal) {
		assert.Equal(t, 4, goal.Id)
		assert.Equal(t, "ddd", goal.Name)
	}

	// dao error
	_, err = s.Create(nil, &models.Goal{
		Id:   100,
		Name: "ddd",
	})
	assert.NotNil(t, err)

	// validation error
	_, err = s.Create(nil, &models.Goal{
		Name: "",
	})
	assert.NotNil(t, err)
}

func TestGoalService_Update(t *testing.T) {
	s := NewGoalService(newMockGoalDAO())
	goal, err := s.Update(nil, 2, &models.Goal{
		Name: "ddd",
	})
	if assert.Nil(t, err) && assert.NotNil(t, goal) {
		assert.Equal(t, 2, goal.Id)
		assert.Equal(t, "ddd", goal.Name)
	}

	// dao error
	_, err = s.Update(nil, 100, &models.Goal{
		Name: "ddd",
	})
	assert.NotNil(t, err)

	// validation error
	_, err = s.Update(nil, 2, &models.Goal{
		Name: "",
	})
	assert.NotNil(t, err)
}

func TestGoalService_Delete(t *testing.T) {
	s := NewGoalService(newMockGoalDAO())
	goal, err := s.Delete(nil, 2)
	if assert.Nil(t, err) && assert.NotNil(t, goal) {
		assert.Equal(t, 2, goal.Id)
		assert.Equal(t, "bbb", goal.Name)
	}

	_, err = s.Delete(nil, 2)
	assert.NotNil(t, err)
}

func TestGoalService_Query(t *testing.T) {
	s := NewGoalService(newMockGoalDAO())
	result, err := s.Query(nil, 1, 2)
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(result))
	}
}

func newMockGoalDAO() goalDAO {
	return &mockGoalDAO{
		records: []models.Goal{
			{Id: 1, Name: "aaa"},
			{Id: 2, Name: "bbb"},
			{Id: 3, Name: "ccc"},
		},
	}
}

type mockGoalDAO struct {
	records []models.Goal
}

func (m *mockGoalDAO) Get(rs app.RequestScope, id int) (*models.Goal, error) {
	for _, record := range m.records {
		if record.Id == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockGoalDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Goal, error) {
	return m.records[offset : offset+limit], nil
}

func (m *mockGoalDAO) Count(rs app.RequestScope) (int, error) {
	return len(m.records), nil
}

func (m *mockGoalDAO) Create(rs app.RequestScope, goal *models.Goal) error {
	if goal.Id != 0 {
		return errors.New("Id cannot be set")
	}
	goal.Id = len(m.records) + 1
	m.records = append(m.records, *goal)
	return nil
}

func (m *mockGoalDAO) Update(rs app.RequestScope, id int, goal *models.Goal) error {
	goal.Id = id
	for i, record := range m.records {
		if record.Id == id {
			m.records[i] = *goal
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockGoalDAO) Delete(rs app.RequestScope, id int) error {
	for i, record := range m.records {
		if record.Id == id {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
