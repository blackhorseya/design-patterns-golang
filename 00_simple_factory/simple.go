package simplefactory

import "fmt"

// Adventurer declare an adventurer's behavior
type Adventurer interface {
	Say(name string) string
}

// NewAdventurer serve caller to create an Adventurer
func NewAdventurer(t string) Adventurer {
	switch t {
	case "archer":
		return &archer{}
	case "warrior":
		return &warrior{}
	default:
		return nil
	}
}

type archer struct {
}

func (a *archer) Say(name string) string {
	return fmt.Sprintf("Hi %s, I'm an archer", name)
}

type warrior struct {
}

func (w *warrior) Say(name string) string {
	return fmt.Sprintf("Hi %s, I'm a warrior", name)
}
