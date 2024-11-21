package receiver

import "fmt"

func (rcv *Receiver) Print() {
	str_names := ""
	str_types := ""
	str_values := "\n" 
	mtd_names := ""

	for i := 0; i < len(rcv.AttrNames); i++ {
		if len(rcv.AttrNames) > 0 {
			str_names = str_names + rcv.AttrNames[i] + ", "
		}
	}
	
	for i := 0; i < len(rcv.AttrTypes); i++ {
		if len(rcv.AttrTypes) > 0 {
			str_types = str_types + rcv.AttrTypes[i].String() + ", "
		}
	}
	
	for i := 0; i < len(rcv.MethodNames); i++ {
		if len(rcv.MethodNames) > 0 {
			mtd_names = mtd_names + rcv.MethodNames[i] + ", "
		}
	}

	for i := 0; i < len(rcv.AttrValues); i++ {
		if len(rcv.AttrValues) > i {
			str_values = str_values + fmt.Sprintf("%v: %v",rcv.AttrNames[i], rcv.AttrValues[i]) + "\n"
		}
	}

	fmt.Println("Receiver name: " + rcv.Name)
	fmt.Println("Methods: " + mtd_names)
	fmt.Println("Attributes: " + str_names)
	fmt.Println("Attributes types: " + str_types)
	fmt.Println("Attributes values: " + str_values)
}