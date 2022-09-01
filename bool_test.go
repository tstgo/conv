/*
 * @Author: leoking
 * @Date: 2022-09-01 19:52:16
 * @LastEditTime: 2022-09-01 20:39:37
 * @LastEditors: leoking
 * @Description:
 */
package conv

import "testing"

func TestBool(t *testing.T) {
	t.Log(Bool("1"))
	t.Log(Bool("0"))
	t.Log(Bool("是"))
	t.Log(Bool("Yes"))
	t.Log(Bool("y"))
	t.Log(Bool("no"))
	t.Log(Bool("True"))
	t.Log(Bool("no"))
}
