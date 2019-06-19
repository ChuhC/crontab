package scheduler

import (
	"fmt"
	"time"

	"github.com/ChuhC/crontab/common"
	"github.com/astaxie/beego/logs"
)

type Scheduler struct {
	JobEventChan    chan *common.JobEvent
	jobPlanTable    map[string]*common.JobSchedulePlan
	jobExecutingJob map[string]*JobExecuteInfo
}

func NewScheduler() *Scheduler {
	sch := &Scheduler{
		JobEventChan: make(chan *common.JobEvent, 1000),
		jobPlanTable: make(map[string]*common.JobSchedulePlan),
	}

	go sch.ScheduleLoop()
	return sch
}

func (s *Scheduler) TryStartJob(jobPlan *common.JobSchedulePlan) {

}

func (s *Scheduler) TrySchedule() time.Duration {

	fmt.Println(" in try schedule", len(s.jobPlanTable))
	//如果任务表为空，随便睡眠多久
	if len(s.jobPlanTable) == 0 {
		return 1 * time.Second
	}

	//遍历所有任务
	var nearTime *time.Time
	for _, jobPlan := range s.jobPlanTable {
		if jobPlan.NextTime.Before(time.Now()) || jobPlan.NextTime.Equal(time.Now()) {
			// TODO： 尝试执行任务
			fmt.Println("执行任务", jobPlan.Job.Name)
			jobPlan.NextTime = jobPlan.Expr.Next(time.Now())
		}

		//统计最近一下要过期的时间
		if nearTime == nil || jobPlan.NextTime.Before(*nearTime) {
			nearTime = &jobPlan.NextTime
		}
	}

	scheduleAfter := (*nearTime).Sub(time.Now())
	return scheduleAfter
}

func (s *Scheduler) ScheduleLoop() {

	fmt.Println("in schedule loop")
	scheduleAfter := s.TrySchedule()

	timer := time.NewTimer(scheduleAfter)
	for {
		select {
		case event := <-s.JobEventChan:
			switch event.EventType {
			case common.JOB_EVENT_SAVE:
				// TODO: do save
				fmt.Println("recv event save. ", event.Job.Name)
				jobPlan, err := common.NewJobSchedulePlan(event.Job)
				if err != nil {
					logs.Error(err.Error())
					break
				}
				fmt.Println(jobPlan.Job.Name)
				fmt.Println(jobPlan.Job.CronExpr)

				s.jobPlanTable[event.Job.Name] = jobPlan

			case common.JOB_EVENT_DELETE:

				fmt.Println("recv event delete!")
				// TODO: do delete
				_, existed := s.jobPlanTable[event.Job.Name]
				if existed {
					delete(s.jobPlanTable, event.Job.Name)
				}

			}
		case <-timer.C:
		}
		scheduleAfter = s.TrySchedule()
		timer.Reset(scheduleAfter)

	}
}

func (s *Scheduler) PushScheduler(event *common.JobEvent) {
	s.JobEventChan <- event

}

func (s *Scheduler) handleScheduler(event *common.JobEvent) {

}
