package babylon

import(
	"time"
	"strings"
	"math/rand"
	"io/ioutil"
)

type Babylon struct {
	Number int //Default Number of words
	Separator string //Seperator for words, default is a space
	dict []string //Word Storage
	path string
}

//Arg is default number of words, can be changed anytime
func NewBabylon(number int) *Babylon {
	babylon := &Babylon{
		Number:number,
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

func (b *Babylon) Babble() string {
	output := make([]string, b.Number)
	
	for i := 0; i < b.Number ; i++ {
		output[i] = b.dict[rand.Intn(len(b.dict) - 1)]
	}
	
	return strings.Join(output, b.Separator)
}

/* Set Random Seed */
func init() {
	rand.Seed(time.Now().UnixNano())
}