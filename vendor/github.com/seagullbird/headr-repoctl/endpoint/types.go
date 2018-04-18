package endpoint

// NewSiteRequest collects the request parameters for the NewSite method.
type NewSiteRequest struct {
	SiteID uint   `json:"site_id"`
	Theme  string `json:"theme"`
}

// NewSiteResponse collects the response values for the NewSite method.
type NewSiteResponse struct {
	Err error `json:"-"`
}

// DeleteSiteRequest collects the request parameters for the DeleteSite method.
type DeleteSiteRequest struct {
	SiteID uint `json:"site_id"`
}

// DeleteSiteResponse collects the response values for the DeleteSite method.
type DeleteSiteResponse struct {
	Err error `json:"-"`
}

// WritePostRequest collects the request parameters for the WritePost method.
type WritePostRequest struct {
	SiteID   uint   `json:"site_id"`
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

// WritePostResponse collects the response values for the WritePost method.
type WritePostResponse struct {
	Err error `json:"-"`
}

// RemovePostRequest collects the request parameters for the RemovePost method.
type RemovePostRequest struct {
	SiteID   uint   `json:"site_id"`
	Filename string `json:"filename"`
}

// RemovePostResponse collects the response values for the RemovePost method.
type RemovePostResponse struct {
	Err error `json:"-"`
}

// ReadPostRequest collects the request parameters for the ReadPost method.
type ReadPostRequest struct {
	SiteID   uint   `json:"site_id"`
	Filename string `json:"filename"`
}

// ReadPostResponse collects the response values for the ReadPost method.
type ReadPostResponse struct {
	Content string `json:"content"`
	Err     error  `json:"-"`
}

// WriteConfigRequest collects the request parameters for the WriteConfig method.
type WriteConfigRequest struct {
	SiteID uint   `json:"site_id"`
	Config string `json:"config"`
}

// WriteConfigResponse collects the response values for the WriteConfig method.
type WriteConfigResponse struct {
	Err error `json:"-"`
}

// ReadConfigRequest collects the request parameters for the ReadConfig method.
type ReadConfigRequest struct {
	SiteID uint `json:"site_id"`
}

// ReadConfigResponse collects the response values for the ReadConfig method.
type ReadConfigResponse struct {
	Config string `json:"config"`
	Err    error  `json:"-"`
}

// UpdateAboutRequest collects the request parameters for the UpdateAbout method.
type UpdateAboutRequest struct {
	SiteID  uint   `json:"site_id"`
	Content string `json:"content"`
}

// UpdateAboutResponse collects the response values for the UpdateAbout method.
type UpdateAboutResponse struct {
	Err error `json:"-"`
}

// ReadAboutRequest collects the request parameters for the ReadAbout method.
type ReadAboutRequest struct {
	SiteID uint `json:"site_id"`
}

// ReadAboutResponse collects the response values for the ReadAbout method.
type ReadAboutResponse struct {
	Content string `json:"content"`
	Err     error  `json:"-"`
}

// ChangeDefaultConfigRequest collects the request parameters for the ChangeDefaultConfig method.
type ChangeDefaultConfigRequest struct {
	SiteID uint   `json:"site_id"`
	Theme  string `json:"theme"`
}

// ChangeDefaultConfigResponse collects the response values for the ChangeDefaultConfig method.
type ChangeDefaultConfigResponse struct {
	Err error `json:"-"`
}
