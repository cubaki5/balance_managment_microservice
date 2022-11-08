// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: report_accounting.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createReportRow = `-- name: CreateReportRow :exec
INSERT INTO report_accounting (
    service_id, date, income
) VALUE (
    ?, ?, ?
    )
`

type CreateReportRowParams struct {
	ServiceID int64
	Date      sql.NullTime
	Income    int32
}

func (q *Queries) CreateReportRow(ctx context.Context, arg CreateReportRowParams) error {
	_, err := q.db.ExecContext(ctx, createReportRow, arg.ServiceID, arg.Date, arg.Income)
	return err
}

const getMonthReport = `-- name: GetMonthReport :many
SELECT service_id, income
FROM report_accounting
WHERE YEAR(date) = ? AND MONTH(date) = ?
`

type GetMonthReportParams struct {
	Date   sql.NullTime
	Date_2 sql.NullTime
}

type GetMonthReportRow struct {
	ServiceID int64
	Income    int32
}

func (q *Queries) GetMonthReport(ctx context.Context, arg GetMonthReportParams) ([]GetMonthReportRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthReport, arg.Date, arg.Date_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMonthReportRow
	for rows.Next() {
		var i GetMonthReportRow
		if err := rows.Scan(&i.ServiceID, &i.Income); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
