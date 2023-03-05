package helper

// PanicIfError is function for panic application
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
