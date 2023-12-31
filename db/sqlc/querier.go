// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteEntry(ctx context.Context, id int64) error
	DeletePost(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	GetAccounts(ctx context.Context, arg GetAccountsParams) ([]Account, error)
	GetAllEntries(ctx context.Context, arg GetAllEntriesParams) ([]Entry, error)
	GetEntriesByAccount(ctx context.Context, accountID int64) ([]Entry, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetPost(ctx context.Context, id int64) (Post, error)
	GetPosts(ctx context.Context, arg GetPostsParams) ([]Post, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateEntry(ctx context.Context, arg UpdateEntryParams) (Entry, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
}

var _ Querier = (*Queries)(nil)
