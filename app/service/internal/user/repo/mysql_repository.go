package repo

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"service/internal/user/domain/user"
)

// MySQLUserRepo ...
type MySQLUserRepo struct {
	db sq.StdSqlCtx
}

var (
	// UserTableName is table name for User
	UserTableName = "user"
	// UserTable is columns for user entity
	UserTable = struct {
		Id     string
		Number string
	}{
		Id:     "id",
		Number: "number",
	}
)

// NewMySQLUserRepo ...
func NewMySQLUserRepo(db *sql.DB) user.Repository {
	return &MySQLUserRepo{
		db: db,
	}
}

// Create ...
func (repo MySQLUserRepo) Create(ctx context.Context, ent *user.User) (err error) {
	query, args, err := sq.Insert(UserTableName).
		Columns(UserTable.Number).
		Values(ent.Number).ToSql()
	if err != nil {
		return
	}
	_, err = repo.db.ExecContext(ctx, query, args...)
	return
}

// Find ...
func (repo MySQLUserRepo) Find(ctx context.Context, id int64) (*user.User, error) {
	builder := sq.Select(
		UserTable.Id,
		UserTable.Number,
	).From(UserTableName).Where(sq.Eq{UserTable.Id: id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := repo.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	list := make([]*user.User, 0)
	for rows.Next() {
		ent := new(user.User)
		if err = rows.Scan(
			&ent.ID,
			&ent.Number,
		); err != nil {
			return nil, err
		}
		list = append(list, ent)
	}
	return list[0], err
}
