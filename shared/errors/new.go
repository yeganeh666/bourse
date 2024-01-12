package errors

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mabna/shared/i18next"
)

func New(ctx context.Context, code codes.Code) *Error {
	return &Error{status.New(code, i18next.ByContext(ctx, CodeToSlug(code)))}
}

func NewWithSlug(ctx context.Context, code codes.Code, slug string) *Error {
	return &Error{status.New(code, i18next.ByContext(ctx, slug))}
}

func NewWithDynamicSlug(ctx context.Context, code codes.Code, slug string, data interface{}) *Error {
	return &Error{status.New(code, i18next.ByContextWithData(ctx, slug, data))}
}
