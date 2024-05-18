package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/coreos/go-semver/semver"
	flag "github.com/spf13/pflag"
)

func main() {
	// parse the flags via pflags
	input := flag.StringP("input", "i", "", "input version")
	bump := flag.StringP("bump", "b", "", "bump version")
	flag.Parse()

	if input == nil || *input == "" {
		slog.Error("Input version is required")
		os.Exit(1)
	}

	if bump == nil || *bump == "" {
		slog.Error("Bump type is required")
		os.Exit(1)
	}

	validBumps := map[string]bool{
		"major": true,
		"minor": true,
		"patch": true,
	}

	if !validBumps[*bump] {
		slog.Error("Invalid bump type", "bump", *bump)
		os.Exit(1)
	}

	addPrefix := false
	if strings.HasPrefix(*input, "v") {
		*input = (*input)[1:]
		addPrefix = true
	}

	parsed, err := semver.NewVersion(*input)
	if err != nil {
		slog.Error("Invalid version", "version", *input, "error", err.Error())
		os.Exit(1)
	}

	switch *bump {
	case "major":
		parsed.Major += 1
		parsed.Minor = 0
		parsed.Patch = 0
	case "minor":
		parsed.Minor += 1
		parsed.Patch = 0
	case "patch":
		parsed.Patch += 1
	}

	updatedVersion := parsed.String()
	if addPrefix {
		updatedVersion = "v" + parsed.String()
	}

	fmt.Printf(updatedVersion + "\n")
	githubOutput := os.Getenv("GITHUB_OUTPUT")
	if githubOutput != "" {
		f, err := os.OpenFile(githubOutput, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			slog.Error("Failed to open file", "file", githubOutput, "error", err.Error())
			os.Exit(1)
		}
		defer f.Close()

		if _, err = f.WriteString(fmt.Sprintf("version=%s", updatedVersion)); err != nil {
			slog.Error("Failed to write to file", "file", githubOutput, "error", err.Error())
			os.Exit(1)
		}
	}
}
