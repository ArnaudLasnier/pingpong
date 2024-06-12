package web

type pathKey string

func (pathValue pathKey) String() string {
	return string(pathValue)
}

func (pathValue pathKey) DynamicSegment() string {
	return "{" + string(pathValue) + "}"
}

const (
	tournamentID pathKey = "tournamentID"
	playerID     pathKey = "playerID"
)
