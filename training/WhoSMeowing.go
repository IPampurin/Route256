/*
Условие задачи
Чем дольше Иван слышал этот звук, тем больше сходил с ума! В отчаянии он решил отложить все свои дела, чтобы наконец вычислить,
кто из окружающих его коллег исподтишка выполняет это действие.
Для этого Иван подошёл к каждому из коллег и попросил высказать своё мнение о сложившейся ситуации.
Ответы коллег он записал в одном из следующих форматов:
1. «A: I am x!» — подозреваемый по имени A утверждает, что выполняет действие x.
2. «A: I am not x!» — подозреваемый по имени A утверждает, что не выполняет действие x.
3. «A: B is x!» — подозреваемый по имени A считает, что подозреваемый B выполняет действие x.
4. «A: B is not x!» — подозреваемый по имени A считает, что подозреваемый B не выполняет действие x.

Каждый из подозреваемых дал не более одного высказывания по каждому подозреваемому, включая себя. Если подозреваемый говорил про себя,
Иван точно записал его ответ в формате
«A: I am x!»
 или
«A: I am not x!».
Также Иван уверен, что не существует подозреваемых, которые не фигурируют в его записях: то есть тех, кто ничего не говорил сам, и про кого ничего не говорили другие.
Теперь, имея набор записей, Иван решил подсчитать очки каждого подозреваемого по следующему алгоритму:

• изначально у каждого подозреваемого 0 очков;
• каждое высказывание «A: I am x!» прибавляет подозреваемому A два очка;
• каждое высказывание «A: I am not x!» отнимает у подозреваемого A одно очко;
• каждое высказывание «A: B is x!» прибавляет подозреваемому B одно очко;
• каждое высказывание «A: B is not x!» отнимает у подозреваемого B одно очко.

В итоге Иван считает, что выполняют действие те люди, кто суммарно набрал наибольшее количество очков в сумме по всем высказываниям.
Количество очков может быть и отрицательным. Помогите Ивану рассчитать, кто выполняет это действие.
Входные данные
Первая строка входных данных содержит натуральное число t (1≤t≤10^5) — количество наборов входных данных.
Следующие строки содержат подряд идущие наборы входных данных. Рассмотрим очередной такой набор.
Первая строка набора входных данных содержит натуральное число n (1≤n≤10^5) — количество высказываний.
Далее идёт n строк в формате, описанном в условии задачи, где A и B обозначают имена подозреваемых, а x — какое-то действие.

Гарантируется, что:
• A, B и x состоят только из букв английского алфавита и не длиннее 10 символов. При этом x содержит только строчные буквы.
	A и B же начинаются с заглавной буквы, а все остальные буквы строчные.
• В рамках одного набора входных данных действие (x) не меняется между высказываниями.
• Сумма n по всем наборам входных данных не превышает 10^5.

Группы тестов

who-is-meowing-groups
Выходные данные
Для каждого набора входных данных выведите имя того, кто по описанному алгоритму набрал наибольшее количество очков, в формате
«A is x.», где A — это имя подозреваемого, а x — действие. Если наибольшее количество очков набрали несколько подозреваемых, выведите их всех в описанном формате в
лексикографическом порядке, разделяя предложения переносами строк.

Пояснение к примерам
Разберём, кому сколько очков принесёт каждое высказывание из первого набора входных данных в первом примере:
1. «Andrew: Boris is meowing!» прибавляет подозреваемому Boris одно очко.
2. «Boris: I am not meowing!» отнимает у подозреваемого Boris одно очко.
3. «Kate: Andrew is meowing!» прибавляет подозреваемому Andrew одно очко.
4. «Kate: Boris is not meowing!» отнимает у подозреваемого Boris одно очко.
5. «Kate: I am meowing!» прибавляет подозреваемому Kate два очка.

В итоге очки распределились между подозреваемыми следующим образом:
• Boris — -1 очко;
• Andrew — 1 очко;
• Kate — 2 очка.

Таким образом, наибольшее количество очков набрал подозреваемый Kate, поэтому нужно вывести «Kate is meowing.».

Во втором наборе входных данных оба подозреваемых (Sedan и Ivan) получили по два очка за высказывания «I am hungry!». Так как у них одинаковое количество очков,
нужно вывести обоих подозреваемых в лексикографическом порядке.

В третьем наборе подозреваемый I получает два очка за высказывание «I am serious!». Обратите внимание, «I» — это допустимое имя подозреваемого.
За аналогичное высказывание получает два очка и подозреваемый H, однако он также сказал и «I is serious!», ссылаясь на подозреваемого I, что добавляет последнему ещё одно очко.
Таким образом, подозреваемый I набрал 3 очка, а подозреваемый H — 2 очка.

Пример теста 1
Входные данные

3
5
Andrew: Boris is meowing!
Boris: I am not meowing!
Kate: Andrew is meowing!
Kate: Boris is not meowing!
Kate: I am meowing!
2
Sedan: I am hungry!
Ivan: I am hungry!
3
I: I am serious!
H: I is serious!
H: I am serious!
Выходные данные

Kate is meowing.
Ivan is hungry.
Sedan is hungry.
I is serious.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// определяем количество групп данных
	scanner.Scan()
	countGroup, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	// считываем ответы по группам
	for i := 0; i < countGroup; i++ {

		// определяем количество ответов в группе
		scanner.Scan()
		countComment, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		// ball - мапа для подсчёта очков по группе ответов
		ball := make(map[string]int)
		// выполняемое действие в каждой группе
		action := ""

		// построчно считываем комментарии в группе и подсчитываем баллы
		for i := 0; i < countComment; i++ {

			scanner.Scan()                                  // считываем строку с ответом
			line := strings.TrimSuffix(scanner.Text(), "!") // удаляем из неё восклицательный знак
			brokenLine := strings.Split(line, ": ")         // делим ответ по двоеточию
			name := brokenLine[0]                           // сохраняем имя отвечающего
			answer := strings.Split(brokenLine[1], " ")     // сохраняем сам ответ по словам

			// если подозреваемого нет в базе,
			if _, ok := ball[name]; !ok {
				ball[name] = 0
			}

			// определяем действие один раз
			if i == 0 {
				action = answer[len(answer)-1]
			}

			if len(answer) == 4 {
				if answer[1] == "am" {
					ball[name]--
				}
				if answer[1] == "is" {
					ball[answer[0]]--
				}
			}
			if len(answer) == 3 {
				if answer[1] == "am" {
					ball[name] += 2
				}
				if answer[1] == "is" {
					ball[answer[0]]++
				}
			}
		}

		// определяем максимум по значению в мапе баллов
		var maxNames []string
		maxValue := math.MinInt

		for key, val := range ball {
			if val > maxValue {
				maxValue = val
				maxNames = []string{key}
			} else if val == maxValue {
				maxNames = append(maxNames, key)
			}
		}

		// сортируем имена с максимумом баллов
		sort.Strings(maxNames)

		// выводим результат по группе
		for _, v := range maxNames {
			fmt.Fprintf(out, "%s is %s.\n", v, action)
		}
	}
}

/*
Ниже приведён вариант с многопоточкой. Протестировать его не удалось - превышен лимит отправки посылок.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Структура для хранения результатов обработки группы
type GroupResult struct {
	maxNames []string
	action   string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Читаем количество групп
	scanner.Scan()
	countGroup, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	wg.Add(countGroup)

	// Канал для получения результатов обработки групп
	results := make(chan GroupResult, countGroup)

	// Читаем все данные заранее
	allData := make([][]string, countGroup)
	for i := 0; i < countGroup; i++ {
		// Читаем количество комментариев в группе
		scanner.Scan()
		countComment, _ := strconv.Atoi(scanner.Text())

		// Читаем все строки группы
		allData[i] = make([]string, countComment)
		for j := 0; j < countComment; j++ {
			scanner.Scan()
			allData[i][j] = scanner.Text()
		}
	}

	// Обрабатываем каждую группу в отдельной горутине
	for i := 0; i < countGroup; i++ {
		go func(groupData []string) {
			defer wg.Done()

			ball := make(map[string]int)
			action := ""

			for _, line := range groupData {
				line = line[:len(line)-1]
				brokenLine := strings.Split(line, ": ")
				name := brokenLine[0]
				answer := strings.Split(brokenLine[1], " ")

				if _, ok := ball[name]; !ok {
					ball[name] = 0
				}

				if action == "" {
					action = answer[len(answer)-1]
				}

				if len(answer) == 4 {
					if answer[1] == "am" {
						ball[name]--
					}
					if answer[1] == "is" {
						ball[answer[0]]--
					}
				}
				if len(answer) == 3 {
					if answer[1] == "am" {
						ball[name] += 2
					}
					if answer[1] == "is" {
						ball[answer[0]]++
					}
				}
			}

			// Находим максимум
			var maxNames []string
			maxValue := math.MinInt
			for key, val := range ball {
				if val > maxValue {
					maxValue = val
					maxNames = []string{key}
				} else if val == maxValue {
					maxNames = append(maxNames, key)
				}
			}

			sort.Strings(maxNames)
			results <- GroupResult{
				maxNames: maxNames,
				action:   action,
			}
		}(allData[i])
	}

	// Ждем завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Выводим результаты по мере их поступления
	for result := range results {
		for _, v := range result.maxNames {
			fmt.Fprintf(out, "%s is %s.\n", v, result.action)
		}
	}
}

*/
