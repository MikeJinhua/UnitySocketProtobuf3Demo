package Test

type TableAAA struct {
	IndexInt int "index"
}

type TableAAAManager struct {
	data map[string]TableAAA
}

var (
	Instance = new(TableAAAManager)
)



