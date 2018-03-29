package endpoint

// NewSiteRequest collects the request parameters for the NewSite method.
type NewSiteRequest struct {
	SiteName string `json:"site_name"`
}

// NewSiteResponse collects the response values for the NewSite method.
type NewSiteResponse struct {
	SiteID uint  `json:"site_id"`
	Err    error `json:"-"`
}

// DeleteSiteRequest collects the request parameters for the DeleteSite method.
type DeleteSiteRequest struct {
	SiteID uint `json:"site_id"`
}

// DeleteSiteResponse collects the response values for the DeleteSite method.
type DeleteSiteResponse struct {
	Err error `json:"-"`
}

// CheckSitenameExistsRequest collects the request parameters for the CheckSitenameExists method.
type CheckSitenameExistsRequest struct {
	Sitename string `json:"sitename"`
}

// CheckSitenameExistsResponse collects the response values for the CheckSitenameExists method.
type CheckSitenameExistsResponse struct {
	Exists bool  `json:"exists"`
	Err    error `json:"-"`
}

// GetSiteIDByUserIDRequest collects the request parameters for the GetSiteIDByUserID method.
type GetSiteIDByUserIDRequest struct {
}

// GetSiteIDByUserIDResponse collects the response values for the GetSiteIDByUserID method.
type GetSiteIDByUserIDResponse struct {
	SiteID uint  `json:"site_id"`
	Err    error `json:"-"`
}

// GetConfigRequest collects the request parameters for the GetConfig method.
type GetConfigRequest struct {
	SiteID uint `json:"site_id"`
}

// GetConfigResponse collects the response values for the GetConfig method.
type GetConfigResponse struct {
	Config string `json:"config"`
	Err    error  `json:"-"`
}

// UpdateConfigRequest collects the request parameters for the UpdateConfig method.
type UpdateConfigRequest struct {
	SiteID uint   `json:"site_id"`
	Config string `json:"config"`
}

// UpdateConfigResponse collects the response values for the UpdateConfig method.
type UpdateConfigResponse struct {
	Err error `json:"-"`
}
