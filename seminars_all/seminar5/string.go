package main
import(
	"fmt"
	"strings"
)

func main() {
	str1 := "ананас"
	str2 := "банан"
	fmt.Println(str1 < str2) //сравнение строк как в С++ по алфавиту(но по факту по Unicode-таблице)
	fmt.Println(len(str1)) //так узнаю кол-во байт, а не длину строки
	fmt.Println(len(str2))
	fmt.Println(len([]rune(str1))) //так узнаю длину

	//конкатенация строк(можно и по-обычному, но сказали это долго и может вызывать проблемы)
	var str12 strings.Builder
	str12.WriteString(str1)
	str12.WriteString(" это вам не ")
	str12.WriteString(str2)
	fmt.Println(str12.String())
	
	//так менять содержимое строки
	str2rune := []rune(str2)
	str2rune[0], str2rune[4] = 'к', 'л'
	str2 = string(str2rune)
	fmt.Println(str2)

}