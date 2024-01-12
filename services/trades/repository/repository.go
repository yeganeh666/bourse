package repository

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mabna/services/trades/configs"
	"mabna/shared/gormext"
)

type Repository struct {
	log *logrus.Logger
	DB  *gorm.DB
}

func NewRepository(log *logrus.Logger, conf *configs.Configs) (*Repository, error) {
	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		conf.Postgres.User, conf.Postgres.Pass, conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.DB)
	db, err := gormext.Open(dsn)
	if err != nil {
		log.WithError(err).Fatal("can not load repository configs")
		return nil, err
	}
	if err = db.Transaction(func(tx *gorm.DB) error {
		if err = gormext.EnableExtensions(tx, gormext.UUIDExtension); err != nil {
			return err
		}
		if err = tx.AutoMigrate(
			new(Instrument),
			new(Trade),
		); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.WithError(err).Fatal("can not migration database")
		return nil, err
	}
	return &Repository{log: log, DB: db}, nil
}
