package main

import "fmt"

const MAX = 12307

func main() {
	var num, res int

	_, err := fmt.Scan(&num)
	if err == nil {
		res, err = Process(num)
	}

	if err == nil {
		fmt.Printf("Число = %v\n", res)
		fmt.Printf("В текстовом виде:\n%v\n", NumIntoWords(res))
	} else {
		fmt.Print(err)
	}
}

func Process(num int) (int, error) {
	for num < MAX {
		fmt.Println(num)
		if num < 0 {
			num = -num
		}
		if num%7 == 0 {
			num *= 39
		}
		if num%9 == 0 {
			num = num*13 + 1
			continue
		} else {
			num = (num + 2) * 3
		}

		if num%13 == 0 && num%9 == 0 {
			return num, fmt.Errorf("service error")
		} else {
			num++
		}
	}

	return num, nil
}

func NumIntoWords(num int) string {
	thousands := num / 1000
	var result string = HundredsIntoWords(thousands, 'f')
	switch thousands % 10 {
	case 1:
		result += " тысяча"
	case 2, 3, 4:
		result += " тысячи"
	default:
		result += " тысяч"
	}

	tail := HundredsIntoWords(num, 'm')
	if tail != "" {
		result += " " + tail
	}

	return CapitalizeFirstLetter(result)
}

func HundredsIntoWords(num int, gender rune) string {
	word_hundreds := []string{"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}
	word_tens := []string{"", "десять", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
	word_teens := []string{"", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
	word_ones := []string{"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"}
	word_ones_female := []string{"", "одна", "две"}

	hundreds := num % 1000 / 100
	tens := num % 100 / 10
	teens := num % 100
	ones := num % 10

	var result string = word_hundreds[hundreds]

	if 10 < teens && teens < 20 {
		result = AddSpaceIfNeed(result)
		result += word_teens[ones]
	} else {
		if tens > 0 {
			result = AddSpaceIfNeed(result)
			result += word_tens[tens]
		}
		if ones > 0 {
			result = AddSpaceIfNeed(result)
			if gender == 'f' && ones <= 2 {
				result += word_ones_female[ones]
			} else {
				result += word_ones[ones]
			}
		}
	}

	return result
}

func AddSpaceIfNeed(s string) string {
	if s != "" {
		return s + " "
	}
	return s
}

func CapitalizeFirstLetter(s string) string {
	runes := []rune(s)
	first_letter := runes[0]
	if 'а' < first_letter && first_letter < 'я' {
		runes[0] += 'А' - 'а'
	}

	return string(runes)
}
