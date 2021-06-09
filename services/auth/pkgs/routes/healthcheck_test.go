package routes_test

import (
	"auth/pkgs/routes"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"pluto/pkgs/client"
	"pluto/pkgs/providers/http"
	"testing"
)

type HealthCheckHandlerUnitTestSuite struct {
	suite.Suite
}

func TestHealthCheckHandlerUnitTestSuite(t *testing.T) {
	suite.Run(t, new(HealthCheckHandlerUnitTestSuite))
}

func (s *HealthCheckHandlerUnitTestSuite) Test_Success() {
	c := client.New("test-client")

	httpProvider := routes.SetupRoute()
	c.RegisterProvider(httpProvider)
	c.Start(false)

	w := httpProvider.NewRecorder()
	req, _ := http.NewRequest(routes.HealthCheckEndpoint, nil)
	httpProvider.ServeHTTP(w, req)

	responseJson := make(map[string]interface{})
	_ = json.Unmarshal([]byte(w.Body.String()), &responseJson)

	s.Equal(200, w.Code)
	s.Equal("healthy", responseJson["status"])
}
