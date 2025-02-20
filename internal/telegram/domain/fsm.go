package domain

import (
	"reflect"
	"sync"
)

type State = string

type AbstractFsm interface {
	Set(userID int64, state State)
	Current(userID int64) (State, bool)
	Finish(userID int64)
	Mutex() *sync.Mutex
	Map() map[int64]map[reflect.Type]interface{}
}

type AbstractDataStorage[T any] interface {
	Set(data T)
	Get() T
}
