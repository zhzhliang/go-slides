// +build OMIT

// START0 OMIT
for {
	select {
	case v1 := <- c1:
		// handle ...
	case v2 := <- c2:
		// handle ...
	}
}
// STOP0 OMIT

// START1 OMIT
for _, s := range []string{"a", "b", "c"} {
	select {
	case ch <- s:
	}
}
// STOP1 OMIT