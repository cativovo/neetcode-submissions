func isValidSudoku(board [][]byte) bool {
	const sectionCount int = 3
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

			if _, ok := cols[j]; !ok {
				cols[j] = make(map[byte]struct{})
			}
			sectionKey := fmt.Sprintf("%d-%d", i/sectionCount, j/sectionCount);
			if _, ok := sections[sectionKey]; !ok {
				sections[sectionKey] = make(map[byte]struct{})
			}

			if _, ok := row[v]; ok {
				return false
			}
			if _, ok := cols[j][v]; ok {
				return false
			}
			if _, ok := sections[sectionKey][v]; ok {
				return false
			}

			row[v] = struct{}{}
			cols[j][v] = struct{}{}
			sections[sectionKey][v] = struct{}{}
		}

		rows[i] = row
	}

	return true

}