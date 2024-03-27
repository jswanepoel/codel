// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: tasks.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (
  type,
  status,
  args,
  results,
  flow_id,
  message
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, created_at, updated_at, type, status, args, results, flow_id, message
`

type CreateTaskParams struct {
	Type    pgtype.Text
	Status  pgtype.Text
	Args    []byte
	Results pgtype.Text
	FlowID  pgtype.Int8
	Message pgtype.Text
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRow(ctx, createTask,
		arg.Type,
		arg.Status,
		arg.Args,
		arg.Results,
		arg.FlowID,
		arg.Message,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Type,
		&i.Status,
		&i.Args,
		&i.Results,
		&i.FlowID,
		&i.Message,
	)
	return i, err
}

const readTasksByFlowId = `-- name: ReadTasksByFlowId :many
SELECT id, created_at, updated_at, type, status, args, results, flow_id, message FROM tasks
WHERE flow_id = $1
`

func (q *Queries) ReadTasksByFlowId(ctx context.Context, flowID pgtype.Int8) ([]Task, error) {
	rows, err := q.db.Query(ctx, readTasksByFlowId, flowID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Type,
			&i.Status,
			&i.Args,
			&i.Results,
			&i.FlowID,
			&i.Message,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTaskResults = `-- name: UpdateTaskResults :one
UPDATE tasks
SET results = $1
WHERE id = $2
RETURNING id, created_at, updated_at, type, status, args, results, flow_id, message
`

type UpdateTaskResultsParams struct {
	Results pgtype.Text
	ID      int64
}

func (q *Queries) UpdateTaskResults(ctx context.Context, arg UpdateTaskResultsParams) (Task, error) {
	row := q.db.QueryRow(ctx, updateTaskResults, arg.Results, arg.ID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Type,
		&i.Status,
		&i.Args,
		&i.Results,
		&i.FlowID,
		&i.Message,
	)
	return i, err
}

const updateTaskStatus = `-- name: UpdateTaskStatus :one
UPDATE tasks
SET status = $1
WHERE id = $2
RETURNING id, created_at, updated_at, type, status, args, results, flow_id, message
`

type UpdateTaskStatusParams struct {
	Status pgtype.Text
	ID     int64
}

func (q *Queries) UpdateTaskStatus(ctx context.Context, arg UpdateTaskStatusParams) (Task, error) {
	row := q.db.QueryRow(ctx, updateTaskStatus, arg.Status, arg.ID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Type,
		&i.Status,
		&i.Args,
		&i.Results,
		&i.FlowID,
		&i.Message,
	)
	return i, err
}