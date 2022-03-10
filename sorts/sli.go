package sorts

import "sort"

type Uint8Slice []uint8
type Int8Slice []int8
type Uint16Slice []uint16
type Int16Slice []int16
type IntSlice []int
type UintSlice []uint
type Int32Slice []int32
type Uint32Slice []uint32
type Int64Slice []int64
type Uint64Slice []uint64
type Float32Slice []float32
type Float64Slice []float64
type StringSlice []string

func SortUint8s(s []uint8) {
	sort.Sort(Uint8Slice(s))
}

func SortInt8s(s []int8) {
	sort.Sort(Int8Slice(s))
}

func SortUint16s(s []uint16) {
	sort.Sort(Uint16Slice(s))
}

func SortInt16s(s []int16) {
	sort.Sort(Int16Slice(s))
}

func SortInts(s []int) {
	sort.Sort(IntSlice(s))
}

func SortUints(s []uint) {
	sort.Sort(UintSlice(s))
}

func SortInt32s(s []int32) {
	sort.Sort(Int32Slice(s))
}

func SortUint32s(s []uint32) {
	sort.Sort(Uint32Slice(s))
}

func SortInt64s(s []int64) {
	sort.Sort(Int64Slice(s))
}

func SortUint64s(s []uint64) {
	sort.Sort(Uint64Slice(s))
}

func SortFloat64s(s []float64) {
	sort.Sort(Float64Slice(s))
}

func SortFloat32s(s []float32) {
	sort.Sort(Float32Slice(s))
}

func SortStrings(s []string) {
	sort.Sort(StringSlice(s))
}

func (u Uint8Slice) Len() int {
	return len(u)
}

func (u Uint8Slice) Less(i, j int) bool {
	return u[i] < u[j]
}

func (u Uint8Slice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (i8s Int8Slice) Len() int {
	return len(i8s)
}

func (i8s Int8Slice) Less(i, j int) bool {
	return i8s[i] < i8s[j]
}

func (i8s Int8Slice) Swap(i, j int) {
	i8s[i], i8s[j] = i8s[j], i8s[i]
}

func (u16s Uint16Slice) Len() int {
	return len(u16s)
}

func (u16s Uint16Slice) Less(i, j int) bool {
	return u16s[i] < u16s[j]
}

func (u16s Uint16Slice) Swap(i, j int) {
	u16s[i], u16s[j] = u16s[j], u16s[i]
}

func (i16s Int16Slice) Len() int {
	return len(i16s)
}

func (i16s Int16Slice) Less(i, j int) bool {
	return i16s[i] < i16s[j]
}

func (i16s Int16Slice) Swap(i, j int) {
	i16s[i], i16s[j] = i16s[j], i16s[i]
}

func (i16s IntSlice) Len() int {
	return len(i16s)
}

func (i16s IntSlice) Less(i, j int) bool {
	return i16s[i] < i16s[j]
}

func (i16s IntSlice) Swap(i, j int) {
	i16s[i], i16s[j] = i16s[j], i16s[i]
}

func (i16s UintSlice) Len() int {
	return len(i16s)
}

func (i16s UintSlice) Less(i, j int) bool {
	return i16s[i] < i16s[j]
}

func (i16s UintSlice) Swap(i, j int) {
	i16s[i], i16s[j] = i16s[j], i16s[i]
}

func (i32s Int32Slice) Len() int {
	return len(i32s)
}

func (i32s Int32Slice) Less(i, j int) bool {
	return i32s[i] < i32s[j]
}

func (i32s Int32Slice) Swap(i, j int) {
	i32s[i], i32s[j] = i32s[j], i32s[i]
}

func (u32s Uint32Slice) Len() int {
	return len(u32s)
}

func (u32s Uint32Slice) Less(i, j int) bool {
	return u32s[i] < u32s[j]
}

func (u32s Uint32Slice) Swap(i, j int) {
	u32s[i], u32s[j] = u32s[j], u32s[i]
}

func (u64s Uint64Slice) Len() int {
	return len(u64s)
}

func (u64s Uint64Slice) Less(i, j int) bool {
	return u64s[i] < u64s[j]
}

func (u64s Uint64Slice) Swap(i, j int) {
	u64s[i], u64s[j] = u64s[j], u64s[i]
}

func (i64s Int64Slice) Len() int {
	return len(i64s)
}

func (i64s Int64Slice) Less(i, j int) bool {
	return i64s[i] < i64s[j]
}

func (i64s Int64Slice) Swap(i, j int) {
	i64s[i], i64s[j] = i64s[j], i64s[i]
}

func (f32s Float32Slice) Len() int {
	return len(f32s)
}

func (f32s Float32Slice) Less(i, j int) bool {
	return f32s[i] < f32s[j]
}

func (f32s Float32Slice) Swap(i, j int) {
	f32s[i], f32s[j] = f32s[j], f32s[i]
}

func (f64s Float64Slice) Len() int {
	return len(f64s)
}

func (f64s Float64Slice) Less(i, j int) bool {
	return f64s[i] < f64s[j]
}

func (f64s Float64Slice) Swap(i, j int) {
	f64s[i], f64s[j] = f64s[j], f64s[i]
}

func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
