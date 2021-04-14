package service

import (
	"encoding/csv"
	"fmt"
	"nowgoal/internal/core/domain"
	"nowgoal/internal/core/ports"
	"nowgoal/pkg/converter"
	"nowgoal/pkg/uidgen"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	defaultLimit = 10
)

type Service struct {
	postgresRepo ports.PostgresRepository
	uidgen       uidgen.UIDGen
	converter    converter.Converter
}

func New(postgresRepo ports.PostgresRepository, _uidgen uidgen.UIDGen, converter converter.Converter) *Service {
	return &Service{
		postgresRepo: postgresRepo,
		uidgen:       _uidgen,
		converter:    converter,
	}
}

// To create new tote ...
func (srv *Service) ReadValueFromCSVFile() (domain.ResultInformation, error) {
	csvFile, err := os.Open("./pkg/assets/onprocess/1513932-1539449.csv")
	if err != nil {
		logrus.Errorln(err)
		return domain.ResultInformation{}, err
	}
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		logrus.Errorln(err)
		return domain.ResultInformation{}, err
	}
	// left_team, 0
	// right_team, 1
	// first_half_score, 2
	// second_half_score, 3
	// final_score, 4

	// early_handicap_home_half, 5
	// early_handicap_handicap_half, 6
	// early_handicap_away_half, 7

	// start_handicap_home_half, 8
	// start_handicap_handicap_half, 9
	// start_handicap_away_half, 10

	// early_over_under_over_half, 11
	// early_over_under_o_u_half, 12
	// early_over_under_under_half, 13

	// start_over_under_over_half, 14
	// start_over_under_o_u_half, 15
	// start_over_under_under_half, 16

	// early_handicap_home_full, 17
	// early_handicap_handicap_full, 18
	// early_handicap_away_full, 19

	// start_handicap_home_full, 20
	// start_handicap_handicap_full, 21
	// start_handicap_away_full, 22

	// early_over_under_over_full, 23
	// early_over_under_o_u_full, 24
	// early_over_under_under_full, 25

	// start_over_under_over_full, 26
	// start_over_under_o_u_full, 27
	// start_over_under_under_full 28
	results := []domain.ResultInformation{}
	for _, line := range csvLines {
		halfStatEarlyHandicapHome := line[5]
		halfStatEarlyHandicapHandicap := line[6]
		halfStatEarlyHandicapAway := line[7]

		halfStatStartHandicapHome := line[8]
		halfStatStartHandicapHandicap := line[9]
		halfStatStartHandicapAway := line[10]

		halfStatEarlyOverUnderOver := line[11]
		halfStatEarlyOverUnderOU := line[12]
		halfStatEarlyOverUnderUnder := line[13]

		halfStatStartOverUnderOver := line[14]
		halfStatStartOverUnderOU := line[15]
		halfStatStartOverUnderUnder := line[16]

		fullStatEarlyHandicapHome := line[17]
		fullStatEarlyHandicapHandicap := line[18]
		fullStatEarlyHandicapAway := line[19]

		fullStatStartHandicapHome := line[20]
		fullStatStartHandicapHandicap := line[21]
		fullStatStartHandicapAway := line[22]

		fullStatEarlyOverUnderOver := line[23]
		fullStatEarlyOverUnderOU := line[24]
		fullStatEarlyOverUnderUnder := line[25]

		fullStatStartOverUnderOver := line[26]
		fullStatStartOverUnderOU := line[27]
		fullStatStartOverUnderUnder := line[28]

		strs := []string{
			halfStatEarlyHandicapHome,
			halfStatEarlyHandicapAway,

			halfStatStartHandicapHome,
			halfStatStartHandicapAway,

			halfStatEarlyOverUnderOver,
			halfStatEarlyOverUnderUnder,

			halfStatStartOverUnderOver,
			halfStatStartOverUnderUnder,

			fullStatEarlyHandicapHome,
			fullStatEarlyHandicapAway,

			fullStatStartHandicapHome,
			fullStatStartHandicapAway,

			fullStatEarlyOverUnderOver,
			fullStatEarlyOverUnderUnder,

			fullStatStartOverUnderOver,
			fullStatStartOverUnderUnder,
		}
		floats := make([]float32, len(strs))
		floats, err = srv.converter.ConvertStringsToFloats32(strs)
		if err != nil {
			logrus.Errorln(err)
			return domain.ResultInformation{}, err
		}
		results = append(results, domain.ResultInformation{
			ResultID:        srv.uidgen.New(),
			LeftTeam:        line[0],
			RightTeam:       line[1],
			FirstHalfScore:  line[2],
			SecondHalfScore: line[3],
			FinalScore:      line[4],
			HalfStat: domain.Stat{
				EarlyHandicap: domain.Handicap{
					Home:     floats[0],
					Handicap: halfStatEarlyHandicapHandicap,
					Away:     floats[1],
				},
				EarlyOverUnder: domain.OverUnder{
					Over:  floats[4],
					OU:    halfStatEarlyOverUnderOU,
					Under: floats[5],
				},
				StartHandicap: domain.Handicap{
					Home:     floats[2],
					Handicap: halfStatStartHandicapHandicap,
					Away:     floats[3],
				},
				StartOverUnder: domain.OverUnder{
					Over:  floats[6],
					OU:    halfStatStartOverUnderOU,
					Under: floats[7],
				},
			},
			FullStat: domain.Stat{
				EarlyHandicap: domain.Handicap{
					Home:     floats[8],
					Handicap: fullStatEarlyHandicapHandicap,
					Away:     floats[9],
				},
				EarlyOverUnder: domain.OverUnder{
					Over:  floats[12],
					OU:    fullStatEarlyOverUnderOU,
					Under: floats[13],
				},
				StartHandicap: domain.Handicap{
					Home:     floats[10],
					Handicap: fullStatStartHandicapHandicap,
					Away:     floats[11],
				},
				StartOverUnder: domain.OverUnder{
					Over:  floats[14],
					OU:    fullStatStartOverUnderOU,
					Under: floats[15],
				},
			},
		})
	}
	err = srv.postgresRepo.Insert3In1Results(results)
	if err != nil {
		return domain.ResultInformation{}, err
	}

	return results[0], nil

}

func (srv *Service) GetStatPattern1(result domain.GetStatsPattern1Request) (domain.GetStatsPattern1Response, error) {
	fmt.Println(result)
	results, err := srv.postgresRepo.FindPattern1(domain.ResultInformation{
		FullStat: domain.Stat{
			EarlyHandicap: domain.Handicap{
				Home:     result.Early.Handicap.Home,
				Handicap: result.Early.Handicap.Handicap,
				Away:     result.Early.Handicap.Away,
			},
			EarlyOverUnder: domain.OverUnder{
				Over:  result.Early.OverUnder.Over,
				OU:    result.Early.OverUnder.OU,
				Under: result.Early.OverUnder.Under,
			},
			StartHandicap: domain.Handicap{
				Home:     result.Start.Handicap.Home,
				Handicap: result.Start.Handicap.Handicap,
				Away:     result.Start.Handicap.Away,
			},
			StartOverUnder: domain.OverUnder{
				Over:  result.Start.OverUnder.Over,
				OU:    result.Start.OverUnder.OU,
				Under: result.Start.OverUnder.Under,
			},
		},
	})
	if err != nil {
		return domain.GetStatsPattern1Response{}, err
	}
	//TODO: sort results ...
	// if go `/` split
	// (x + (x * 0.03)) = x + 3%, 			n <= x
	// (x - (x * 0.03)) = x - 3%,           n >= x
	//TODO: Data formating ...
	// data over  ...

	// data equal ...
	// data lower  ...
	return domain.GetStatsPattern1Response{
		LenOverResult: len(results),
		OverResults:   results,
	}, nil
}
