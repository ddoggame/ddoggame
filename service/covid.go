package service

type CovidSummaryResponse struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}

var ageGroupBoundaries = []int{30, 60}

const (
	AgeGroup0to30  = "0-30"
	AgeGroup31to60 = "31-60"
	AgeGroup60plus = "60+"
	AgeGroupNA     = "N/A"
)

func GetAgeGroup(age int) string {
	if age >= 0 && age <= 30 {
		return AgeGroup0to30
	} else if age >= 31 && age <= 60 {
		return AgeGroup31to60
	} else if age > 60 {
		return AgeGroup60plus
	} else {
		return AgeGroupNA
	}
}
