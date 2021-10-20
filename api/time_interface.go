package api

import "github/vonuki/template_api/api/api_exemplar"

type TimeAPI interface {
	GetTime(string) (string, error)
}

func GetCurrentTimeAPI() TimeAPI {
	return api_exemplar.NewBestTimeAPI()
}
