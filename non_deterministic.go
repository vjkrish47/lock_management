type TryLock struct {
	ch chan struct{}
}
func NewTryLock() *TryLock {
	l := &TryLock{ch: make(chan struct{}, 1)}
	l.ch <  struct{}{}
	return l
}
func (l *TryLock) TryLock() bool {
	select {
	case < l.ch:
		return true
	default:
		return false
	}
}
func (l *TryLock) Unlock() {
	l.ch <  struct{}{}
}
func worker(id int, lock *TryLock, wg *sync.WaitGroup) {
	defer wg.Done()
	if !lock.TryLock() {
		fmt.Println("Worker", id, "exits: lock not acquired")
		return
	}
	fmt.Println("Worker", id, "acquired lock")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Worker", id, "releasing lock")
	lock.Unlock()
}
func main() {
	lock := NewTryLock()
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, lock, &wg)
	}
	wg.Wait()
}
