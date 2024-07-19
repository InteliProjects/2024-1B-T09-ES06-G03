package geolocation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/config"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/utils"
	"github.com/gorilla/mux"
)

// GeoLocationResponse representa a resposta da API de geolocalização
type GeoLocationResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat  float64 `json:"lat"`
				Long float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}

// GeoHandler é o handler para os endpoints de geolocalização
type GeoHandler struct {
	getCoordinatesFunc func(string) (float64, float64, error)
}

// NewGeoHandler cria um novo GeoHandler
func NewGeoHandler() *GeoHandler {
	return &GeoHandler{
		getCoordinatesFunc: getCoordinates,
	}
}

// RegisterRoutes registra as rotas de geolocalização no roteador
func (h *GeoHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/geocode", h.handleGeoCode).Methods(http.MethodGet)
}

// handleGeoCode manipula as solicitações de geocodificação
// @Summary Obtém as coordenadas de um endereço
// @Description Retorna a latitude e longitude de um endereço fornecido
// @Tags geolocation
// @Produce json
// @Param address query string true "Endereço para geocodificação"
// @Success 200 {object} map[string]float64
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /geocode [get]
func (h *GeoHandler) handleGeoCode(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid address"))
		return
	}

	lat, long, err := h.getCoordinatesFunc(address)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response := map[string]float64{
		"lat":  lat,
		"long": long,
	}
	utils.WriteJSON(w, http.StatusOK, response)
}

// getCoordinates obtém as coordenadas de um endereço usando a API do Google Maps
func getCoordinates(address string) (float64, float64, error) {
	apiKey := config.Envs.GoogleMapsAPIKey
	escapedAddress := url.QueryEscape(address)
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", escapedAddress, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	fmt.Println("Google Maps API response:", string(body))

	var geoResp GeoLocationResponse

	if err := json.Unmarshal(body, &geoResp); err != nil {
		return 0, 0, err
	}

	if len(geoResp.Results) > 0 {
		return geoResp.Results[0].Geometry.Location.Lat, geoResp.Results[0].Geometry.Location.Long, nil
	}

	return 0, 0, fmt.Errorf("no results found")
}
