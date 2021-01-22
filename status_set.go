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
