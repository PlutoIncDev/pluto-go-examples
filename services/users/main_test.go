package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type MainUnitTestSuite struct {
	suite.Suite
}

func TestMainUnitTestSuite(t *testing.T) {
	suite.Run(t, new(MainUnitTestSuite))
}

//func (s *MainUnitTestSuite) Test_HttpEndpoint_HealthCheck() {
//	client := RunClient()
//}
