package cli

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
)

var (
	styleTitle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("81"))
	styleFrame   = lipgloss.NewStyle().Foreground(lipgloss.Color("60"))
	styleDim     = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	styleSuccess = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	styleWarning = lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
	styleError   = lipgloss.NewStyle().Foreground(lipgloss.Color("203"))
	styleSelect  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("231")).Background(lipgloss.Color("61"))
	styleInput   = lipgloss.NewStyle().Foreground(lipgloss.Color("81"))
	border       = lipgloss.RoundedBorder()
)

const (
	minimumWidth  = 60
	minimumHeight = 20
)

func fixed(value string, width int) string {
	if width <= 0 {
		return ""
	}
	value = ansi.Truncate(value, width, "…")
	if current := lipgloss.Width(value); current < width {
		value += strings.Repeat(" ", width-current)
	}
	return value
}

func selectedLine(value string, width int, selected bool) string {
	value = fixed(value, width)
	if selected {
		return styleSelect.Render(value)
	}
	return value
}

func visibleWindow(cursor, count, height int) (int, int) {
	if count <= 0 || height <= 0 {
		return 0, 0
	}
	cursor = max(0, min(cursor, count-1))
	start := max(0, cursor-height/2)
	if start+height > count {
		start = max(0, count-height)
	}
	return start, min(count, start+height)
}

func pane(title string, focused bool, width, height int, rows []string) []string {
	if width < 1 || height < 1 {
		return nil
	}
	heading := styleDim.Render(title)
	if focused {
		heading = styleTitle.Render(title)
	}
	result := make([]string, height)
	result[0] = fixed(" "+heading, width)
	for index := 1; index < height; index++ {
		row := ""
		if index-1 < len(rows) {
			row = rows[index-1]
		}
		result[index] = fixed(row, width)
	}
	return result
}

func joinPanes(height int, panes ...[]string) []string {
	rows := make([]string, height)
	for row := 0; row < height; row++ {
		var line strings.Builder
		for index, panel := range panes {
			if index > 0 {
				line.WriteString(styleFrame.Render("│"))
			}
			if row < len(panel) {
				line.WriteString(panel[row])
			}
		}
		rows[row] = line.String()
	}
	return rows
}

func wrapLines(value string, width int) []string {
	if width < 1 {
		return nil
	}
	return strings.Split(ansi.Wordwrap(value, width, ""), "\n")
}

// safeText removes terminal control sequences from backend and user-controlled
// text before the application applies its own styles.
func safeText(value string, multiline bool) string {
	value = ansi.Strip(value)
	var out strings.Builder
	for _, character := range value {
		switch {
		case character == '\n' && multiline:
			out.WriteRune(character)
		case character == '\t' && multiline:
			out.WriteRune(character)
		case character == '\n' || character == '\r' || character == '\t':
			out.WriteByte(' ')
		case character < 0x20 || character >= 0x7f && character <= 0x9f:
			continue
		default:
			out.WriteRune(character)
		}
	}
	return out.String()
}

func safeLine(value string) string { return safeText(value, false) }

func safeMultiline(value string) string { return safeText(value, true) }
