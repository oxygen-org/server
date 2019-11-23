/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 22:29:04
 * @LastEditors: Please set LastEditors
 */
package apiv2

import (
	"github.com/oxygen-org/server/api/apiv1"

	"github.com/labstack/echo"
)

// Version get current version
func Version(e echo.Context) error {
	return apiv1.Version(e)
}
