package repo

var (
	Word = make(map[string]int)
)

type GetWordParam struct {
	Limit  int32
	Page   int32
	Search string
}

type GetWord struct {
	Word  string
	Point int32
}

type GetWordResult struct {
	Words []*GetWord
	Count int32
}

type WordStorageI interface {
	Create(body map[string]int) error
	GetAll(param *GetWordParam) (*GetWordResult, error)
}
