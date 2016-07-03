// sentence cutter.
package cutter

import "strings"

func Sentences(p string, punctuation bool) (s []string) {
	var start, offset int
	var word, sentence string

	size := len(p)

	for k, b := range p {
		word = string(b)
		offset += len(word)

		if strings.Contains(punctuations, word) {
			sentence = p[start:k]
			start = k + len(word)

			if len(sentence) == 0 {
				continue
			}

			if punctuation {
				sentence += word
			}

			s = append(s, sentence)

			continue
		}

		// End without punctuation.
		if start < size && offset == size && !strings.Contains(punctuations, word) {
			s = append(s, p[start:])
			break
		}
	}
	return
}
