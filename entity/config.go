package entity

type ConfigCronJob struct {
	LogFile *string   `json:"log_file"`
	Scripts []CronJob `json:"scripts"`
}

type CronJob struct {
	Spec                 string   `json:"spec"`
	Command              []string `json:"command"`
	TimeOut              int64    `json:"timeout"`
	MultiProcessingLimit int64    `json:"multi_processing_limit"`
	IsRunning            bool
	MultiProcessingCount int64
}
