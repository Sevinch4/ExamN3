package postgres

import (
	"context"
	"exam3/config"
	"exam3/pkg/logger"
	"exam3/storage"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
)

type Store struct {
	db  *pgxpool.Pool
	cfg config.Config
	log logger.ILogger
}

func New(ctx context.Context, cfg config.Config, log logger.ILogger) (storage.IStorage, error) {
	url := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=disable`,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Error("error is while parsing config", logger.Error(err))
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Error("error is while new with config", logger.Error(err))
		return nil, err
	}

	// migration up
	m, err := migrate.New("file://migrations/postgres/", url)
	if err != nil {
		log.Error("error is while migrating", logger.Error(err))
		return Store{}, err
	}

	if err = m.Up(); err != nil {
		log.Warning("migration up", logger.Error(err))
		if !strings.Contains(err.Error(), "no change") {
			version, dirty, err := m.Version()
			log.Info("version and dirty", logger.Any("version", version), logger.Any("dirty", dirty))
			if err != nil {
				log.Error("err in checking version and dirty", logger.Error(err))
				return nil, err
			}

			if dirty {
				version--
				if err = m.Force(int(version)); err != nil {
					log.Error("ERR in making force", logger.Error(err))
					return nil, err
				}
			}
			log.Warning("WARNING in migrating", logger.Error(err))
			return nil, err
		}
	}

	return Store{
		db:  pool,
		cfg: cfg,
	}, nil
}

func (s Store) Close() {
	s.db.Close()
}

func (s Store) Book() storage.IBook {
	return NewBookRepo(s.db, s.log)
}
