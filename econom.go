/*
Важная информация:
Коды цифр и символов
0-48
1-49
2-50
3-51
4-52
5-53
6-54
7-55
8-56
9-57
( - 40
) - 41
* - 42
+ - 43
"-" - 45
@-64
#-35
$-36
*/
package main

import "fmt"

var MyArray []string

//поиск знаков @(собакен) #(тюрьма) $(шашлык)
func action(a byte) bool {
	if a == 35 {
		return true
	}
	if a == 36 {
		return true
	}
	if a == 64 {
		return true
	}
	return false
}

//поиск букв
func letter(a byte) bool {
	if (a >= 97) && (a <= 122) {
		return true
	}
	return false
}

//поиск следующего выражения по скобкам
func next(Stroka string) string {
	var t, k, l int
	t = 0
	k = 0
	l = len(Stroka)
	for i := 0; i < l; i++ {
		if letter(Stroka[i]) {
			t++
		} else if action(Stroka[i]) {
			t = t - 2
		} else if Stroka[i] == 40 {
			k++
		} else if Stroka[i] == 41 {
			k--
			t++
		}
		if (k == 0) && (t == 1) {
			return Stroka[i+1:]
		}
	}
	return ""
}

func parcer(Stroka string) string {
	var a byte
	var b string
	a = Stroka[0]
	if a == 40 {
		return parcer(Stroka[1:])
	} else if letter(a) {
		return Stroka[:1]
	} else if action(a) {
		b = Stroka[:1] + parcer(Stroka[1:]) + parcer(next(Stroka[1:]))
		l := len(MyArray)
		t := true
		for i := 26; i < l; i++ {
			if t && (MyArray[i] == b) {
				t = false
			}
		}
		if t {
			MyArray = append(MyArray, b)
		}
		return (b)
	}
	return ""
}

func main() {
	var f, n int //длинна массива
	var Stroka string
	MyArray = append(MyArray, "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z")
	/*
		//определение кодов символов
		var n string
		var num byte
		fmt.Scan(&n)
		num = n[0]
		fmt.Print(num)
	*/
	f = len(MyArray)
	fmt.Scanln(&Stroka)
	Stroka = parcer(Stroka)
	n = len(MyArray)
	/*
		for i:=0;i<26;i++{
			fmt.Println(MyArray)
		}
	*/
	fmt.Println(n - f)
	/*
	   x	0
	   ($xy)	1
	   ($(@ab)c)	2
	   (#i($jk))	2
	   (#($ab)($ab))	2
	   (@(#ab)($ab))	3
	   (#($a($b($cd)))(@($b($cd))($a($b($cd)))))	5
	   (#($(#xy)($(#ab)(#ab)))(@z($(#ab)(#ab))))	6
	*/
}
