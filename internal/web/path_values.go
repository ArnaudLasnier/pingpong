package web

type pathValue string

func (pathValue pathValue) String() string {
	return string(pathValue)
}

func (pathValue pathValue) Segment() string {
	return "{" + string(pathValue) + "}"
}

const (
	tournamentID pathValue = "tournamentID"
)
