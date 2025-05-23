package utils

import (
	"encoding/json"
	"sync"
)

type Map struct {
	// 读写互斥锁（读写锁）
	sync.RWMutex
	m map[interface{}]interface{}
}

func (m *Map) init() {
	if m.m == nil {
		m.m = make(map[interface{}]interface{})
	}
}

func (m *Map) UnsafeGet(key interface{}) interface{} {
	if m.m == nil {
		return nil
	} else {
		return m.m[key]
	}
}

func (m *Map) Get(key interface{}) interface{} {
	m.RLock()
	defer m.RUnlock()
	return m.UnsafeGet(key)
}

func (m *Map) UnsafeSet(key, value interface{}) {
	m.init()
	m.m[key] = value
}

func (m *Map) Set(key interface{}, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.UnsafeSet(key, value)
}

func (m *Map) TestAndSet(key interface{}, value interface{}) interface{} {
	m.Lock()
	defer m.Unlock()

	m.init()

	if v, ok := m.m[key]; ok {
		return v
	} else {
		m.m[key] = value
		return nil
	}
}

func (m *Map) UnsafeDel(key interface{}) {
	m.init()
	delete(m.m, key)
}

func (m *Map) Del(key interface{}) {
	m.Lock()
	defer m.Unlock()
	m.UnsafeDel(key)
}

func (m *Map) UnsafeLen() int {
	if m.m == nil {
		return 0
	} else {
		return len(m.m)
	}
}

func (m *Map) Len() int {
	m.RLock()
	defer m.RUnlock()
	return m.UnsafeLen()
}

func (m *Map) UnsafeRange(f func(interface{}, interface{}) bool) {
	if m.m == nil {
		return
	}
	for k, v := range m.m {
		if !f(k, v) {
			break
		}
	}
}

func (m *Map) RLockRange(f func(interface{}, interface{}) bool) {
	m.RLock()
	defer m.RUnlock()
	m.UnsafeRange(f)
}

func (m *Map) LockRange(f func(interface{}, interface{}) bool) {
	m.Lock()
	defer m.Unlock()
	m.UnsafeRange(f)
}

func MapToJSONString(param map[string]interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

func JsonStringToMap(jsonString string) (res map[string]interface{}) {
	_ = json.Unmarshal([]byte(jsonString), &res)
	return res
}

func GetSwitchFromOptions(Options map[string]interface{}, key string) bool {
	flag, ok := Options[key]
	if !ok {
		return true
	}
	if intVal, ok := flag.(int); ok && intVal == 1 {
		return true
	}
	return false
}
