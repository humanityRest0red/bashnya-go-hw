package main

import "fmt"

const MAX = 12307

func main() {
	var num, res int

	fmt.Println("Введите число:")
	_, err := fmt.Scan(&num)
	if err == nil {
		res, err = Process(num)
	}

	if err == nil {
		fmt.Printf("Результат: %v - %v\n", res, NumIntoWords(res))
	} else {
		fmt.Print(err)
	}
}

func Process(num int) (int, error) {
	for num < MAX {
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
	milliards := num / 1_000_000_000
	remainder := num % 1_000_000_000
	millions := remainder / 1_000_000
	remainder = remainder % 1_000_000
	thousands := remainder / 1_000

	var result string

	tail := HundredsIntoWords(milliards, 'm')
	tail = ConcatWord(tail, "миллиард", 'm', milliards)
	result = ConcatTail(result, tail)

	tail = HundredsIntoWords(millions, 'm')
	tail = ConcatWord(tail, "миллион", 'm', millions)
	result = ConcatTail(result, tail)

	tail = HundredsIntoWords(thousands, 'f')
	tail = ConcatWord(tail, "тысяч", 'f', thousands)
	result = ConcatTail(result, tail)

	tail = HundredsIntoWords(num, 'm')
	result = ConcatTail(result, tail)

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

func ConcatWord(s, word string, gender rune, num int) string {
	if num > 0 {
		s += " " + word
		if 10 <= num && num <= 20 {
			if gender == 'm' {
				s += "ов"
			}
		} else {
			switch num % 10 {
			case 1:
				if gender == 'f' {
					s += "а"
				}
			case 2, 3, 4:
				if gender == 'f' {
					s += "и"
				} else {
					s += "а"
				}
			default:
				if gender == 'm' {
					s += "ов"
				}
			}
		}
	}
	return s
}

func AddSpaceIfNeed(s string) string {
	if s != "" {
		return s + " "
	}
	return s
}

func ConcatTail(s, tail string) string {
	if tail != "" {
		if s != "" {
			s += " "
		}
		s += tail
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
