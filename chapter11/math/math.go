package math

// Найти среднее в массиве чисел.
func Average(xs []float64) float64 {
    total := float64(0)
    for _, x := range xs {
        total += x
    }
    return total / float64(len(xs))
}

// Найти максимальное число в массиве.
func Max(xs []float64) float64  {
  max := xs[0]

  for _, value := range xs {
    if value > max {
      max = value
    }
  }

  return max
}

// Найти минимальное число в массиве.
func Min(xs []float64) float64  {
  min := xs[0]

  for _, value := range xs {
    if value < min {
      min = value
    }
  }

  return min
}
