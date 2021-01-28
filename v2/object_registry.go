package vlc

// #include <stdlib.h>
import "C"
import (
	"sync"
	"unsafe"
)

type objectID = unsafe.Pointer

type objectContext struct {
	refs uint
	data interface{}
}

type objectRegistry struct {
	sync.RWMutex

	contexts map[objectID]*objectContext
}

func newObjectRegistry() *objectRegistry {
	return &objectRegistry{
		contexts: map[objectID]*objectContext{},
	}
}

func (or *objectRegistry) get(id objectID) (interface{}, bool) {
	if id == nil {
		return nil, false
	}

	or.RLock()
	ctx, ok := or.contexts[id]
	or.RUnlock()

	if !ok {
		return nil, false
	}
	return ctx.data, ok
}

func (or *objectRegistry) add(data interface{}) objectID {
	or.Lock()

	var id objectID = C.malloc(C.size_t(1))
	or.contexts[id] = &objectContext{
		refs: 1,
		data: data,
	}

	or.Unlock()
	return id
}

func (or *objectRegistry) incRefs(id objectID) {
	if id == nil {
		return
	}

	or.Lock()

	ctx, ok := or.contexts[id]
	if ok {
		ctx.refs++
	}

	or.Unlock()
}

func (or *objectRegistry) decRefs(id objectID) {
	if id == nil {
		return
	}

	or.Lock()

	ctx, ok := or.contexts[id]
	if ok {
		ctx.refs--
		if ctx.refs == 0 {
			delete(or.contexts, id)
			C.free(id)
		}
	}

	or.Unlock()
}
