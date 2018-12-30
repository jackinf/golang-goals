package daos

import (
	"golang-restful-starter-kit/app"
	"golang-restful-starter-kit/models"
)

// GoalDAO persists goal data in database
type GoalDAO struct{}

// NewGoalDAO creates a new GoalDAO
func NewGoalDAO() *GoalDAO {
	return &GoalDAO{}
}

// Get reads the goal with the specified ID from the database.
func (dao *GoalDAO) Get(rs app.RequestScope, id int) (*models.Goal, error) {
	var goal models.Goal
	err := rs.Tx().Select().Model(id, &goal)
	return &goal, err
}

// Create saves a new goal record in the database.
// The Goal.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *GoalDAO) Create(rs app.RequestScope, goal *models.Goal) error {
	goal.Id = 0
	return rs.Tx().Model(goal).Insert()
}

// Update saves the changes to an goal in the database.
func (dao *GoalDAO) Update(rs app.RequestScope, id int, goal *models.Goal) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	goal.Id = id
	return rs.Tx().Model(goal).Exclude("Id").Update()
}

// Delete deletes an goal with the specified ID from the database.
func (dao *GoalDAO) Delete(rs app.RequestScope, id int) error {
	goal, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(goal).Delete()
}

// Count returns the number of the goal records in the database.
func (dao *GoalDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("goal").Row(&count)
	return count, err
}

// Query retrieves the goal records with the specified offset and limit from the database.
func (dao *GoalDAO) Query(rs app.RequestScope, offset, limit int) ([]models.Goal, error) {
	goals := []models.Goal{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&goals)
	return goals, err
}
