/*
Условие задачи
Склад маркетплейса имеет форму прямоугольника и разделен на клетки площадью 1×1. Общая площадь склада n×m,
где n — количество строк, m — количество столбцов, пронумерованных от 1 до n и от 1 до m соответственно.
Известно, что все клетки склада с чётными индексами строки и столбца заняты стойками, остальные свободны для перемещения.
Также известно что склад состоит из нечётного количества строк и столбцов.
На двух различных свободных для перемещения клетках расположены роботы моделей A и B.
Роботы могут перемещаться в соседнюю по вертикали или горизонтали свободную клетку.
Вам необходимо построить два непересекающихся маршрута, которые приведут одного из роботов в верхнюю левую клетку склада (1;1),
а другого — в нижнюю правую (n;m).
Обратите внимание, минимизировать длину маршрутов роботов не требуется. Какой из роботов окажется в верхней левой клетке, а какой в правой нижней — не важно.
Путь робота должен быть простым, то есть робот не может посещать клетку, в которую он уже перемещался.

Входные данные
Первая строка содержит целое число t (1≤t≤100) — количество наборов входных данных.
Далее следует описание наборов входных данных.
Первая строка каждого набора входных данных содержит два нечетных целых числа (3≤n,m<100) — количество строк и столбцов склада.
Следующие n строк каждого набора входных данных содержат по m символов в каждой — описание схемы склада. Стойки обозначаются символом #,
свободные клетки символом ., а роботы — символами A и B.

Выходные данные
Для каждого набора входных данных выведите n строк по m символов в каждой. Путь робота A обозначьте символами a, а путь робота B — символами b.

Пояснения к первому примеру:
A и B — это роботы, a и b — это их путь, стрелками показано как они двигаются.
Точки — это пустые клетки. Решетки — это стойки (занятые клетки).

Пример теста 1
Входные данные
1
3 3
B..
.#.
..A

Выходные данные
B..
.#.
..A

Пример теста 2
Входные данные
2
5 5
.....
.#A#.
...B.
.#.#.
.....
7 9
.........
.#.#.#.#.
..AB.....
.#.#.#.#.
.........
.#.#.#.#.
.........

Выходные данные
aaa..
.#A#.
...Bb
.#.#b
....b
aaa......
.#a#.#.#.
..ABb....
.#.#b#.#.
....b....
.#.#b#.#.
....bbbbb
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

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// определяем количество групп входных данных
	scanner.Scan()
	countGroup, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	// считываем ответы по группам
	for group := 0; group < countGroup; group++ {

		// определяем размеры склада (n и m)
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

		// склад
		storage := make([][]string, n, n)

		// начальные координаты гоботов
		var xA, yA, xB, yB int

		// построчно считываем входные данные
		for i := 0; i < n; i++ {

			storage[i] = make([]string, m, m)

			scanner.Scan() // считываем строку с описанием
			for j, value := range scanner.Text() {
				storage[i][j] = string(value)
				if string(value) == "A" {
					xA, yA = i, j
				}
				if string(value) == "B" {
					xB, yB = i, j
				}
			}
		}

		// если А в верхнем левом углу и В в нижнем правом углу
		if (xA == 0 && yA == 0) && (xB == n-1 && yB == m-1) {
			outputing(out, storage, n, m)
			continue
		}
		// если B в верхнем левом углу и A в нижнем правом углу
		if (xA == n-1 && yA == m-1) && (xB == 0 && yB == 0) {
			outputing(out, storage, n, m)
			continue
		}

		var xa, ya, xb, yb int

		if yA < yB { // если А левее В
			if yA%2 != 0 || xA == 0 { // если для А столбик сверху (или снизу) или А в верхней строке
				ya = yA - 1 // делаем шаг влево
				xa = xA
			} else { // если для А нет столбика сверху и А не в верхней строке
				ya = yA
				xa = xA - 1 // делаем шаг вверх
			}

			upRunning(storage, xa, ya, "a")

			xa = 0
			leftRunning(storage, xa, ya, "a")

			if yB%2 != 0 || xB == n-1 { // если для В столбик снизу (или сверху) или B в нижней строке
				yb = yB + 1 // делаем шаг вправо
				xb = xB
			} else { // если для B нет столбика снизу и В не в нижней строке
				yb = yB
				xb = xB + 1 // делаем шаг вниз
			}

			downRunning(storage, xb, yb, "b")

			xb = len(storage) - 1
			rightRunning(storage, xb, yb, "b")
		}

		if yA == yB { // если А и В посередине

		}

		if yB < yA { // если В левее А

			if yB%2 != 0 || xB == 0 { // если для B столбик сверху (или снизу) или B в верхней строке
				yb = yB - 1 // делаем шаг влево
				xb = xB
			} else { // если для B нет столбика сверху и B не в верхней строке
				yb = yB
				xb = xB - 1 // делаем шаг вверх
			}

			upRunning(storage, xb, yb, "b")

			xb = 0
			leftRunning(storage, xb, yb, "b")

			if yA%2 != 0 || xA == n-1 { // если для A столбик снизу (или сверху) или A в нижней строке
				ya = yA + 1 // делаем шаг вправо
				xa = xA
			} else { // если для A нет столбика снизу и A не в нижней строке
				ya = yA
				xa = xA + 1 // делаем шаг вниз
			}

			downRunning(storage, xa, ya, "a")

			xa = len(storage) - 1
			rightRunning(storage, xa, ya, "a")
		}

		// выводим результат по группе данных
		outputing(out, storage, n, m)
	}
}

// upRunning описывает движение вверх до упора
func upRunning(storage [][]string, x, y int, letter string) {
	for i := x; i >= 0; i-- {
		storage[i][y] = letter
	}
}

// downRunning описывает движение вниз до упора
func downRunning(storage [][]string, x, y int, letter string) {
	for i := x; i < len(storage); i++ {
		storage[i][y] = letter
	}
}

// leftRunning описывает движение влево до упора
func leftRunning(storage [][]string, x, y int, letter string) {
	for i := y; i >= 0; i-- {
		storage[x][i] = letter
	}
}

// rightRunning описывает движение вправо до упора
func rightRunning(storage [][]string, x, y int, letter string) {
	for i := y; i < len(storage[x]); i++ {
		storage[x][i] = letter
	}
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
