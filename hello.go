package hello

type Hello struct {
	Word string
}

func (h *Hello) Print() {
	print(h.get())
}

func (h *Hello) get() string {
	return "Hello " + h.Word
}
