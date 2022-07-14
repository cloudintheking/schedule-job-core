/**
 * @Author: cloudintheking
 * @Description:
 * @File: job
 * @Version: 1.0.0
 * @Date: 2022/7/13 15:24
 */
package cik_job

import (
	"encoding/json"
	"github.com/robfig/cron/v3"
)

type JobRunner interface {
	Call()           //执行函数
	GetCron() string //cron表达式
}

func JobSerialize(runner JobRunner) ([]byte, error) {
	return json.Marshal(runner)
}

func JobDeserialize(data []byte, target JobRunner) error {
	return json.Unmarshal(data, target)
}

func AddRunner(c *cron.Cron, runner JobRunner) (cron.EntryID, error) {
	if entryId, err := c.AddFunc(runner.GetCron(), runner.Call); err != nil {
		return entryId, err
	} else {
		return entryId, nil
	}
}
