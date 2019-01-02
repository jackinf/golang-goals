package apis

import (
	"github.com/jackinf/golang-goals/app"
	"github.com/jackinf/golang-goals/models"
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
)

type (
	// goalService specifies the interface for the goal service needed by goalResource.
	goalService interface {
		Get(rs app.RequestScope, id int) (*models.Goal, error)
		Query(rs app.RequestScope, offset, limit int) ([]models.Goal, error)
		Count(rs app.RequestScope) (int, error)
		Create(rs app.RequestScope, model *models.Goal) (*models.Goal, error)
		Update(rs app.RequestScope, id int, model *models.Goal) (*models.Goal, error)
		Delete(rs app.RequestScope, id int) (*models.Goal, error)
	}

	// goalResource defines the handlers for the CRUD APIs.
	goalResource struct {
		service goalService
	}
)

// ServeGoal sets up the routing of goal endpoints and the corresponding handlers.
func ServeGoalResource(rg *routing.RouteGroup, service goalService) {
	r := &goalResource{service}
	rg.Get("/goals/<id>", r.get)
	rg.Get("/goals", r.query)
	rg.Post("/goals", r.create)
	rg.Put("/goals/<id>", r.update)
	rg.Delete("/goals/<id>", r.delete)
}

func (r *goalResource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Get(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *goalResource) query(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	count, err := r.service.Count(rs)
	if err != nil {
		return err
	}
	paginatedList := getPaginatedListFromRequest(c, count)
	items, err := r.service.Query(app.GetRequestScope(c), paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return err
	}
	paginatedList.Items = items
	return c.Write(paginatedList)
}

func (r *goalResource) create(c *routing.Context) error {
	var model models.Goal
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(app.GetRequestScope(c), &model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *goalResource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	rs := app.GetRequestScope(c)

	model, err := r.service.Get(rs, id)
	if err != nil {
		return err
	}

	if err := c.Read(model); err != nil {
		return err
	}

	response, err := r.service.Update(rs, id, model)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r *goalResource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	response, err := r.service.Delete(app.GetRequestScope(c), id)
	if err != nil {
		return err
	}

	return c.Write(response)
}
