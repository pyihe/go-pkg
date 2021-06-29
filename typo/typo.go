package typo

func Bool(b bool) *bool {
	var ptr = new(bool)
	*ptr = b
	return ptr
}

func String(s string) *string {
	var ptr = new(string)
	*ptr = s
	return ptr
}

func Byte(b byte) *byte {
	var ptr = new(byte)
	*ptr = b
	return ptr
}

func Uint8(u uint8) *uint8 {
	var ptr = new(uint8)
	*ptr = u
	return ptr
}

func Int8(i int8) *int8 {
	var ptr = new(int8)
	*ptr = i
	return ptr
}

func Int16(i int16) *int16 {
	var ptr = new(int16)
	*ptr = i
	return ptr
}

func Uint16(u uint16) *uint16 {
	var ptr = new(uint16)
	*ptr = u
	return ptr
}

func Int(i int) *int {
	var ptr = new(int)
	*ptr = i
	return ptr
}

func Uint(u uint) *uint {
	var ptr = new(uint)
	*ptr = u
	return ptr
}

func Int32(i int32) *int32 {
	var ptr = new(int32)
	*ptr = i
	return ptr
}

func Uint32(u uint32) *uint32 {
	var ptr = new(uint32)
	*ptr = u
	return ptr
}

func Int64(i int64) *int64 {
	var ptr = new(int64)
	*ptr = i
	return ptr
}

func Uint64(u uint64) *uint64 {
	var ptr = new(uint64)
	*ptr = u
	return ptr
}

func Float32(f float32) *float32 {
	var ptr = new(float32)
	*ptr = f
	return ptr
}

func Float64(f float64) *float64 {
	var ptr = new(float64)
	*ptr = f
	return ptr
}
