package gemini

import (
	"errors"
	"fmt"
	log "github.com/cihub/seelog"
	"reflect"
	//tcp "game/common/server/tcp"
)

type GeminiHandler struct {
	defaultHandler func(*GeminiRequest)
	handlers       map[string]func(*GeminiRequest)
}

func NewGeminiHandler(defaultHandler func(*GeminiRequest)) *GeminiHandler {
	return &GeminiHandler{
		defaultHandler,
		make(map[string]func(*GeminiRequest)),
	}
}

// register handlers with the functions defined by <public func(*GeminiRequest)>
func (h *GeminiHandler) RegHandlers(handler interface{}) error {
	t := reflect.TypeOf(handler)
	v := reflect.ValueOf(handler)
	for i := 0; i < t.NumMethod(); i++ {
		if mt := t.Method(i); mt.PkgPath == "" {
			vt := v.Method(i)
			vi := vt.Interface()
			if f, ok := vi.(func(*GeminiRequest)); ok {
				h.regHandler(mt.Name, f)
			}
		}
	}
	return nil
}

func (h *GeminiHandler) SetDefaultHandler(handler func(*GeminiRequest)) {
	h.defaultHandler = handler
}

func (h *GeminiHandler) regHandler(name string, handler func(*GeminiRequest)) error {
	if _, ok := h.handlers[name]; ok {
		return errors.New(fmt.Sprintf("handler already existed, name=", name))
	}
	h.handlers[name] = handler
	log.Debugf("register handler, name=%s", name)
	return nil
}

func (h *GeminiHandler) invoke(req *GeminiRequest) error {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()

	if handler, ok := h.handlers[req.name]; ok {
		handler(req)
		return nil
	}
	if h.defaultHandler != nil {
		h.defaultHandler(req)
		return nil
	}

	return errors.New(fmt.Sprintf("no such handler[%s] or no default handler"))
}
