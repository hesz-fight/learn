package leftrotate

func leftRotate(str string, n int) string {
	bs := []byte(str)
	ans := []byte{}
	ans = append(ans, bs[n:]...)
	ans = append(ans, bs[:n]...)

	return string(ans)
}
