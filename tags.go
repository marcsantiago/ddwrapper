package ddwrapper

type Tags []string

func (t Tags) Pair() []string {
	// array needs to be even in len
	if len(t)&1 == 1 {
		return t
	}
	tag := make([]string, len(t)/2)
	for i := 0; i < len(t)/2; i++ {
		tag[i] = t[i*2] + ":" + t[i*2+1]
	}
	return tag
}
