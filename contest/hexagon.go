/*
Условие задачи

Изобразите шестиугольник заданного размера, используя символы «_» — нижнее подчёркивание, «/» — прямой слеш, «\» — обратный слеш.
Размер шестиугольника определяется двумя параметрами:
width — количество символов «_» на верхней и нижней сторонах шестиугольника;
height — количество прямых «/» или обратных «\» слешей на каждой из четырёх боковых сторон шестиугольника.

Входные данные
В первой строке дано натуральное число t (1≤t≤500) — количество шестиугольников, которые нужно изобразить.
В каждой из следующих t строк дано по два натуральных числа через пробел width[i], height[i] (1≤width[i],height[i]≤50) — требуемая ширина и высота i-го шестиугольника.
Гарантируется, что в рамках генерации понадобится вывести не более 3000 символов.
Выходные данные
Выведите в ответе t шестиугольников, руководствуясь примерами, где i−й шестиугольник имеет размер width[i] и height[i].
Используйте нужное количество пробелов:
• Между символами «/» и «\». Для нижней стороны шестиугольника вместо пробелов используйте «_» в количестве width[i].
• От начала каждой строки до первого символа «/», «\» или «_». Для самой длинной строки в шестиугольнике пробелы в начале не нужны.
После последнего символа «_», «/» или «\» в каждой строке используйте перенос строки без дополнительных пробелов.
Шестиугольники должны идти один за другим через один перенос строки. Между шестиугольниками не должно быть пустых строк.
Вывод должен оканчиваться одним переносом строки.

Пример теста 1
Входные данные

3
1 1
2 1
5 3
Выходные данные

 _
/ \
\_/
 __
/  \
\__/
   _____
  /     \
 /       \
/         \
\         /
 \       /
  \_____/
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

// inpTwoInt считывает строку, в которой вводятся два числа int
func inpTwoInt(sc *bufio.Scanner) (int, int) {

	// считываем строку с будем надеяться двумя числами
	sc.Scan()
	aWithb := strings.Split(sc.Text(), " ")
	a, err := strconv.Atoi(aWithb[0])
	b, err := strconv.Atoi(aWithb[1])
	if err != nil {
		log.Fatal(err)
	}

	return a, b
}

// fieldInit инициализирует массив
func fieldInit(h, w int) [][]string {

	field := make([][]string, 2*h+1, 2*h+1)

	for i := 0; i < len(field); i++ {
		field[i] = make([]string, 2*h+w, 2*h+w)
		for j := 0; j < len(field[i]); j++ {
			field[i][j] = " "
		}
	}

	return field
}

// hexagonCompletion рисует шестиугольник в массиве
func hexagonCompletion(h, w int, field *[][]string) {

	// определим условное начало шестиугольника
	x := h + 1
	y := 0

	prefix := 0 // смещение для отрисовки

	// идём по крышке
	for i := x - 1; i >= x-h-1; i-- {
		for j := y; j <= y+2*h+w-1; j++ {
			if i >= 0 && i < len(*field) && j >= 0 && j < len((*field)[0]) {
				if i != x-h-1 && j == y+prefix {
					(*field)[i][j] = "/"
				}
				if i != x-h-1 && y+prefix < j && j < y+2*h+w-1-prefix {
					(*field)[i][j] = " "
				}
				if i != x-h-1 && j == y+2*h+w-1-prefix {
					(*field)[i][j] = "\\"
				}
				if i == x-h-1 && y+prefix < j && j < y+2*h+w-1-prefix {
					(*field)[i][j] = "_"
				}
			}
		}
		prefix++
	}

	prefix = 0

	// идём по донышку
	for i := x; i <= x+h-1; i++ {
		for j := y; j <= y+2*h+w-1; j++ {
			if i >= 0 && i < len(*field) && j >= 0 && j < len((*field)[0]) {
				if j == y+prefix {
					(*field)[i][j] = "\\"
				}
				if i != x+h-1 && y+prefix < j && j < y+2*h+w-1-prefix {
					(*field)[i][j] = " "
				}
				if j == y+2*h+w-1-prefix {
					(*field)[i][j] = "/"
				}
				if i == x+h-1 && y+prefix < j && j < y+2*h+w-1-prefix {
					(*field)[i][j] = "_"
				}
			}
		}
		prefix++
	}
}

// outputing выводит результат
func outputing(out *bufio.Writer, arr *[][]string) {

	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len((*arr)[i]); j++ {
			fmt.Fprintf(out, "%v", (*arr)[i][j])
		}
		fmt.Fprint(out, "\n")
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// определяем количество шестиугольников
	scanner.Scan()
	countHex, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	// считываем данные по группам
	for h := 0; h < countHex; h++ {

		// определяем основание width и полувысоту height текущего шестиугольника
		width, height := inpTwoInt(scanner)

		// field это массив для хранения символов, из которых состоит шестиугольник
		field := fieldInit(height, width)

		hexagonCompletion(height, width, &field)

		// выводим результат по группе
		outputing(out, &field)
	}
}
