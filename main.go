package main

type Hello struct {
	Word string
}

func (h *Hello) Print() {
	print("Hello" + h.Word)
}
