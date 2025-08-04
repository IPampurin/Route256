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

// inputing считывает данные по шестиугольнику
func inputing(sc *bufio.Scanner) (int, int) {

	sc.Scan()
	nm := strings.Split(sc.Text(), " ")

	// основание шестиугольника
	width, err := strconv.Atoi(nm[0])
	// полувысота шестиугольника
	height, err := strconv.Atoi(nm[1])
	if err != nil {
		log.Fatal(err)
	}

	return width, height
}

// hexagonCompletion формирует массив с шестиугольником
func hexagonCompletion(n, m, height int) [][]string {

	// field это массив для хранения символов, из которых состоит шестиугольник
	field := make([][]string, n, n)

	prefix := height + 1

	for i := 0; i < n; i++ {
		field[i] = make([]string, m, m)
		if i <= n/2 {
			prefix--
		}
		if i > n/2 {
			prefix++
		}
		for j := 0; j < m; j++ {
			if i == 0 && j < prefix {
				field[i][j] = ` `
			}
			if i == 0 && prefix <= j && j <= m-1-prefix {
				field[i][j] = `_`
			}
			if i > 0 && i <= n/2 {
				if j < prefix {
					field[i][j] = ` `
				}
				if j == prefix {
					field[i][j] = `/`
				}
				if prefix < j && j < m-1-prefix {
					field[i][j] = ` `
				}
				if j == m-1-prefix {
					field[i][j] = `\`
				}
			}
			if i > n/2 {
				if j < prefix-1 {
					field[i][j] = ` `
				}
				if j == prefix-1 {
					field[i][j] = `\`
				}
				if prefix-1 < j && j < m-prefix {
					field[i][j] = ` `
				}
				if j == m-prefix {
					field[i][j] = `/`
				}
				if i == len(field)-1 && j < prefix-1 {
					field[i][j] = ` `
				}
				if i == len(field)-1 && prefix-1 < j && j < m-prefix {
					field[i][j] = `_`
				}
			}
		}
	}

	return field
}

// outputing выводит результат
func outputing(out *bufio.Writer, arr [][]string, n, m int) {

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fprintf(out, "%s", arr[i][j])
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

		// определяем параметры текущего шестиугольника
		width, height := inputing(scanner)

		// высота поля под шестиугольник
		n := 2*height + 1
		// длина поля под шестиугольник
		m := width + 2*height

		// определяем массив с шестиугольником
		field := hexagonCompletion(n, m, height)

		// выводим результат по группе
		outputing(out, field, n, m)
	}
}
