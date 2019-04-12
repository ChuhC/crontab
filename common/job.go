package app

type Job struct {
	Name        string    `form:"name"`
	CronExpr    string    `form:"crontab"`
	Command     string    `form:"command"`
}
