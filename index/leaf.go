package index

// data page "leaf node"
type d struct {
	c int
	d [2*kd + 1]de
	n *d
	p *d
}

// data page (leaf nodes's) data
type de struct {
	k interface{} /*K*/
	v interface{} /*V*/
}

func (l *d) mvL(r *d, c int) {
	copy(l.d[l.c:], r.d[:c])
	copy(r.d[:], r.d[c:r.c])
	l.c += c
	r.c -= c
}

func (l *d) mvR(r *d, c int) {
	copy(r.d[c:], r.d[:r.c])
	copy(r.d[:c], l.d[l.c-c:])
	r.c += c
	l.c -= c
}
