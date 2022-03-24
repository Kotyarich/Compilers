package dfa

func toTreeMap(t *tree) map[int]*tree {
	m := make(map[int]*tree)
	m[t.Pos] = t

	for _, c := range t.Children {
		treeMapRec(c, m)
	}

	return m
}

func treeMapRec(t *tree, m map[int]*tree) {
	m[t.Pos] = t

	for _, c := range t.Children {
		treeMapRec(c, m)
	}
}
