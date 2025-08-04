/*
Условие задачи
В этой задаче вам необходимо реализовать валидацию корректности карты для стратегической компьютерной игры.
Карта состоит из гексагонов (шестиугольников), каждый из которых принадлежит какому-то региону карты.
В файлах игры карта представлена как n строк по m символов в каждой (строки и символы в них нумеруются с единицы).
Каждый нечетный символ каждой четной строки и каждый четный символ каждой нечетной строки — точка (символ «.» с ASCII кодом 46);
все остальные символы соответствуют гексагонам и являются заглавными буквами латинского алфавита. Буква указывает на то, какому региону принадлежит гексагон.
Регионы R, G, V, Y и B окрашены в красный, зеленый, фиолетовый, желтый и синийцвет, соответственно.
Вы должны проверить, что каждый регион карты является одной связной областью. Иными словами, не должно быть двух гексагонов,
принадлежащих одному и тому же региону, которые не соединены другими гексагонами этого же региона.
Неполные решения этой задачи (например, недостаточно эффективные) могут быть оценены частичным баллом.

Входные данные
В первой строке задано одно целое число t (1≤t≤100) — количество наборов входных данных.
Первая строка набора входных данных содержит два целых числа n и m (2≤n,m≤20) — количество строк и количество символов в каждой строке в описании карты.
Далее следуют n строк по m символов в каждой — описание карты. Каждый нечетный символ каждой четной строки и каждый четный символ
каждой нечетной строки — точка (символ «.» с ASCII кодом 46); все остальные символы соответствуют гексагонам и являются заглавными буквами латинского алфавита.
Первые два набора входных данных из примера показаны на второй картинке в условии.

Выходные данные
На каждый набор входных данных выведите ответ в отдельной строке — YES, если каждый регион карты представляет связную область, или NO, если это не так.

Пример теста 1
Входные данные

3
3 7
R.R.R.G
.Y.G.G.
B.Y.V.V
4 8
Y.R.B.B.
.R.R.B.V
B.R.B.R.
.B.B.R.R
2 7
G.B.R.G
.G.G.G.
Выходные данные

YES
NO
YES
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Hexagon описывает ячейку карты
type Hexagon struct {
	DotOrColor string // значение в ячейке
	Mark       bool   // отметка о посещении
}

// outputing выводит результат
func outputing(out *bufio.Writer, mes string) {

	fmt.Fprintf(out, "%s\n", mes)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// определяем количество групп входных данных
	scanner.Scan()
	t, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	// считываем ответы по группам
	for group := 1; group <= t; group++ {

		// определяем размеры карты (n, m)
		scanner.Scan()
		nWithm := strings.Split(scanner.Text(), " ")

		n, err := strconv.Atoi(nWithm[0])
		if err != nil {
			log.Fatal(err)
		}
		m, err := strconv.Atoi(nWithm[1])
		if err != nil {
			log.Fatal(err)
		}

		// field это карта
		field := make([][]Hexagon, n, n)
		// colorCounter мапа для подсчёта ячеек с одним цветом
		colorCounter := make(map[string]int)
		// linkCounter мапа для подсчёта связей между ячейками одного цвета
		linkCounter := make(map[string]int)

		mes := "YES" // итоговое сообщение

		// построчно считываем входные данные
		for i := 0; i < n; i++ {
			field[i] = make([]Hexagon, m, m)
			// считываем строку с описанием
			scanner.Scan()
			for j, value := range scanner.Text() {
				// записываем поле структуры с цветом или точкой
				field[i][j].DotOrColor = string(value)
				// если попадается точка, то уходим за следующим элементом
				if field[i][j].DotOrColor == "." {
					continue
				}
				key := field[i][j].DotOrColor // переназовём для удобства
				// исходя из условий доступности в карте проверяем количество одноцветных связей и увеличиваем счётчики
				if i == 0 && j == 0 {
					if _, ok := colorCounter[key]; !ok {
						colorCounter[key] = 1
						linkCounter[key] = 1
					}
				}

				if i == 0 && 1 < j {
					if _, ok := colorCounter[key]; !ok {
						colorCounter[key] = 1
						linkCounter[key] = 1
					} else {
						colorCounter[key]++
						if field[i][j].DotOrColor == field[i][j-2].DotOrColor && field[i][j-2].Mark == false {
							linkCounter[key]++
							field[i][j-2].Mark = true
						}
					}
				}

				if 0 < i && i%2 != 0 {
					if _, ok := colorCounter[key]; !ok {
						colorCounter[key] = 1
						linkCounter[key] = 1
					} else {
						colorCounter[key]++
						if 0 < j && j < m-1 {
							if field[i][j].DotOrColor == field[i-1][j-1].DotOrColor && field[i-1][j-1].Mark == false {
								linkCounter[key]++
								field[i-1][j-1].Mark = true
							}
							if field[i][j].DotOrColor == field[i-1][j+1].DotOrColor && field[i-1][j+1].Mark == false {
								linkCounter[key]++
								field[i-1][j+1].Mark = true
							}
						}
						if 1 < j {
							if field[i][j].DotOrColor == field[i][j-2].DotOrColor && field[i][j-2].Mark == false {
								linkCounter[key]++
								field[i][j-2].Mark = true
							}
						}
						if j == m-1 {
							if field[i][j].DotOrColor == field[i-1][j-1].DotOrColor && field[i-1][j-1].Mark == false {
								linkCounter[key]++
								field[i-1][j-1].Mark = true
							}
						}
					}
				}

				if 0 < i && i%2 == 0 {
					if _, ok := colorCounter[key]; !ok {
						colorCounter[key] = 1
						linkCounter[key] = 1
					} else {
						colorCounter[key]++
						if j != m-1 {
							if field[i][j].DotOrColor == field[i-1][j+1].DotOrColor && field[i-1][j+1].Mark == false {
								linkCounter[key]++
								field[i-1][j+1].Mark = true
							}
						}
						if 1 < j {
							if field[i][j].DotOrColor == field[i-1][j-1].DotOrColor && field[i-1][j-1].Mark == false {
								linkCounter[key]++
								field[i-1][j-1].Mark = true
							}
							if field[i][j].DotOrColor == field[i][j-2].DotOrColor && field[i][j-2].Mark == false {
								linkCounter[key]++
								field[i][j-2].Mark = true
							}
						}
					}
				}
			}
			// как же я ненавижу эти ступенчатые условия
		}

		// если количество ячеек одного цвета не равно количеству связей между ячейками одного цвета,
		// значит какие-то ячейки оторваны от коллектива
		for key, value := range colorCounter {
			if linkCounter[key] != value {
				mes = "NO"
				break
			}
		}

		// выводим результат по группе данных
		outputing(out, mes)
	}
}
