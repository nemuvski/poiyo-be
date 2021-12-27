package api

import (
	"net/http"
	customMiddleware "poiyo-be/src/middleware"
	"poiyo-be/src/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// PostAccount /authでユーザーを登録、または取得するAPI.
func PostAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		m := new(model.AuthPostRequest)
		c.Bind(m)

		if err := c.Validate(m); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

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

// DeleteAccount /accounts/:aidでアカウントをID指定で削除するAPI.
func DeleteAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		accountId := c.Param("aid")

		// パスパラメータについてバリデーション.
		params := model.DeleteAccountPathParameter{Aid: accountId}
		if err := c.Validate(params); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		tx := c.Get(customMiddleware.TxKey).(*gorm.DB)
		account := model.Account{}
		result := tx.Where("account_id = ?", accountId).Delete(&account)
		if result.RowsAffected == 0 {
			return c.NoContent(http.StatusNoContent)
		}
		return c.NoContent(http.StatusOK)
	}
}
