package config

import "path/filepath"

var (
	// PORT is the serving port, this value is set during compile time
	PORT = "unset"
	// DATADIR is the root directory for persistent volume
	DATADIR = "/data"
	// SITESDIR contains all user sites
	SITESDIR = filepath.Join(DATADIR, "sites")
)
