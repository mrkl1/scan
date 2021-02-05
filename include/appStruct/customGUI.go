package appStruct

import "github.com/therecipe/qt/widgets"

type CustomProgressBar struct {
	widgets.QProgressBar

	_ func() `constructor:"init"`

	_ func(int) `signal:"valueChangedFromGoroutine"`
}

type CustomLabel struct {
	widgets.QLabel

	_ func() `constructor:"init"`

	_ func(string) `signal:"updateTextFromGoroutine"`
	_ func() `signal:"adjustSizeFromGoroutine"`
}

type CustomButton struct {
	widgets.QPushButton


	_ func() `constructor:"init"`

	_ func(bool) `signal:"setEnabledFromGoroutine"`
}

func (c *CustomLabel) init() {
	c.ConnectUpdateTextFromGoroutine(c.SetText)
	c.ConnectAdjustSizeFromGoroutine(c.AdjustSize)
}

func (p *CustomProgressBar) init() {
	p.ConnectValueChangedFromGoroutine(p.SetValue)
}

func (b *CustomButton) init() {
	b.ConnectSetEnabledFromGoroutine(b.SetEnabled)
}