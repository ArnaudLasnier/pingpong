package web

type pathValue string

func (pathValue pathValue) String() string {
	return string(pathValue)
}

func (pathValue pathValue) DynamicSegment() string {
	return "{" + string(pathValue) + "}"
}

const (
	tournamentID pathValue = "tournamentID"
	playerID     pathValue = "playerID"
)
