package number_of_island

func solution(input [][]byte) int {
	if len(input) == 0 {
		return 0
	}

	r := len(input)
	c := len(input[0])
	x := make([]int, 0, r*c)
	y := make([]int, 0, r*c)

	var add = func(i, j int) {
		x = append(x, i)
		y = append(y, j)
		input[i][j] = '0'
	}

	var pop = func() (int, int) {
		i := x[0]
		x = x[1:]
		j := y[0]
		y = y[1:]
		return i, j
	}

	var bfs = func(i, j int) int {
		if input[i][j] == '0' {
			return 0
		}

		add(i, j)

		for len(x) > 0 {
			i, j = pop()

			if 0 <= i-1 && input[i-1][j] == '1' {
				add(i-1, j)
			}
			if 0 <= j-1 && input[i][j-1] == '1' {
				add(i, j-1)
			}
			if i+1 < r && input[i+1][j] == '1' {
				add(i+1, j)
			}
			if j+1 < c && input[i][j+1] == '1' {
				add(i, j+1)
			}
		}

		return 1
	}

	res := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			res += bfs(i, j)
		}
	}

	return res
}
