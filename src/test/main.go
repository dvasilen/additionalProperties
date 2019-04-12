package main

import (
	"fmt"
	"net/http"
)

// @title AdditionalProperties
// @version 1.0
// @description Example for additionalProperties in query

// host localhost:8089
// @BasePath /test
func main() {
	fmt.Println("Hello")
}

// @Param   q query {object} model.PropMap false "The optional query parameters"


// @Summary Retrieves logs
// @Description Retrieves logs
// @Tags Jobs
// @ID jobLogs
// @Accept   json
// @Produce  plain
// @Param   id path string  true  "The job guid per RFC 4122"
// @Success 200 {object}  model.PropMap "Just to show it works for Success"
// @Router /v2/jobs/{id}/logs [get]
func JobLogs(w http.ResponseWriter, r *http.Request) {

}
