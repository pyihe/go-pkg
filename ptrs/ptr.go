package ptrs

/*
返回各种基本数据类型的指针
*/

// Bool 返回bool指针
func Bool(b bool) *bool {
	var ptr = new(bool)
	*ptr = b
	return ptr
}

// Uint8 返回uint8类型的指针
func Uint8(u uint8) *uint8 {
	var ptr = new(uint8)
	*ptr = u
	return ptr
}

// Uint16 返回uint16类型的指针
func Uint16(u uint16) *uint16 {
	var ptr = new(uint16)
	*ptr = u
	return ptr
}

// Uint32 返回uint32类型的指针
func Uint32(u uint32) *uint32 {
	var ptr = new(uint32)
	*ptr = u
	return ptr
}

// Uint64 返回uint64类型的指针
func Uint64(u uint64) *uint64 {
	var ptr = new(uint64)
	*ptr = u
	return ptr
}

// Int8 返回int8类型的指针
func Int8(i int8) *int8 {
	var ptr = new(int8)
	*ptr = i
	return ptr
}

// Int16 返回int16类型的指针
func Int16(i int16) *int16 {
	var ptr = new(int16)
	*ptr = i
	return ptr
}

// Int32 返回int32类型的指针
func Int32(i int32) *int32 {
	var ptr = new(int32)
	*ptr = i
	return ptr
}

// Int64 返回int64类型的指针
func Int64(i int64) *int64 {
	var ptr = new(int64)
	*ptr = i
	return ptr
}

// Float32 返回float32类型的指针
func Float32(f float32) *float32 {
	var ptr = new(float32)
	*ptr = f
	return ptr
}

// Float64 返回float64类型的指针
func Float64(f float64) *float64 {
	var ptr = new(float64)
	*ptr = f
	return ptr
}

// String 返回string类型的指针
func String(s string) *string {
	var ptr = new(string)
	*ptr = s
	return ptr
}

// Int 返回int类型的指针
func Int(i int) *int {
	var ptr = new(int)
	*ptr = i
	return ptr
}

// Uint 返回uint类型的指针
func Uint(u uint) *uint {
	var ptr = new(uint)
	*ptr = u
	return ptr
}

// Byte 返回byte类型的指针
func Byte(b byte) *byte {
	var ptr = new(byte)
	*ptr = b
	return ptr
}
