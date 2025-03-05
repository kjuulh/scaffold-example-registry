package externalhttp_test

import (
	"testing"

	"github.com/kjuulh/scaffold/internal/tests"
)

func TestScaffold(t *testing.T) {
	tests.
		Test(t, "scaffold").
		ScaffoldTest("scaffold itself",
			func(fixture *tests.ScaffoldFixture) {
				fixture.WithVariable("name", "some-name")
			},
		)
}
