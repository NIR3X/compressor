package compressor

import (
	"unsafe"

	"github.com/NIR3X/varsizedint"
)

const (
	MinChunkSize                 = 3
	ChunkSame, ChunkMixed uint64 = 0, 1
)

func GetNextChunkSame(src []uint8, pos uint64, left, right *uint64) bool {
	var (
		chunkValue int16  = -1
		chunkSize  uint64 = 0
		srcSize           = uint64(len(src))
	)

	for i := pos; i < srcSize; i++ {
		var curr = src[i]
		if chunkValue == int16(curr) {
			chunkSize++
		}

		if chunkValue != int16(curr) || i == srcSize-1 {
			if chunkSize >= MinChunkSize {
				if chunkValue == int16(curr) {
					*right = i + 1
				} else {
					*right = i
				}
				return true
			}
			chunkValue = int16(curr)
			chunkSize = 1
			*left = i
		}
	}

	return false
}

func PutChunkMixed(src []uint8, i *uint64, iEnd uint64, dest []uint8, destPos *uint64) {
	var (
		i_             = *i
		destPos_       = *destPos
		chunkMixedSize = iEnd - i_
	)
	if chunkMixedSize > 0 {
		destPos_ += uint64(varsizedint.Encode(
			dest[destPos_:],
			(chunkMixedSize<<1)|ChunkMixed,
		))
		for j := uint64(0); j < chunkMixedSize; j++ {
			dest[destPos_] = src[i_]
			i_++
			destPos_++
		}
		*i = i_
		*destPos = destPos_
	}
}

func Compress(src []uint8, pDest *[]uint8) uint64 {
	var (
		destPos uint64 = 0
		srcSize        = uint64(len(src))
		dest           = make([]uint8, srcSize*2)
	)
	for i := uint64(0); i < srcSize; i++ {
		var left, right uint64 = 0, 0
		if GetNextChunkSame(src, i, &left, &right) {
			PutChunkMixed(src, &i, left, dest, &destPos)
			var chunkSameSize = right - left
			destPos += uint64(varsizedint.Encode(
				dest[destPos:],
				(chunkSameSize<<1)|ChunkSame,
			))
			dest[destPos] = src[left]
			destPos += uint64(unsafe.Sizeof(src[left]))
			i += chunkSameSize - 1
		} else {
			PutChunkMixed(src, &i, srcSize, dest, &destPos)
			*pDest = dest
			return destPos
		}
	}
	*pDest = dest
	return destPos
}

func Decompress(src, dest []uint8) uint64 {
	var (
		destPos uint64 = 0
		srcSize        = uint64(len(src))
	)
	for i := uint64(0); i < srcSize; {
		var (
			chunkSize = varsizedint.Decode(src[i:])
			chunkType = chunkSize & 1
		)
		var chunkSizeSize = uint64(varsizedint.ParseSize(src[i:]))
		i += chunkSizeSize
		chunkSize >>= 1
		switch chunkType {
		case ChunkSame:
			{
				var chunkSameValue = src[i]
				for j := uint64(0); j < chunkSize; j++ {
					dest[destPos] = chunkSameValue
					destPos++
				}
				i++
			}
		case ChunkMixed:
			for j := uint64(0); j < chunkSize; j++ {
				dest[destPos] = src[i]
				destPos++
				i++
			}
		}
	}
	return destPos
}
