package git

import (
	"sync"
)

type _tagCache struct {
	sync.RWMutex
	m map[sha1]*Tag
}

var tagCache _tagCache

func (c *_tagCache) get(tagID sha1) *Tag {
	c.RLock()
	result := c.m[tagID]
	c.RUnlock()
	return result
}

func (c *_tagCache) set(tagID sha1, tag *Tag) {
	c.Lock()
	if c.m == nil {
		c.m = make(map[sha1]*Tag)
	}
	c.m[tagID] = tag
	c.Unlock()
}
