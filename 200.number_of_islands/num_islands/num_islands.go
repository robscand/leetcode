package num_islands

import "fmt"

func NumIslands(grid [][]byte) int {
 	counter := 0 // число островов
 	if len(grid) > 0 {
		rows := len(grid)
		columns := len(grid[0])
		for row := 0; row < rows; row++ {
			for col := 0; col < columns; col++ {
				if grid[row][col] == '1' {
					// в первой ячейке есть 1 - тут начало острова
					// рекурсивный обход острова
					walkIsland(grid, row, col)
					counter++
				}
			}
		}
	}
	return counter
}

// Функция walkIsland позволяет обойти весь "остров", обнуляя при этом его "кусочки"
// чтобы в цикле было возможно найти новое вхождение в другой остров
func walkIsland(g [][]byte, r int, c int) {
	if (r >= 0 && r < len(g)) && (c >= 0 && c < len(g[0])) { // нет выхода за границы данной карты
		if g[r][c] != '0' {
			// ячейка является частью острова, т.к. тут нет нуля; также это условие того, что тут уже были
			g[r][c] = '0' // пометить ячейку как воду, т.е. теперь это не остров
			// рекурсивный проход по всем ячейкам острова по вертикали и по горизонтали
			walkIsland(g, r+1, c)
			walkIsland(g, r-1, c)
			walkIsland(g, r, c+1)
			walkIsland(g, r, c-1)
		}
	}
}

// Функция выводит отображение острова в консоль
func PrintIsland(grid [][]byte) int {
	counter := 0 // число островов
	for row := range grid{
		for col := range grid[1] {
			fmt.Print(grid[row][col])
		}
		fmt.Println()
	}
	return counter
}
