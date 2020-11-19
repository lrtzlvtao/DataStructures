package sort

// PriorityQueue struct
type PriorityQueue struct {
	queue []int64
	count int
}

// Len return length
func (pq *PriorityQueue) Len() int {
	return pq.count
}

// Sort process sort
func (pq *PriorityQueue) Sort() {

}

// Exch exchange elements
func (pq *PriorityQueue) Exch(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
}

// Less value compare
func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.queue[i] < pq.queue[j]
}

// Swim element swim check
func (pq *PriorityQueue) Swim(k int) {}

// Sink element sink check
func (pq *PriorityQueue) Sink(k int) {}

// Insert add element
func (pq *PriorityQueue) Insert(v int64) {}

// DelMax delete max element
func (pq *PriorityQueue) DelMax() (int64, error) {
	return 0, nil
}

// DelMin delete min element
func (pq *PriorityQueue) DelMin() (int64, error) {
	return 0, nil
}
