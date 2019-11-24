package main

import (
	"flag"
	"os"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	tests "github.com/kapustkin/go_guard/pkg/integration-tests"
)

func TestMain(m *testing.M) {
	flag.Parse()

	status := godog.RunWithOptions("integration", func(s *godog.Suite) {
		tests.FeatureContext(s)
	}, godog.Options{
		Output:    colors.Colored(os.Stdout),
		Format:    "pretty",
		Paths:     []string{"features"},
		Randomize: 0,
	})

	if st := m.Run(); st > status {
		status = st
	}
	//nolint:wsl
	os.Exit(status)
}
