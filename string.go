/*
 * @Author: leoking
 * @Date: 2022-09-01 16:46:48
 * @LastEditTime: 2022-09-01 16:47:58
 * @LastEditors: leoking
 * @Description:
 */
package conv

import "fmt"

/*
*
  - @description:
    任意类型转string
  - @return {*}
*/
func String[T any](v T) string {
	return fmt.Sprintf("%v", v)
}
