package methodsandinterfaces

import "fmt"

type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", i[0], i[1], i[2], i[3])
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negitive number : %d", int(e))
}

func Sqrt(arg int) (*int, ErrNegativeSqrt) {
	return nil, ErrNegativeSqrt(arg)
}
