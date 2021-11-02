package metrics

import "github.com/gin-gonic/gin"

type PageViewSerializer struct {
	C	*gin.Context
	PageViewModel
}

type PageViewsSerializer struct {
	C	*gin.Context
	PageViews []PageViewModel
}

type PageViewResponse struct {
	ID			uint		`json:"-"`
	Url			string		`json:"url"`
	CreatedAt	string		`json:"createdAt"`
}

func (s *PageViewSerializer) Response() PageViewResponse {
	response := PageViewResponse{
		ID:			s.ID,
		Url:		s.url,
		CreatedAt: 	s.CreatedAt.UTC().Format("2021-01-02T15:04:05.999Z"),
	}
	return response
}

func (s *PageViewsSerializer) Response() []PageViewResponse {
	response := []PageViewResponse{}
	for _, pageView := range s.PageViews {
		serializer := PageViewSerializer{s.C, pageView}
		response = append(response, serializer.Response())
	}
	return response
}
