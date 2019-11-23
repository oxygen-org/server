/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 22:28:53
 * @LastEditors: Please set LastEditors
 */
package apiv1

import (
	"net/http"
	"github.com/oxygen-org/server/db"

	"github.com/labstack/echo"
)

// Version get current version
func Version(e echo.Context) error{
	var version db.Version
	db.DB.Last(&version)
	msg := map[string]interface{}{"last_version": version.VersionNo, "release_note": version.ReleaseNote}
	return e.JSON(http.StatusOK, map[string]interface{}{"result": 0, "message": msg})
}
