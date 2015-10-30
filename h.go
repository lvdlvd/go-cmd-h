// h is a unix line filter that makes large integers readable by inserting k,m,g,t,p... before groups of 3 digits.
// Usage:
//      lvd$ cmd_that_generates_text --with-large=numbers | h | less
//      lvd$ echo bla:123456789 | h
//      bla:123m456k789
package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func isdigit(c byte) bool { return '0' <= c && c <= '9' }

func main() {

	wr := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := wr.Flush(); err != nil {
			log.Fatal("Flushing:", err)
		}
	}()

	w := func(c byte) {
		if err := wr.WriteByte(c); err != nil {
			log.Fatal("Writing:", err)
		}
	}

	const (
		INIT = iota
		DIGITS
		TOOLONG
	)
	state := INIT

	var buf [21]byte
	b := buf[:0]

	r := bufio.NewReader(os.Stdin)
	for {
		c, err := r.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Reading:", err)
		}

		switch state {
		case INIT:
			if !isdigit(c) {
				w(c)
				break
			}

			b = append(b, c)
			state = DIGITS

		case DIGITS:
			if !isdigit(c) {
				w(b[0])
				for i, c := range b[1:] {
					if (len(b)-i)%3 == 1 {
						w(" kmgtpe"[(len(b)-i)/3])
					}
					w(c)
				}
				b = b[:0]
				w(c)
				state = INIT
				break
			}

			if len(b) < cap(b) {
				b = append(b, c)
				break
			}

			for _, c := range b {
				w(c)
			}
			b = b[:0]
			w(c)
			state = TOOLONG

		case TOOLONG:
			w(c)
			if !isdigit(c) {
				state = INIT
			}

		}
	}

}
