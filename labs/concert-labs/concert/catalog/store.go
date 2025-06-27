package catalog

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

var (
	concertBatchKeyFmt = "concert_batch:%s"
)

type Store struct {
	db      *sqlx.DB
	dbCache *sq.StmtCache
	redis   redis.UniversalClient
}

func NewStore(db *sqlx.DB, redis redis.UniversalClient) *Store {
	return &Store{
		db:      db,
		dbCache: sq.NewStmtCache(db),
		redis:   redis,
	}
}

func (s *Store) Clear() error {
	return s.dbCache.Clear()
}

func (s *Store) FindAllConcert(ctx context.Context, opts ...ListOption) ([]Concert, string, error) {
	options := &ListOptions{Limit: 10}
	for _, o := range opts {
		o(options)
	}

	nextPage := pageToken{page: options.Page + 1}.encode()
	var concerts []Concert

	sb := sq.StatementBuilder.RunWith(s.dbCache)
	selectConcerts := sb.
		Select("c.id", "c.name", "c.slug", "c.description", "c.status", "c.announced_at").
		From("concerts c").
		Where(sq.Eq{"c.deleted_at": nil, "c.status": ConcertStatusPublished}).
		OrderBy("c.announced_at DESC").
		Offset(uint64(options.GetOffset())).
		Limit(uint64(options.Limit)).
		PlaceholderFormat(sq.Dollar)

	rows, err := selectConcerts.QueryContext(ctx)
	if err != nil {
		return nil, "", err
	}

	for rows.Next() {
		var c Concert
		if err := rows.Scan(&c.ID, &c.Name, &c.Slug, &c.Description, &c.Status, &c.AnnouncedAt); err != nil {
			return nil, "", err
		}
		if options.Preload {
			batchs, _, err := s.FindAllBatchesByConcertID(ctx, c.ID.String())
			if err != nil {
				return nil, "", err
			}
			c.Batches = batchs
		}
		concerts = append(concerts, c)
	}
	return concerts, nextPage, nil
}

func (s *Store) FindConcertByID(ctx context.Context, id string) (*Concert, error) {
	c := Concert{}
	sb := sq.StatementBuilder.RunWith(s.dbCache)
	selectConcert := sb.
		Select("c.id", "c.name", "c.slug", "c.description", "c.status", "c.announced_at").
		From("concerts c").
		Where(sq.Eq{"c.deleted_at": nil, "c.id": id, "c.status": ConcertStatusPublished}).
		PlaceholderFormat(sq.Dollar)

	if err := selectConcert.QueryRowContext(ctx).Scan(&c.ID, &c.Name, &c.Slug, &c.Description, &c.Status, &c.AnnouncedAt); err != nil {
		return nil, err
	}

	selectBatches := sb.
		Select("id", "name", "max_capacity", "available_tickets", "price", "currency", "start_time", "end_time", "version").
		From("concert_batchs").
		Where(sq.Eq{"concert_id": c.ID, "deleted_at": nil, "status": BatchStatusPublished}).
		PlaceholderFormat(sq.Dollar)

	rows, err := selectBatches.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	var batchs []Batch
	for rows.Next() {
		var s Batch
		if err := rows.Scan(&s.ID, &s.Name, &s.MaxCapacity, &s.AvailableTickets, &s.Price, &s.Currency, &s.StartTime, &s.EndTime, &s.Version); err != nil {
			return nil, err
		}
		batchs = append(batchs, s)
	}
	c.Batches = batchs
	return &c, nil
}

func (s *Store) CreateConcert(ctx context.Context, concert *Concert) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	sb := sq.StatementBuilder.RunWith(tx)
	insertConcert := sb.
		Insert("concerts").
		Columns("id", "name", "slug", "description", "status", "announced_at", "created_at", "updated_at").
		Values(concert.ID, concert.Name, concert.Slug, concert.Description, concert.Status, concert.AnnouncedAt, concert.CreatedAt, concert.UpdatedAt).
		PlaceholderFormat(sq.Dollar)

	insertBatches := sb.
		Insert("concert_batchs").
		Columns("id", "name", "max_capacity", "available_tickets", "price", "currency", "start_time", "end_time", "concert_id", "created_at", "updated_at", "status").
		PlaceholderFormat(sq.Dollar)
	for _, s := range concert.Batches {
		insertBatches = insertBatches.Values(s.ID, s.Name, s.MaxCapacity, s.AvailableTickets, s.Price, s.Currency, s.StartTime, s.EndTime, concert.ID, s.CreatedAt, s.UpdatedAt, s.Status)
	}

	if _, err := insertConcert.ExecContext(ctx); err != nil {
		tx.Rollback()
		return err
	}
	if _, err := insertBatches.ExecContext(ctx); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (s *Store) FindAllBatchesByConcertID(ctx context.Context, concertID string, opts ...ListOption) ([]Batch, string, error) {
	options := &ListOptions{Limit: 10}
	for _, o := range opts {
		o(options)
	}

	nextPage := pageToken{page: options.Page + 1}.encode()
	var batchs []Batch
	sb := sq.StatementBuilder.RunWith(s.dbCache)
	selectBatches := sb.
		Select("id", "name", "max_capacity", "available_tickets", "price", "currency", "start_time", "end_time", "version").
		From("concert_batchs").
		Where(sq.Eq{"concert_id": concertID, "deleted_at": nil, "status": 1}).
		OrderBy("created_at DESC").
		Offset(uint64(options.GetOffset())).
		Limit(uint64(options.Limit)).
		PlaceholderFormat(sq.Dollar)

	rows, err := selectBatches.QueryContext(ctx)
	if err != nil {
		return nil, "", err
	}
	for rows.Next() {
		var s Batch
		if err := rows.Scan(&s.ID, &s.Name, &s.MaxCapacity, &s.AvailableTickets, &s.Price, &s.Currency, &s.StartTime, &s.EndTime, &s.Version); err != nil {
			return nil, "", err
		}
		batchs = append(batchs, s)
	}
	return batchs, nextPage, nil
}
