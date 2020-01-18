package sysinit

//在main函数调用之前只会调用一次
func init() {
	sysinit()
	dbinit()
}
