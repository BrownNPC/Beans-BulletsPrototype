// This file is auto generated. DO NOT EDIT
package t

var _cachePlayer cache[Player]

func (p *Player) onDestroy() {
	_cachePlayer.Store(p)
}

var _cacheThingRef cache[ThingRef]

func (p *ThingRef) onDestroy() {
	_cacheThingRef.Store(p)
}

var _cacheThings cache[Things]

func (p *Things) onDestroy() {
	_cacheThings.Store(p)
}

// New creates a new Thing. It tries reusing the object from a cache to avoid allocations.
func New[T any](v T) *T {
	switch arg := any(v).(type) {
	case Player:
		return _cachePlayer.New(arg).(*T)
	case ThingRef:
		return _cacheThingRef.New(arg).(*T)
	case Things:
		return _cacheThings.New(arg).(*T)
	default:
		return &v
	}
}
