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


// This just feels like a curl app dam it

// Adds the path to router's path which checks if the path is avaaible to the user
// this should return the success of the method
func (watcher *Watcher) AddWatchPoint(watchPoint *WatchPoint) {
	watcher.database.InsertWatchPoint(watchPoint)
}

// Remove the path that was added to the router
func (watcher *Watcher) RemoveWatchPath(watchPoint *WatchPoint) {
	watch.database.RemoveWatchPoint(watchPoint)
}

func (watcher *Watcher) Start() {
	
	for  {
		
	}

}
