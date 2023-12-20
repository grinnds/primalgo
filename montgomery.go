package primalgo

import "math/big"

var (
	intOne = big.NewInt(1)
	intTwo = big.NewInt(2)
)

func MontgomeryLadder(P, Q, n, N *big.Int) *big.Int {
	double := func(v, k *big.Int) *big.Int {
		//
		//  V(2k) = V(k)² - 2Qᵏ
		//
		t := new(big.Int).Mul(v, v)
		t2 := new(big.Int).Exp(Q, k, N)
		t2.Mul(t2, intTwo)
		t.Sub(t, t2)
		return t
	}
	add := func(vk, vk1, k *big.Int) *big.Int {
		//
		//  V(k+1) + V(k) = V(k+1) V(k) - Qᵏ P
		//
		t := new(big.Int).Mul(vk, vk1)
		t2 := new(big.Int).Exp(Q, k, N)
		t2.Mul(t2, P)
		t.Sub(t, t2)
		return t
	}

	vk := big.NewInt(2)
	vk1 := new(big.Int).Set(P)
	k := big.NewInt(0)
	for i := n.BitLen(); i >= 0; i-- {
		if n.Bit(i) == 0 {
			// k' = 2k
			vk1 = add(vk, vk1, k)
			vk = double(vk, k)

			k.Mul(k, intTwo)
		} else {
			// k' = 2k+1
			vk = add(vk, vk1, k)
			vk1 = double(vk1, k.Add(k, intOne))

			k.Mul(k, intTwo).Sub(k, intOne)
		}

		vk.Mod(vk, N)
		vk1.Mod(vk1, N)
	}
	return vk
}
