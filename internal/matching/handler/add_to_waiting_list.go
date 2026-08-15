package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sahar-mirtalebi/quiz-battle/internal/auth/middleware"
	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/dto"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/httpmessage"
)

func (h Handler) addToWaitingList(c echo.Context) error {
	userID, err := middleware.UserID(c)
	if err != nil {
		return err
	}

	var req dto.AddToWaitingListInput
	err = c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	req.UserID = userID

	fieldError, err := h.matchingValidator.ValidateAddToWaitingListInput(req)
	if err != nil {
		msg, code := httpmessage.ErrorCodeAndMsg(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldError,
		})
	}

	resp, err := h.matchingSvc.AddToWaitingList(c.Request().Context(), req)
	if err != nil {
		msg, code := httpmessage.ErrorCodeAndMsg(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
