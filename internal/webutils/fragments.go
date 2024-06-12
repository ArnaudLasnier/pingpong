package webutils

import "path"

type Fragment string

func (fragment Fragment) String() string {
	return string(fragment)
}

func (fragment Fragment) IDSelector() string {
	return "#" + string(fragment)
}

func (fragment Fragment) Endpoint() string {
	return "/" + string(fragment)
}

func (fragment Fragment) GetEndpoint() string {
	return "GET /" + string(fragment)
}

func (fragment Fragment) GetEndpointWithPathValues(values ...PathKey) string {
	var dynamicSegments []string
	for _, value := range values {
		dynamicSegments = append(dynamicSegments, value.DynamicSegment())
	}
	suffix := path.Join(dynamicSegments...)
	return "GET /" + string(fragment) + "/" + suffix
}

func (fragment Fragment) PostEndpoint() string {
	return "POST /" + string(fragment)
}
