/*
Условие задачи
В Nozo разрабатывается новая компьютерная игра — «Колонизация 8». Вам поручено реализовать поиск пути для сухопутных юнитов.
Поле игры состоит из шестиугольников, которыми можно полностью покрыть бесконечную плоскость.
Все шестиугольники на поле игры делятся на 2 типа — суша и море. Каждый шестиугольник, у которого есть 6 сторон, считается сушей.
Остальные шестиугольники — морем.

Вам дана таблица из символов, состоящая из n строк. Каждая строка содержит m символов.
Границы шестиугольников обозначаются символами «\» — обратный слеш, «/» — прямой слеш и «_» — нижнее подчёркивание. Остальные символы поля — пробелы.
Все шестиугольники на поле имеют одинаковый размер. Обозначим за ширину шестиугольника количество подряд идущих символов «_» в его верхней и нижней сторонах.
Аналогично назовём высотой шестиугольника количество символов «/» в левой верхней и правой нижней сторонах.

Определите, есть ли путь между двумя шестиугольниками суши, который проходит только по шестиугольникам суши.
Из одного шестиугольника суши можно перейти в любой из шести соседних шестиугольников суши. Шестиугольники называются соседними, если имеют общую границу.

Входные данные
Каждый тест состоит из нескольких наборов входных данных.
Первая строка каждого теста содержит целое число t (1≤t≤100) — количество наборов входных данных.
Далее следует описание наборов входных данных.
Первая строка каждого набора входных данных содержит два целых числа n и m (3≤n,m≤100) — количество строк и столбцов, из которых состоит поле.
Следующие n строк каждого набора входных данных содержат по m символов — поле с шестиугольниками.
Следующие две строки каждого набора входных данных содержат по два целых числа x,y (1≤x≤n, 1≤y≤m) — координаты точки: номер строки и номер столбца.

Гарантируется:
• Обе координаты принадлежат к некоторым шестиуголькам суши, возможно — к одному и тому же.
• Данная координата не является границей никаких шестиугольников, то есть символ на данной координате — пробел.
• В первой и последней строках,  а также в первом и последнем столбцах есть хотя бы один непробельный символ. Следовательно, в этом поле есть хотя бы один шестиугольник суши.
• В этой таблице символов некоторые пробелы можно заменить на символы «_», «/», «\» так, чтобы получилась регулярная сетка из шестиугольников одинакового размера.
• Высота и ширина каждого шестиугольника не больше 10.

Выходные данные
Для каждого набора входных данных выведите YES, если между двумя шестиугольниками есть путь, который проходит только по шестиугольникам суши, иначе выведите NO.

Пример теста 1
Входные данные

4
3 3
 _
/ \
\_/
2 2
2 2
4 9
 _   _
/ \ / \_
\_/ \_/ \
      \_/
2 2
3 8
12 11
     _   _
   _/ \_/ \
  / \_/ \_/
  \_/ \_/
 _/ \_/ \_
/ \_/ \_/ \
\_/ \_/ \_/
/ \ / \
\_/ \_/  _
/ \_/ \ / \
\_/ \_/ \_/
  \_/
10 2
6 10
5 5
   _
 _/ \
/ \_/
\_/ \
  \_/
2 4
3 2

Выходные данные
YES
NO
YES
YES

Пример теста 2
Входные данные

3
5 5
  _
 / \
/   \
\   /
 \_/
2 3
4 4
6 7
 __
/  \__
\__/  \
/  \__/
\__/  \
   \__/
2 3
5 5
12 16
 ____      ____
/    \____/    \
\____/    \____/
     \____/
 ____      ____
/    \    /    \
\____/    \____/

 ____
/    \____
\____/    \
     \____/
6 14
10 4

Выходные данные
YES
YES
NO
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

// Dot описывает ячейку поля
type Dot struct {
	Symbol   string // какой символ был на вводе
	NamePeak int    // свяжем точку массива с вершиной графа по номеру вершины
}

// Point содержит координаты сетки шестиугольников
type Point struct {
	x int // индекс строки условного начала шестиугольника
	y int // индеск столбца условного начала шестиугольника
}

// Peak описывает вершину графа, в которую мы превращаем шестиугольник суши или моря
type Peak struct {
	Name  int              // имя (или номер) вершиины графа (идентификатор шестиугольника)
	Badge string           // признак отнесения к суше или морю (earth - земля, water - вода)
	Link  map[int]struct{} // связи с суседями
	Mark  bool             // отметка о посещении вершины (шестиугольника)
}

// inputCalc объединяет логику работы с данными
func inputCalc(sc *bufio.Scanner, out *bufio.Writer) {

	// считываем количество групп данных
	sc.Scan()
	groupCount, err := strconv.Atoi(sc.Text())
	if err != nil {
		log.Fatal(err)
	}

	// считываем и обрабатываем данные по группам
	for group := 1; group <= groupCount; group++ {

		// зададим итоговое сообщение
		message := "NO"

		// считываем количество строк и столбцов
		n, m := inpTwoInt(sc)

		// поле для отображения карты
		field := make([][]Dot, n, n)

		// в point запишем первую попавшуюся точку условного начала шестиугольника для последующего формирования сетки шестиугольников
		var point Point
		flag := true

		// построчно сканируем ввод и посимвольно вписываем в field
		for i := 0; i < len(field); i++ {
			field[i] = make([]Dot, m, m)
			sc.Scan()
			for j, val := range sc.Text() {
				dot := Dot{
					Symbol:   string(val),
					NamePeak: -1, // при заполнении массива все точки свяжем с той вершиной, которой не будет
				}
				field[i][j] = dot
				// если попадается признак левой оконечности шестиугольника, записываем координаты точки
				if flag && i > 0 && field[i][j].Symbol == "\\" && field[i-1][j].Symbol == "/" {
					point = Point{
						x: i,
						y: j,
					}
					flag = false
				}
			}
		}

		// считываем координаты стартовой точки
		startX, startY := inpTwoInt(sc)
		// считываем координаты финальной точки
		finishX, finishY := inpTwoInt(sc)

		// height полувысота шестиугольника, width длина основания шестиугольника
		// определяем параметры сетки шестиугольников исходя из предположения, что на поле есть хотя бы один шестиугольник
		height, width := sizeHex(point.x, point.y, &field)

		netPoints := netCoordinates(point.x, point.y, height, width, &field)

		// peaks это набор вершин графа, граф, где вершина это шестиугольник ячеек поля, объединённых одним признаком (суша или море)
		peaks := make([]Peak, 0)

		// номер вершины для удобства построения связей
		var currentNamePeak int
		// проходим по координатам вершин предполагаемых шестиугольников и после валидации доводим граф до ума
		for z, hex := range peaks {
			if validHex(height, width, hex.x, hex.y, &field) {
				peaks[z].Badge = "earth"
				hexPrintEarth(height, width, hex.x, hex.y, hex.Name, &field)
				// заполняем связи между вершинами с учётом шага сетки и доступности в массиве
				if hex.y-2*(height+width) >= 0 {
					currentNamePeak = field[hex.x][hex.y-2*(height+width)].NamePeak
					if currentNamePeak != -1 {
						peaks[z].Link[currentNamePeak] = struct{}{} // вносим в связи вершину левее
						peaks[currentNamePeak].Link[z] = struct{}{} // для вершины левее вносим в связи текущую вершину
					}
				}
				if hex.x-2*height >= 0 {
					currentNamePeak = field[hex.x-2*height][hex.y].NamePeak
					if currentNamePeak != -1 {
						peaks[z].Link[currentNamePeak] = struct{}{} // вносим в связи вершину выше
						peaks[currentNamePeak].Link[z] = struct{}{} // для вершины выше вносим в связи текущую вершину
					}
				}
				if hex.y-(height+width) >= 0 && hex.x-height >= 0 {
					currentNamePeak = field[hex.x-height][hex.y-(height+width)].NamePeak
					if currentNamePeak != -1 {
						peaks[z].Link[currentNamePeak] = struct{}{} // вносим в связи вершину выше и левее
						peaks[currentNamePeak].Link[z] = struct{}{} // для вершины выше и левее вносим в связи текущую вершину
					}
				}
				if hex.y+(height+width) < len(field[0]) && hex.x-height >= 0 {
					if currentNamePeak != -1 {
						peaks[z].Link[currentNamePeak] = struct{}{} // вносим в связи вершину выше и правее
						peaks[currentNamePeak].Link[z] = struct{}{} // для вершины выше и правее вносим в связи текущую вершину
					}
				}
			}
		}

		// определяем стартовую и финишную вершины графа
		startPeak := field[startX][startY].NamePeak
		finishPeak := field[finishX][finishY].NamePeak

		// выводим поле по группе
		outputing(out, peaks)
	}
}

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

// sizeHex определяет параметры шестиугольной сетки в нашем специально обученном массиве
func sizeHex(x, y int, field *[][]Dot) (int, int) {

	x--    // перемещаемся на нижний левый элемент верхней крышки шестиугольника
	h := 1 // полувысота в этом случае == 1
	// шагаем вверх по левой стороне крышки шестиугольника пока не нащупаем самый верх
	for (*field)[x-1][y+1].Symbol != "_" {
		h++
		x--
		y++
	}

	// перемещаемся на крайний левый элемент верха крышки шестиугольника
	x--
	y++
	// в этом случае длина основания шестиугольника == 0
	w := 0
	// шагаем вправо по верху крышки шестиугольника пока не нащупаем её край
	for (*field)[x][y].Symbol == "_" {
		w++
		y++
	}

	return h, w
}

func netCoordinates(x, y, height, width int, field *[][]Dot) []Point {

	// набор координат точек условного начала шестиугольников
	points := make([]Point, 0)

	return points
}

// hexPrintEarth меняет принадлежность точек массива с несуществующей вершины на конкретную вершину графа
func hexPrintEarth(h, w, x, y, numPeak int, field *[][]Dot) {

	// если есть выход за границы массива, то ничего не меняем - это кусочек моря
	if y+2*h+w-1 > len((*field)[0])-1 {
		return
	}

	prefix := -1
	// идём по крышке
	for i := x - 1; i >= x-h-1; i-- {
		prefix++
		for j := y; j <= y+2*h+w; j++ {
			if i != x-h-1 && j == y+prefix {
				(*field)[i][j].NamePeak = numPeak
			}
			if i != x-h-1 && y+prefix < j && j < y+2*h+w-prefix {
				(*field)[i][j].NamePeak = numPeak
			}
			if i != x-h-1 && j == y+2*h+w-1-prefix {
				(*field)[i][j].NamePeak = numPeak
			}
			if i == x-h-1 && y+prefix <= j && j <= y+2*h+w-1-prefix {
				(*field)[i][j].NamePeak = numPeak
			}
		}
	}
	// идём по донышку
	prefix = -1
	for i := x; i <= x+h-1; i++ {
		prefix++
		for j := y; j <= y+2*h+w; j++ {
			if j == y+prefix {
				(*field)[i][j].NamePeak = numPeak
			}
			if i != x+h-1 && y+prefix < j && j < y+2*h+w-prefix {
				(*field)[i][j].NamePeak = numPeak
			}
			if j == y+2*h+w-1-prefix {
				(*field)[i][j].NamePeak = numPeak
			}
			if i == x+h-1 && y+prefix < j && j < y+2*h+w-1-prefix {
				(*field)[i][j].NamePeak = numPeak
			}
		}
	}
}

// validHex проверяет наличие всех шести сторон шестиугольника определённых размеров по заданным координатам в массиве
func validHex(h, w, x, y int, field *[][]Dot) bool {

	// если при проверке шестиугольника есть выход за границы массива, то ничего не рисуем
	if y+2*h+w-1 > len((*field)[0])-1 {
		return false
	}

	prefix := -1
	// проверяем крышку
	for i := x - 1; i >= x-h-1; i-- {
		prefix++
		for j := y; j <= y+2*h+w; j++ {
			if i != x-h-1 && j == y+prefix {
				if (*field)[i][j].Symbol != "/" {
					return false
				}
			}
			if i != x-h-1 && j == y+2*h+w-1-prefix {
				if (*field)[i][j].Symbol != "\\" {
					return false
				}
			}
			if i == x-h-1 && y+prefix <= j && j <= y+2*h+w-1-prefix {
				if (*field)[i][j].Symbol != "_" {
					return false
				}
			}
		}
	}
	// проверяем донышко
	prefix = -1
	for i := x; i <= x+h-1; i++ {
		prefix++
		for j := y; j <= y+2*h+w; j++ {
			if j == y+prefix {
				if (*field)[i][j].Symbol != "\\" {
					return false
				}
			}
			if j == y+2*h+w-1-prefix {
				if (*field)[i][j].Symbol != "/" {
					return false
				}
			}
			if i == x+h-1 && y+prefix < j && j < y+2*h+w-1-prefix {
				if (*field)[i][j].Symbol != "_" {
					return false
				}
			}
		}
	}

	return true
}

/*
// outputing выводит результат
func outputing(out *bufio.Writer, message string) {

	fmt.Fprintf(out, "%s", message)

}
*/

func outputing(out *bufio.Writer, arr []Peak) {

	for i := 0; i < len(arr); i++ {
		fmt.Fprintf(out, "%v\n", arr[i])
	}
}

func main() {

	// определяем ввод
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	// определяем вывод
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	inputCalc(scanner, out)
}
