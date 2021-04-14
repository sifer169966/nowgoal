package ports

import "nowgoal/internal/core/domain"

type Service interface {
	ReadValueFromCSVFile() (domain.ResultInformation, error)
	GetStatPattern1(result domain.GetStatsPattern1Request) (domain.GetStatsPattern1Response, error)
}
