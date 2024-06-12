package webutils

type PathKey string

func (pathValue PathKey) String() string {
	return string(pathValue)
}

func (pathValue PathKey) DynamicSegment() string {
	return "{" + string(pathValue) + "}"
}
