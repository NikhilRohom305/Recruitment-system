package middleware

import (
	"Recruitment-Managment-system/services"
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func AdminAuth(next http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract token from request header or cookie
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid authorization format. Expected: Bearer <token>", http.StatusUnauthorized)
			return
		}
		token := authHeader[len("Bearer "):]
		if token == "" {
			http.Error(w, "Token is empty", http.StatusUnauthorized)
			return
		}

		//	tokenWithBearer := AppendBearerToken(authHeader)
		// Validate token and extract user ID and role
		userID, role, err := validateToken(token)
		if err != nil {
			http.Error(w, "only admin user can access this API", http.StatusUnauthorized)
			return
		}

		// Set user ID and role in request context
		ctx := context.WithValue(r.Context(), "userID", userID)
		ctx = context.WithValue(ctx, "role", role)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	}
}

func validateToken(tokenString string) (userID string, role string, err error) {
	tokenObject, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(services.SecretKey), nil
	})
	if err != nil {
		return
	}
	claims, ok := tokenObject.Claims.(jwt.MapClaims)
	if !ok {
		return
	}
	role = string(claims["Role"].(string))
	userIDInterface, ok := claims["user_id"]
	if !ok {
		return
	}
	userID, ok = userIDInterface.(string)
	if !ok {
		return
	}
	//userID = string(claims["user_id"]).(string)
	if !ok || (role != "admin") {
		err = errors.New("user is not admin or superadmin")
		//logger.WithField("err", err.Error()).Error(" user is not admin or superadmin")
		return
	}
	return
}

func ApplicantAuth(next http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract token from request header or cookie
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid authorization format. Expected: Bearer <token>", http.StatusUnauthorized)
			return
		}
		token := authHeader[len("Bearer "):]
		if token == "" {
			http.Error(w, "Token is empty", http.StatusUnauthorized)
			return
		}

		//	tokenWithBearer := AppendBearerToken(authHeader)
		// Validate token and extract user ID and role
		userID, role, err := validateApplicant(token)
		if err != nil {
			http.Error(w, "only applicant can access this API", http.StatusUnauthorized)
			return
		}

		// Set user ID and role in request context
		ctx := context.WithValue(r.Context(), "userID", userID)
		ctx = context.WithValue(ctx, "role", role)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	}
}

func validateApplicant(tokenString string) (userID string, role string, err error) {
	tokenObject, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(services.SecretKey), nil
	})
	if err != nil {
		return
	}
	claims, ok := tokenObject.Claims.(jwt.MapClaims)
	if !ok {
		return
	}
	role = string(claims["Role"].(string))
	userIDInterface, ok := claims["user_id"]
	if !ok {
		return
	}
	userID, ok = userIDInterface.(string)
	if !ok {
		return
	}
	//userID = string(claims["user_id"]).(string)
	if !ok || (role != "Applicant") {
		err = errors.New("user is not admin or superadmin")
		//logger.WithField("err", err.Error()).Error(" user is not admin or superadmin")
		return
	}
	return
}
