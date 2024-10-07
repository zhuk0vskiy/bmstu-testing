package repository_test

import (
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"sync"
	"testing"
)

func Test_Runner(t *testing.T) {
	//db, ids := utils.NewTestStorage()
	//defer utils.DropTestStorage(db)
	//

	t.Parallel()

	wg := &sync.WaitGroup{}
	suits := []runner.TestSuite{
		&UserRepositorySuite{},
		&ProducerRepositorySuite{},
		&InstrumentalistRepositorySuite{},
		&RoomRepositorySuite{},
		&EquipmentRepositorySuite{},
		&ReserveRepositorySuite{},
		&StudioRepositorySuite{},
	}
	wg.Add(len(suits))

	for _, s := range suits {
		go func() {
			suite.RunSuite(t, s)
			wg.Done()
		}()
	}

	wg.Wait()
}
