package statistic

// Average функция рассчёта среднего
func Average(xs []float64) float64 {
    total := float64(0)
    for _, x := range xs {
        total += x
    }
    return total / float64(len(xs))
}

// Summa функция рассчёта суммы
func Summa(s []float64) float64 {
    total := float64(0)
    for _, x := range s {
        total += x
    }
    return total
}