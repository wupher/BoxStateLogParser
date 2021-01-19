package main

type StatusSet map[int]struct{}

func (set StatusSet) Add(status int) {
	set[status] = struct{}{}
}

func (set StatusSet) Remove(status int) {
	delete(set, status)
}

func (set StatusSet) Has(status int) bool {
	_, ok := set[status]
	return ok
}
