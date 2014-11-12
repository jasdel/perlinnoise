package perlinnoise

// Interface for collections which can be shuffled.
type Shufflable interface {
	Len() int
	Swap(i, j int)
}

// Method for generating random values, i.g math/rand rand.Intn
type RandEng interface {
	Intn(i int) int
}

// Shuffles the collection.
// Based on the Fisher-Yates modern shuffle method
// http://en.wikipedia.org/wiki/Fisher-Yates_shuffle#The_modern_algorithm
func Shuffle(s Shufflable, r RandEng) {
	for end := s.Len() - 1; end > 0; end-- {
		i := r.Intn(end + 1)
		s.Swap(i, end)
	}
}
