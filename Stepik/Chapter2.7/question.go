package main

import "fmt"

func main() {
	v := 5
	//в переменной лежит значение 5
	fmt.Println("адрес V", &v)
	p := &v
	//в переменную р положили указатель на переменную v, &v - значит взять адрес v и положить этот адрес в p
	//р - становится типом указатель, * и лежит там теперь адрес типа - 0х000002 (например)
	//fmt.Print(*p, " ")
	//печатаем *р - то есть пойдем по адресу который там лежит и напечатаем то что есть по этому адресу
	//попробуйте в этом месте напечатать просто р, типа так
	fmt.Println("адрес который лежит в Р", p)
	//он напечатает именно адрес, который лежит в р и он будет такой же, который мы печатали раннее со строкой "адрес V"
	changePointer(p)
	//fmt.Print(*p)
	fmt.Println("адрес который лежит в Р после функции", p)
	// поэтому тут указатель все на ту же v которая 5
	// !!!!!!ВАЖНО!!!!!!  Если вы хотите повыводить что в какой переменной лежит. Если это указатель как Р, то нужно выводить просто Р
	// А если вы хотите вывести именно адрес переменной то &p. Потому что Р это указатель конкретно тут в main и печатая Р вы напечатаете
	// ТО ЧТО ЛЕЖИТ ВНУТРИ Р(адрес V), а если вы напечатаете &P, то он выведет адрес САМОЙ ЭТОЙ Р.
}

func changePointer(p *int) {
	v := 3
	// мы в другой функции, инициализируем ДРУГУЮ v под которую выделен ДРУГОЙ адрес. Потому что мы в ДРУГОЙ функции
	// можно на каждую переменную что в main, что тут, написать
	fmt.Println("адрес переменной V в changePointer", &v)
	// и посмотреть что адрес v в main и тут разные
	p = &v
	fmt.Println("адрес переменной Р в changePointer", &p)
	// тут тоже самое мы в ДРУГОЙ функции инициализировали р, точно так же можно вывести принтом адрес той р в main и этой р и они будут разные
	// и в этом случае мы переменной р в ЭТОЙ функции присвоили адрес v которая тоже в ЭТОЙ функции
	// если написать *р = v, мы пойдем по адресу, который в р, придем в переменную v = 5 и ТУДА положим 3, значение которое у нас в v этой функции
	// то есть мы изменим изначальную v
	// и второй раз fmt.Print(*p) после функции выведет 3, все по тому же адресу.
}