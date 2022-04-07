package config

// CurrentConfiguration represents currently running configuration
var CurrentConfiguration = configuration{
	routes: make(map[string]Route, 0),
}
