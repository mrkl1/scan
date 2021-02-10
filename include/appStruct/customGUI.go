package appStruct

import "github.com/therecipe/qt/widgets"

type CustomProgressBar struct {
	widgets.QProgressBar

	_ func() `constructor:"init"`

	_ func(int) `signal:"valueChangedFromGoroutine"`
}

func (p *CustomProgressBar) init() {
	p.ConnectValueChangedFromGoroutine(p.SetValue)
}

type CustomLabel struct {
	widgets.QLabel

	_ func() `constructor:"init"`

	_ func(string) `signal:"updateTextFromGoroutine"`
	_ func() `signal:"adjustSizeFromGoroutine"`
}

func (c *CustomLabel) init() {
	c.ConnectUpdateTextFromGoroutine(c.SetText)
	c.ConnectAdjustSizeFromGoroutine(c.AdjustSize)
}

type CustomButton struct {
	widgets.QPushButton

	_ func() `constructor:"init"`

	_ func(bool) `signal:"setEnabledFromGoroutine"`
}

func (b *CustomButton) init() {
	b.ConnectSetEnabledFromGoroutine(b.SetEnabled)
}

type CustomTreeWidget struct {
	widgets.QTreeWidget

	_ func() `constructor:"init"`

	_ func( widgets.QTreeWidgetItem_ITF) `signal:"addTopLevelItemFromGoroutine"`
}

func (b *CustomTreeWidget) init() {
	b.ConnectAddTopLevelItemFromGoroutine(b.AddTopLevelItem)
}

//type CustomMessageBox struct {
//	widgets.QMessageBox
//	_ func() `constructor:"init"`
//
//	_ func() `signal:"showFromGoroutine"`
//
//}
//
//func (m CustomMessageBox) init(){
//	m.ConnectShowFromGorutine(m.Show)
//}



