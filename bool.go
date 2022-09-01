/*
 * @Author: leoking
 * @Date: 2022-09-01 19:46:33
 * @LastEditTime: 2022-09-01 20:39:20
 * @LastEditors: leoking
 * @Description:
 */
package conv

import (
	"strings"
)

type Boolable interface {
	string | bool | uint8 | uint16 | uint | uint32 | uint64 | int8 | int16 | int | int32 | int64 | float32 | float64
}

/*
*
  - @description:
    将常见取值转为bool
  - @param {any} v
  - @return {*}
*/
func Bool(v any) bool {
	if tmp, ok := v.(string); ok {
		v = strings.Title(tmp)
		// fmt.Println(v, tmp)
	}
	if v == "是" || v == "Yes" || v == "Y" || v == 1 || v == "1" || v == true || v == "True" {
		return true
	}
	return false
}
