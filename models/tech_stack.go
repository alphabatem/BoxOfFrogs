package models

import "fmt"

type TechStack struct {
	Frontend *Frontend `json:"frontend"`
	Backend  *Backend  `json:"backend"`
	Database *Database `json:"database"`
}

type Frontend struct {
	Framework string `json:"framework"`
}

func (f *Frontend) String() string {
	s := fmt.Sprintf("Frontend: %s", f.Framework)
	return s
}

type Backend struct {
	Language string `json:"language"`
}

func (b *Backend) String() string {
	s := fmt.Sprintf("Backend: %s", b.Language)
	return s
}

type Database struct {
	Driver string `json:"driver"`
}

func (d *Database) String() string {
	s := fmt.Sprintf("Database: %s", d.Driver)
	return s
}

func (s *TechStack) String() string {
	str := ""
	if s.Frontend != nil {
		str = fmt.Sprintf("%s\n%s", str, s.Frontend.String())
	}

	if s.Backend != nil {
		str = fmt.Sprintf("%s\n%s", str, s.Backend.String())
	}

	if s.Database != nil {
		str = fmt.Sprintf("%s\n%s", str, s.Database.String())
	}

	return str
}
