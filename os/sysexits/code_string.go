// Code generated by "stringer -type=Code"; DO NOT EDIT.

package sysexits

import "strconv"

const (
	_Code_name_0 = "OK"
	_Code_name_1 = "UsageDataErrorNoInputNoUserNoHostUnavailableSoftwareOSErrorOSFileCantCreateToErrorTempFailureProtocolNoPermissionConfig"
)

var (
	_Code_index_1 = [...]uint8{0, 5, 14, 21, 27, 33, 44, 52, 59, 65, 75, 82, 93, 101, 113, 119}
)

func (i Code) String() string {
	switch {
	case i == 0:
		return _Code_name_0
	case 64 <= i && i <= 78:
		i -= 64
		return _Code_name_1[_Code_index_1[i]:_Code_index_1[i+1]]
	default:
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}