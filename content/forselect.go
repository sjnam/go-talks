// START1 OMIT
for { // 무한 반복 또는 특정 범위에 대한 루프
	select {
		// 채널에 대한 작업
	}
}
// STOP1 OMIT

// START2 OMIT
for _, s := range []string{"a", "b", "c"} {
	select {
	case <-done:
		return
	case stringStream <- s:
	}
}
// STOP2 OMIT

// START3 OMIT
for {
	select {
	case <-done:
		return
	default:
	}
	// 선점 불가능한 작업 수행
}
// STOP3 OMIT

// START4 OMIT
for {
	select {
	case <-done:
		return
	default:
		// 선점 불가능한 작업 수행
	}
}
// STOP4 OMIT
