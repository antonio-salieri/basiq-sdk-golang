package Services

import (
	"errors"
	"strings"
	"time"
)

type JobStep struct {
	Title  string                 `json:"title"`
	Status string                 `json:"status"`
	Result map[string]interface{} `json:"result"`
}

type JobLinks struct {
	Self   string `json:"self"`
	Source string `json:"source"`
}

type Job struct {
	Id      string    `json:"id"`
	Steps   []JobStep `json:"steps"`
	Created string    `json:"created"`
	Updated string    `json:"updated"`
	Links   JobLinks  `json:"links"`
	Service *ConnectionService
}

func (j *Job) GetConnectionId() string {
	if j.Links.Source == "" {
		return ""
	}

	return j.Links.Source[strings.LastIndex(j.Links.Source, "/")+1:]
}

func (j *Job) WaitForCredentials(interval int64, timeout int64) (Connection, error) {
	var data Connection
	intervalDuration := time.Duration(interval) * time.Millisecond
	end := time.Now().Add(time.Duration(timeout) * time.Second)

	time.Sleep(intervalDuration)

	for {
		current := time.Now()
		if current.After(end) {
			return data, errors.New("Timeout")
		}

		job, err := j.Service.GetJob(j.Id)
		if err != nil {
			return data, nil
		}

		if job.Steps[0].Status == "failed" {
			return data, errors.New("Credentials failure")
		} else if job.Steps[0].Status == "success" {
			return j.Service.GetConnection(job.GetConnectionId())
		}

		time.Sleep(intervalDuration)
	}

}
