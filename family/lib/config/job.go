package config

type JobConfig struct {
	Topic         string `toml:"topic"`
	NsqHost       string `toml:"nsq_host"`
	JobCenterHost string `toml:"job_center_host"`
}
