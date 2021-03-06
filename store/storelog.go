/*
Copyright 2017 by GoSpider author.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package store

import (
	"github.com/op/go-logging"
	"os"
)

// 全局日志
var Logger = logging.MustGetLogger("GoStore")

// 格式
var format = logging.MustStringFormatter(
	"%{color}%{time:2006-01-02 15:04:05.000} %{longpkg}:%{longfunc} [%{level:.5s}]:%{color:reset} %{message}",
)

// level name you can refer
var LevelNames = []string{
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
}

// init log record
// 初始化日志
func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
	logging.SetLevel(logging.INFO, "GoStore")
}

// 设置日志级别
// set log level
func SetLogLevel(level string) {
	lvl, _ := logging.LogLevel(level)
	logging.SetLevel(lvl, "GoStore")
}

// 返回全局对象
// return global log
func Log() *logging.Logger {
	return Logger
}
