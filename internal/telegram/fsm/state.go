package fsm

import (
	"reflect"
	"sync"
)

type State = string

type userStateStorage struct {
	mu      sync.Mutex
	Storage map[int64]State
}

type MultiTypeMemoryFSM struct {
	mu          sync.Mutex
	fsms        map[int64]map[reflect.Type]interface{} // Хранит FSM для разных типов
	userStorage userStateStorage
}

func NewMultiTypeFSM() *MultiTypeMemoryFSM {
	return &MultiTypeMemoryFSM{
		fsms: make(map[int64]map[reflect.Type]interface{}),
		userStorage: userStateStorage{
			mu:      sync.Mutex{},
			Storage: make(map[int64]State),
		},
	}
}

func (m *MultiTypeMemoryFSM) Set(userID int64, state State) {
	m.userStorage.mu.Lock()
	defer m.userStorage.mu.Unlock()
	m.userStorage.Storage[userID] = state
}

func (m *MultiTypeMemoryFSM) Current(userID int64) (State, bool) {
	state, ok := m.userStorage.Storage[userID]
	return state, ok
}
func (m *MultiTypeMemoryFSM) Finish(userID int64) {
	m.userStorage.mu.Lock()
	defer m.userStorage.mu.Unlock()
	delete(m.userStorage.Storage, userID)

	delete(m.fsms, userID)
}

type Fsm interface {
	Set(userID int64, state State)
	Current(userID int64) (State, bool)
	Finish(userID int64)
}

type DataStorage[T any] interface {
	SetData(userID int64, data T)
	GetData(userID int64) (T, bool)
}

type MemoryStorage[T any] struct {
	mu      sync.Mutex
	Storage map[int64]T
}

func NewMemoryFsm[T any]() *MemoryStorage[T] {
	return &MemoryStorage[T]{
		Storage: make(map[int64]T),
	}
}

func (m *MemoryStorage[T]) SetData(userID int64, data T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Storage[userID] = data
}

func (m *MemoryStorage[T]) GetData(userID int64) (T, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	data, ok := m.Storage[userID]
	return data, ok
}

func GetFSM[T any](multiFsm *MultiTypeMemoryFSM, userID int64) DataStorage[T] {
	multiFsm.mu.Lock()
	defer multiFsm.mu.Unlock()

	t := reflect.TypeOf((*T)(nil)).Elem()
	if fsm, ok := multiFsm.fsms[userID][t]; ok {
		return fsm.(DataStorage[T])
	}

	newFsm := NewMemoryFsm[T]()
	if multiFsm.fsms[userID] == nil {
		multiFsm.fsms[userID] = make(map[reflect.Type]interface{})
	}
	multiFsm.fsms[userID][t] = newFsm
	return newFsm
}
