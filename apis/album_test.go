package apis

import (
	"net/http"
	"testing"

	"golang-restful-starter-kit/daos"
	"golang-restful-starter-kit/services"
	"golang-restful-starter-kit/testdata"
)

func TestAlbum(t *testing.T) {
	testdata.ResetDB()
	router := newRouter()
	ServeAlbumResource(&router.RouteGroup, services.NewAlbumService(daos.NewAlbumDAO()))

	notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	nameRequiredError := `{"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"name","error":"cannot be blank"}]}`

	runAPITests(t, router, []apiTestCase{
		{"t1 - get an album", "GET", "/albums/2", "", http.StatusOK, `{"id":2,"name":"Accept"}`},
		{"t2 - get a nonexisting album", "GET", "/albums/99999", "", http.StatusNotFound, notFoundError},
		{"t3 - create an album", "POST", "/albums", `{"name":"Qiang"}`, http.StatusOK, `{"id": 276, "name":"Qiang"}`},
		{"t4 - create an album with validation error", "POST", "/albums", `{"name":""}`, http.StatusBadRequest, nameRequiredError},
		{"t5 - update an album", "PUT", "/albums/2", `{"name":"Qiang"}`, http.StatusOK, `{"id": 2, "name":"Qiang"}`},
		{"t6 - update an album with validation error", "PUT", "/albums/2", `{"name":""}`, http.StatusBadRequest, nameRequiredError},
		{"t7 - update a nonexisting album", "PUT", "/albums/99999", "{}", http.StatusNotFound, notFoundError},
		{"t8 - delete an album", "DELETE", "/albums/2", ``, http.StatusOK, `{"id": 2, "name":"Qiang"}`},
		{"t9 - delete a nonexisting album", "DELETE", "/albums/99999", "", http.StatusNotFound, notFoundError},
		{"t10 - get a list of albums", "GET", "/albums?page=3&per_page=2", "", http.StatusOK, `{"page":3,"per_page":2,"page_count":138,"total_count":275,"items":[{"id":6,"name":"Antônio Carlos Jobim"},{"id":7,"name":"Apocalyptica"}]}`},
	})
}
