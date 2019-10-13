// Note: this implementation is not thread safe!

package singleton

// Singleton defines the singleton interface
// that the caller of this package interacts with
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

var instance *singleton

// GetInstance returns an instance of
// the Singleton interface
func GetInstance() Singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
