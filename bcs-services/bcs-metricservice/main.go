/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"bk-bcs/bcs-common/common"
	"bk-bcs/bcs-common/common/blog"
	"bk-bcs/bcs-common/common/check"
	"bk-bcs/bcs-services/bcs-metricservice/app"
	"bk-bcs/bcs-services/bcs-metricservice/app/config"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	conf := config.ParseConfig()

	blog.InitLogs(conf.LogConfig)
	defer blog.CloseLogs()

	//pid
	if err := common.SavePid(conf.ProcessConfig); err != nil {
		blog.Error("fail to save pid: err:%s", err.Error())
	}

	if err := app.Run(conf); nil != err {
		check.Fail(err.Error())
		blog.Fatal(err)
	}
}
