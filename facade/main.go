package main

import "fmt"

type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{width, height, make([]rune, width*height)}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

type Viewport struct {
	buffer *Buffer
	offset int
}

func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{buffer: buffer}
}

func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

type Console struct {
	buffers   []*Buffer
	viewports []*Viewport
	offset    int
}

// NewConsole creates a new console with a buffer and a viewport
/**
 * The Console class is a facade that provides a simple interface to the client. It hides the complexity of the Buffer and
 * Viewport classes. The client can interact with the Console class without knowing the details of the Buffer and Viewport
 * classes.
 */
func NewConsole() *Console {
	b := NewBuffer(200, 150)
	v := NewViewport(b)
	return &Console{
		buffers:   []*Buffer{b},
		viewports: []*Viewport{v},
		offset:    0,
	}
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}

/**
 * The client can use the NewConsole, to get rid of the complexity of the Buffer and Viewport classes.
 * If they want, they can use the Buffer and Viewport classes directly.
 * The facade pattern is useful when you want to provide a simple interface to a complex subsystem.
 */
func main() {
	c := NewConsole()
	u := c.GetCharacterAt(1)
	fmt.Println(u)
}
