//Copyright 2017-2019 Baidu Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http: //www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package models

import "rasp-cloud/tools"


func TaskCleanUpHosts() {
	// init crontab
	timer1 := &tools.CronTabTime{
		Hour: 0,
		Min:  0,
		Sec:  0,
		Nsec: 0,
	}
	tools.CronTabTimer(CleanOfflineHosts, timer1, 1)
}