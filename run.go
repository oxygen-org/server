/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 22:07:33
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"

	conf "github.com/oxygen-org/server/config"
)

func main() {
	// defer db.DB.Close()

	hostPORT := conf.CONFIG.Get("HOST_PORT").MustInt()
	r := setupRouter()
	r.Logger.Info(" Listen and Server in ", hostPORT)
	r.Logger.Fatal(r.Start(fmt.Sprintf(":%v", hostPORT)))
}
