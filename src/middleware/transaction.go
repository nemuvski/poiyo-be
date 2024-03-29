package middleware

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	TxKey = "Tx"
)

// Transaction トランザクション処理.
func Transaction(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tx := db.Begin()
			c.Set(TxKey, tx)
			err := next(c)
			if err != nil {
				tx.Rollback()
				return err
			}
			tx.Commit()

			return nil
		}
	}
}
