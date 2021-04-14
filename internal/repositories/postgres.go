package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"nowgoal/internal/core/domain"
	"nowgoal/internal/rdbms/postgresql/nowgoal/public/model"
	"nowgoal/internal/rdbms/postgresql/nowgoal/public/table"
	"nowgoal/pkg/converter"

	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/sirupsen/logrus"
)

type Postgres struct {
	DB        *sql.DB
	converter converter.Converter
}

func NewPostgres(DB *sql.DB, converter converter.Converter) *Postgres {
	createTable(DB)
	return &Postgres{
		DB:        DB,
		converter: converter,
	}
}

func createTable(db *sql.DB) error {
	var err error
	const createTable3In1Results = `CREATE TABLE IF NOT EXISTS Public."three_in_one_results"(
		result_id varchar(200) PRIMARY KEY NOT NULL,
		left_team varchar(200) NOT NULL,
		right_team varchar(200) NOT NULL,
		first_half_score varchar(100) NOT NULL,
		second_half_score varchar(100) NOT NULL,
		final_score varchar(100) NOT NULL,
		early_handicap_home_half real NOT NULL,
		early_handicap_handicap_half varchar(100) NOT NULL,
		early_handicap_away_half real NOT NULL,
		early_over_under_over_half real NOT NULL,
		early_over_under_o_u_half varchar(100) NOT NULL,
		early_over_under_under_half real NOT NULL,
		start_handicap_home_half real NOT NULL,
		start_handicap_handicap_half varchar(100) NOT NULL,
		start_handicap_away_half real NOT NULL,
		start_over_under_over_half real NOT NULL,
		start_over_under_o_u_half varchar(100) NOT NULL,
		start_over_under_under_half real NOT NULL,
		early_handicap_home_full real NOT NULL,
		early_handicap_handicap_full varchar(100) NOT NULL,
		early_handicap_away_full real NOT NULL,
		early_over_under_over_full real NOT NULL,
		early_over_under_o_u_full varchar(100) NOT NULL,
		early_over_under_under_full real NOT NULL,
		start_handicap_home_full real NOT NULL,
		start_handicap_handicap_full varchar(100) NOT NULL,
		start_handicap_away_full real NOT NULL,
		start_over_under_over_full real NOT NULL,
		start_over_under_o_u_full varchar(100) NOT NULL,
		start_over_under_under_full real NOT NULL
	);`
	_, err = db.Exec(createTable3In1Results)
	if err != nil {
		log.Fatalln("Cannot create table 3_in_1_results got error: ", err)
		return err
	}

	return nil
}

// now results 61171 ...
func (r *Postgres) Insert3In1Results(results []domain.ResultInformation) error {
	//resultSchemas := []model.ThreeInOneResults{}
	tx, err := r.DB.Begin()
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	var count int
	for _, v := range results {
		stmt := table.ThreeInOneResults.INSERT(
			table.ThreeInOneResults.AllColumns,
		).MODEL(model.ThreeInOneResults{
			ResultID:                  v.ResultID,
			LeftTeam:                  v.LeftTeam,
			RightTeam:                 v.RightTeam,
			FirstHalfScore:            v.FirstHalfScore,
			SecondHalfScore:           v.SecondHalfScore,
			FinalScore:                v.FinalScore,
			EarlyHandicapHomeHalf:     v.HalfStat.EarlyHandicap.Home,
			EarlyHandicapHandicapHalf: v.HalfStat.EarlyHandicap.Handicap,
			EarlyHandicapAwayHalf:     v.HalfStat.EarlyHandicap.Away,
			EarlyOverUnderOverHalf:    v.HalfStat.EarlyOverUnder.Over,
			EarlyOverUnderOUHalf:      v.HalfStat.EarlyOverUnder.OU,
			EarlyOverUnderUnderHalf:   v.HalfStat.EarlyOverUnder.Under,
			StartHandicapHomeHalf:     v.HalfStat.StartHandicap.Home,
			StartHandicapHandicapHalf: v.HalfStat.StartHandicap.Handicap,
			StartHandicapAwayHalf:     v.HalfStat.StartHandicap.Away,
			StartOverUnderOverHalf:    v.HalfStat.StartOverUnder.Over,
			StartOverUnderOUHalf:      v.HalfStat.StartOverUnder.OU,
			StartOverUnderUnderHalf:   v.HalfStat.StartOverUnder.Under,
			EarlyHandicapHomeFull:     v.FullStat.EarlyHandicap.Home,
			EarlyHandicapHandicapFull: v.FullStat.EarlyHandicap.Handicap,
			EarlyHandicapAwayFull:     v.FullStat.EarlyHandicap.Away,
			EarlyOverUnderOverFull:    v.FullStat.EarlyOverUnder.Over,
			EarlyOverUnderOUFull:      v.FullStat.EarlyOverUnder.OU,
			EarlyOverUnderUnderFull:   v.FullStat.EarlyOverUnder.Under,
			StartHandicapHomeFull:     v.FullStat.StartHandicap.Home,
			StartHandicapHandicapFull: v.FullStat.StartHandicap.Handicap,
			StartHandicapAwayFull:     v.FullStat.StartHandicap.Away,
			StartOverUnderOverFull:    v.FullStat.StartOverUnder.Over,
			StartOverUnderOUFull:      v.FullStat.StartOverUnder.OU,
			StartOverUnderUnderFull:   v.FullStat.StartOverUnder.Under,
		})
		//logrus.Info(stmt.DebugSql())
		_, err = stmt.Exec(tx)
		if err != nil {
			logrus.Errorln(err)
			tx.Rollback()
			return err
		}
		count++
		fmt.Println("inserted count ", count)
	}
	err = tx.Commit()
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	//logrus.Info(stmt.DebugSql())

	return nil
}

func (r *Postgres) FindPattern1(result domain.ResultInformation) ([]domain.ResultInformation, error) {
	earlyHomePlusThreePercentage := result.FullStat.EarlyHandicap.Home * 1.03 // x  + 3%
	earlyHomeMinusThreePercentage := result.FullStat.EarlyHandicap.Home * 0.97

	ealyAwayPlusThreePercentage := result.FullStat.EarlyHandicap.Away * 1.03
	ealyAwayMinusThreePercentage := result.FullStat.EarlyHandicap.Away * 0.97

	earlyOverPlusThreePercentage := result.FullStat.EarlyOverUnder.Over * 1.03
	earlyOverMinusThreePercentage := result.FullStat.EarlyOverUnder.Over * 0.97

	earlyUnderPlusThreePercentage := result.FullStat.EarlyOverUnder.Under * 1.03
	earlyUnderMinusThreePercentage := result.FullStat.EarlyOverUnder.Under * 0.97

	startHomePlusThreePercentage := result.FullStat.StartHandicap.Home * 1.03
	startHomeMinusThreePercentage := result.FullStat.StartHandicap.Home * 0.97

	startAwayPlusThreePercentage := result.FullStat.StartHandicap.Away * 1.03
	startAwayMinusThreePercentage := result.FullStat.StartHandicap.Away * 0.97

	startOverPlusThreePercentage := result.FullStat.StartOverUnder.Over * 1.03
	startOverMinusThreePercentage := result.FullStat.StartOverUnder.Over * 0.97

	startUnderPlusThreePercentage := result.FullStat.StartOverUnder.Under * 1.03
	startUnderMinusThreePercentage := result.FullStat.StartOverUnder.Under * 0.97

	stmt := jet.SELECT(
		table.ThreeInOneResults.AllColumns,
	).
		FROM(
			table.ThreeInOneResults,
		).
		WHERE(
			table.ThreeInOneResults.EarlyHandicapHomeFull.GT_EQ(jet.Float(float64(earlyHomeMinusThreePercentage))).
				AND(table.ThreeInOneResults.EarlyHandicapHomeFull.LT_EQ(jet.Float(float64(earlyHomePlusThreePercentage)))).
				AND(table.ThreeInOneResults.EarlyHandicapHandicapFull.EQ(jet.String(result.FullStat.EarlyHandicap.Handicap))).
				AND(table.ThreeInOneResults.EarlyHandicapAwayFull.GT_EQ(jet.Float(float64(ealyAwayMinusThreePercentage)))).
				AND(table.ThreeInOneResults.EarlyHandicapAwayFull.LT_EQ(jet.Float(float64(ealyAwayPlusThreePercentage)))).
				AND(table.ThreeInOneResults.EarlyOverUnderOverFull.GT_EQ(jet.Float(float64(earlyOverMinusThreePercentage)))).
				AND(table.ThreeInOneResults.EarlyOverUnderOverFull.LT_EQ(jet.Float(float64(earlyOverPlusThreePercentage)))).
				AND(table.ThreeInOneResults.EarlyOverUnderOUFull.EQ(jet.String(result.FullStat.EarlyOverUnder.OU))).
				AND(table.ThreeInOneResults.EarlyOverUnderUnderFull.GT_EQ(jet.Float(float64(earlyUnderMinusThreePercentage)))).
				AND(table.ThreeInOneResults.EarlyOverUnderUnderFull.LT_EQ(jet.Float(float64(earlyUnderPlusThreePercentage)))).
				AND(table.ThreeInOneResults.StartHandicapHomeFull.GT_EQ(jet.Float(float64(startHomeMinusThreePercentage)))).
				AND(table.ThreeInOneResults.StartHandicapHomeFull.LT_EQ(jet.Float(float64(startHomePlusThreePercentage)))).
				AND(table.ThreeInOneResults.StartHandicapHandicapFull.EQ(jet.String(result.FullStat.StartHandicap.Handicap))).
				AND(table.ThreeInOneResults.StartHandicapAwayFull.GT_EQ(jet.Float(float64(startAwayMinusThreePercentage)))).
				AND(table.ThreeInOneResults.StartHandicapAwayFull.LT_EQ(jet.Float(float64(startAwayPlusThreePercentage)))).
				AND(table.ThreeInOneResults.StartOverUnderOverFull.GT_EQ(jet.Float(float64(startOverMinusThreePercentage)))).
				AND(table.ThreeInOneResults.StartOverUnderOverFull.LT_EQ(jet.Float(float64(startOverPlusThreePercentage)))).
				AND(table.ThreeInOneResults.StartOverUnderOUFull.EQ(jet.String(result.FullStat.StartOverUnder.OU))).
				AND(table.ThreeInOneResults.StartOverUnderUnderFull.GT_EQ(jet.Float(float64(startUnderMinusThreePercentage)))).
				AND(table.ThreeInOneResults.StartOverUnderUnderFull.LT_EQ(jet.Float(float64(startUnderPlusThreePercentage)))),
		)

	var dest []model.ThreeInOneResults

	err := stmt.Query(r.DB, &dest)
	if err != nil {
		logrus.Errorln(err)
		return []domain.ResultInformation{}, err
	}
	results := []domain.ResultInformation{}
	for _, v := range dest {
		results = append(results, domain.ResultInformation{
			ResultID:        v.ResultID,
			LeftTeam:        v.LeftTeam,
			RightTeam:       v.RightTeam,
			FirstHalfScore:  v.FirstHalfScore,
			SecondHalfScore: v.SecondHalfScore,
			FinalScore:      v.FinalScore,
			HalfStat: domain.Stat{
				EarlyHandicap: domain.Handicap{
					Home:     v.EarlyHandicapHomeHalf,
					Handicap: v.EarlyHandicapHandicapHalf,
					Away:     v.EarlyHandicapAwayHalf,
				},
				EarlyOverUnder: domain.OverUnder{
					Over:  v.EarlyOverUnderOverHalf,
					OU:    v.EarlyOverUnderOUHalf,
					Under: v.EarlyOverUnderUnderHalf,
				},
				StartHandicap: domain.Handicap{
					Home:     v.StartHandicapHomeHalf,
					Handicap: v.StartHandicapHandicapHalf,
					Away:     v.StartHandicapAwayHalf,
				},
				StartOverUnder: domain.OverUnder{
					Over:  v.StartOverUnderOverHalf,
					OU:    v.StartOverUnderOUHalf,
					Under: v.StartOverUnderUnderHalf,
				},
			},
			FullStat: domain.Stat{
				EarlyHandicap: domain.Handicap{
					Home:     v.EarlyHandicapHomeFull,
					Handicap: v.EarlyHandicapHandicapFull,
					Away:     v.EarlyHandicapAwayFull,
				},
				EarlyOverUnder: domain.OverUnder{
					Over:  v.EarlyOverUnderOverFull,
					OU:    v.EarlyOverUnderOUFull,
					Under: v.EarlyOverUnderUnderFull,
				},
				StartHandicap: domain.Handicap{
					Home:     v.StartHandicapHomeFull,
					Handicap: v.StartHandicapHandicapFull,
					Away:     v.StartHandicapAwayFull,
				},
				StartOverUnder: domain.OverUnder{
					Over:  v.StartOverUnderOverFull,
					OU:    v.StartOverUnderOUFull,
					Under: v.StartOverUnderUnderFull,
				},
			},
		})
	}
	return results, nil
}
