package eventsub

import "sync"

var maxIdsStored = 20

type idChecker struct {
	ids        []string
	lastStored int
	lock       sync.Mutex
}

func (i *idChecker) Add(new string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	if len(i.ids) < maxIdsStored {
		i.lastStored++
		i.ids = append(i.ids, new)
	} else {
		i.lastStored = (i.lastStored + 1) % maxIdsStored
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
