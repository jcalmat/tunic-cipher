package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type ClickableCanvas struct {
	widget.DisableableWidget

	// actions
	OnTapped func() `json:"-"`

	hovered, focused bool
	tapAnim          *fyne.Animation

	// resource
	Image *canvas.Image `json:"image"`
}

func NewClickableCanvas(tapped func()) *ClickableCanvas {
	c := &ClickableCanvas{
		OnTapped: tapped,
	}
	c.ExtendBaseWidget(c)
	return c
}

func NewClickableCanvasWithImage(img *canvas.Image, tapped func()) *ClickableCanvas {
	c := NewClickableCanvas(tapped)
	c.Image = img
	return c
}

func (c *ClickableCanvas) CreateRenderer() fyne.WidgetRenderer {
	v := container.NewBorder(nil, nil, nil, nil, c.Image)
	return widget.NewSimpleRenderer(v)
}

// TypedKey is a hook called by the input handling logic on key events if this object is focused.
func (c *ClickableCanvas) TypedKey(ev *fyne.KeyEvent) {
	if ev.Name == fyne.KeySpace {
		c.Tapped(nil)
	}
}

func (c *ClickableCanvas) Tapped(ev *fyne.PointEvent) {
	if c.OnTapped != nil {
		c.OnTapped()
	}
}

// MouseIn is called when a desktop pointer enters the widget
func (c *ClickableCanvas) MouseIn(*desktop.MouseEvent) {
	c.hovered = true
	c.Refresh()
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (c *ClickableCanvas) MouseMoved(*desktop.MouseEvent) {
}

// MouseOut is called when a desktop pointer exits the widget
func (c *ClickableCanvas) MouseOut() {
	c.hovered = false
	c.Refresh()
}

// SetImage updates the wudget image - pass nil to hide an icon
func (c *ClickableCanvas) SetImage(img *canvas.Image) {
	// c.propertyLock.Lock()
	c.Image = img
	// c.propertyLock.Unlock()

	c.Refresh()
}
