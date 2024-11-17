package internal

//go:generate mockgen -mock_names historyLogger=MockHistoryLoggerGomock -typed -package internal  -destination zzz_history_logger_mock_gomock.go --source ./historyLogger.go
//go:generate minimock -g -i .historyLogger -o ./zzz_history_logger_mock_minimock.go

type historyLogger interface {
	Log(s *Shipment) error
}
