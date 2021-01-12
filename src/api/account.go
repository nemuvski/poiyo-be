package api

import (
	"net/http"
	customMiddleware "poiyo-be/src/middleware"
	"poiyo-be/src/model"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func PostAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		m := new(model.Account)
		c.Bind(&m)

		tx := c.Get(customMiddleware.TxKey).(*gorm.DB)
		account := model.Account{}

		// アカウントデータが存在するか（登録済みか）を確認.
		result := tx.Where("service_type = ? AND service_id = ?", m.ServiceType, m.ServiceId).First(&account)
		if result.RowsAffected > 0 {
			// もしメールアドレスが変わっていた場合は更新する.
			if account.Email != m.Email {
				tx.Model(&account).UpdateColumn("email", m.Email)
			}
			return c.JSON(http.StatusOK, account)
		}

		account.ServiceType = m.ServiceType
		account.ServiceId = m.ServiceId
		account.Email = m.Email
		tx.Create(&account)

		return c.JSON(http.StatusCreated, account)
	}
}
