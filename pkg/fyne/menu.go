package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

func makeMenu(a fyne.App) *fyne.MainMenu {
	themeItem := fyne.NewMenuItem("Dark Theme", nil)
	themeItem.Checked = a.Settings().ThemeVariant() == theme.VariantDark
	file := fyne.NewMenu("File", themeItem)
	main := fyne.NewMainMenu(file)
	themeItem.Action = func() {
		themeItem.Checked = !themeItem.Checked
		if themeItem.Checked {
			a.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark})
		} else {
			a.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight})
		}
		main.Refresh()
	}

	return main
}
