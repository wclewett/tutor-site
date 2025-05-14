package session

import (
  "sync"
  "fmt"
)
var lock = &sync.Mutex{}
var singleManager *Manager

type single struct{}

type Manager struct {
  cookieName  string
  mut         sync.Mutex
  provider    Provider
  maxlife     int64
}

type Provider interface {
  SessionInit(sid string) (Session, error)
  SessionRead(sid string) (Session, error)
  SessionDestroy(sid string) error
  SessionGC(maxLifeTime int64)
}

type Session interface {
  Set(key, value interface{}) error 
  Get(key interface{}) interface{}
  Delete(key interface{}) error    
  SessionID() string
}

func getManager() *Manager {
  if singleManager == nil {
    lock.Lock()
    defer lock.Unlock()
    if singleManager == nil {
      fmt.Println("Creating single instance now.")
      singleManager = &Manager{}
    } else {
      fmt.Println("Single instance already created.")
    }
  } else {
    fmt.Println("Single instance already created.")
  }

  return singleManager
}
