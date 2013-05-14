package lapack

import (
	"github.com/gonum/blas"
)

type Job byte
type Range byte
type Type byte
type Sort byte
type Info byte
type Select func(float64, float64) bool
type SelectG func(float64, float64, float64) bool

type Float64 interface {
	Dlacpy(ul blas.Uplo, m int, n int, a []float64, lda int, b []float64, ldb int)

	//Missing: bdsdc bdsqr disna gbbrd gbcon gbequ gbequb gbrfs gbrfsx gbsvx gbsvxx
	Dgbsv(n int, kl int, ku int, nrhs int, ab []float64, ldab int, ipiv []int32, b []float64, ldb int) Info
	Dgbtrf(m int, n int, kl int, ku int, ab []float64, ldab int, ipiv []int32) Info

	//Missing: gebak gebal gebrd gecon geequ geequb geesx geev geevx gehrd gejsv
	Dgbtrs(ta blas.Transpose, n int, kl int, ku int, nrhs int, ab []float64, ldab int, ipiv []int32, b []float64, ldb int) Info
	Dgees(jobvs Job, s Sort, sel Select, n int, a []float64, lda int, sdim int, wr []float64, wi []float64, vs []float64, ldvs int) Info

	//Missing: gelsd gelss gelsy geqlf geqpf geqrfp gerfs gerfsx gerqf
	Dgelqf(m int, n int, a []float64, lda int, tau []float64) Info
	Dgels(ta blas.Transpose, m int, n int, nrhs int, a []float64, lda int, b []float64, ldb int) Info
	Dgeqp3(m int, n int, a []float64, lda int, jpvt []int, tau []float64) Info
	Dgeqrf(m int, n int, a []float64, lda int, tau []float64) Info

	//Missing: gesvj gesvx gesvxx ggbak ggbal ggesx
	Dgesdd(jobz Job, m int, n int, a []float64, lda int, s []float64, u []float64, ldu int, vt []float64, ldvt int) Info
	Dgesv(n int, nrhs int, a []float64, lda int, ipiv []int32, b []float64, ldb int) Info
	Dgesvd(jobu Job, jobvt Job, m int, n int, a []float64, lda int, s []float64, U []float64, ldu int, vt []float64, ldvt int) Info
	Dgetrf(m int, n int, a []float64, lda int, ipiv []int32) Info
	Dgetri(n int, a []float64, lda int, ipiv []int32) Info
	Dgetrs(ta blas.Transpose, n int, nrhs int, a []float64, lda int, ipiv []int32, b []float64, ldb int) Info
	Dgges(jobvsl Job, jobvsr Job, s Sort, sel SelectG, n int, a []float64, lda int, b []float64, ldb int, sdim int, alphar []float64, alphai []float64, beta []float64, vsl []float64, ldvsl int, vsr []float64, ldvsr int) Info

	//Missing: ggev ggevx ggglm gghrd gglse ggqrf ggrqf ggsvd ggsvp gtcon gtrfs gtsvx
	Dgtsv(n int, nrhs int, dl []float64, d []float64, du []float64, b []float64, ldb int) Info

	//Missing: hgeqz hsein hseqr opgtr opmtr orgbr orghr orgql orgrq
	Dgttrf(n int, dl []float64, d []float64, du []float64, du2 []float64, ipiv []int32) Info
	Dgttrs(ta blas.Transpose, n int, nrhs int, dl []float64, d []float64, du []float64, du2 []float64, ipiv []int32, b []float64, ldb int) Info
	Dorglq(m int, n int, k int, a []float64, lda int, tau []float64) Info
	Dorgqr(m int, n int, k int, a []float64, lda int, tau []float64) Info

	//Missing: orgtr ormbr ormhr ormql ormrq ormrz ormtr pbcon pbequ pbrfs pbstf
	Dormlq(s blas.Side, ta blas.Transpose, m int, n int, k int, a []float64, lda int, tau []float64, c []float64, ldc int) Info
	Dormqr(s blas.Side, ta blas.Transpose, m int, n int, k int, a []float64, lda int, tau []float64, c []float64, ldc int) Info

	//Missing: pbsvx pftrf pftri pftrs pocon poequ poequb porfs porfsx
	Dpbsv(ul blas.Uplo, n int, kd int, nrhs int, ab []float64, ldab int, b []float64, ldb int) Info
	Dpbtrf(ul blas.Uplo, n int, kd int, ab []float64, ldab int) Info
	Dpbtrs(ul blas.Uplo, n int, kd int, nrhs int, ab []float64, ldab int, b []float64, ldb int) Info
	Dposv(ul blas.Uplo, n int, nrhs int, a []float64, lda int, b []float64, ldb int) Info

	//Missing: posvx posvxx ppcon ppequ pprfs ppsv ppsvx pptrf pptri pptrs
	Dpotrf(ul blas.Uplo, n int, a []float64, lda int) Info
	Dpotri(ul blas.Uplo, n int, a []float64, lda int) Info
	Dpotrs(ul blas.Uplo, n int, nrhs int, a []float64, lda int, b []float64, ldb int) Info

	//Missing: pstrf ptcon pteqr ptrfs ptsvx sbev sbevd sbevx sbgst sbgv
	Dptsv(n int, nrhs int, d []float64, e []float64, b []float64, ldb int) Info
	Dpttrf(n int, d []float64, e []float64) Info
	Dpttrs(n int, nrhs int, d []float64, e []float64, b []float64, ldb int) Info

	//Missing: sbgvd sbgvx sbtrd sfrk spcon spev spevd spevx spgst spgv spgvd spgvx sprfs spsv

	//Missing: spsvx sptrd sptrf sptri sptrs stebz stedc stegr stein stemr steqr sterf stev

	//Missing: stevd stevr stevx sycon syequb sygst sygvd sygvx
	Dsyev(jobz Job, ul blas.Uplo, n int, a []float64, lda int, w []float64) Info
	Dsyevd(jobz Job, ul blas.Uplo, n int, a []float64, lda int, w []float64) Info
	Dsyevr(jobz Job, r Range, ul blas.Uplo, n int, a []float64, lda int, vl float64, vu float64, il int, iu int, abstol float64, m int, w []float64, z []float64, ldz int, isuppz []int) Info
	Dsyevx(jobz Job, r Range, ul blas.Uplo, n int, a []float64, lda int, vl float64, vu float64, il int, iu int, abstol float64, m int, w []float64, z []float64, ldz int, ifail []int) Info
	Dsygv(t Type, jobz Job, ul blas.Uplo, n int, a []float64, lda int, b []float64, ldb int, w []float64) Info

	//Missing: syrfs syrfsx sysvx sysvxx sytrd tbcon tbrfs tfsm
	Dsysv(ul blas.Uplo, n int, nrhs int, a []float64, lda int, ipiv []int32, b []float64, ldb int) Info
	Dsytrf(ul blas.Uplo, n int, a []float64, lda int, ipiv []int32) Info
	Dsytri(ul blas.Uplo, n int, a []float64, lda int, ipiv []int32) Info
	Dsytrs(ul blas.Uplo, n int, nrhs int, a []float64, lda int, ipiv []int32, b []float64, ldb int) Info
	Dtbtrs(ul blas.Uplo, ta blas.Transpose, d blas.Diag, n int, kd int, nrhs int, ab []float64, ldab int, b []float64, ldb int) Info

	//Missing: tftri tfttp tfttr tgevc tgexc tgsen tgsja tgsna tgsyl tpcon tprfs tptri tptrs

	//Missing: tpttf tpttr trcon trevc trexc trrfs trsen trsna trsyl trttf trttp tzrzf
	Dtrtrs(ul blas.Uplo, ta blas.Transpose, d blas.Diag, n int, nrhs int, a []float64, lda int, b []float64, ldb int) Info
	Dtrtri(ul blas.Uplo, d blas.Diag, n int, a []float64, lda int) Info
}
