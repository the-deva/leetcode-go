/*
2257 : https://leetcode.com/problems/count-unguarded-cells-in-the-grid/description
*/
package main

func countUnguarded(m, n int, guards, walls [][]int) (result int) {
	// Build the grid g
	g := make([][]int, m)
	for i := range g {
		g[i] = make([]int, n)
	}
	for _, p := range guards {
		g[p[0]][p[1]] = 1 // Mark guards
	}
	for _, p := range walls {
		g[p[0]][p[1]] = 2 // Mark walls
	}

	// Simulate row
	for _, row := range g {
		for j := 0; j < n; j++ {
			if row[j] == 2 {
				continue
			}
			st, has1 := j, false
			for ; j < n && row[j] != 2; j++ {
				if row[j] == 1 {
					has1 = true
				}
			}
			if has1 {
				for ; st < j; st++ {
					if row[st] == 0 {
						row[st] = -1
					}
				}
			}
		}
	}

	// Simulate column
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if g[i][j] == 2 {
				continue
			}
			st, has1 := i, false
			for ; i < m && g[i][j] != 2; i++ {
				if g[i][j] == 1 {
					has1 = true
				}
			}
			if has1 {
				for ; st < i; st++ {
					if g[st][j] == 0 {
						g[st][j] = -1
					}
				}
			}
		}
	}

	// Count
	for _, row := range g {
		for _, v := range row {
			if v == 0 {
				result++
			}
		}
	}
	return
}

func main() {

}
