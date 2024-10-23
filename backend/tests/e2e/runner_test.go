//go:build e2e_test

package e2e

import (
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"sync"
	"testing"
)

func TestRunner(t *testing.T) {
	t.Parallel()

	wg := &sync.WaitGroup{}
	suits := []runner.TestSuite{
		&E2ESuite{},
	}
	wg.Add(len(suits))

	for _, s := range suits {
		go func(s runner.TestSuite) {
			suite.RunSuite(t, s)
			wg.Done()
		}(s)
	}

	wg.Wait()
}
