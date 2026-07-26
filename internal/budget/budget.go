// Package budget provides deterministic o200k_base token accounting.
package budget

import (
	"errors"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/tiktoken-go/tokenizer"
)

var (
	codecOnce sync.Once
	codec     tokenizer.Codec
	codecErr  error
)

func encoding() (tokenizer.Codec, error) {
	codecOnce.Do(func() {
		codec, codecErr = tokenizer.Get(tokenizer.O200kBase)
	})
	return codec, codecErr
}

// Count returns the exact number of o200k_base tokens in text.
func Count(text string) (int, error) {
	enc, err := encoding()
	if err != nil {
		return 0, err
	}
	return enc.Count(text)
}

// Truncate returns a valid UTF-8 prefix bounded by both tokens and bytes.
// The returned token count describes the returned prefix, not either limit.
func Truncate(text string, tokenLimit, byteLimit int) (prefix string, tokens int, truncated bool, err error) {
	if !utf8.ValidString(text) {
		return "", 0, false, errors.New("token budget input is not valid UTF-8")
	}
	if tokenLimit <= 0 || byteLimit <= 0 {
		return "", 0, text != "", nil
	}

	bounded := text
	if len(bounded) > byteLimit {
		bounded = validPrefix(bounded, byteLimit)
		truncated = true
	}
	enc, err := encoding()
	if err != nil {
		return "", 0, false, err
	}

	probeLimit := 4 << 10
	if tokenLimit <= (int(^uint(0)>>1))/(8) && tokenLimit*8 > probeLimit {
		probeLimit = tokenLimit * 8
	}
	probeLimit = min(probeLimit, len(bounded))
	candidate := validPrefix(bounded, probeLimit)
	var ids []uint
	var pieces []string
	for {
		ids, pieces, err = enc.Encode(candidate)
		if err != nil {
			return "", 0, false, err
		}
		if len(ids) > tokenLimit || len(candidate) == len(bounded) {
			break
		}
		next := min(len(bounded), max(len(candidate)+1, len(candidate)*2))
		candidate = validPrefix(bounded, next)
	}
	if len(ids) <= tokenLimit {
		return candidate, len(ids), truncated || len(candidate) != len(text), nil
	}

	truncated = true
	prefix = validPrefix(strings.Join(pieces[:tokenLimit], ""), len(candidate))
	tokens, err = enc.Count(prefix)
	if err != nil {
		return "", 0, false, err
	}
	for tokens > tokenLimit && prefix != "" {
		_, size := utf8.DecodeLastRuneInString(prefix)
		prefix = prefix[:len(prefix)-size]
		tokens, err = enc.Count(prefix)
		if err != nil {
			return "", 0, false, err
		}
	}
	return prefix, tokens, true, nil
}

func validPrefix(text string, limit int) string {
	if limit <= 0 {
		return ""
	}
	prefix := text
	if limit < len(prefix) {
		prefix = prefix[:limit]
	}
	for !utf8.ValidString(prefix) {
		prefix = prefix[:len(prefix)-1]
	}
	return prefix
}

// Paginate groups complete records by the token count of render(records).
// A record that exceeds the budget by itself remains as the sole page record.
func Paginate[T any](records []T, requestedPage, tokenLimit int, render func([]T) (string, error)) ([]T, int, error) {
	if len(records) == 0 {
		return nil, 0, nil
	}
	starts := []int{0}
	start := 0
	for end := 1; end <= len(records); end++ {
		representation, err := render(records[start:end])
		if err != nil {
			return nil, 0, err
		}
		used, err := Count(representation)
		if err != nil {
			return nil, 0, err
		}
		if end-start > 1 && used > tokenLimit {
			start = end - 1
			starts = append(starts, start)
		}
	}

	totalPages := len(starts)
	if requestedPage < 1 || requestedPage > totalPages {
		return nil, totalPages, nil
	}
	start = starts[requestedPage-1]
	end := len(records)
	if requestedPage < totalPages {
		end = starts[requestedPage]
	}
	return append([]T(nil), records[start:end]...), totalPages, nil
}
