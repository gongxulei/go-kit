package container

import (
	"errors"
	"fmt"
	"sync"
)

type Container interface {
	Bind(abstract string, entity interface{})
	Invoke(abstract string)
}

type ServiceContainer struct {
	Bindings  map[string]interface{}
	lockMutex sync.RWMutex
}

func (c *ServiceContainer) Bind(abstract string, entity interface{}) {
	c.lockMutex.Lock()
	defer c.lockMutex.Unlock()
	c.Bindings[abstract] = entity
}

func (c *ServiceContainer) Invoke(abstract string) (entity interface{}, err error) {
	c.lockMutex.RLock()
	defer c.lockMutex.RUnlock()
	var ok bool
	entity, ok = c.Bindings[abstract]
	if !ok {
		err = errors.New(fmt.Sprintf("abstract: %s not exist", abstract))
	}
	return
}
