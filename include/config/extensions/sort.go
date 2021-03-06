package extensions


func (e EXTensions) Len() int           { return len(e) }
func (e EXTensions) Less(i, j int) bool { return e[i].Ext < e[j].Ext }
func (e EXTensions) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }