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

func verify_existence_by_hash(seqs_hash_map map[uint64]bool, hash uint64) bool {
	_, ok := seqs_hash_map[hash]

	// Chave não existe
	if !ok {
		return false
	}

	return true
}

func GetAllHashesFromSequences(sqs []*Sequence) map[uint64]bool {
	hash_map := make(map[uint64]bool, 0)

	for _, sq := range sqs {
		hash_map[sq.SequenceID] = true
	}

	return hash_map
}

func UpdateHashMap(hash_map map[uint64]bool, sqs []*Sequence) map[uint64]bool {
	new_hash_map := GetAllHashesFromSequences(sqs)

	for key, _ := range new_hash_map {
		_, exist := hash_map[key]
		if !exist {
			hash_map[key] = true
		}
	}

	return hash_map
}