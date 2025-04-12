package main

var ignoreList = []string{
	"-completions", // shell completions
	"-rule",        // opensnitch rules

	// systemd stuff
	".service",
	".automount",
	".device",
	".mount",
	".nspawn",
	".path",
	".preset",
	".scope",
	".slice",
	".socket",
	".swap",
	".target",
	".timer",
	"X-Restart-Triggers",
	"-unit-",
}
