package cats

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	nest "github.com/vanyastar/nest"
)

type CatsService struct {
	mu   sync.Mutex
	name string
}

var GetCatsService = &CatsService{
	// Initiate cat name
	name: "Tom",
}

func (this *CatsService) GetName() string {
	return "My name is " + this.name
}

func (this *CatsService) SetName(name string) error {
	if len(name) <= 0 {
		return errors.New("Name is empty")
	}
	this.mu.Lock()
	defer this.mu.Unlock()
	this.name = name
	return nil
}

func (this *CatsService) pushNotification(c *nest.Ctx) error {
	// If you need to close the connection use this cancel()
	ctx, cancel := context.WithCancel(c.Req().Context())
	c.Req().Request = c.Req().WithContext(ctx)

	testMap := map[string]string{
		"Name": this.GetName(),
	}

	workerDone := make(chan bool)
	go func(workerDone chan<- bool) {
		for i := 0; i <= 10; i++ {
			// Simulate some work
			if err := c.Send(testMap); err != nil {
				break
			}
			if err := c.Flush(); err != nil {
				break
			}
			time.Sleep(time.Second * 1)
		}
		cancel()
		workerDone <- true
	}(workerDone)

	<-c.Req().Context().Done()
	// Your logic after disconnection

	<-workerDone
	return c.Error(http.StatusOK, "") // or nil
}
