package git

import (
	"sync"
)

type _commitCache struct {
	sync.RWMutex
	m map[sha1]*Commit
}

var commitCache _commitCache

func (c *_commitCache) get(commitId sha1) *Commit {
	c.RLock()
	result := c.m[commitId]
	c.RUnlock()
	return result
}

func (c *_commitCache) set(commitId sha1, commit *Commit) {
	c.Lock()
	if c.m == nil {
		c.m = make(map[sha1]*Commit)
	}
	c.m[commitId] = commit
	c.Unlock()
}
