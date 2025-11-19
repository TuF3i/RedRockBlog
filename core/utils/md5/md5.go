package md5

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
)

func GenMD5(userID int64) string {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, uint64(userID))
	result := buf[:n]

	hash := md5.Sum(result)
	md5String := hex.EncodeToString(hash[:])
	return md5String
}
