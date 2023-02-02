package controller

import (
	"sync"

	"github.com/asdine/storm/v3"
)

type (
	Controller struct {
		port string
		db   *storm.DB
		sync.Mutex
	}
)

func NewController(port string, db *storm.DB) *Controller {
	ctrl := Controller{
		port: port,
		db:   db,
	}

	return &ctrl
}

func (c *Controller) Load() error {
	c.Lock()
	defer c.Unlock()

	// do the magic during loads

	return nil
}

// Start is non-blocking
func (c *Controller) Start() {
	c.Lock()
	defer c.Unlock()

	c.start()
}

func (c *Controller) start() {

	// call any additionl handler -> must be non-blocking
	//go c.XXXX.ListenAndServ()
}

// Stop is non-blocking
func (c *Controller) Stop() {
	c.Lock()
	defer c.Unlock()

	c.stop()
}

func (c *Controller) stop() {

	// shutdown all additional handler -> must be non-blocking
	//c.XXXX.Shutdown()
}
func (c *Controller) Port() string {
	return c.port
}
func (c *Controller) Db() *storm.DB {
	return c.db
}
