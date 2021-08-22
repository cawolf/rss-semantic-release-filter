package main

import (
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	Feeds []FeedConfiguration `yaml:"feeds" validate:"required,min=1,dive"`
}

type FeedConfiguration struct {
	Url          string           `yaml:"url" validate:"required,url"`
	MinimumLevel MinimumLevelType `yaml:"minimum_level" validate:"required,oneof=major minor patch"`
}

type MinimumLevelType string

const (
	Major MinimumLevelType = "major"
	Minor MinimumLevelType = "minor"
	Patch MinimumLevelType = "patch"
)

func (configuration *Configuration) Read(directory string) *Configuration {
	yamlFile, err := ioutil.ReadFile(directory + "config.yaml")
	if err != nil {
		logger.Fatalf("no configuration file found - '%v'", err)
	}
	err = yaml.Unmarshal(yamlFile, configuration)
	if err != nil {
		logger.Fatalf("could not parse configuration file - '%v'", err)
	}

	configuration.Validate()

	logger.Debug(configuration)

	return configuration
}

func (configuration *Configuration) Validate() {
	v := validator.New()
	err := v.Struct(configuration)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			switch e.ActualTag() {
			case "required":
				logger.Errorf("field '%s' is required, but is missing", e.Field())
			case "url":
				logger.Errorf("field '%s' does not contain a valid URL: %s", e.Field(), e.Value())
			case "oneof":
				logger.Errorf("field '%s' must be one of [%s], but is: %s", e.Field(), e.Param(), e.Value())
			case "min":
				logger.Errorf("field '%s' must contain at least one feed entry", e.Field())
			}
		}
		logger.Fatal("configuration file is invalid")
	}
}
