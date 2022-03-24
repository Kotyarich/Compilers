package dfa

import (
	"cc/lab/set"
)

type MinDFA struct {
	q  set.IntSet
	f  set.IntSet
	q0 int
	t set.ByteSet
	d []Tran
}

type Tran struct {
	Symbol byte
	From   int
	To     int
}

func ToMinDFA(dfa tempDFA) MinDFA {
	result := MinDFA{t: dfa.t}

	stateMap := make(map[*state]int)

	curState := 0
	for _, q := range dfa.q {
		if _, ok := stateMap[q]; !ok {
			stateMap[q] = curState
			result.q.Add(curState)
			curState++
		}

		if q.value.Equals(dfa.q0) {
			result.q0 = stateMap[q]
		}
	}

	for _, f := range dfa.f {
		for k, v := range stateMap {
			if k.value.Equals(f) {
				result.f.Add(v)
			}
		}
	}

	for _, t := range dfa.d {
		founds := 0
		tran := Tran{Symbol: t.symbol}
		for k, v := range stateMap {
			if k.value.Equals(t.state) {
				tran.From = v
				founds++
			}
			if k.value.Equals(t.destState) {
				tran.To = v
				founds++
			}
			if founds == 2 {
				break
			}
		}
		result.d = append(result.d, tran)
	}

	return result
}

func Minimise(dfa MinDFA) []set.IntSet {
	P := []set.IntSet{dfa.f, dfa.q.Subtract(dfa.f)}
	W := []set.IntSet{dfa.f, dfa.q.Subtract(dfa.f)}

	for len(W) > 0 {
		A := W[0]
		W = W[1:]

		for _, c := range dfa.t {
			// let X be the set of states for which a transition on c leads to a state in A
			var X set.IntSet
			for _, t := range dfa.d {
				if t.Symbol == c && A.Contains(t.To) {
					X = append(X, t.From)
				}
			}

			// for each set Y in P for which X ∩ Y is nonempty and Y \ X is nonempty
			newP := make([]set.IntSet, 0, cap(P))
			for _, Y := range P {
				intersection := X.Intersect(Y)
				if len(intersection) == 0 {
					newP = append(newP, Y)
					continue
				}

				subtraction := Y.Subtract(X)
				if len(subtraction) == 0 {
					newP = append(newP, Y)
					continue
				}

				// replace Y in P by the two sets X ∩ Y and Y \ X
				newP = append(newP, intersection, subtraction)

				// if Y is in W replace by the same two sets
				foundY := false
				for i, sets := range W {
					if sets.Equals(Y) {
						foundY = true
						W[i] = intersection
						W = append(W, subtraction)
						break
					}
				}

				if !foundY {
					if len(intersection) <= len(subtraction) {
						W = append(W, intersection)
					} else {
						W = append(W, subtraction)
					}
				}
			}
			P = newP
		}
	}

	return P
}
