/*
 * @Author: leoking
 * @Date: 2022-09-01 16:46:48
 * @LastEditTime: 2022-09-01 16:48:26
 * @LastEditors: leoking
 * @Description:
 */
package conv

import "testing"

func TestString(t *testing.T) {
	t.Log(String(1))
	t.Log(String(1.2))
	t.Log(String("xxxxx"))
	t.Log(String(&struct {
		Name string
		Age  int
	}{"10", 20}))
}
