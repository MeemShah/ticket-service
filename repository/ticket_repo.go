package repo

import (
	"context"
	"log/slog"
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
