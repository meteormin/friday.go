package task

import (
	"errors"
	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
	"github.com/meteormin/friday.go/internal/core"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Job struct {
	gorm.Model
	JobID      string       `json:"jobID" gorm:"not null;uniqueIndex"`
	Name       string       `json:"name" gorm:"not null"`
	Tags       []JobTag     `json:"tags" gorm:"foreignKey:JobID"`
	JobCounter []JobCounter `json:"jobCounter" gorm:"foreignKey:JobID"`
	JobTiming  []JobTiming  `json:"jobTiming" gorm:"foreignKey:JobID"`
}

type JobTiming struct {
	gorm.Model
	JobID   string    `json:"jobID" gorm:"not null;uniqueIndex"`
	StartAt time.Time `json:"startAt" gorm:"not null"`
	EndAt   time.Time `json:"endAt" gorm:"not"`
}

type JobCounter struct {
	gorm.Model
	JobID  string `json:"jobID" gorm:"not null;uniqueIndex"`
	Status string `json:"status" gorm:"not null"`
	Error  string `json:"error" gorm:"not null"`
}

type JobTag struct {
	gorm.Model
	JobID uint   `json:"jobID" gorm:"not null;uniqueIndex"`
	Tag   string `json:"tag" gorm:"not null"`
}

func NewJobTags(tags []string) []JobTag {
	t := make([]JobTag, 0, len(tags))
	for _, tag := range tags {
		t = append(t, JobTag{
			Tag: tag,
		})
	}

	return t
}

type JobRepository interface {
	All() []Job
	FindByJobID(jobID uuid.UUID) (*Job, error)
	IncrementJob(id uuid.UUID, name string, tags []string, status gocron.JobStatus)
	RecordJobTiming(startAt, endAt time.Time, id uuid.UUID, name string, tags []string)
	RecordJobTimingWithStatus(startAt, endAt time.Time, id uuid.UUID, name string, tags []string, status gocron.JobStatus, err error)
}

type JobRepositoryImpl struct {
	db *gorm.DB
}

func (j *JobRepositoryImpl) All() []Job {
	var jobs []Job
	err := j.db.Preload(clause.Associations).
		Find(&jobs).Error

	if err != nil {
		return make([]Job, 0)
	}

	return jobs
}

func (j *JobRepositoryImpl) Save(jobEntity Job) (*Job, error) {
	err := j.db.Save(&jobEntity).Error
	if err != nil {
		return nil, err
	}

	return &jobEntity, nil
}

func (j *JobRepositoryImpl) FindByJobID(jobID uuid.UUID) (*Job, error) {
	ent := Job{}
	err := j.db.Preload(clause.Associations).
		Where("job_id = ?", jobID).
		First(&ent).Error

	if err != nil {
		return nil, err
	}

	return &ent, nil
}

func (j *JobRepositoryImpl) FindByName(name string) (*Job, error) {
	ent := Job{}
	err := j.db.Where("name = ?", name).First(&ent).Error
	if err != nil {
		return nil, err
	}

	return &ent, nil
}

func (j *JobRepositoryImpl) IncrementJob(id uuid.UUID, name string, tags []string, status gocron.JobStatus) {
	job, err := j.FindByJobID(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		job = &Job{
			JobID: id.String(),
			Name:  name,
			Tags:  NewJobTags(tags),
		}
	} else if err != nil {
		core.GetLogger().Error(err)
		return
	}

	job.JobCounter = append(job.JobCounter, JobCounter{
		JobID:  id.String(),
		Status: string(status),
	})

	_, err = j.Save(*job)
	if err != nil {
		core.GetLogger().Error(err)
	}
}

func (j *JobRepositoryImpl) RecordJobTiming(startAt, endAt time.Time, id uuid.UUID, name string, tags []string) {
	job, err := j.FindByJobID(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		job = &Job{
			JobID: id.String(),
			Name:  name,
			Tags:  NewJobTags(tags),
		}
	} else if err != nil {
		core.GetLogger().Error(err)
		return
	}

	job.JobTiming = append(job.JobTiming, JobTiming{
		JobID:   id.String(),
		StartAt: startAt,
		EndAt:   endAt,
	})

	_, err = j.Save(*job)
	if err != nil {
		core.GetLogger().Error(err)
	}
}

func (j *JobRepositoryImpl) RecordJobTimingWithStatus(startAt, endAt time.Time, id uuid.UUID, name string, tags []string, status gocron.JobStatus, err error) {
	job, gormErr := j.FindByJobID(id)
	if gormErr != nil && errors.Is(gormErr, gorm.ErrRecordNotFound) {
		job = &Job{
			JobID:      id.String(),
			Name:       name,
			Tags:       NewJobTags(tags),
			JobCounter: make([]JobCounter, 0),
			JobTiming:  make([]JobTiming, 0),
		}

	}

	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	job.JobCounter = append(job.JobCounter, JobCounter{
		JobID:  id.String(),
		Status: string(status),
		Error:  errMsg,
	})

	core.GetLogger().Debugf("JobID: %s", job.JobID)
	core.GetLogger().Debugf("Job: %s", job.Name)
	core.GetLogger().Debugf("Status: %s", status)
	core.GetLogger().Debugf("Start: %s, End: %s", startAt, endAt)
	core.GetLogger().Debugf("Error: %s", err)

	_, err = j.Save(*job)
	if err != nil {
		core.GetLogger().Error(err)
	}
}

func NewJobRepository(db *gorm.DB) JobRepository {
	return &JobRepositoryImpl{db: db}
}
