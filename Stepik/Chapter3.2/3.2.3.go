/*
Для решения данной задачи вам понадобится пакет strconv, возможно использовать пакеты strings или encoding/csv, или даже bufio -
вы не ограничены в выборе способа решения задачи. В решениях мы поделимся своими способами решения этой задачи, предлагаем вам сделать то же самое.

В привычных нам редакторах электронных таблиц присутствует удобное представление числа с разделителем разрядов в виде пробела,
кроме того в России целая часть от дробной отделяется запятой. Набор таких чисел был экспортирован в формат CSV,
где в качестве разделителя используется символ ";".

На стандартный ввод вы получаете 2 таких вещественных числа, в качестве результата требуется вывести частное от деления первого числа на второе
с точностью до четырех знаков после "запятой" (на самом деле после точки, результат не требуется приводить к исходному формату).

P.S. небольшое отступление, связанное с чтением из стандартного ввода.
Кто-то захочет использовать для этого пакет bufio.Reader. Это вполне допустимый вариант,
но если вы резонно обрабатываете ошибку метода ReadString('\n'), то получаете ошибку EOF, а точнее (io.EOF - End Of File).
На самом деле это не ошибка, а состояние, означающее, что файл (а os.Stdin является файлом) прочитан до конца.
Чтобы ошибка была обработана правильно, вы можете поступить так:

if err != nil && err != io.EOF {
	...
}

Sample Input:
1 149,6088607594936;1 179,0666666666666

Sample Output:
0.9750
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	text, _, _ := bufio.NewReader(os.Stdin).ReadLine()

	str := strings.ReplaceAll(string(text), " ", "")
	str = strings.ReplaceAll(str, ",", ".")
	strArr := strings.Split(str, ";")

	num1, _ := strconv.ParseFloat(strArr[0], 64)
	num2, _ := strconv.ParseFloat(strArr[1], 64)
	fmt.Printf("%.4f", num1/num2)
}