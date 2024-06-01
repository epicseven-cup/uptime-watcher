package watcher

import (
	"time"
	"github.com/labstack/echo/v4"
)

type Watcher struct {
	path     map[string]time.Ticker
	database *Database
	logger   log.Log
}

func NewWatcher() *Watcher {
	path := map[string]time.Ticker{}
	return &Watcher{}
}

// Adds the path to router's path which checks if the path is avaaible to the user
// this should return the success of the method
func (watcher *Watcher) addWatchPoint(watchPoint *WatchPoint) bool {
	return router.database.InsertWatchPoint(&watchPoint)
}

// Remove the path that was added to the router
func (watcher *Watcher) removePath(path string) bool {}
