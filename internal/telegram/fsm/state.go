package fsm

import (
	"TelegramBot/internal/telegram/domain"
	"reflect"
	"sync"
)

type userStateStorage struct {
	mu      sync.Mutex
	Storage map[int64]domain.State
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
			Storage: make(map[int64]domain.State),
		},
	}
}

func (m *MultiTypeMemoryFSM) Set(userID int64, state domain.State) {
	m.userStorage.mu.Lock()
	defer m.userStorage.mu.Unlock()
	m.userStorage.Storage[userID] = state
}

func (m *MultiTypeMemoryFSM) Current(userID int64) (domain.State, bool) {
	state, ok := m.userStorage.Storage[userID]
	return state, ok
}
func (m *MultiTypeMemoryFSM) Finish(userID int64) {
	m.userStorage.mu.Lock()
	defer m.userStorage.mu.Unlock()
	delete(m.userStorage.Storage, userID)

	delete(m.fsms, userID)
}

func (m *MultiTypeMemoryFSM) Mutex() *sync.Mutex {
	return &m.mu
}

func (m *MultiTypeMemoryFSM) Map() map[int64]map[reflect.Type]interface{} {
	return m.fsms
}

type MemoryStorage[T any] struct {
	value T
}

func NewMemoryFsm[T any]() *MemoryStorage[T] {
	return &MemoryStorage[T]{}
}

func (m *MemoryStorage[T]) Set(data T) {
	m.value = data
}

func (m *MemoryStorage[T]) Get() T {
	return m.value
}

func GetFSM[T any](ctx domain.AbstractContext) domain.AbstractDataStorage[T] {
	multiFsm := ctx.FSM()
	multiFsm.Mutex().Lock()
	fsms := multiFsm.Map()
	defer multiFsm.Mutex().Unlock()

	t := reflect.TypeOf((*T)(nil)).Elem()
	if fsm, ok := fsms[ctx.UserID()][t]; ok {
		return fsm.(domain.AbstractDataStorage[T])
	}

	newFsm := NewMemoryFsm[T]()
	if fsms[ctx.UserID()] == nil {
		fsms[ctx.UserID()] = make(map[reflect.Type]interface{})
	}
	fsms[ctx.UserID()][t] = newFsm
	return newFsm
}
