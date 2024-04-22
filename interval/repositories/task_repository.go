package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-jwt/interval/models"

	"github.com/doug-martin/goqu/v9"
)

type TaskRepository interface {
	Create(ctx context.Context, task *models.Task) error
	GetById(ctx context.Context, id int) (*models.Task, error)
	Update(ctx context.Context, task *models.Task) error
	Delete(ctx context.Context, id int) error
	GetByUserID(ctx context.Context, userID int) ([]*models.Task, error)
}

type taskRepository struct {
	db *DB
}

func NewTaskRepository(db *DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(ctx context.Context, task *models.Task) error {
	query := goqu.Insert("task").
		Rows(goqu.Record{
			"name":          task.Name,
			"statusid":      task.Status.ID,
			"projectid":     task.Project.ID,
			"assignedforid": task.AssignedFor.ID,
			"datecreated":   task.DateCreated,
		}).
		Returning("id")

	insertSQL, args, err := query.ToSQL()
	fmt.Println(insertSQL, args)
	if err != nil {
		return err
	}

	err = r.db.Conn.QueryRowContext(ctx, insertSQL, args...).Scan(&task.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) GetById(ctx context.Context, taskID int) (*models.Task, error) {
	if taskID < 1 {
		return nil, errors.New("record not found")
	}

	query := goqu.From("task").
		Where(goqu.Ex{"id": taskID})

	insertSql, args, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	task := models.Task{
		Status:      &models.Status{},
		Project:     &models.Project{},
		AssignedFor: &models.User{},
	}
	err = r.db.Conn.QueryRowContext(ctx, insertSql, args...).Scan(
		&task.ID,
		&task.Name,
		&task.Status.ID,
		&task.Project.ID,
		&task.AssignedFor.ID,
		&task.DateCreated,
	)
	if err != nil {
		if errors.Is(err, errors.New("no rows in result set")) {
			return nil, nil
		}
		return nil, fmt.Errorf("error scanning row: %v", err)
	}

	return &task, nil
}

func (r *taskRepository) Update(ctx context.Context, task *models.Task) error {
	if _, err := r.GetById(ctx, int(task.ID)); err != nil {
		return err
	}

	query := goqu.Update("task").
		Set(goqu.Record{
			"name":          task.Name,
			"statusid":      task.Status.ID,
			"projectid":     task.Project.ID,
			"assignedforid": task.AssignedFor.ID,
			"datecreated":   task.DateCreated,
		}).
		Where(goqu.Ex{"id": task.ID})

	updateSQL, args, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = r.db.Conn.ExecContext(ctx, updateSQL, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) Delete(ctx context.Context, taskID int) error {
	if _, err := r.GetById(ctx, int(taskID)); err != nil {
		return err
	}

	query := goqu.Delete("task").
		Where(goqu.Ex{"id": taskID})

	deleteSQL, args, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = r.db.Conn.ExecContext(ctx, deleteSQL, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) GetByUserID(ctx context.Context, userID int) ([]*models.Task, error) {
	query := goqu.From("task").
		Join(goqu.I("user").As("u"), goqu.On(goqu.Ex{"task.assignedforid": goqu.I("u.id")})).
		Select(goqu.I("task.*")).
		Where(goqu.I("u.id").Eq(userID))

	insertSql, args, err := query.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	rows, err := r.db.Conn.QueryContext(ctx, insertSql, args...)
	if err != nil {
		return nil, err
	}

	return r.fetchTasks(rows)
}

func (r *taskRepository) fetchTasks(rows *sql.Rows) ([]*models.Task, error) {
	var tasks []*models.Task

	for rows.Next() {
		task := &models.Task{
			Status:      &models.Status{},
			Project:     &models.Project{},
			AssignedFor: &models.User{},
		}

		err := rows.Scan(
			&task.ID,
			&task.Name,
			&task.Status.ID,
			&task.Project.ID,
			&task.AssignedFor.ID,
			&task.DateCreated,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch row: %v", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occured while fetching rows: %v", err)
	}

	return tasks, nil
}
