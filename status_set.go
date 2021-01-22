package main

import "sort"

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

type CompanySet map[string]struct{}

func (set CompanySet) Add(companyCode string) {
	set[companyCode] = struct{}{}
}

func (set CompanySet) Remove(companyCode string) {
	delete(set, companyCode)
}

func (set CompanySet) Has(companyCode string) bool {
	_, ok := set[companyCode]
	return ok
}

// A data structure to hold a key/value pair.
type Pair struct {
	CompanyCode string
	ErrCount    int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].ErrCount < p[j].ErrCount }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// A function to turn a map into a PairList, then sort and return it.
func sortByErrCount(errorCompany map[string]int) PairList {
	pl := make(PairList, len(errorCompany))
	i := 0
	for k, v := range errorCompany {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}
