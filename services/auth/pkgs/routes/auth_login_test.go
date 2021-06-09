package routes_test

import (
	"auth/pkgs/routes"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"github.com/plutoincdev/pluto-go/pkgs/client"
	"github.com/plutoincdev/pluto-go/pkgs/providers/http"
	"testing"
)

type AuthLoginHandlerUnitTestSuite struct {
	suite.Suite
}

func TestAuthLoginHandlerUnitTestSuite(t *testing.T) {
	suite.Run(t, new(AuthLoginHandlerUnitTestSuite))
}

func (s *AuthLoginHandlerUnitTestSuite) Test_CorrectMethod() {
	s.Equal("POST", routes.AuthLoginEndpoint.Method)
}

func (s *AuthLoginHandlerUnitTestSuite) Test_Success() {
	c := client.New("test-client")

	httpProvider := routes.SetupRoute()
	c.RegisterProvider(httpProvider)
	c.Start(false)

	w := httpProvider.NewRecorder()
	req, _ := http.NewRequest(routes.AuthLoginEndpoint, nil)
	httpProvider.ServeHTTP(w, req)

	responseJson := make(map[string]interface{})
	_ = json.Unmarshal([]byte(w.Body.String()), &responseJson)

	s.Equal(200, w.Code)
	s.Equal("<JWT HERE>", responseJson["jwt"])
}
