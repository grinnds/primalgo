package primalgo

import (
	"fmt"
	"math/big"
	"math/bits"
)

type MorrisonNumber struct {
	k, n *big.Int
	val  *big.Int
}

func (mn MorrisonNumber) String() string {
	return fmt.Sprintf("%s", mn.val.String())
}

func NewMorrisonNumber(k, n uint64) (*MorrisonNumber, error) {
	if k < 1 {
		return nil, fmt.Errorf("expected k positive integer, got %d", k)
	}
	if n < 2 {
		return nil, fmt.Errorf("expected n > 1, got %d", n)
	}

	powTwo := bits.TrailingZeros64(k)
	k >>= powTwo

	kBig := new(big.Int).SetUint64(k)
	nBig := new(big.Int).SetUint64(n)
	nBig.Lsh(nBig, uint(powTwo))

	val := new(big.Int).Exp(intTwo, nBig, nil)

	if val.Cmp(kBig) != 1 {
		return nil, fmt.Errorf("expected 2ⁿ > k, got 2ⁿ <= k")
	}

	val.Mul(val, kBig)
	val.Sub(val, intOne)

	return &MorrisonNumber{
		k:   kBig,
		n:   nBig,
		val: val,
	}, nil
}

func MorrisonTest(N *MorrisonNumber) bool {
	// https://eprint.iacr.org/2023/195
	var P, Q *big.Int

	// Mersenne Prime
	twoSqrt := new(big.Int).Add(N.n, intOne)
	twoSqrt.Div(twoSqrt, intTwo)
	twoSqrt.Exp(intTwo, twoSqrt, N.val)

	P = twoSqrt
	Q = new(big.Int).Neg(intOne)
	Q.Mod(Q, N.val)

	n := new(big.Int).Set(N.val)
	n.Add(n, intOne)
	n.Div(n, intTwo)

	vn := MontgomeryLadder(P, Q, n, N.val)

	// V(n) = 0 mod N
	if vn.BitLen() == 0 {
		return true
	}
	return false
}
