package dijkstra

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

func stringToEdgeWeightMatrix(s string) ([][]int, error) {
	subs := strings.Split(s, "\n")
	n, _ := strconv.Atoi(subs[0])
	fmt.Println(n)
	edges := make([][]int, n)
	for i := range edges {
		edges[i] = make([]int, n)
		for j := range edges[i] {
			edges[i][j] = math.MinInt
		}
	}

	for i, s := range subs {
		if i == 0 {
			continue
		}
		fromto := strings.Split(s, " ")
		//fmt.Println(fromto)
		if len(fromto) != 3 {
			return nil, errors.New("bad formatting")
		}
		from, err1 := strconv.Atoi(fromto[0])
		to, err2 := strconv.Atoi(fromto[1])
		w, err3 := strconv.Atoi(fromto[2])
		if err1 != nil || err2 != nil || err3 != nil {
			return nil, errors.New("bad formatting")
		}
		if from < 0 || from >= n || to < 0 || to >= n || w < 0 {
			return nil, errors.New("invalid range of from to weight")
		}
		edges[from][to] = w
	}
	return edges, nil
}

func TestEdgeMatrixFromString(t *testing.T) {
	s := `5
0 1 1
1 2 7`
	edges, err := stringToEdgeWeightMatrix(s)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(edges)

	if len(edges) != 5 {
		t.Fatalf("not size 5")
	}
	for i := range edges {
		if len(edges[i]) != 5 {
			t.Fatalf("not size 5")
		}
		for j := range edges[i] {
			if i == 0 && j == 1 && edges[i][j] != 1 {
				t.Fatalf("incorrect weight for 01")
			} else {
				continue
			}
			if i == 1 && j == 2 && edges[i][j] != 7 {
				t.Fatalf("incorrect weight for 12")
			} else {
				continue
			}

			if edges[i][j] != math.MinInt {
				t.Fatalf("not math.MinInt")
			}
		}
	}
}

func TestDijkstra(t *testing.T) {
	s := `5
0 1 1
1 2 1
2 3 1
3 4 2`
	e, err := stringToEdgeWeightMatrix(s)
	if err != nil {
		t.Fatalf(err.Error())
	}
	o := dijkstra(5, e, 0, 4)
	if o != 5 {
		t.Fatalf("bad result")
	}
}
