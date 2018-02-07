package endpoint

type NewSiteRequest struct {
	Email		string
	SiteName	string
}

type NewSiteResponse struct {
	Err		error 	`json:"-"`
}
