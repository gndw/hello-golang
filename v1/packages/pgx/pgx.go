package pgx

import (
	"context"
	"hello-golang/v1/functions/builder/fxb"
	"hello-golang/v1/services/db"
	"log"

	"github.com/jackc/pgx/v4"
	"go.uber.org/fx"
)

type Instance struct {

}

func GetDb(lc fx.Lifecycle, oe *fxb.OnError) (db.Interface, error) {
	ins := &Instance{}

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/database_name")
	if err != nil {
		return nil, err
	}

	oe.Append(func ()  {
		log.Println("close db conn")
		conn.Close(context.Background())
	})

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return conn.Close(ctx)
		},
	})

	panic("yo")
	

	return ins, nil
}