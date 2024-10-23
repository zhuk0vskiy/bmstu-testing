//go:build unit_test

package unit_tests

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
		&AuthSuite{
			JwtKey: "asdhjkashjks",
		},
		&AuthRepoSuite{},

		&CommentSuite{},
		&CommentRepoSuite{},

		&SaladTypeSuite{},
		&SaladTypeRepoSuite{},

		&RecipeSuite{},
		&RecipeRepoSuite{},

		&KeywordsValidatorSuite{},
		&KeywordsRepoSuite{},

		&MeasurementSuite{},
		&MeasurementRepoSuite{},

		&RecipeStepInteractorSuite{},
		&RecipeStepSuite{},
		&RecipeStepRepoSuite{},

		&IngredientTypeSuite{},
		&IngredientTypeRepoSuite{},

		&SaladInteractorSuite{},
		&SaladSuite{},
		&SaladRepoSuite{},

		&UserSuite{},
		&UserRepoSuite{},

		&UrlValidatorSuite{},
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
