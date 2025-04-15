package repo

import (
	"context"
	"fmt"
	"log/slog"
	"ticket-service/controller/utils"
	"ticket-service/dto"
	"ticket-service/logger"
	"ticket-service/ticket"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type TicketRepo interface {
	ticket.TicketRepo
}

type ticketRepo struct {
	table string
	db    *sqlx.DB
	psql  sq.StatementBuilderType
}

func NewTicketRepo(db *DB) TicketRepo {
	return &ticketRepo{
		table: "tickets",
		db:    db.Db,
		psql:  db.psql,
	}
}

func (r *ticketRepo) Get(ctx context.Context, ticketId string) (*dto.Ticket, error) {
	query, args, err := r.psql.
		Select("*").
		From(r.table).
		Where(sq.Eq{"id": ticketId}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var t dto.Ticket
	if err := r.db.GetContext(ctx, &t, query, args...); err != nil {
		slog.Error("failed to get ticket", logger.Extra(map[string]any{
			"err":      err.Error(),
			"ticketId": ticketId,
		}))
		return nil, err
	}

	return &t, nil
}

func (r *ticketRepo) Create(ctx context.Context, limit int, category string, price int, status string) error {
	query := r.psql.Insert(r.table).
		Columns("category", "price", "purchased_by", "status", "is_active")

	for i := 0; i < limit; i++ {
		query = query.Values(category, price, "", status, true)
	}

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *ticketRepo) Update(ctx context.Context, columns []string, values []any, ticketId string) error {
	if len(columns) == 0 || len(values) == 0 || len(columns) != len(values) {
		return fmt.Errorf("columns and values must be non-empty and of the same length")
	}

	query := r.psql.Update(r.table)
	for i, col := range columns {
		query = query.Set(col, values[i])
	}

	query = query.Where(sq.Eq{"id": ticketId})
	sqlStr, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *ticketRepo) GetTickets(ctx context.Context, params utils.PaginationParams) ([]*dto.Ticket, error) {
	query := r.psql.Select("*").
		From(r.table).
		Where(sq.Eq{"is_active": true})

	for field, values := range params.Filters {

		if len(values[0]) == 0 {
			continue
		}

		if field == "price" {
			query = query.Where(sq.LtOrEq{field: values[0]})
			continue
		}

		if len(values) == 1 {
			query = query.Where(sq.Eq{field: values[0]})
		} else if len(values) > 1 {
			query = query.Where(sq.Eq{field: values})
		}
	}

	if params.SortBy != "" {
		order := "ASC"
		if params.SortOrder == "desc" {
			order = "DESC"
		}
		query = query.OrderBy(fmt.Sprintf("%s %s", params.SortBy, order))
	}

	offset := (params.Page - 1) * params.Limit
	query = query.Limit(uint64(params.Limit)).Offset(uint64(offset))

	sqlStr, args, err := query.ToSql()
	if err != nil {
		slog.Error("failed to build query", logger.Extra(map[string]any{"error": err.Error()}))
		return nil, err
	}

	var tickets []*dto.Ticket
	if err := r.db.SelectContext(ctx, &tickets, sqlStr, args...); err != nil {
		slog.Error("failed to execute query", logger.Extra(map[string]any{
			"error": err.Error(),
			"query": sqlStr,
			"args":  args,
		}))
		return nil, err
	}

	return tickets, nil
}
