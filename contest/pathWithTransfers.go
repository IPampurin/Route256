/*
Условие задачи
В Nozo разрабатывается новая компьютерная игра — «Колонизация 8». Вам поручено реализовать поиск пути для универсальных юнитов.
Поле игры состоит из шестиугольников, которыми можно полностью покрыть бесконечную плоскость. Все шестиугольники на поле игры делятся на 2 типа — суша и море.
Каждый шестиугольник, у которого есть 6 сторон, считается сушей. Остальные шестиугольники — морем.
Вам дана таблица из символов, состоящая из n строк. Каждая строка содержит m символов. Границы шестиугольников обозначаются символами «\» — обратный слеш,
«/» — прямой слеш и «_» — нижнее подчёркивание. Остальные символы поля — пробелы.
Все шестиугольники на поле имеют одинаковый размер. Обозначим за ширину шестиугольника количество подряд идущих символов «_» в его верхней и нижней сторонах.
Аналогично назовём высотой шестиугольника количество символов «/» в левой верхней и правой нижней сторонах.

Вам даны два шестиугольника суши. Определите минимальное количество раз, которое вам придётся переместиться из шестиугольника суши в шестиугольник моря
и наоборот на пути между двумя данными шестиугольниками. Количество перемещений между шестиугольниками одного типа не имеет значения.
Например, между двумя шестиугольниками суши. Из одного шестиугольника можно перейти в любой из шести соседних шестиугольников.

Шестиугольники называются соседними, если имеют общую границу. Карта для игры бесконечная, и юниты могут выходить за границы поля.
Все шестиугольники, входящие в таблицу символов частично или не входящие совсем, являются шестиугольниками моря.

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

Высота и ширина каждого шестиугольника не больше 10.

Выходные данные
Для каждого набора входных данных выведите минимальное количество раз, которое вам придётся переместиться из шестиугольника суши в шестиугольник моря
и наоборот на пути между двумя данными шестиугольниками.

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
4 13
 _   _
/ \ / \_   _
\_/ \_/ \ / \
      \_/ \_/
2 2
3 12
15 15
 _           _
/ \    _    / \
\_/   / \   \_/
   _  \_/
  / \_/ \_
  \_/ \_/ \_
 _/ \_/ \_/ \
/ \_/  _  \_/
\_/ \ / \ / \
  \_/ \_/ \_/
  / \_   _/ \
  \_/ \_/ \_/
 _  \_/ \_/ \
/ \   \_/ \_/
\_/     \_/
2 2
9 8

Выходные данные
0
2
2
4

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
0
0
2
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Dot описывает ячейку поля
type Dot struct {
	Symbol   string // какой символ был на вводе
	NamePeak int    // свяжем точку массива с вершиной графа по номеру вершины
}

// Peak описывает вершину графа, в которую мы превращаем шестиугольник суши, моря или окружающее поле
type Peak struct {
	Link  map[int]struct{} // связи с суседями
	Badge string           // признак отнесения к суше или морю (earth - земля, water - вода)
	Name  int              // имя (или номер) вершиины графа (идентификатор шестиугольника)
	x     int              // индекс строки условного начала шестиугольника
	y     int              // индеск столбца условного начала шестиугольника
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

		// считываем количество строк и столбцов
		n, m := inpTwoInt(sc)

		var x, y int // x,y координаты, от которых будем строить сетку шестиугольников
		need := true // need признак необходимости поиска x и y

		// поле для отображения карты
		field := make([][]Dot, n, n)

		// построчно сканируем ввод и посимвольно вписываем в field
		for i := 0; i < len(field); i++ {
			field[i] = make([]Dot, m, m)
			sc.Scan()
			for j, val := range sc.Text() {
				dot := Dot{
					Symbol:   string(val),
					NamePeak: -1, // при первоначальном заполнении массива все точки свяжем с не существующей вершиной графа
				}
				field[i][j] = dot
				// если попадается признак левой оконечности шестиугольника, запоминаем её координаты
				if need && i > 0 && field[i][j].Symbol == "\\" && field[i-1][j].Symbol == "/" {
					x, y = i, j
					need = false
				}
			}
		}

		// считываем координаты стартовой и финальной точек
		startX, startY := inpTwoInt(sc)
		finishX, finishY := inpTwoInt(sc)

		// попробуем ускорить работу программы
		if startX == finishX && startY == finishY {
			outputing(out, 0)
			return
		}

		// height полувысота шестиугольника, width длина основания шестиугольника
		// определяем параметры сетки шестиугольников исходя из предположения, что на поле есть хотя бы один шестиугольник
		height, width := sizeHex(x, y, &field)

		// расширим начальное поле, чтобы влезали окружающие шестиугольники моря
		newField, offsetX, offsetY := newFieldInit(height, width, &field)

		// приведём координаты для отсчёта сетки шестиугольников к расширенному полю
		x += offsetX
		y += offsetY

		// получаем перечень вершин графа с начальными значениями
		graph := netAndGraphInit(height, width, x, y, &newField)

		// уточняем признак отнесения вершин графа и заполняем внутренности шестиугольников номерами вершин, которые эти внутренности отражают
		for key, _ := range graph {
			if validHexEarth(height, width, graph[key].x, graph[key].y, &newField) {
				graph[key].Badge = "earth"
			}
			hexNumingPrint(height, width, graph[key].x, graph[key].y, graph[key].Name, &newField)
			linkBuilder(height, width, &newField, graph[key])
		}

		startPeak := newField[startX+offsetX-1][startY+offsetY-1].NamePeak
		finishPeak := newField[finishX+offsetX-1][finishY+offsetY-1].NamePeak

		// попробуем ускорить работу программы
		if startPeak == finishPeak {
			outputing(out, 0)
			return
		}

		//	outputingNumPeak(out, newField)
		outputingGraph1(out, graph, &newField)
		outputingGraph(out, graph)
		//	outputingSymbol(out, newField)
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

// newFieldInit расширяет первоначальную карту, чтобы на неё вместились окружающие шестиугольники моря
func newFieldInit(h, w int, field *[][]Dot) ([][]Dot, int, int) {

	// учтём смещение
	offsetX := (h + 1)   // смещение по строкам в связи с увеличением массива
	offsetY := (2*h + w) // смещение по столбцам в связи с увеличением массива

	newField := make([][]Dot, len(*field)+2*offsetX, len(*field)+2*offsetX)
	for i := 0; i < len(newField); i++ {
		newField[i] = make([]Dot, len((*field)[0])+2*offsetY, len((*field)[0])+2*offsetY)
		for j := 0; j < len(newField[i]); j++ {
			dot := Dot{
				Symbol:   " ",
				NamePeak: -1,
			}
			newField[i][j] = dot
		}
	}

	// допишем в новую карту старые данные с учётом смещения
	for i := 0; i < len(*field); i++ {
		for j := 0; j < len((*field)[i]); j++ {
			newField[i+offsetX][j+offsetY] = (*field)[i][j]
		}
	}

	return newField, offsetX, offsetY
}

// netAndGraphInit строит сетку точек начал шестиугольников на увеличенном поле и привязывает к каждой точке вершину графа
func netAndGraphInit(h, w, xOne, yOne int, newField *[][]Dot) map[int]*Peak {

	// graph это набор вершин графа, где вершина это шестиугольник ячеек поля, объединённых одним признаком (суша или море)
	graph := make(map[int]*Peak)

	// координаты сетки по условно нечётным рядам
	xTwo := xOne + h
	yTwo := yOne + h + w

	// смещаемся в верхний левый угол сетки для условно чётных рядов с учётом припуска на окружающие шестиугольники бесконечного моря
	for xOne-2*h >= 0 {
		xOne -= 2 * h
	}
	for yOne-2*(h+w) >= 0 {
		yOne -= 2 * (h + w)
	}

	// смещаемся в верхний левый угол сетки для условно нечётных рядов с учётом припуска на окружающие шестиугольники бесконечного моря
	for xTwo-2*h >= 0 {
		xTwo -= 2 * h
	}
	for yTwo-2*(h+w) >= 0 {
		yTwo -= 2 * (h + w)
	}

	// шагаем по сетке по условно чётным рядам и к каждой точке привязываем вершину графа
	for i := xOne; i < len((*newField)); i += 2 * h {
		for j := yOne; j < len((*newField)[i]); j += 2 * (h + w) {
			peak1 := Peak{
				Link:  make(map[int]struct{}),
				Badge: "water",
				Name:  len(graph),
				x:     i,
				y:     j,
				Mark:  false,
			}
			graph[peak1.Name] = &peak1              // добавляем предварительные данные по шестиугольнику в текущем ряду
			(*newField)[i][j].NamePeak = peak1.Name // указываем номер вершины в точке начала шестиугольника
		}
	}

	// шагаем по сетке по условно чётным рядам и к каждой точке привязываем вершину графа
	for i := xTwo; i < len((*newField)); i += 2 * h {
		for j := yTwo; j < len((*newField)[i]); j += 2 * (h + w) {
			peak2 := Peak{
				Link:  make(map[int]struct{}),
				Badge: "water",
				Name:  len(graph),
				x:     i,
				y:     j,
				Mark:  false,
			}
			graph[peak2.Name] = &peak2              // добавляем предварительные данные по шестиугольнику в соседнем ряду
			(*newField)[i][j].NamePeak = peak2.Name // указываем номер вершины в точке начала шестиугольника
		}
	}

	return graph
}

// validHexEarth проверяет наличие всех шести сторон шестиугольника определённых размеров по заданным координатам в массиве
func validHexEarth(h, w, x, y int, newField *[][]Dot) bool {

	// если при проверке шестиугольника есть выход за правую границу массива, то валидировать нечего
	if y+2*h+w-1 > len((*newField)[0])-1 {
		return false
	}
	// если при проверке шестиугольника есть выход за верхнюю или нижнюю границы массива, то валидировать нечего
	if x-h-1 < 0 || x+h-1 > len(*newField)-1 {
		return false
	}

	// проверяем наличие правой нижней стороны
	if (*newField)[x][y+2*h+w-1].Symbol != "/" {
		return false
	}
	// проверяем наличие правой верхней стороны
	if (*newField)[x-1][y+2*h+w-1].Symbol != "\\" {
		return false
	}
	// проверяем наличие верхней стороны
	if (*newField)[x-h-1][y+h].Symbol != "_" {
		return false
	}
	// проверяем наличие нижней стороны
	if (*newField)[x+h-1][y+h].Symbol != "_" {
		return false
	}

	return true
}

// hexNumingPrint уточняет номера вершин внутри шестиугольника
func hexNumingPrint(h, w, x, y, numPeak int, newField *[][]Dot) {

	prefix := 0
	// идём по крышке
	for i := x - 1; i >= x-h-1; i-- {
		for j := y; j <= y+2*h+w-1; j++ {
			if i >= 0 && i < len(*newField) && j >= 0 && j < len((*newField)[0]) {
				if i != x-h-1 && y+prefix < j && j < y+2*h+w-1-prefix {
					(*newField)[i][j].NamePeak = numPeak
				}
			}
		}
		prefix++
	}

	// идём по донышку
	prefix = 0
	for i := x; i <= x+h-1; i++ {
		for j := y; j <= y+2*h+w-1; j++ {
			if i >= 0 && i < len(*newField) && j >= 0 && j < len((*newField)[0]) {
				if i != x+h-1 && y+prefix < j && j < y+2*h+w-1-prefix {
					(*newField)[i][j].NamePeak = numPeak
				}
			}
		}
		prefix++
	}
}

// linkBuilder корректирует мапу связей вершины графа
func linkBuilder(h, w int, field *[][]Dot, peak *Peak) {

	// переобозначим для удобства
	x := peak.x
	y := peak.y

	// добавляем в мапу связей вершины данные по соседним шести шестиугольникам

	var num int // номер вершины шестиугольника-соседа
	// для шестиугольника выше описанного в peak
	if x-2*h >= 0 {
		if num = (*field)[x-2*h][y].NamePeak; num != -1 {
			peak.Link[num] = struct{}{}
		}
	}
	// для шестиугольника ниже, чем в peak
	if x+2*h < len(*field) {
		if num = (*field)[x+2*h][y].NamePeak; num != -1 {
			peak.Link[num] = struct{}{}
		}
	}
	// для шестиугольника выше и левее
	if x-h >= 0 && y-(h+w) >= 0 {
		if num = (*field)[x-h][y-(h+w)].NamePeak; num != -1 {
			peak.Link[num] = struct{}{}
		}
	}
	// для шестиугольника ниже и левее
	if x+h < len(*field) && y-(h+w) >= 0 {
		if num = (*field)[x+h][y-(h+w)].NamePeak; num != -1 {
			peak.Link[num] = struct{}{}
		}
	}
	// для шестиугольника выше и правее
	if x-h >= 0 && y+(h+w) < len((*field)[0]) {
		if num = (*field)[x-h][y+(h+w)].NamePeak; num != -1 {
			peak.Link[num] = struct{}{}
		}
	}
	// для шестиугольника ниже и правее
	if x+h < len(*field) && y+(h+w) < len((*field)[0]) {
		if num = (*field)[x+h][y+(h+w)].NamePeak; num != -1 {
			peak.Link[num] = struct{}{}
		}
	}
}

// для теста
func hexPrint(h, w, x, y, numPeak int, newField *[][]Dot) {

	prefix := 0
	// идём по крышке
	for i := x - 1; i >= x-h-1; i-- {
		for j := y; j <= y+2*h+w-1; j++ {
			if i >= 0 && i < len(*newField) && j >= 0 && j < len((*newField)[0]) {
				if i != x-h-1 && j == y+prefix {
					(*newField)[i][j].Symbol = "/"
				}
				if i != x-h-1 && y+prefix < j && j < y+2*h+w-1-prefix {
					(*newField)[i][j].Symbol = " "
				}
				if i != x-h-1 && j == y+2*h+w-1-prefix {
					(*newField)[i][j].Symbol = "\\"
				}
				if i == x-h-1 && y+prefix <= j && j <= y+2*h+w-1-prefix {
					(*newField)[i][j].Symbol = "_"
				}
			}
		}
		prefix++
	}
	// идём по донышку
	prefix = 0
	for i := x; i <= x+h-1; i++ {
		for j := y; j <= y+2*h+w-1; j++ {
			if i >= 0 && i < len(*newField) && j >= 0 && j < len((*newField)[0]) {
				if j == y+prefix {
					(*newField)[i][j].Symbol = "\\"
				}
				if i != x+h-1 && y+prefix < j && j < y+2*h+w-1-prefix {
					(*newField)[i][j].Symbol = " "
				}
				if j == y+2*h+w-1-prefix {
					(*newField)[i][j].Symbol = "/"
				}
				if i == x+h-1 && y+prefix < j && j < y+2*h+w-1-prefix {
					(*newField)[i][j].Symbol = "_"
				}
			}
		}
		prefix++
	}
}

func outputingGraph(out *bufio.Writer, graph map[int]*Peak) {

	keys := slices.Sorted(maps.Keys(graph))

	for _, val := range keys {
		fmt.Fprintf(out, "У вершины %v связи: %v\n", graph[val].Name, graph[val].Link)
	}
	fmt.Fprint(out, "\n")
}

func outputingGraph1(out *bufio.Writer, graph map[int]*Peak, newField *[][]Dot) {

	arr := make([][]string, len(*newField), len(*newField))
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]string, len((*newField)[0]))
		for j := 0; j < len(arr[i]); j++ {
			if (*newField)[i][j].Symbol == " " {
				arr[i][j] = "   "
			} else {
				arr[i][j] = " " + (*newField)[i][j].Symbol
			}
		}
	}

	keys := slices.Sorted(maps.Keys(graph))

	for _, val := range keys {

		arr[graph[val].x][graph[val].y] = strconv.Itoa(graph[val].Name)
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Fprintf(out, "%v", arr[i][j])
		}
		fmt.Fprint(out, "\n")
	}
	fmt.Fprint(out, "\n")

	fmt.Fprint(out, "\n")
}

/*
func outputingNumPeak(out *bufio.Writer, arr [][]Dot) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Fprintf(out, "%v", arr[i][j].NamePeak)
		}
		fmt.Fprint(out, "\n")
	}
	fmt.Fprint(out, "\n")
}
*/
/*
func outputingSymbol(out *bufio.Writer, arr [][]Dot) {

		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr[i]); j++ {
				fmt.Fprintf(out, "%v", arr[i][j].Symbol)
			}
			fmt.Fprint(out, "\n")
		}
		fmt.Fprint(out, "\n")
	}
*/

// outputing выводит результат
func outputing(out *bufio.Writer, message int) {

	fmt.Fprintf(out, "%v\n", message)
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
