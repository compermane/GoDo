package sequence

import (
	"fmt"
	"math/rand"

	"github.com/compermane/ic-go/pkg/domain/receiver"
	"github.com/compermane/ic-go/pkg/domain/testfunction"
)

type Sequence struct {
	SequenceID 	uint64
	Functions 	[]*testfunction.TestFunction
	Receivers 	[]*receiver.Receiver
}

/* Creates a new sequence.
 * :param functions: list of sequence functions
 * :param receivers: list of receivers associated with the functions (can be nil)
 * :returns: pointer to a new sequence
 */
func NewSequence(functions []*testfunction.TestFunction, receivers []*receiver.Receiver) *Sequence {
	str := ""
	for _, fn := range functions {
		str += fn.Name
		for _, arg_value := range fn.ArgValues {
			str += "(" + fmt.Sprintf("%+v", arg_value) + ")"
		}
	}
	fmt.Println(str)
	sq_id := create_id(str)
	fmt.Println(sq_id)
	return &Sequence{
		SequenceID: sq_id,
		Functions: functions,
		Receivers: receivers,
	}
}

func ChooseRandom(seqs []*Sequence) *Sequence {
	return seqs[rand.Intn(len(seqs))]
}

/* Verify if a sequence alredy exists on a list of sequences.
 * :param seqs: List of sequences
 * :param seq: Sequence to be tested
 * :returns: true if it exists, false otherwise
 */
func VerifyExistence(seqs []*Sequence, seq *Sequence) bool {
	for _, sq := range seqs {
		if verify_equal_sequences(seq, sq) {
			return true
		}
	}

	return false
}

/* Verify if it exists a duplicate in a list of sequences.
 * :param seqs: List to be tested
 * :returns: true if it exists, false otherwise
 */
func VerifyDuplicate(seqs []*Sequence) bool {
	for i := 0; i < len(seqs) - 1; i++ {
		for j := i + 1; j < len(seqs); j++ {
			if verify_equal_sequences(seqs[i], seqs[j]) {
				return true
			}
		}
	}

	return false
}

/* Concatenates one sequence to another.
 * :param seq: sequence concatenated to the sq receiver
 * :returns: none
 */
func (sq *Sequence) AppendSequence(seq *Sequence) {
	sq.Functions  = append(sq.Functions, seq.Functions...)
	sq.Receivers  = append(sq.Receivers, seq.Receivers...)
	sq.SequenceID = update_id(sq, seq)
}

/* Represents sq sequence functions as a string.
 * :returns: string containing all sequence functions
 */
func (sq *Sequence) String() string {
	str := "[ "

	for _, fn := range sq.Functions {
		str += fn.Name + " "
	}

	str += "]"
	return str
}