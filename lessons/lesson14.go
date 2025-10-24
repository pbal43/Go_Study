package lessons

import (
	"errors"
	"fmt"
	"log"
	"math"
	"sort"
)

var ErrNotQuadro = errors.New("not quadro")
var ErrNotRealRoots = errors.New("not real roots")

func humanBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func Add(a, b int) int {
	return a + b
}

func Quadro(a, b, c float64) ([]float64, error) {
	if a == 0 {
		return nil, ErrNotQuadro
	}
	d := b*b - 4*a*c
	if d < 0 {
		return nil, ErrNotRealRoots
	}
	if math.Abs(d) <= 1e-15 {
		x := -b / (2 * a)
		return []float64{x}, nil
	}
	sqrt := math.Sqrt(d)
	x1 := (-b + sqrt) / (2 * a)
	x2 := (-b - sqrt) / (2 * a)
	roots := []float64{x1, x2}
	sort.Float64s(roots)
	return roots, nil
}

func Lesson14() {
	//debug.SetGCPercent(100) // пытаемся управлять поведением сборщика мусора - при росте кучи на 100 процентов
	//const holdCount = 50
	//const allocSize = 1 << 20
	//ring := make([][]int, holdCount)
	//
	//fmt.Println("GO GC TEST")
	//
	//stop := time.After(15 * time.Second)
	//ticker := time.NewTicker(1 * time.Second)
	//
	//i := 0
	//for {
	//	select {
	//	case <-stop:
	//		fmt.Println("TEST STOP")
	//		return
	//	case <-ticker.C:
	//		var m runtime.MemStats // переменная, в которой хранится статистика по памяти
	//		runtime.ReadMemStats(&m)
	//		fmt.Printf("[%s] HeapAlloc = %s HeapSys = %s NumGC = %d PauseTotal = %dms NextGC = %s\n",
	//			time.Now().Format("2006-01-02 15:04:05"),
	//			humanBytes(m.HeapAlloc),
	//			humanBytes(m.HeapSys),
	//			m.NumGC,
	//			m.PauseTotalNs/1e6,
	//			humanBytes(m.NextGC))
	//	default:
	//		b := make([]int, allocSize)
	//		b[0] = 1
	//		ring[i%holdCount] = b
	//		i++
	//	}
	//}

	roots, err := Quadro(1, -2, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("roots: %v\n", roots)
}
