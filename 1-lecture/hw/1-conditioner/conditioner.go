package main

import (
	"fmt"
	"os"
	"strconv"
)

// В офисе, где работает программист Петр, установили кондиционер нового типа. Этот кондиционер отличается особой простотой в управлении. У кондиционера есть всего лишь два управляемых параметра: желаемая температура и режим работы.
// Кондиционер может работать в следующих четырех режимах:
// «freeze» — охлаждение. В этом режиме кондиционер может только уменьшать температуру. Если температура в комнате и так не больше желаемой, то он выключается.
// «heat» — нагрев. В этом режиме кондиционер может только увеличивать температуру. Если температура в комнате и так не меньше желаемой, то он выключается.
// «auto» — автоматический режим. В этом режиме кондиционер может как увеличивать, так и уменьшать температуру в комнате до желаемой.
// «fan» — вентиляция. В этом режиме кондиционер осуществляет только вентиляцию воздуха и не изменяет температуру в комнате.
// Кондиционер достаточно мощный, поэтому при настройке на правильный режим работы он за час доводит температуру в комнате до желаемой.
// Требуется написать программу, которая по заданной температуре в комнате troom, установленным на кондиционере желаемой температуре tcond и режиму работы определяет температуру, которая установится в комнате через час.

// Первая строка входного файла содержит два целых числа troom, и tcond, разделенных ровно одним пробелом (–50 ≤ troom ≤ 50, –50 ≤ tcond ≤ 50).
// Вторая строка содержит одно слово, записанное строчными буквами латинского алфавита — режим работы кондиционера.

// Выходной файл должен содержать одно целое число — температуру, которая установится в комнате через час.

func main() {
	var tt string 
	fmt.Fscan(os.Stdin, tt)

	var mode string
	fmt.Fscan(os.Stdin, mode)

	res := SetTemperature(tt, mode)
	fmt.Fprint(os.Stdout, res)
}

func SetTemperature(tt, mode string) int {
	var troom int 
	var tcond int

	var spaceIdx int = -1
	for i, ch := range tt {
		if ch == ' ' {
			spaceIdx = i
		}
	}

	if spaceIdx < 0 {
		return 0
	}

	troom, err := strconv.Atoi(tt[:spaceIdx])
	if err != nil {
		return 0
	}
	tcond, err = strconv.Atoi(tt[spaceIdx + 1:])
	if err != nil {
		return 0
	}

	if troom < -50 || troom > 50 || tcond < -50 || tcond > 50 {
		return 0
	}
	
	switch mode {
	case "freeze":
		if tcond > troom {
			return troom
		}
		return tcond
	case "heat":
		if tcond < troom {
			return troom
		}
		return tcond
	case "auto":
		return tcond
	case "fan":
		return troom
	default:
		return 0
	}
}