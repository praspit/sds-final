package common

type Response struct {
	Data string `json:"data"`
}

type ListResponse struct {
	Data []string `json:"data"`
}

type MapResponse struct {
	Data map[string]string `json:"data"`
}

type MapOfMapResponse struct {
	Data map[string]map[string]string `json:"data"`
}
