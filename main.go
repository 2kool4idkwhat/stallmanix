package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

var (
	totalPkgs    int
	totalGnuPkgs int
)

func isIgnoredPkg(pkg string) bool {
	for _, ignoredPkg := range ignoreList {
		if strings.Contains(pkg, ignoredPkg) {
			slog.Debug(fmt.Sprintf("Ignoring %s because it contains %s", pkg, ignoredPkg))
			return true
		}
	}

	return false
}

func isGnuPkg(pkg string) bool {
	for _, gnuPkg := range listOfGnuPkgs {
		if strings.Contains(pkg, "-"+gnuPkg+"-") {
			slog.Info(fmt.Sprintf("Counting %s as %s", pkg, gnuPkg))
			return true
		}
	}

	return false
}

func logLevel(logLevelFlag string) slog.Level {
	switch logLevelFlag {
	case "warn":
		return slog.LevelWarn
	case "info":
		return slog.LevelInfo
	case "debug":
		return slog.LevelDebug
	default:
		slog.Error(fmt.Sprintf("%s isn't a valid log level value", logLevelFlag))
		os.Exit(1)
		return slog.LevelWarn
	}
}

func percentOf(part int, total int) float64 {
	return (float64(part) * float64(100)) / float64(total)
}

func main() {
	var profilePath string
	flag.StringVar(&profilePath, "path", "/run/current-system", "the path of the nix profile to check")

	var logLevelFlag string
	flag.StringVar(&logLevelFlag, "log-level", "warn", "log level. Possible values: warn, info, debug")

	flag.Parse()

	slog.SetLogLoggerLevel(logLevel(logLevelFlag))

	rawStorePaths, err := exec.Command("nix", "path-info", "-r", profilePath).Output()
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't get the list of store paths: %v", err))
		os.Exit(1)
	}
	storePaths := strings.SplitSeq((string(rawStorePaths)), "\n")

	for path := range storePaths {
		if isIgnoredPkg(path) {
			continue
		}
		totalPkgs++

		if isGnuPkg(path) {
			totalGnuPkgs++
		}
	}

	fmt.Printf("Around %d out of %d (%.2f%%) packages in your closure are GNU software\n", totalGnuPkgs, totalPkgs, percentOf(totalGnuPkgs, totalPkgs))
}
