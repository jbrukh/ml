package ml

import "testing"
import "fmt"

const ERR = 1e-5

func TestSLR_slope1(t *testing.T) {
    x := []float64{0,1,2,3}
    y := x
    SLR(t, x, y, 0, 1)
}

func TestSLR_slope0(t *testing.T) {
    x := []float64{0,1,2,3}
    y := []float64{3,3,3,3}
    SLR(t, x, y, 3, 0)
}

func TestSLR(t *testing.T) {
    x := []float64{0,1.5,2.5,3.1}
    y := []float64{3.1,3.0,2.99,2.8}
    SLR(t, x, y, 3.11745, -0.08166)
}


func BenchmarkSLR_10k(b *testing.B) {
    benchmarkSLR(10000, b)
}
func BenchmarkSLR_20k(b *testing.B) {
    benchmarkSLR(20000, b)
}
func BenchmarkSLR_30k(b *testing.B) {
    benchmarkSLR(30000, b)
}
func BenchmarkSLR_40k(b *testing.B) {
    benchmarkSLR(40000, b)
}
func BenchmarkSLR_50k(b *testing.B) {
    benchmarkSLR(50000, b)
}
func BenchmarkSLR_100k(b *testing.B) {
    benchmarkSLR(100000, b)
}
func BenchmarkSLR_200k(b *testing.B) {
    benchmarkSLR(200000, b)
}
func BenchmarkSLR_500k(b *testing.B) {
    benchmarkSLR(500000, b)
}
func benchmarkSLR(iter int, b *testing.B) {
	b.StopTimer()
	x := make([]float64, iter)
	y := make([]float64, iter)
	for i := 0; i < iter; i++ {
		x[i] = float64(i)
		y[i] = float64(2*i)
	}
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        SimpleLinearRegression(x, y)
    }
}

func SLR(t *testing.T, x, y []float64, theta0, theta1 float64) {
	t0, t1 := SimpleLinearRegression(x, y)
	fmt.Printf("theta0 = %v, theta1 = %v\n", theta0, theta1)
	if !within(t0, theta0, ERR) || !within(t1, theta1, ERR) {
    	t.Fail()
	}
}

func within(x, y, err float64) bool {
	if x > y {
		return x - y <= err
	} else {
		return y - x <= err
	}
	return false
}

