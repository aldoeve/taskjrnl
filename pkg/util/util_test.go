package util

import (
	"flag"
	"os"
	"taskjrnl/internal/config"
	"taskjrnl/internal/consts"
	"taskjrnl/internal/schema"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ArgsAfterKeyword_onlyKeywordPresent(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	additionalArgs := []string{"keyword"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{consts.AppName}
	os.Args = append(os.Args, additionalArgs...)

	flag.Parse()

	parsed := ArgsAfterKeyword()
	expected := []string{}

	assert.Equal(t, expected, parsed)
}

func Test_ArgsAfterKeyword_keywordWithTokens(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	additionalArgs := []string{"keyword", "token1"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{consts.AppName}
	os.Args = append(os.Args, additionalArgs...)

	flag.Parse()

	parsed := ArgsAfterKeyword()
	expected := additionalArgs[1:]

	assert.Equal(t, expected, parsed)
}

func Test_CalculateImportance_default(t *testing.T) {
	currentTime := time.Now()
	priority := "L"

	task := schema.Tasks{
		Id:                 0,
		Name:               consts.AppName,
		Tag:                nil,
		DateCreated:        currentTime.String(),
		Priority:           &priority,
		ImportanceVariance: 0,
	}

	value := CalculateImportance(&task)

	assert.Equal(t, 0, value)
}

func Test_CalculateImportance_24HrsPassed(t *testing.T) {
	fiveHoursAgo := time.Now().Add(-24 * time.Hour)
	priority := "L"

	task := schema.Tasks{
		Id:                 0,
		Name:               consts.AppName,
		Tag:                nil,
		DateCreated:        fiveHoursAgo.Format(config.TimeFormat),
		Priority:           &priority,
		ImportanceVariance: 0,
	}

	value := CalculateImportance(&task)

	assert.Equal(t, 1, value)
}

func Test_CalculateImportance_differentPriorites(t *testing.T) {
	currentTime := time.Now()
	priority := "L"

	task := schema.Tasks{
		Id:                 0,
		Name:               consts.AppName,
		Tag:                nil,
		DateCreated:        currentTime.Format(config.TimeFormat),
		Priority:           &priority,
		ImportanceVariance: 0,
	}

	value := CalculateImportance(&task)

	assert.Equal(t, 0, value)

	priority = "M"

	value = CalculateImportance(&task)

	assert.Equal(t, 1000, value)

	priority = "H"

	value = CalculateImportance(&task)

	assert.Equal(t, 2000, value)
}

func Test_CalculateImportance_variance(t *testing.T) {
	currentTime := time.Now()
	priority := "L"

	task := schema.Tasks{
		Id:                 0,
		Name:               consts.AppName,
		Tag:                nil,
		DateCreated:        currentTime.Format(config.TimeFormat),
		Priority:           &priority,
		ImportanceVariance: 3234989234,
	}

	value := CalculateImportance(&task)

	assert.Equal(t, 3234989234, value)
}

func Test_CalculateImportance_everything(t *testing.T) {
	threeDaysAgo := time.Now().Add((-24 * 3) * time.Hour)
	priority := "M"

	task := schema.Tasks{
		Id:                 0,
		Name:               consts.AppName,
		Tag:                nil,
		DateCreated:        threeDaysAgo.Format(config.TimeFormat),
		Priority:           &priority,
		ImportanceVariance: 21,
	}

	value := CalculateImportance(&task)

	assert.Equal(t, 1024, value)
}
