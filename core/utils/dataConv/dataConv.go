package dataConv

import "strconv"

func InitConv() *DataConv {
	return &DataConv{}
}

func (root *DataConv) Str2uint(obj string) uint {
	intNum, _ := strconv.Atoi(obj)
	return uint(intNum)
}

func (root *DataConv) Uint2str(obj uint) string {
	return strconv.Itoa(int(obj))
}
