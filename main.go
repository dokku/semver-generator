package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	go_semver "github.com/coreos/go-semver/semver"
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

	parsed, err := go_semver.NewVersion(*input)
	if err != nil {
		slog.Error("Invalid version", "version", *input, "error", err.Error())
		os.Exit(1)
	}

	switch *bump {
	case "major":
		parsed.Major += 1
	case "minor":
		parsed.Minor += 1
	case "patch":
		parsed.Patch += 1
	}

	if addPrefix {
		fmt.Printf("v" + parsed.String() + "\n")
	} else {
		fmt.Printf(parsed.String() + "\n")
	}
}
