package endpoint

type NewSiteRequest struct {
	Email		string	`json:"email"`
	SiteName	string	`json:"site_name"`
}

type NewSiteResponse struct {
	Err		error 	`json:"-"`
}
