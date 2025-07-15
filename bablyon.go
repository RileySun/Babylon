package babylon

import(
	"time"
	"strings"
	"math/rand"
	"io/ioutil"
)

type Babylon struct {
	Separator string //Seperator for words, default is a space
	dict []string //Word Storage
	path string
}

//Arg is default number of words, can be changed anytime
func NewBabylon() *Babylon {
	babylon := &Babylon{
		Separator:" ", 
	}
	
	err := babylon.LoadDict("/usr/share/dict/words")
	if err != nil {
		panic(err)
	}
	
	return babylon
}

func (b *Babylon) LoadDict(path string) error {
	//So much faster than any other option, dont bother with os.Open or Buffio
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	
	//Split
	b.dict = strings.Split(string(f[:]), "\n")
	
	return nil
}

func (b *Babylon) Babble(number int) string {
	return strings.Join(b.BabbleSlice(number), b.Separator)
}

func (b *Babylon) BabbleSlice(number int) []string {
	output := make([]string, number)
	
	for i := 0; i < number ; i++ {
		output[i] = b.dict[rand.Intn(len(b.dict) - 1)]
	}
	
	return output
}

/* Set Random Seed */
func init() {
	rand.Seed(time.Now().UnixNano())
}