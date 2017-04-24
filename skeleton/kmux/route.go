package kmux

import (
	"net/http"
	"strings"
)

type Handle func(http.ResponseWriter, *http.Request, Params)

type route struct {
	allRegistered map[string]bool
	statics       map[string]Handle
	// :id
	dynamics []*routePattern
	// // *file
	// matchAlls []routePattern
}

func newRoute() *route {
	r := route{}
	r.allRegistered = make(map[string]bool)
	r.statics = make(map[string]Handle)
	return &r
}

type routePattern struct {
	parts        []string
	partsCount   int
	pathVariable Params
	handle       Handle
}

type Params map[int]Param

type Param struct {
	Key   string
	Value string
}

func (r *route) add(pattern string, handle Handle) {

	if r.allRegistered == nil {
		r.allRegistered = make(map[string]bool)
	}
	if r.allRegistered[pattern] {
		panic("pattern" + pattern + "have registered")
	}
	dynamic := buildRoutePattern(pattern)
	if dynamic == nil {
		r.statics[pattern] = handle
	} else {
		dynamic.handle = handle
		r.dynamics = append(r.dynamics, dynamic)
	}
	r.allRegistered[pattern] = true
}

// buildRoutePattern retrun routePattern, isStatic
func buildRoutePattern(pattern string) *routePattern {
	// static
	if pattern == "/" {
		return nil
	}

	p := &routePattern{}
	parts := strings.Split(pattern, "/")
	p.parts = parts[1 : len(parts)-1]
	p.partsCount = len(p.parts)
	p.pathVariable = make(map[int]Param)
	for i, segment := range p.parts {
		if len(segment) < 1 {
			panic("segment of pattern should at least own one character")
		}
		if segment[0] == ':' {
			if len(segment) < 2 {
				panic("invalid segment: " + segment)
			}
			p.pathVariable[i] = Param{Key: segment[1:]}
		}
	}
	if len(p.pathVariable) == 0 {
		return nil
	}
	return p
}

func (rp *routePattern) match(pattern string) bool {
	segments := strings.Split(pattern, "/")
	segments = segments[1 : len(segments)-1]
	if len(segments) != rp.partsCount {
		return false
	}
	params := make(map[int]Param, len(rp.pathVariable))
	for i, seg := range segments {
		if rp.parts[i] == seg {
			continue
		}
		param, found := rp.pathVariable[i]
		if !found {
			return false
		}
		param.Value = seg
		params[i] = param
	}
	rp.pathVariable = params
	return true
}

func (r *route) getHandle(pattern string) (Handle, Params) {
	h, ok := r.statics[pattern]
	if ok {
		return h, nil
	}
	for _, rp := range r.dynamics {
		if rp.match(pattern) {
			return rp.handle, rp.pathVariable
		}
	}
	return nil, nil
}
