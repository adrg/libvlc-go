package vlc

import "sync"

type objectID uint64

type objectContext struct {
	refs uint
	data interface{}
}

type objectRegistry struct {
	sync.RWMutex

	contexts map[objectID]*objectContext
	sequence objectID
}

func newObjectRegistry() *objectRegistry {
	return &objectRegistry{
		contexts: map[objectID]*objectContext{},
	}
}

func (or *objectRegistry) get(id objectID) (interface{}, bool) {
	if id == 0 {
		return nil, false
	}

	or.RLock()
	defer or.RUnlock()

	ctx, ok := or.contexts[id]
	return ctx.data, ok
}

func (or *objectRegistry) add(data interface{}) objectID {
	or.Lock()
	defer or.Unlock()

	or.sequence++
	or.contexts[or.sequence] = &objectContext{
		refs: 1,
		data: data,
	}

	return or.sequence
}

func (or *objectRegistry) incRefs(id objectID) {
	if id == 0 {
		return
	}

	or.Lock()
	defer or.Unlock()

	ctx, ok := or.contexts[id]
	if !ok {
		return
	}
	ctx.refs++
}

func (or *objectRegistry) decRefs(id objectID) {
	if id == 0 {
		return
	}

	or.Lock()
	defer or.Unlock()

	ctx, ok := or.contexts[id]
	if !ok {
		return
	}

	ctx.refs--
	if ctx.refs == 0 {
		delete(or.contexts, id)
	}
}
