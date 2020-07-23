package main
import(
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int{
	m := make(map[string]int)
	words := strings.Fields(s)
	for i,_:= range words{
		word := words[i]
		count,there := m[word]
		
		if there{
			count = m[word]
			m[word] = count +1
		}else{
			m[word] = 1
		}
	}
	return m
}

func main(){
	wc.Test(WordCount)
}
