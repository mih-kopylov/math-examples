package main

import (
	"flag"
	"maps"

	"github.com/joomcode/errorx"
)

var (
	ErrProfileNamespace = errorx.NewNamespace("profile")
	ErrNoProfileName    = errorx.NewType(ErrProfileNamespace, "NoProfileName")
	ErrProfileNotFound  = errorx.NewType(ErrProfileNamespace, "ProfileNotFound")
)

func ReadProfile(printer Printer) (*ProfileParams, error) {
	app, err := ReadParams()
	if err != nil {
		return nil, err
	}
	profileName := flag.String("p", "", "Profile to use")
	flag.Parse()

	if *profileName == "" {
		return nil, ErrNoProfileName.New("Profile name is not pass. Use -p flag")
	}

	profile, exists := app.Profiles[*profileName]
	if !exists {
		return nil, ErrProfileNotFound.New(
			"no profile found %v; known profiles are: %v", *profileName, maps.Keys(app.Profiles),
		)
	}

	printer.Println("Добро пожаловать, %v!", *profileName)
	return &profile, nil
}
