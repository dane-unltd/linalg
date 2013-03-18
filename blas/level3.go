package blas

import "github.com/dane-unltd/linalg/dense"

//import "fmt"

func MulD(M ...MatD) *dense.MatD {
	numM := len(M)
	if numM == 0 {
		return nil
	}
	if numM == 1 {
		m, n := M[0].Size()
		arr, use := M[0].ArrayD()
		return dense.FromArrayD(arr, use, m, n)
	}
	m := make([]int, numM)
	n := make([]int, numM)
	stride := make([]int, numM)

	minDim, _ := M[0].Size()
	ixMinDim := 0
	for i, v := range M {
		m[i], n[i] = v.Size()
		if m[i] < minDim {
			minDim = m[i]
			ixMinDim = i
		}
		stride[i] = v.Stride()
	}
	if n[len(M)-1] < minDim {
		ixMinDim = len(M)
	}
	for i := 0; i < numM-1; i++ {
		if n[i] != m[i+1] {
			panic("dimension mismatch")
		}
	}

	var arrRight []float64
	strRight := 0
	var arrLeft []float64
	strLeft := 0

	if numM-ixMinDim > 0 {
		j := ixMinDim

		strideCurr := stride[j]
		var input []float64
		output := make([]float64, 9)
		arr, use := M[j].ArrayD()
		if use {
			input = arr
		} else {
			input = make([]float64, len(arr))
			copy(input, arr)
		}

		j += 1

		for j < numM {
			if len(output) < m[ixMinDim]*n[j] {
				output = make([]float64, m[0]*n[j])
			}
			arrM, _ := M[j].ArrayD()
			dgemm("N", "N", m[ixMinDim], n[j], n[j-1], 1, input,
				strideCurr, arrM, stride[j], 0, output, m[0])

			temp := input
			input = output
			output = temp
			strideCurr = m[ixMinDim]
			j++
		}
		arrRight = input
		strRight = strideCurr
	}
	if ixMinDim > 0 {
		j := ixMinDim - 1

		strideCurr := stride[j]
		var input []float64
		output := make([]float64, 9)
		arr, use := M[j].ArrayD()
		if use {
			input = arr
		} else {
			input = make([]float64, len(arr))
			copy(input, arr)
		}

		j--

		for j >= 0 {
			if len(output) < m[j]*n[ixMinDim-1] {
				output = make([]float64, m[j]*n[ixMinDim-1])
			}
			arrM, _ := M[j].ArrayD()
			dgemm("N", "N", m[j], n[ixMinDim-1], m[j+1], 1, arrM,
				stride[j], input, strideCurr, 0, output, m[j])

			temp := input
			input = output
			output = temp
			strideCurr = m[j]
			j--
		}
		arrLeft = input
		strLeft = strideCurr
	}

	if arrLeft == nil {
		return dense.FromArrayD(arrRight, false, m[0], n[numM-1])
	}
	if arrRight == nil {
		return dense.FromArrayD(arrLeft, false, m[0], n[numM-1])
	}

	output := make([]float64, m[0]*n[numM-1])
	dgemm("N", "N", m[0], n[numM-1], m[ixMinDim], 1, arrLeft,
		strLeft, arrRight, strRight, 0, output, m[0])

	return dense.FromArrayD(output, true, m[0], n[numM-1])
}
