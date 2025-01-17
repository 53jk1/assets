package report

import "fmt"

type Service struct {
	errors     int
	totalFiles int
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) IncErrors() {
	s.errors++
}

func (s *Service) IncTotalFiles() {
	s.totalFiles++
}

func (s Service) IsFailed() bool {
	return s.errors > 0
}

func (s Service) GetReport() string {
	return fmt.Sprintf("Total files: %d, errors: %d", s.totalFiles, s.errors)
}
