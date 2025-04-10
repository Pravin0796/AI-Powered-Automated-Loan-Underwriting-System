package middleware

import (
	"context"
	"strings"

	"AI-Powered-Automated-Loan-Underwriting-System/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var rolePermissions = map[string][]string{
	"admin":       {"ApproveLoan", "RejectLoan", "ViewLoans"},
	"underwriter": {"ApproveLoan", "ViewLoans"},
	"customer":    {"ApplyLoan", "ViewOwnLoan"},
}

// Context keys (to avoid string key collisions)
type contextKey string

const (
	ContextUserIDKey contextKey = "user_id"
	ContextRoleKey   contextKey = "role"
)

// JWTAuthInterceptor validates JWT and enforces role-based access control
func JWTAuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Missing metadata")
	}

	authHeader := md["authorization"]
	if len(authHeader) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "Missing token")
	}

	tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
	claims, err := utils.ValidateJWT(tokenString)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token: %v", err)
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, status.Errorf(codes.PermissionDenied, "Role not found in token")
	}

	if !isAuthorized(role, info.FullMethod) {
		return nil, status.Errorf(codes.PermissionDenied, "Access denied for role: %s", role)
	}

	// Add user claims to context
	ctx = context.WithValue(ctx, ContextUserIDKey, claims["user_id"])
	ctx = context.WithValue(ctx, ContextRoleKey, role)

	// Proceed with the request using the updated context
	return handler(ctx, req)
}

// isAuthorized checks if a role has access to the requested method
func isAuthorized(role, method string) bool {
	allowedMethods, exists := rolePermissions[role]
	if !exists {
		return false
	}

	for _, allowed := range allowedMethods {
		if strings.Contains(method, allowed) {
			return true
		}
	}
	return false
}
