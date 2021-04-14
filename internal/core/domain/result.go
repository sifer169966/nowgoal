package domain

type ResultInformation struct {
	ResultID        string `json:"result_id"`
	LeftTeam        string `json:"left_team"`
	RightTeam       string `json:"right_team"`
	FirstHalfScore  string `json:"first_half_score"`
	SecondHalfScore string `json:"second_half_score"`
	FinalScore      string `json:"final_score"`
	HalfStat        Stat   `json:"half_stat"`
	FullStat        Stat   `json:"full_stat"`
}

type GetStatsPattern1Request struct {
	Early HandicapOverUnder `json:"early"`
	Start HandicapOverUnder `json:"start"`
}

type GetStatsPattern1Response struct {
	LenOverResult int                 `json:"len_over_result"`
	LenLessResult int                 `json:"len_less_result"`
	OverResults   []ResultInformation `json:"over_results"`
	LessResults   []ResultInformation `json:"less_resutls"`
}

type HandicapOverUnder struct {
	Handicap  Handicap  `json:"handicap"`
	OverUnder OverUnder `json:"over_under"`
}

// first half ...
type Stat struct {
	EarlyHandicap  Handicap  `json:"early_handicap"`
	EarlyOverUnder OverUnder `json:"early_over_under"`
	StartHandicap  Handicap  `json:"start_handicap"`
	StartOverUnder OverUnder `json:"start_over_under"`
}
type Handicap struct {
	Home     float32 `json:"home"`
	Handicap string  `json:"handicap"`
	Away     float32 `json:"away"`
}
type OverUnder struct {
	Over  float32 `json:"over"`
	OU    string  `json:"ou"`
	Under float32 `json:"under"`
}

// handicap
// 0
// 0/0.5 = 0.25
// 0.5 = 0.5
// 0.5/1 = 0.75

// 1 = 1
// 1/1.5 = 1.25
// 1.5 =
// 1.5/2

// 2/2.5 2.25
// 2.5

// diff home, away 2%
// diff over, under 2%
