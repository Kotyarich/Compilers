package dfa

import (
	"cc/lab/queue"
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

// При построении ДКА каждое состояние является множеством номеров
// листьев в синтаксическом дереве регулярки, перехождим к формату ДКА,
// в котором каждое состояние - уникальное число
func toMinDFA(dfa tempDFA) MinDFA {
	result := MinDFA{t: dfa.t}

	stateMap := make(map[*state]int)
	// Маппинг состояний
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
	// Маппинг конечных состояний
	for _, f := range dfa.f {
		for k, v := range stateMap {
			if k.value.Equals(f) {
				result.f.Add(v)
			}
		}
	}
	// Маппинг переходов
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

func swapState(P []set.IntSet, state int, from int, to int) []set.IntSet {
	for i, v := range P[from] {
		if v == state {
			P[from] = append(P[from][:i], P[from][i+1:]...)
			break
		}
	}

	P[to] = append(P[to], state)
	return P
}

func mapClasses(P []set.IntSet) map[int]int {
	stateClass := make(map[int]int)

	for classIndex, _ := range P {
		for _, v := range P[classIndex] {
			stateClass[v] = classIndex
		}
	}

	return stateClass
}

func mapTransicions(trans []Tran) map[byte]map[int]*set.IntSet {
	inv := make(map[byte]map[int]*set.IntSet)

	for _, t := range trans {
		if _, ok := inv[t.Symbol]; !ok {
			inv[t.Symbol] = make(map[int]*set.IntSet)
		}
		if _, ok := inv[t.Symbol][t.To]; !ok {
			inv[t.Symbol][t.To] = &set.IntSet{}
		}
		inv[t.Symbol][t.To].Add(t.From)
	}

	return inv
}

func MinimiseOptimal(dfa MinDFA) []set.IntSet {
	P := []set.IntSet{dfa.f, dfa.q.Subtract(dfa.f)}

	stateClass := mapClasses(P)
	inv := mapTransicions(dfa.d)

	charQueue := queue.Queue{}
	setQueue := queue.Queue{}

	for _, c := range dfa.t {
		charQueue.Push(c)
		charQueue.Push(c)
		setQueue.Push(dfa.f)
		setQueue.Push(dfa.q.Subtract(dfa.f))
	}

	for !charQueue.Empty() {
		C := setQueue.Pop().(set.IntSet)
		a := charQueue.Pop().(byte)

		involved := make(map[int]*set.IntSet, 0)
		for _, q := range C {
			rs, ok := inv[a][q]
			if !ok {
				continue
			}

			for _, r := range *rs {
				i := stateClass[r]
				if _, ok := involved[i]; !ok {
					involved[i] = &set.IntSet{}
				}
				involved[i].Add(r)
			}
		}

		for i := range involved {
			if involved[i].Size() < len(P[i]) {
				P = append(P, set.IntSet{})
				j := len(P) - 1

				for _, r := range *involved[i] {
					P = swapState(P, r, i, j)
				}

				if len(P[j]) > len(P[i]) {
					P[j], P[i] = P[i], P[j]
				}

				for _, r := range P[j] {
					stateClass[r] = j
				}

				for _, c := range dfa.t {
					charQueue.Push(c)
					setQueue.Push(P[j])
				}
			}
		}
	}

	return P
}