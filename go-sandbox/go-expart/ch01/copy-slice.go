src := []int{1,2,3,4,5}

dst := make([]int, len(src))
copy(dst, src)

fmt.Println(dst, len(dst), cap(dst))