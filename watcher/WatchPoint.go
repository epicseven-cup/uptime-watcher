package watcher

import (
	"uuid"
)

type WatchPoint struct {
	url string
	watchPointID uuid.UUID
}

// Getter for these property of watchpoint

func (watchPoint *WatchPoint) GetUrl() string {
	return watchPoint.url
}


func (watchPoint *WatchPoint) GetWatchPointID() uuid.UUID {
	return watchPoint.watchPointID
}
