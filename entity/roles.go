package entity

import "fmt"

type Villager struct {
}

func (v *Villager) SayHello() {
	fmt.Println("hello I'm a villager")
}

type Werewolf struct {
	Name string
}

func (w *Werewolf) SayHello() {
	fmt.Println("hello I'm a werewolf")
}

type Seer struct {
}

func (s *Seer) SayHello() {
	fmt.Println("hello I'm a seer")
}

type Healer struct {
}

func (h *Healer) SayHello() {
	fmt.Println("hello I'm a healer")
}

type PlayerRole interface{
	SayHello()
}

func CreateRolesSlice(n int) []PlayerRole {
	roles := []PlayerRole{new(Werewolf), new(Werewolf), new(Healer), new(Seer)}
	n = n-4
	i := 0
	for (i < n) {
		roles = append(roles, new(Villager))
		i++
	}
	return roles
}
