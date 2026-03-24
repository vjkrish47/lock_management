type Lock struct {
	mu    sync.Mutex
	queue []int
	owner int
}

func NewLock() *Lock {
	return &Lock{owner: -1}
}

func (l *Lock) Acquire(id int) {
	l.mu.Lock()
	l.queue = append(l.queue, id)
	for l.owner != -1 || l.queue[0] != id {
		l.mu.Unlock()
		l.mu.Lock()
	}
	l.queue = l.queue[1:]
	l.owner = id
	l.mu.Unlock()
}

func (l *Lock) Release(id int) {
	l.mu.Lock()
	if l.owner == id {
		l.owner = -1
	}
	l.mu.Unlock()
}

func worker(id int, l *Lock, wg *sync.WaitGroup) {
	defer wg.Done()
	l.Acquire(id)
	fmt.Println("Worker", id, "entered")
	l.Release(id)
}

func main() {
	l := NewLock()
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, l, &wg)
	}
	wg.Wait()
}

