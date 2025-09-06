package services

import "log/slog"

type SlogAdapter struct {
	logger *slog.Logger
}

func (s *SlogAdapter) Info(msg string, args ...any) {
	s.logger.Info(msg, args...)
}

func (s *SlogAdapter) Error(msg string, args ...any) {
	s.logger.Error(msg, args...)
}

func (s *SlogAdapter) Debug(msg string, args ...any) {
	s.logger.Debug(msg, args...)
}

func (s *SlogAdapter) Warn(msg string, args ...any) {
	s.logger.Warn(msg, args...)
}

func NewSlogAdapter(log *slog.Logger) *SlogAdapter {
	return &SlogAdapter{
		logger: log,
	}
}
