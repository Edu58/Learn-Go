package sets

type Set struct {
	intergerMap map[int]bool
}

func (set *Set) New() {
	set.intergerMap = make(map[int]bool)
}

func (set *Set) FindEelement(el int) bool {
	return set.intergerMap[el]
}

func (set *Set) AddEelement(el int, val bool) {
	if !set.FindEelement(el) {
		set.intergerMap[el] = val
	}
}

func (set *Set) DeleteElement(el int) {
	exists := set.FindEelement(el)
	if exists {
		delete(set.intergerMap, el)
	}
}

func (set *Set) Intersection(anotherSet *Set) *Set {
	var intersectionSet = &Set{}
	intersectionSet.New()
	for key, val := range set.intergerMap {
		if anotherSet.intergerMap[key] {
			intersectionSet.AddEelement(key, val)
		}
	}
	return intersectionSet
}
