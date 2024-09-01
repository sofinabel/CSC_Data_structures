package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pack struct {
	begin int
	end   int
}

type Queue struct {
	items []Pack
}

func NewQueue() *Queue {
	return &Queue{items: []Pack{}}
}

func (q *Queue) Enqueue(item Pack) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() Pack {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) First() Pack {
	item := q.items[0]
	return item
}

func (q *Queue) Last() Pack {
	item := q.items[q.Size()-1]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

/*в очереди храним времена окончания обработки пакетов.
Как пакет пришёл - выкидываем из очереди все пакеты, успевшие к этому моменту обработаться,
проверяем очередь на свободное место и вычисляем требуемое время начала обработки пакетов.
Это может быть или время прихода пакета (очередь пуста, обработка началась сразу),
или время окончания обработки последнего пакета из очереди.*/

func main() {
	var size, n int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &size, &n)
	beginTime := make([]int, n)
	duration := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &beginTime[i], &duration[i])
	}
	queue := NewQueue()
	if n == 0 {
		return
	}
	result := make([]int, n)
	queue.Enqueue(Pack{begin: beginTime[0], end: beginTime[0] + duration[0]})
	result[0] = queue.Last().begin
	for i := 1; i < n; i++ {
		//выкидываем все пакеты, успевшие обработаться
		for !queue.IsEmpty() && queue.First().end <= beginTime[i] {
			queue.Dequeue()
		}

		if queue.Size() < size {
			if queue.IsEmpty() {
				queue.Enqueue(Pack{begin: beginTime[i], end: beginTime[i] + duration[i]})
			} else {
				last := queue.Last().end
				queue.Enqueue(Pack{begin: last, end: last + duration[i]})
			}
			result[i] = queue.Last().begin
		} else {
			result[i] = -1
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, result[i])
	}
}
