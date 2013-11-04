package p2

type T1 struct {
	Exported   int
	unexported int
}

type T2 struct {
	unexported int
	Exported   int
}

type T3 struct {
	// Doc comment.
	Exported   int
	unexported int
}

type T4 struct {
	unexported int
	// Doc comment.
	Exported   int
}
