package fileProcessor

type StringInt struct {
	Word  string
	Count int
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []StringInt, elementToFind StringInt) (int, bool) {
	for i, item := range slice {
		if item.Word == elementToFind.Word {
			return i, true
		}
	}
	return -1, false
}

//Find elements from one slice in other slice and when if finds matching elements it updates it
func FindAndUpdate(source []StringInt, target []StringInt) []StringInt {
	for _, element := range source {
		j, exist := Find(target, element)
		if !exist {
			target = append(target, element)
		} else {
			target[j].Count = target[j].Count + element.Count
		}
	}
	return target
}
