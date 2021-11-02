package metrics


type PageViewModelValidator struct {
	PageView struct {
		Url string `form:"url" json:"url"`
	} `json:"pageView"`
	pageViewModel PageViewModel `json:"-"`
}

func NewPageViewModelValidator() PageViewModelValidator {
	return PageViewModelValidator{}
}