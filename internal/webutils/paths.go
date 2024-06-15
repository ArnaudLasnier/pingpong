package webutils

type PathKey string

func (pathKey PathKey) String() string {
	return string(pathKey)
}

func (pathKey PathKey) DynamicSegment() string {
	return "{" + string(pathKey) + "}"
}
