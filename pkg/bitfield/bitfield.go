package bitfield

import (
	"fmt"
)

// Implementations for uint8, uint16
type Bitfield8 uint8
type Bitfield16 uint16

// Writes a true value to the bits at the specified positions.
func (b *Bitfield8) ManySet(pos ...uint) {
	for _, p := range pos {
		b.Set(p)
	}
}

// Writes a false value to the bits at the specified positions.
func (b *Bitfield8) ManyClear(pos ...uint) {
	for _, p := range pos {
		b.Clear(p)
	}
}

// Toggles the value on the bits at the specified positions.
func (b *Bitfield8) ManyToggle(pos ...uint) {
	for _, p := range pos {
		b.Toggle(p)
	}
}

// Writes a true value to the bit at the specified position.
func (b *Bitfield8) Set(pos uint) {
	*b |= 1 << pos
}
func (b *Bitfield16) Set(pos uint) {
	*b |= 1 << pos
}

// Writes a false value to the bit at the specified position.
func (b *Bitfield8) Clear(pos uint) {
	*b &^= 1 << pos
}
func (b *Bitfield16) Clear(pos uint) {
	*b &^= 1 << pos
}

// Toggles the value on the bit at the specified position.
func (b *Bitfield8) Toggle(pos uint) {
	*b ^= 1 << pos
}
func (b *Bitfield16) Toggle(pos uint) {
	*b ^= 1 << pos
}

// Returns the value on the bit at the specified position.
func (b Bitfield8) Read(pos uint) bool {
	return b&(1<<pos) != 0
}
func (b Bitfield16) Read(pos uint) bool {
	return b&(1<<pos) != 0
}

// Reads the bitfield as a string
func (b Bitfield8) String() string {
	return fmt.Sprintf("%08b", b)
}
func (b Bitfield16) String() string {
	return fmt.Sprintf("%16b", b)
}
