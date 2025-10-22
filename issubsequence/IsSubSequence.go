package issubsequence

import "hometask2/queue1"

func IsSubSequence1(a, b string) bool {
	q := queue1.NewQueue[rune]()
	for _, el := range a {
		q.Enqueue(el)
	}
	for _, el := range b {
		val, err := q.Peek()
		if val == el && err == nil {
			q.Dequeue()
		}
	}
	return q.IsEmpty()
}

func IsSubSequence2(a, b string) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
		}
		j++
	}
	return i == len(a)
}
