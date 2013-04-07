package matrix

//import "fmt"

func (C *MatD) Dgemm(alpha, beta float64, A, B MatDable) *MatD {
	m, k := A.Size()
	_, n := B.Size()
	strA := A.Stride()
	strB := B.Stride()
	strC := C.stride

	Ops.Dgemm(A.IsTr(), B.IsTr(), m, n, k, alpha, A.ArrayD(),
		strA, B.ArrayD(), strB, beta, C.data, strC)
	return C
}

func MulD(M ...MatDable) *MatD {
	numM := len(M)
	if numM == 0 {
		return nil
	}
	if numM == 1 {
		m, n := M[0].Size()
		arr := M[0].ArrayD()
		return FromArrayD(arr, false, m, n)
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
	trLeft := false
	var arrLeft []float64
	strLeft := 0
	trRight := false

	if numM-ixMinDim > 0 {
		j := ixMinDim

		strideCurr := stride[j]
		var input []float64
		output := make([]float64, 9)
		arr := M[j].ArrayD()
		input = make([]float64, len(arr))
		copy(input, arr)
		trRight = M[j].IsTr()

		j += 1

		for j < numM {
			if len(output) < m[ixMinDim]*n[j] {
				output = make([]float64, m[0]*n[j])
			}
			Ops.Dgemm(trRight, M[j].IsTr(), m[ixMinDim], n[j], n[j-1], 1, input,
				strideCurr, M[j].ArrayD(), stride[j], 0, output, m[0])

			trRight = false
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
		arr := M[j].ArrayD()
		input = make([]float64, len(arr))
		copy(input, arr)
		trLeft = M[j].IsTr()

		j--

		for j >= 0 {
			if len(output) < m[j]*n[ixMinDim-1] {
				output = make([]float64, m[j]*n[ixMinDim-1])
			}
			Ops.Dgemm(M[j].IsTr(), trLeft, m[j], n[ixMinDim-1], m[j+1], 1, M[j].ArrayD(),
				stride[j], input, strideCurr, 0, output, m[j])

			trLeft = false
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
		return FromArrayD(arrRight, false, m[0], n[numM-1])
	}
	if arrRight == nil {
		return FromArrayD(arrLeft, false, m[0], n[numM-1])
	}

	output := make([]float64, m[0]*n[numM-1])
	Ops.Dgemm(trLeft, trRight, m[0], n[numM-1], m[ixMinDim], 1, arrLeft,
		strLeft, arrRight, strRight, 0, output, m[0])

	return FromArrayD(output, false, m[0], n[numM-1])
}
