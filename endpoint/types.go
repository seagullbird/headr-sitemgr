package endpoint

type NewSiteRequest struct {
	UserId   uint   `json:"user_id"`
	SiteName string `json:"site_name"`
}

type NewSiteResponse struct {
	SiteId uint  `json:"site_id"`
	Err    error `json:"-"`
}

type DeleteSiteRequest struct {
	SiteId uint `json:"site_id"`
}

type DeleteSiteResponse struct {
	Err error `json:"-"`
}
