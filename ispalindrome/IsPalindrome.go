package ispalindrome

import (
	"hometask2/deque"
	"hometask2/stack1"
)

func IsPalindrome1(s string) bool {
	stack := stack1.NewStack[rune]()
	for _, char := range s {
		stack.Push(char)
	}
	for _, char := range s {
		val, err := stack.Pop()
		if val != char || err != nil {
			return false
		}
	}
	return true
}

func IsPalindrome2(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func IsPalindrome3(s string) bool {
	d := deque.NewDeque[rune]()
	for _, char := range s {
		d.PushBack(char)
	}
	for d.Size > 1 {
		front, err := d.Front()
		if err != nil {
			return false
		}
		back, err := d.Back()
		if err != nil {
			return false
		}
		if front != back {
			return false
		}
		d.PopFront()
		d.PopBack()
	}
	return true
}
