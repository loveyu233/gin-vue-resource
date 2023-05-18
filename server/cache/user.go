package cache

func VerCodeIsOk(email, x string) bool {
	if x == Get(email) {
		Del(email)
		return true
	}
	return false
}
