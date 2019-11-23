/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 21:47:55
 * @LastEditors: your name
 */
package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	sj "github.com/bitly/go-simplejson"
)

//CONFIG JSON配置
var CONFIG = &sj.Json{}

func init() {
	fmt.Println("config init...")
	dat, _ := ioutil.ReadFile("config.json")
	config, _ := sj.NewJson(dat)
	env := strings.ToTitle(os.Getenv("NITROGEN-ENV"))
	if env == "" {
		env = "Dev"
	}
	CONFIG = config.Get(env)
}
