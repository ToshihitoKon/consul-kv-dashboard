// generated by stringer -type=Status; DO NOT EDIT

package main

import "fmt"

const _Status_name = "SuccessInfoWarningDanger"

var _Status_index = [...]uint8{7, 11, 18, 24}

func (i Status) String() string {
	if i < 0 || i >= Status(len(_Status_index)) {
		return fmt.Sprintf("Status(%d)", i)
	}
	hi := _Status_index[i]
	lo := uint8(0)
	if i > 0 {
		lo = _Status_index[i-1]
	}
	return _Status_name[lo:hi]
}
