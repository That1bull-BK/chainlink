package web

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/smartcontractkit/chainlink/core/auth"
	"github.com/smartcontractkit/chainlink/core/services/chainlink"
	clsession "github.com/smartcontractkit/chainlink/core/sessions"
	"github.com/smartcontractkit/chainlink/core/utils"
	webauth "github.com/smartcontractkit/chainlink/core/web/auth"
	"github.com/smartcontractkit/chainlink/core/web/presenters"
)

// UserController manages the current Session's User.
type UserController struct {
	App chainlink.Application
}

// UpdatePasswordRequest defines the request to set a new password for the
// current session's User.
type UpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

// Create creates a new API user with provided context arugments.
func (c *UserController) Create(ctx *gin.Context) {
	// TODO: Andrew - STUB
	return
}

// Update changes an API user's role.
func (c *UserController) Update(ctx *gin.Context) {
	// TODO: Andrew - STUB
	return
}

// Delete deletes an API user and any sessions by email
func (c *UserController) Delete(ctx *gin.Context) {
	// TODO: Andrew - STUB
	return
}

// UpdatePassword changes the password for the current User.
func (c *UserController) UpdatePassword(ctx *gin.Context) {
	var request UpdatePasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		jsonAPIError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	sessionUser, ok := webauth.GetAuthenticatedUser(ctx)
	if !ok {
		jsonAPIError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to obtain current user from context"))
		return
	}
	user, err := c.App.SessionORM().FindUser(sessionUser.Email)
	if err != nil {
		jsonAPIError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to obtain current user record: %+v", err))
		return
	}
	if !utils.CheckPasswordHash(request.OldPassword, user.HashedPassword) {
		jsonAPIError(ctx, http.StatusConflict, errors.New("old password does not match"))
		return
	}
	if err := c.updateUserPassword(ctx, &user, request.NewPassword); err != nil {
		jsonAPIError(ctx, http.StatusInternalServerError, err)
		return
	}

	jsonAPIResponse(ctx, presenters.NewUserResource(user), "user")
}

// NewAPIToken generates a new API token for a user overwriting any pre-existing one set.
func (c *UserController) NewAPIToken(ctx *gin.Context) {
	var request clsession.ChangeAuthTokenRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		jsonAPIError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	sessionUser, ok := webauth.GetAuthenticatedUser(ctx)
	if !ok {
		jsonAPIError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to obtain current user from context"))
		return
	}
	user, err := c.App.SessionORM().FindUser(sessionUser.Email)
	if err != nil {
		jsonAPIError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to obtain current user record: %+v", err))
		return
	}
	if !utils.CheckPasswordHash(request.Password, user.HashedPassword) {
		jsonAPIError(ctx, http.StatusUnauthorized, errors.New("incorrect password"))
		return
	}
	newToken := auth.NewToken()
	if err := c.App.SessionORM().SetAuthToken(&user, newToken); err != nil {
		jsonAPIError(ctx, http.StatusInternalServerError, err)
		return
	}

	jsonAPIResponseWithStatus(ctx, newToken, "auth_token", http.StatusCreated)
}

// DeleteAPIToken deletes and disables a user's API token.
func (c *UserController) DeleteAPIToken(ctx *gin.Context) {
	var request clsession.ChangeAuthTokenRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		jsonAPIError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	sessionUser, ok := webauth.GetAuthenticatedUser(ctx)
	if !ok {
		jsonAPIError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to obtain current user from context"))
		return
	}
	user, err := c.App.SessionORM().FindUser(sessionUser.Email)
	if err != nil {
		jsonAPIError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to obtain current user record: %+v", err))
		return
	}
	if !utils.CheckPasswordHash(request.Password, user.HashedPassword) {
		jsonAPIError(ctx, http.StatusUnauthorized, errors.New("incorrect password"))
		return
	}
	if err := c.App.SessionORM().DeleteAuthToken(&user); err != nil {
		jsonAPIError(ctx, http.StatusInternalServerError, err)
		return
	}
	{
		jsonAPIResponseWithStatus(ctx, nil, "auth_token", http.StatusNoContent)
	}
}

func getCurrentSessionID(ctx *gin.Context) (string, error) {
	session := sessions.Default(ctx)
	sessionID, ok := session.Get(webauth.SessionIDKey).(string)
	if !ok {
		return "", errors.New("unable to get current session ID")
	}
	return sessionID, nil
}

func (c *UserController) updateUserPassword(ctx *gin.Context, user *clsession.User, newPassword string) error {
	sessionID, err := getCurrentSessionID(ctx)
	if err != nil {
		return err
	}
	orm := c.App.SessionORM()
	if err := orm.ClearNonCurrentSessions(sessionID); err != nil {
		return fmt.Errorf("failed to clear non current user sessions: %+v", err)
	}
	if err := orm.SetPassword(user, newPassword); err != nil {
		return fmt.Errorf("failed to update current user password: %+v", err)
	}
	return nil
}
