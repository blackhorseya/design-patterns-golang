package singleton

import "sync"

// Singleton declare singleton object
type Singleton struct {

}

var singleton *Singleton
var once sync.Once

// GetInstance serve caller to get singleton object
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}
