package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/rs/zerolog"
)

type App struct {
	Logger *zerolog.Logger

	App    fyne.App
	Window fyne.Window
}

type Fyne struct {
}

func New(logger *zerolog.Logger) App {
	return App{
		Logger: logger,
		App:    app.New(),
	}
}
func (a *App) MakeTray() {
	logger := a.Logger.With().Str("method", "MakeTray").Logger()

	if desk, ok := a.App.(desktop.App); ok {
		h := fyne.NewMenuItem("Hello", func() {})
		h.Icon = theme.HomeIcon()
		menu := fyne.NewMenu("Hello World", h)
		h.Action = func() {
			logger.Debug().Msg("System tray menu tapped")
			h.Label = "Welcome"
			menu.Refresh()
		}
		desk.SetSystemTrayMenu(menu)
	}
}

func (a *App) CreateWindow() {
	a.Window = a.App.NewWindow("Tunic Translator")

	hello := widget.NewLabel("Hello Fyne!")
	a.Window.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	// TODO: set main menu

	a.Window.SetMaster()

	content := container.NewStack()
	title := widget.NewLabel("Welcome to Tunic Translator")
	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")

	setView := func(t View) {
		title.SetText(t.Title)
		intro.SetText(t.Intro)

		content.Objects = []fyne.CanvasObject{t.View(a.Window)}
		content.Refresh()
	}

	view := container.NewBorder(container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)

	// split screen with navigation
	split := container.NewHSplit(makeNav(setView), view)
	split.Offset = 0.2
	a.Window.SetContent(split)

	// TODO: try to store the size in preferences
	a.Window.Resize(fyne.NewSize(float32(a.App.Preferences().FloatWithFallback(preferenceWidth, 1280)),
		float32(a.App.Preferences().FloatWithFallback(preferenceHeight, 720))))

}

func makeNav(setView func(view View)) fyne.CanvasObject {
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return ViewIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := ViewIndex[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := Views[uid]
			if !ok {
				fyne.LogError("Missing view panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
		},
		OnSelected: func(uid string) {
			if t, ok := Views[uid]; ok {
				a.Preferences().SetString(preferenceCurrentView, uid)
				setView(t)
			}
		},
	}

	// TODO: move it to window menu
	themes := container.NewGridWithColumns(2,
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark})
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight})
		}),
	)

	// load saved view
	currentPref := a.Preferences().StringWithFallback(preferenceCurrentView, "transcriptor") //TODO: change to welcome
	tree.Select(currentPref)

	return container.NewBorder(nil, themes, nil, nil, tree)
}

func (a *App) Run() {
	a.Window.ShowAndRun()
}
