package pubsub

import "sync"

type idChecker struct {
	ids        []string
	maxIds     int
	lastStored int
	lock       sync.Mutex
}

func (i *idChecker) Add(new string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	if len(i.ids) < i.maxIds {
		i.lastStored++
		i.ids = append(i.ids, new)
	} else {
		i.lastStored = (i.lastStored + 1) % i.maxIds
		i.ids[i.lastStored] = new
	}
}

func (i *idChecker) Has(n string) bool {
	i.lock.Lock()
	defer i.lock.Unlock()
	for _, e := range i.ids {
		if e == n {
			return true
		}
	}
	return false
}

func newIdChecker(maxIds int) idChecker {
	return idChecker{
		ids:        make([]string, 0),
		maxIds:     maxIds,
		lastStored: 0,
	}
}
