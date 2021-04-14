package ports

import "nowgoal/internal/core/domain"

type PostgresRepository interface {
	Insert3In1Results(results []domain.ResultInformation) error
	FindPattern1(result domain.ResultInformation) ([]domain.ResultInformation, error)
}
