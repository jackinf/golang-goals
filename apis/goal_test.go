package apis

import (
	"net/http"
	"testing"

	"golang-restful-starter-kit/daos"
	"golang-restful-starter-kit/services"
	"golang-restful-starter-kit/testdata"
)

func TestGoal(t *testing.T) {
	testdata.ResetDB()
	router := newRouter()
	ServeGoalResource(&router.RouteGroup, services.NewGoalService(daos.NewGoalDAO()))

	notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	nameRequiredError := `{"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"name","error":"cannot be blank"}]}`

	runAPITests(t, router, []apiTestCase{
		{"t1 - get an artist", "GET", "/goals/2", "", http.StatusOK, `{"id":2,"name":"Accept"}`},
		{"t2 - get a nonexisting artist", "GET", "/goals/99999", "", http.StatusNotFound, notFoundError},
		{"t3 - create an artist", "POST", "/goals", `{"name":"Qiang"}`, http.StatusOK, `{"id": 276, "name":"Qiang"}`},
		{"t4 - create an artist with validation error", "POST", "/goals", `{"name":""}`, http.StatusBadRequest, nameRequiredError},
		{"t5 - update an artist", "PUT", "/goals/2", `{"name":"Qiang"}`, http.StatusOK, `{"id": 2, "name":"Qiang"}`},
		{"t6 - update an artist with validation error", "PUT", "/goals/2", `{"name":""}`, http.StatusBadRequest, nameRequiredError},
		{"t7 - update a nonexisting artist", "PUT", "/goals/99999", "{}", http.StatusNotFound, notFoundError},
		{"t8 - delete an artist", "DELETE", "/goals/2", ``, http.StatusOK, `{"id": 2, "name":"Qiang"}`},
		{"t9 - delete a nonexisting artist", "DELETE", "/goals/99999", "", http.StatusNotFound, notFoundError},
		{"t10 - get a list of goals", "GET", "/goals?page=3&per_page=2", "", http.StatusOK, `{"page":3,"per_page":2,"page_count":138,"total_count":275,"items":[{"id":6,"name":"Antônio Carlos Jobim"},{"id":7,"name":"Apocalyptica"}]}`},
	})
}
