package md5

import (
	"encoding/binary"
	"math"
)

func Hash(in []byte) [16]byte {
	// Initialise the table of constants
	var ts [64]uint32
	for i := 0; i < 64; i++ {
		ts[i] = uint32(4294967296 * math.Abs(math.Sin(float64(i+1))))
	}

	paddedMsg := padMsg(in)

	var (
		a uint32 = 0x67452301
		b uint32 = 0xEFCDAB89
		c uint32 = 0x98BADCFE
		d uint32 = 0x10325476
	)

	// Process the message in blocks of 512 bits
	for i := 0; i < len(paddedMsg); i += 64 {
		a, b, c, d = processBlock(ts, paddedMsg[i:i+64], a, b, c, d)
	}

	// Reassemble the processed 32 bit vars
	digest := make([]byte, 16)
	binary.LittleEndian.PutUint32(digest[0:4], a)
	binary.LittleEndian.PutUint32(digest[4:8], b)
	binary.LittleEndian.PutUint32(digest[8:12], c)
	binary.LittleEndian.PutUint32(digest[12:16], d)

	var res [16]byte
	copy(res[:], digest)

	return res
}

// padMsg pads out a input message with one followed by zeros such that the length of the resulting message
// is a multiple of 512.
func padMsg(msg []byte) []byte {
	msgLen := len(msg)

	// Add single 1 bit
	msg = append(msg, 0x80)

	// Pad out message with 0's until message length is less than 512-64 bits
	for len(msg)%64 != 56 {
		msg = append(msg, 0x00)
	}

	// Append the length of the message as 64 bits
	lenBytes := make([]byte, 8)
	msgLenInBits := uint64(msgLen * 8)
	binary.LittleEndian.PutUint64(lenBytes, msgLenInBits)
	msg = append(msg, lenBytes...)

	return msg
}

func processBlock(ts [64]uint32, block []byte, a, b, c, d uint32) (uint32, uint32, uint32, uint32) {
	x := make([]uint32, 16)
	for i := 0; i < 16; i++ {
		x[i] = binary.LittleEndian.Uint32(block[i*4 : (i+1)*4])
	}

	aa, bb, cc, dd := a, b, c, d

	s := []uint32{
		7, 12, 17, 22,
		5, 9, 14, 20,
		4, 11, 16, 23,
		6, 10, 15, 21,
	}

	for i := 0; i < 64; i++ {
		var f uint32
		var g int

		switch {
		case i < 16:
			f = (b & c) | (^b & d)
			g = i

		case i < 32:
			f = (d & b) | (^d & c)
			g = (5*i + 1) % 16

		case i < 48:
			f = b ^ c ^ d
			g = (3*i + 5) % 16

		default:
			f = c ^ (b | ^d)
			g = (7 * i) % 16
		}

		f = f + a + ts[i] + x[g]
		a = d
		d = c
		c = b
		b = b + leftRotate(f, s[(i%4)+(i/16)*4])
	}

	a += aa
	b += bb
	c += cc
	d += dd

	return a, b, c, d
}

func leftRotate(x, c uint32) uint32 {
	return (x << c) | (x >> (32 - c))
}
