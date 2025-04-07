package sequence

import (
	"fmt"
	"hash/fnv"
)


func create_id(str string) uint64 {
	hasher := fnv.New64a()
	hasher.Write([]byte(str))

	return uint64(hasher.Sum64())
}

func update_id(sq1, sq2 *Sequence) uint64 {
	str := ""

	for _, fn := range sq1.Functions {
		str += fn.Name
		for _, fn_args := range fn.ArgValues {
			str += fmt.Sprintf("%+v", fn_args)
		}
	}

	for _, fn := range sq1.Functions {
		str += fn.Name
		for _, fn_args := range fn.ArgValues {
			str += fmt.Sprintf("%+v", fn_args)
		}
	}

	return create_id(str)
}

func verify_equal_sequences(sq1, sq2 *Sequence) bool {
	if sq1 == nil || sq2 == nil {
		return false
	}
	
	return sq1.SequenceID == sq2.SequenceID
}