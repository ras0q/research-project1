package exactmatch

import "github.com/ras0q/research-project1/retriever"

type exactmatchRetriever struct {
	dict retriever.Dictionary
}

func NewExactMatchRetriever(dict retriever.Dictionary) retriever.Retriever {
	return &exactmatchRetriever{dict}
}

func (r *exactmatchRetriever) Retrieve(req string) string {
	if res, ok := r.dict[req]; ok {
		return res
	}

	return "I don't know."
}
