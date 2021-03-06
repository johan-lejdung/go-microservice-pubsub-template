package goservice

import (
	"database/sql"
)

// Services contains methods for the Actions
type Services interface {
	TestFunction() error
}

// Service implementes these methods
type Service struct {
	Db *sql.DB `inject:""`
}

// compile-time interface implementation check
var _ Services = &Service{}

// TestFunction will always return nil, replace with real function
func (s *Service) TestFunction() error {
	return nil
}
