package flatten

func FlattenRecursive(slice []interface{}, depth int) []interface{} {
	newSlice := []interface{} {}

	for i := 0; i < len(slice); i++ {
		element, ok := slice[i].(int)

		if ok {
			newSlice = append(newSlice, element)
		} else {
			nestedSlice := slice[i].([]interface{})
			for i := 0; i < len(nestedSlice); i++ {
				newSlice = append(newSlice, nestedSlice[i])
			}
		}
	}

	depth--
	if depth <= 0 {
		return newSlice
	} else {
		return FlattenRecursive(newSlice, depth)
	}
}

/*

 {0,  {1, 2},  {{{3, 4}}}}, 2)


 new slice
	
	iterate over elements
		if slice
			iterate over slice
				append elements to new slice
		else
			append element to new slice
	
	depth--
	if depth == 0 { return new slice }
	else return flatten(newSlice, depth)



*/