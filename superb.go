package superb

import "math/rand"

// Words is a collection of awesome words
var Words = []string{
	"ace",
	"amazing",
	"astonishing",
	"astounding",
	"awe-inspiring",
	"awesome",
	"badass",
	"beautiful",
	"bedazzling",
	"bee's knees",
	"best",
	"breathtaking",
	"brilliant",
	"cat's meow",
	"cat's pajamas",
	"classy",
	"cool",
	"dandy",
	"dazzling",
	"delightful",
	"divine",
	"doozie",
	"epic",
	"excellent",
	"exceptional",
	"exquisite",
	"extraordinary",
	"fabulous",
	"fantastic",
	"fantabulous",
	"fine",
	"finest",
	"first-class",
	"first-rate",
	"flawless",
	"funkadelic",
	"geometric",
	"glorious",
	"gnarly",
	"good",
	"grand",
	"great",
	"groovy",
	"groundbreaking",
	"hunky-dory",
	"impeccable",
	"impressive",
	"incredible",
	"kickass",
	"kryptonian",
	"laudable",
	"legendary",
	"lovely",
	"luminous",
	"magnificent",
	"majestic",
	"marvelous",
	"mathematical",
	"mind-blowing",
	"neat",
	"outstanding",
	"peachy",
	"perfect",
	"phenomenal",
	"pioneering",
	"polished",
	"posh",
	"praiseworthy",
	"premium",
	"priceless",
	"prime",
	"primo",
	"rad",
	"remarkable",
	"riveting",
	"sensational",
	"shining",
	"slick",
	"smashing",
	"solid",
	"spectacular",
	"splendid",
	"stellar",
	"striking",
	"stunning",
	"stupendous",
	"stylish",
	"sublime",
	"super",
	"super-duper",
	"super-excellent",
	"superb",
	"superior",
	"supreme",
	"swell",
	"terrific",
	"tiptop",
	"top-notch",
	"transcendent",
	"tremendous",
	"ultimate",
	"unreal",
	"well-made",
	"wicked",
	"wonderful",
	"wondrous",
	"world-class",
}

var (
	wordsLen int32
	r        *rand.Rand
)

func init() {
	r = rand.New(rand.NewSource(123457))
	wordsLen = int32(len(Words))
}

// Next returns one superb word randomly
func Next() string {
	return Words[r.Int31n(wordsLen)]
}

// Is checks if a given words is superb
func Is(s string) bool {
	for i := range Words {
		if Words[i] == s {
			return true
		}
	}
	return false
}
