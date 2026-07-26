package render

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type Document struct {
	Front   any
	Body    string
	IsError bool
}

func (d Document) String() (string, error) {
	var out strings.Builder
	if d.Front != nil {
		raw, err := yaml.Marshal(d.Front)
		if err != nil {
			return "", fmt.Errorf("render frontmatter: %w", err)
		}
		out.WriteString("---\n")
		out.Write(raw)
		out.WriteString("---\n")
		if d.Body != "" {
			out.WriteByte('\n')
		}
	}
	out.WriteString(d.Body)
	if d.Body != "" && !strings.HasSuffix(d.Body, "\n") {
		out.WriteByte('\n')
	}
	if out.Len() > 1<<20 {
		return "", fmt.Errorf("rendered output exceeds 1 MiB")
	}
	return out.String(), nil
}

func Fence(content, language string) string {
	longest := 2
	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimLeft(line, " \t")
		n := 0
		for n < len(trimmed) && trimmed[n] == '~' {
			n++
		}
		if n > longest {
			longest = n
		}
	}
	fence := strings.Repeat("~", longest+1)
	var out strings.Builder
	out.WriteString(fence)
	out.WriteString(language)
	out.WriteByte('\n')
	out.WriteString(content)
	if !strings.HasSuffix(content, "\n") {
		out.WriteByte('\n')
	}
	out.WriteString(fence)
	return out.String()
}
