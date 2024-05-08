package format

import (
	"fmt"

	"github.com/aarongodin/vpm/pkg/pack"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var HeaderStyle = lipgloss.NewStyle().Padding(0, 1).Bold(true)
var RowStyle = lipgloss.NewStyle().Padding(0, 1)
var InfoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))

func ShowPackageList(packs []pack.Pack) fmt.Stringer {
	t := table.New().
		StyleFunc(func(row, _ int) lipgloss.Style {
			switch {
			case row == 0:
				return HeaderStyle
			default:
				return RowStyle
			}
		}).
		Headers("name", "group", "load", "remote", "head")

	for _, p := range packs {
		t.Row(p.Name, p.Group, p.Load, p.RemoteURL, p.Head[:7])
	}

	return t
}
