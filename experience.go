package tisp

type ExperienceType string

const (
	Junior      ExperienceType = "junior"
	Middle      ExperienceType = "middle"
	Senior      ExperienceType = "senior"
	SuperSenior ExperienceType = "super-senior"
	God         ExperienceType = "god"
)

var experienceIndexes = map[uint]ExperienceType{
	0: Junior,
	1: Middle,
	2: Senior,
	3: SuperSenior,
	4: God,
}

func (exp *ExperienceType) Nerf() {
	var i uint
	var e ExperienceType
	for i, e = range experienceIndexes {
		if e == *exp {
			break
		}
	}

	if i == 0 {
		*exp = experienceIndexes[i]
	} else {
		*exp = experienceIndexes[i-1]
	}
}

func (exp *ExperienceType) Promote() {
	var i uint
	var e ExperienceType
	for i, e = range experienceIndexes {
		if e == *exp {
			break
		}
	}

	if i == uint(len(experienceIndexes)-1) {
		*exp = experienceIndexes[i]
	} else {
		*exp = experienceIndexes[i-1]
	}
}
