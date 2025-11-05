package tokeniser

import (
	"fmt"
	"io"
	"log/slog"
	"strings"
)

type tokeniser struct {
	Source string
	Tokens []string
	Runes  []rune
	Reader *strings.Reader
	Error  error
}

func Tokenise(query string) ([]string, error) {
	slog.Debug("> tokeniser.Tokenise")
	t := new(tokeniser)
	t.Source = query
	t.Tokenise()
	return t.Tokens, t.Error
}

func (t *tokeniser) Tokenise() {
	slog.Debug("Creating reader")
	t.Reader = strings.NewReader(t.Source)
	for {
		ch := t.Next()
		if t.Error != nil {
			return
		}
		if ch == -1 {
			break
		}
		switch ch {
		case ' ', '\n':
			t.NewWord()
		case '=', '(', ')':
			t.OnItsOwn(ch)
		case '-':
			// check if the next character is also a - for a comment
			ch2 := t.Next()
			if t.Error != nil {
				return
			}
			if ch2 == '-' {
				slog.Debug("Comment loop")
				// read until newline
				for {
					ch2 := t.Next()
					if t.Error != nil {
						return
					}
					if ch2 == '\n' {
						slog.Debug("Newline in comment loop")
						break
					}
				}
			} else {
				t.Unread() // put ch2 back
				t.Append(ch)
			}
		default:
			t.Append(ch)
		}
	}
	t.NewWord()
}

func (t *tokeniser) NewWord() {
	slog.Debug("> tokeniser.NewWord")
	if len(t.Runes) > 0 {
		slog.Debug("Appending " + string(t.Runes))
		t.Tokens = append(t.Tokens, string(t.Runes))
		t.Discard()
	}
}

func (t *tokeniser) OnItsOwn(ch rune) {
	slog.Debug("> tokeniser.OnItsOwn")
	t.NewWord()
	t.Append(ch)
	t.NewWord()
}

func (t *tokeniser) ReadUntil(ch rune) {
	slog.Debug("> tokeniser.ReadUntil(" + string(ch) + ")")
	for {
		ch2 := t.Next()
		if t.Error != nil {
			return
		}
		if ch2 == ch {
			break
		} else {
			t.Append(ch2)
		}
	}
}

func (t *tokeniser) Append(ch rune) {
	slog.Debug("> tokeniser.Append")
	t.Runes = append(t.Runes, ch)
}

func (t *tokeniser) Discard() {
	slog.Debug("> tokeniser.Discard")
	t.Runes = nil
}

func (t *tokeniser) Next() rune {
	slog.Debug("> tokeniser.Next")
	ch, _, err := t.Reader.ReadRune()

	if err != nil {
		if err == io.EOF {
			slog.Debug("EOF")
			return -1
		}
		t.Error = err
		slog.Debug("Errored")
		return -1
	}
	slog.Debug("Read: " + string(ch))
	return ch
}

func (t *tokeniser) Unread() {
	slog.Debug("> tokeniser.Unread")
	t.Reader.UnreadRune()
}

func Test(query string) {

	tokens, err := Tokenise(query)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		for _, t := range tokens {
			fmt.Println(t)
		}
	}
}
