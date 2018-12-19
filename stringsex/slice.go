package stringsex

// Differences return two slices differences
func Differences(sliceA, sliceB []string) ([]string, []string, []string) {
	// Slice A, B are not empty
	if len(sliceA) > 0 && len(sliceB) > 0 {
		// New differences and intersection slices
		diffA := make([]string, 0)
		inter := make([]string, 0)
		diffB := make([]string, 0)
		// Interate slice A
		for _, strA := range sliceA {
			// New intersection flag
			isInter := false
			// Interate slice B
			for _, strB := range sliceB {
				// Is intersection
				if strA == strB {
					// Set isInter to true
					isInter = true
					// Append to inter
					inter = append(inter, strA)
					break
				}
			}
			if !isInter {
				// Append to diffA
				diffA = append(diffA, strA)
			}
		}

		// Check inter
		if len(inter) == 0 {
			// Intersection is empnty
			diffB = sliceB
		} else if len(inter) != len(sliceB) {
			// Intersection is not empty
			for i, strB := range sliceB {
				// New different flag
				isInter := false
				// New intersect count
				count := 0
				// Interate slice inter
				for _, strI := range inter {
					if strB == strI {
						count++
						isInter = true
						break
					}
				}

				if !isInter {
					// Append to diffB
					diffB = append(diffB, strB)
				} else if count == len(inter) {
					// All intersections be found, append sliceB remaining strs to diffB
					diffB = append(diffB, sliceB[i+1:len(sliceB)]...)
					break
				}
			}
		}

		return diffA, inter, diffB
	}

	// 1. Slice A is not empty and slice B is empty
	// 2. Slice A is empty and slice B is not empty
	// 3. Slice A is empty and slice B is empty
	return sliceA, []string{}, sliceB
}
