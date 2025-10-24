package lessons

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAdd(t *testing.T) {
	t.Parallel() // параллельный запуск
	x1, x2 := 5, 5
	expected := 10

	result := Add(x1, x2)
	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestQuadro(t *testing.T) {
	type want struct {
		roots      []float64
		rootsCount int
		err        error
	}

	type test struct {
		name string
		a    float64
		b    float64
		c    float64
		want want
	}

	tests := []test{
		{
			name: "two roots",
			a:    1,
			b:    -3,
			c:    2,
			want: want{
				roots:      []float64{1, 2},
				rootsCount: 2,
				err:        nil,
			},
		},
		{
			name: "not Quadro",
			a:    0,
			b:    5,
			c:    7,
			want: want{
				roots:      nil,
				rootsCount: 0,
				err:        ErrNotQuadro,
			},
		},
		{
			name: "not realRoots",
			a:    2,
			b:    1,
			c:    1,
			want: want{
				roots:      nil,
				rootsCount: 0,
				err:        ErrNotRealRoots,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			roots, err := Quadro(tc.a, tc.b, tc.c)
			//if err != tc.want.err {
			//	t.Error("Expected", tc.want.err, "Got", err, "\n", err)
			//}
			//if len(roots) != tc.want.rootsCount {
			//	t.Error("Expected", tc.want.rootsCount, "Got", len(roots))
			//}
			//for i := range roots {
			//	if roots[i] != tc.want.roots[i] {
			//		t.Error("Expected", tc.want.roots[i], "Got", roots[i])
			//	}
			//}
			// или через ассерты
			assert.Equal(t, tc.want.roots, roots)
			assert.Equal(t, tc.want.err, err)
			assert.Equal(t, tc.want.rootsCount, len(roots))
		})
	}
}

func BenchmarkQuadroSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Quadro(1, -3, 2)
		if err != nil {
			return
		}
	}
}

func BenchmarkQuadroBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Quadro(1e6, -3e6, 2e6)
		if err != nil {
			return
		}
	}
}
