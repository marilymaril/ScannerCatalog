package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// @title Scanner Catalog API
// @version 1.0
// @description This is a sample server for a Scanner Catalog API.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

// Scanner represents a scanner item in the catalog.
type Scanner struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var scanners = []Scanner{
	{ID: 1, Name: "Scanner A", Type: "Flatbed"},
	{ID: 2, Name: "Scanner B", Type: "Sheetfed"},
}

// @Summary Get scanner item by ID
// @Description Get scanner item details by numeric ID.
// @Tags scanners
// @Accept json
// @Produce json
// @Param id path int true "Scanner ID"
// @Success 200 {object} Scanner
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /scanners/{id} [get]
func getItemByID(w http.ResponseWriter, r *http.Request) {
	// Remove the route prefix to get just the ID portion from the URL path.
	idStr := strings.TrimPrefix(r.URL.Path, "/scanners/")

	// Convert the ID string to an integer.
	id, err := strconv.Atoi(idStr)

	// If the ID is not a valid positive integer, return a 400 Bad Request.
	if err != nil || id < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Code: 400, Message: "invalid scanner id"})
		return
	}

	// Search the sample scanner list for a matching ID.
	for _, scanner := range scanners {
		if scanner.ID == id {
			// Found a matching scanner: set JSON content type and write the scanner.
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(scanner)
			return
		}
	}

	// No item matched the requested ID, so return 404 Not Found.
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ErrorResponse{Code: 404, Message: "item not found"})
}

func main() {

}
