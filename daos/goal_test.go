package daos

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang-restful-starter-kit/app"
	"golang-restful-starter-kit/models"
	"golang-restful-starter-kit/testdata"
)

func TestGoalDAO(t *testing.T) {
	db := testdata.ResetDB()
	dao := NewGoalDAO()

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			goal, err := dao.Get(rs, 2)
			assert.Nil(t, err)
			if assert.NotNil(t, goal) {
				assert.Equal(t, 2, goal.Id)
			}
		})
	}

	{
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			goal := &models.Goal{
				Id:   1000,
				Name: "tester",
			}
			err := dao.Create(rs, goal)
			assert.Nil(t, err)
			assert.NotEqual(t, 1000, goal.Id)
			assert.NotZero(t, goal.Id)
		})
	}

	{
		// Update
		testDBCall(db, func(rs app.RequestScope) {
			goal := &models.Goal{
				Id:   2,
				Name: "tester",
			}
			err := dao.Update(rs, goal.Id, goal)
			assert.Nil(t, err)
		})
	}

	{
		// Update with error
		testDBCall(db, func(rs app.RequestScope) {
			goal := &models.Goal{
				Id:   2,
				Name: "tester",
			}
			err := dao.Update(rs, 99999, goal)
			assert.NotNil(t, err)
		})
	}

	{
		// Delete
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, 2)
			assert.Nil(t, err)
		})
	}

	{
		// Delete with error
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, 99999)
			assert.NotNil(t, err)
		})
	}

	{
		// Query
		testDBCall(db, func(rs app.RequestScope) {
			goals, err := dao.Query(rs, 1, 3)
			assert.Nil(t, err)
			assert.Equal(t, 3, len(goals))
		})
	}

	{
		// Count
		testDBCall(db, func(rs app.RequestScope) {
			count, err := dao.Count(rs)
			assert.Nil(t, err)
			assert.NotZero(t, count)
		})
	}
}
