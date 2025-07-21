package sequence

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/compermane/ic-go/pkg/domain/testfunction"
)

type Sequence struct {
	SequenceID 		uint64
	Functions 		[]*testfunction.TestFunction
	ReturnedValues	map[string][]reflect.Value
	ExtensibleFlag  []bool
}

/* Creates a new sequence.
 * :param functions: list of sequence functions
 * :param receivers: list of receivers associated with the functions (can be nil)
 * :returns: pointer to a new sequence
 */
func NewSequence(functions []*testfunction.TestFunction) *Sequence {
	str := ""
	for _, fn := range functions {
		str += fn.Name
		for _, arg_value := range fn.ArgValues {
			str += "(" + fmt.Sprintf("%+v", arg_value) + ")"
		}
	}

	sq_id := create_id(str)

	returned_values := make(map[string][]reflect.Value, 0)
	extensible_flag := make([]bool, 0)

	return &Sequence{
		SequenceID: sq_id,
		Functions: functions,
		ReturnedValues: returned_values,
		ExtensibleFlag: extensible_flag,
	}
}

/* Choses a random sequence.
 * :param seqs: List of sequences
 * :returns: Selected sequence
 */
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

func VerifyExistenceByHash(seqs_hash_map map[uint64]bool, seq_hash uint64) bool {
	return verify_existence_by_hash(seqs_hash_map, seq_hash)
}
/* Concatenates one sequence to another.
 * :param seq: sequence concatenated to the sq receiver
 * :returns: new sequence
 */
func (sq *Sequence) AppendSequence(seq *Sequence) *Sequence {
	newSeq := &Sequence{
		Functions: append([]*testfunction.TestFunction{}, sq.Functions...),
		ReturnedValues: make(map[string][]reflect.Value),
	}

	for chave, valores := range sq.ReturnedValues {
		newSlice := make([]reflect.Value, len(valores))
		copy(newSlice, valores)
		newSeq.ReturnedValues[chave] = newSlice
	}

	new_slice := make([]bool, len(sq.ExtensibleFlag))
	copy(new_slice, sq.ExtensibleFlag)
	newSeq.ExtensibleFlag = new_slice

	newSeq.Functions = append(newSeq.Functions, seq.Functions...)

	// Atualiza o ID da sequência, se necessário, usando a nova sequência como base
	newSeq.SequenceID = update_id(newSeq, seq)

	// Aqui você pode, opcionalmente, ajustar ou limpar outros estados que não deseja herdar

	return newSeq
}

/* Puts into a sequence sq the values returned from a reflect.Call execution
 * :param return_type: type of param value
 * :param value: retuned value
 */
func (sq *Sequence) AppendReturnedValue(return_type string, value reflect.Value) {
	_, exist := sq.ReturnedValues[return_type]

	if exist {
		sq.ReturnedValues[return_type] = append(sq.ReturnedValues[return_type], value)
	} else {
		sq.ReturnedValues[return_type] = make([]reflect.Value, 0)
		sq.ReturnedValues[return_type] = append(sq.ReturnedValues[return_type], value)
	}
}

/* Yields a returned value from the list of returned values of a sequence sq based on its type.
 * Should be used in sequence extension.
 * :param value_type: type of the desired random value
 * :returns: a chosen reflect.Value and a flag that indicates if it was successful (true) or not (false)
 */
func (sq *Sequence) GetRandomReturnedValue(value_type string) (reflect.Value, bool) {
	val, ok := sq.ReturnedValues[value_type]
	if !ok {
		return reflect.Zero(reflect.TypeOf((*any)(nil))), false
	}
	
	slice := reflect.ValueOf(val)
	if slice.Kind() != reflect.Slice {
		return reflect.Zero(reflect.TypeOf((*any)(nil))), false
	}
	if slice.Len() == 0 {
		return reflect.Zero(reflect.TypeOf((*any)(nil))), false
	}
	
	source := rand.NewSource(time.Now().UnixNano())
	rng	   := rand.New(source)

	random_index := rng.Intn(slice.Len())

	// fmt.Printf("a: %v\nb: %v\n", len(sq.ExtensibleFlag), slice.Len())
	// for len(sq.ExtensibleFlag) > 0 && !sq.ExtensibleFlag[random_index] {
	// 	fmt.Println("BRUH")
	// 	random_index = rng.Intn(slice.Len())
	// }

	result := slice.Index(random_index)

	if result.Type() == reflect.TypeOf(reflect.Value{}) {
		result = result.Interface().(reflect.Value)
	}

	return result, true
	
}

func (sq *Sequence) ApplyExtensibleFlags(ret_type string, ret_value reflect.Value) {
	error_type  := reflect.TypeOf((*error)(nil)).Elem()
	not_equal   := true
	not_nil		:= true
	not_error   := true
	vals, exist := sq.ReturnedValues[ret_type]

	// o' = o for some o returned from the sequence
	if exist {
		for _, value := range vals {
			if !value.IsValid() || !value.CanInterface() || !ret_value.IsValid() || !ret_value.CanInterface() {
				continue
			}

			if reflect.DeepEqual(value.Interface(), ret_value.Interface()) {
				not_equal = false
			}

			if value.Kind() == reflect.Ptr && value.IsValid() && !value.IsNil() && ret_value.Kind() == reflect.Ptr && ret_value.IsValid() && !ret_value.IsNil() {
				v1 := value.Elem()
				v2 := value.Elem()

				if v1.IsValid() && v1.CanInterface() && v2.IsValid() && v2.CanInterface() {
					if reflect.DeepEqual(v1.Interface(), v2.Interface()) {
						not_equal = false
					}
				}
			}
		}
	}

	switch ret_value.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface,
		reflect.Map, reflect.Ptr, reflect.Slice:
			if !ret_value.IsValid() || ret_value.IsNil() {
				not_nil = false
	}
	}

	if ret_value.Type().Implements(error_type) {
		not_error = false
	}
	
	if not_equal && not_nil && not_error {
		sq.ExtensibleFlag = append(sq.ExtensibleFlag, true)
	} else {
		sq.ExtensibleFlag = append(sq.ExtensibleFlag, false)
	}
}

/* Represents sq sequence functions as a string.
 * :returns: string containing all sequence functions
 */
func (sq *Sequence) String() string {
	str := "[ "

	for _, fn := range sq.Functions {
		str += fn.Name + " ( "

		for _, arg := range fn.ArgValues {
			if arg.IsValid() {
				str += fmt.Sprintf("%+v ", arg.Interface())
			} else {
				str += fmt.Sprintf("nil ")
			}
		}

		str += ")"
	}

	str += "]"
	return str
}