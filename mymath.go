package mymath

import "math"

// Sqrt вычисляет квадратный корень из числа
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Ceil возвращает наименьшее целое число, большее или равное x
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Floor возвращает наибольшее целое число, меньшее или равное x
func Floor(x float64) float64 {
	return math.Floor(x)
}

// Pow возвращает x в степени y
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Max возвращает большее из двух чисел
func Max(x, y float64) float64 {
	return math.Max(x, y)
}

// Min возвращает меньшее из двух чисел
func Min(x, y float64) float64 {
	return math.Min(x, y)
}

// Abs возвращает абсолютное значение числа
func Abs(x float64) float64 {
	return math.Abs(x)
}

// Sin вычисляет синус угла (в радианах)
func Sin(x float64) float64 {
	return math.Sin(x)
}

// Cos вычисляет косинус угла (в радианах)
func Cos(x float64) float64 {
	return math.Cos(x)
}

// Tan вычисляет тангенс угла (в радианах)
func Tan(x float64) float64 {
	return math.Tan(x)
}
