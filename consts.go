package main 

import "fmt" 

/* Константы хранят некоторые данные, но их нельзя изменять, тк устанавливаются они один раз. Их вычесление производится один раз. */ 

const pi float64 = 3.1415 /* Если мы попробуем изменить pi, то получим ошибку. Также можно определить несколько констант сразу (по аналогии с var). 

Если у константы не указан тип, то он выводится неявно на основании значения: */ const n = 5 //Тип int

/* Необходимо обязательно инициализировать константу начальным значением при объявлении. Следующие определения констант являются недопустимыми, тк нет инициализации: */ 

const d 
const n int //НЕЛЬЗЯ! 

/* Если определяется последовательность констант, то инициализацию значением можно опустить для всех констант, кроме первой. В этом случае константа без значения полчит значение предыдущей константы: */ 

const (
    a = 1
    b
    c
    d = 3
    f
)
fmt.Println(a, b, c, d, f)  // 1, 1, 1, 3, 3

/* Константы можно инициализировать только константными значениями, например, литералами типа чисел или строк, или значениями других констант. Но инициализировать константу значением переменной мы не можем: */ 

var m int = 7
 const k = m  // ! Ошибка: m - переменная
const s = 5   // Норм: 5 - числовая константа
const n = s   // Норм: s - константа