package readall

import "io"

// Replaces ioutil.ReadAll, when you know the (uncompressed) length.
func ReadAll(r io.Reader, s int) ([]byte, error) {
	var n,m int
	var e error
	b := make([]byte,s)
	for l:=s; l>0; l-=m {
		m, e = r.Read(b[n:s])
		n += m
		if e == io.EOF {
			return b[0:n], nil
		}
		if e != nil {
			return b, e
		}
	}
	return b, nil
}