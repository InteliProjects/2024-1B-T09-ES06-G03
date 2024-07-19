package geolocation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func getCoordinatesMock(address string) (float64, float64, error) {
	if address == "1600 Amphitheatre Parkway, Mountain View, CA" {
		return 37.4226277, -122.0841644, nil
	}
	return 0, 0, fmt.Errorf("location not found")
}

func TestGeoHandler(t *testing.T) {
	handler := &GeoHandler{
		getCoordinatesFunc: getCoordinatesMock,
	}
	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	t.Run("Get Coordinates - Valid Address", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/geocode?address=1600+Amphitheatre+Parkway,+Mountain+View,+CA", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response map[string]float64
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 37.4226277, response["lat"])
		assert.Equal(t, -122.0841644, response["long"])
	})

	t.Run("Get Coordinates - Invalid Address", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/geocode?address=Invalid+Address", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		var response map[string]string
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "location not found", response["error"])
	})

	t.Run("Get Coordinates - Missing Address", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/geocode", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		var response map[string]string
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "invalid address", response["error"])
	})
}
