## тут весь код за исключением того, который генерит qtmoc в файле AppGui/mainWindow/mainThreadUpdater.go
```
type updateHelper struct {
	core.QObject

	_ func(f func()) `signal:"runUpdate,auto`
}


func (*updateHelper) runUpdate(f func()) { f() }

var UpdateHelper = NewUpdateHelper(nil)
```
автогенерящиеся функции эти
```
UpdateHelper.RunUpdate(func())
var UpdateHelper = NewUpdateHelper(nil)

```
## Список вопросов в файле questions.md


## Список багов в файле bugs.md

