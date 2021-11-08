package metrics

import "github.com/gin-gonic/gin"

type HitSerializer struct {
	C	*gin.Context
	Hit
}

type HitsSerializer struct {
	C	*gin.Context
	Hits []Hit
}

type HitResponse struct {
	ID			uint		`json:"id"`
	Url			string		`json:"url"`
	CreatedAt	string		`json:"createdAt"`
}

func (s *HitSerializer) Response() HitResponse {
	response := HitResponse{
		ID:			s.ID,
		Url:		s.Url,
		CreatedAt: 	s.CreatedAt.UTC().Format("2021-01-02T15:04:05.999Z"),
	}
	return response
}

func (s *HitsSerializer) Response() []HitResponse {
	response := []HitResponse{}
	for _, hit := range s.Hits {
		serializer := HitSerializer{s.C, hit}
		response = append(response, serializer.Response())
	}
	return response
}
