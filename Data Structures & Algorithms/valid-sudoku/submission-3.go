func isValidSudoku(board [][]byte) bool {
	for i := range board[0] {
		fmt.Print(i, " ")
	}
	fmt.Println()
	for _, v := range board {
		for _, vv := range v {
			fmt.Print(string(vv), " ")
		}
		fmt.Println()
	}

	sections := make(map[string]map[byte]struct{})
	rows := make(map[int]map[byte]struct{})
	cols := make(map[int]map[byte]struct{})

	for i := 0; i < len(board); i++ {
		row := make(map[byte]struct{})

		for j := 0; j < len(board[i]); j++ {
			v := board[i][j]
			if v == '.' {
				continue
			}

			if _, ok := row[v]; ok {
				return false
			}
			row[v] = struct{}{}

			if _, ok := cols[j]; !ok {
				cols[j] = make(map[byte]struct{})
			}
			if _, ok := cols[j][v]; ok {
				return false
			}
			cols[j][v] = struct{}{}

			const sectionCount int = 3
			sectionKey := fmt.Sprintf("%d-%d", i/sectionCount, j/sectionCount);
			if _, ok := sections[sectionKey]; !ok {
				sections[sectionKey] = make(map[byte]struct{})
			}
			 if _, ok := sections[sectionKey][v]; ok {
				 return false
			 }
			sections[sectionKey][v] = struct{}{}
		}

		rows[i] = row
	}

	for k, v := range sections {
		fmt.Println(k)
		for kk := range v {
			fmt.Print(string(kk), " ")
		}
		fmt.Println()
	}

	fmt.Println("len", len(sections))


	return true

}