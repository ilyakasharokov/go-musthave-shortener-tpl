package middlewares

import (
	"context"
	"fmt"
	helpers "ilyakasharokov/internal/app/encryptor"
	"net/http"

	"github.com/google/uuid"
)

// CookieUserIDName define cookie name for uuid
const CookieUserIDName = "user_id"

// ContextType set context name for user id
type ContextType string

var UserIDCtxName ContextType = "ctxUserId"

func CookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate new uuid
		userID := uuid.New().String()
		// Check if set cookie
		if cookieUserID, err := r.Cookie(CookieUserIDName); err == nil {
			// logger.Info("cookieUserId", zap.String("cookieUserId", cookieUserID.Value))
			_ = helpers.Decode(cookieUserID.Value, &userID)
		}
		// Generate hash from userId
		encoded, err := helpers.Encode(userID)
		// logger.Info("User ID", zap.String("ID", userID))
		// logger.Info("User encoded", zap.String("Encoded", encoded))
		if err == nil {
			cookie := &http.Cookie{
				Name:  CookieUserIDName,
				Value: encoded,
				Path:  "/",
			}
			http.SetCookie(w, cookie)
		} else {
			fmt.Println(err)
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), UserIDCtxName, userID)))
	})
}
