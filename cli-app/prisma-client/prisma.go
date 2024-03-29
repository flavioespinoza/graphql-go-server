// Code generated by Prisma CLI (https://github.com/prisma/prisma). DO NOT EDIT.

package prisma

import (
	"context"
	"errors"

	"github.com/prisma/prisma-client-lib-go"

	"github.com/machinebox/graphql"
)

var ErrNoResult = errors.New("query returned no result")

func Str(v string) *string { return &v }
func Int32(v int32) *int32 { return &v }
func Bool(v bool) *bool    { return &v }

type BatchPayloadExec struct {
	exec *prisma.BatchPayloadExec
}

func (exec *BatchPayloadExec) Exec(ctx context.Context) (BatchPayload, error) {
	bp, err := exec.exec.Exec(ctx)
	return BatchPayload(bp), err
}

type BatchPayload struct {
	Count int64 `json:"count"`
}

type Aggregate struct {
	Count int64 `json:"count"`
}

type Client struct {
	Client *prisma.Client
}

type Options struct {
	Endpoint string
	Secret   string
}

func New(options *Options, opts ...graphql.ClientOption) *Client {
	endpoint := DefaultEndpoint
	secret := Secret
	if options != nil {
		endpoint = options.Endpoint
		secret = options.Secret
	}
	return &Client{
		Client: prisma.New(endpoint, secret, opts...),
	}
}

func (client *Client) GraphQL(ctx context.Context, query string, variables map[string]interface{}) (map[string]interface{}, error) {
	return client.Client.GraphQL(ctx, query, variables)
}

var DefaultEndpoint = "http://localhost:4466"
var Secret = ""

func (client *Client) Todo(params TodoWhereUniqueInput) *TodoExec {
	ret := client.Client.GetOne(
		nil,
		params,
		[2]string{"TodoWhereUniqueInput!", "Todo"},
		"todo",
		[]string{"id", "createdAt", "text"})

	return &TodoExec{ret}
}

type TodoesParams struct {
	Where   *TodoWhereInput   `json:"where,omitempty"`
	OrderBy *TodoOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) Todoes(params *TodoesParams) *TodoExecArray {
	var wparams *prisma.WhereParams
	if params != nil {
		wparams = &prisma.WhereParams{
			Where:   params.Where,
			OrderBy: (*string)(params.OrderBy),
			Skip:    params.Skip,
			After:   params.After,
			Before:  params.Before,
			First:   params.First,
			Last:    params.Last,
		}
	}

	ret := client.Client.GetMany(
		nil,
		wparams,
		[3]string{"TodoWhereInput", "TodoOrderByInput", "Todo"},
		"todoes",
		[]string{"id", "createdAt", "text"})

	return &TodoExecArray{ret}
}

type TodoesConnectionParams struct {
	Where   *TodoWhereInput   `json:"where,omitempty"`
	OrderBy *TodoOrderByInput `json:"orderBy,omitempty"`
	Skip    *int32            `json:"skip,omitempty"`
	After   *string           `json:"after,omitempty"`
	Before  *string           `json:"before,omitempty"`
	First   *int32            `json:"first,omitempty"`
	Last    *int32            `json:"last,omitempty"`
}

func (client *Client) TodoesConnection(params *TodoesConnectionParams) TodoConnectionExec {
	panic("not implemented")
}

func (client *Client) CreateTodo(params TodoCreateInput) *TodoExec {
	ret := client.Client.Create(
		params,
		[2]string{"TodoCreateInput!", "Todo"},
		"createTodo",
		[]string{"id", "createdAt", "text"})

	return &TodoExec{ret}
}

type TodoUpdateParams struct {
	Data  TodoUpdateInput      `json:"data"`
	Where TodoWhereUniqueInput `json:"where"`
}

func (client *Client) UpdateTodo(params TodoUpdateParams) *TodoExec {
	ret := client.Client.Update(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[3]string{"TodoUpdateInput!", "TodoWhereUniqueInput!", "Todo"},
		"updateTodo",
		[]string{"id", "createdAt", "text"})

	return &TodoExec{ret}
}

type TodoUpdateManyParams struct {
	Data  TodoUpdateManyMutationInput `json:"data"`
	Where *TodoWhereInput             `json:"where,omitempty"`
}

func (client *Client) UpdateManyTodoes(params TodoUpdateManyParams) *BatchPayloadExec {
	exec := client.Client.UpdateMany(
		prisma.UpdateParams{
			Data:  params.Data,
			Where: params.Where,
		},
		[2]string{"TodoUpdateManyMutationInput!", "TodoWhereInput"},
		"updateManyTodoes")
	return &BatchPayloadExec{exec}
}

type TodoUpsertParams struct {
	Where  TodoWhereUniqueInput `json:"where"`
	Create TodoCreateInput      `json:"create"`
	Update TodoUpdateInput      `json:"update"`
}

func (client *Client) UpsertTodo(params TodoUpsertParams) *TodoExec {
	uparams := &prisma.UpsertParams{
		Where:  params.Where,
		Create: params.Create,
		Update: params.Update,
	}
	ret := client.Client.Upsert(
		uparams,
		[4]string{"TodoWhereUniqueInput!", "TodoCreateInput!", "TodoUpdateInput!", "Todo"},
		"upsertTodo",
		[]string{"id", "createdAt", "text"})

	return &TodoExec{ret}
}

func (client *Client) DeleteTodo(params TodoWhereUniqueInput) *TodoExec {
	ret := client.Client.Delete(
		params,
		[2]string{"TodoWhereUniqueInput!", "Todo"},
		"deleteTodo",
		[]string{"id", "createdAt", "text"})

	return &TodoExec{ret}
}

func (client *Client) DeleteManyTodoes(params *TodoWhereInput) *BatchPayloadExec {
	exec := client.Client.DeleteMany(params, "TodoWhereInput", "deleteManyTodoes")
	return &BatchPayloadExec{exec}
}

type TodoOrderByInput string

const (
	TodoOrderByInputIDAsc         TodoOrderByInput = "id_ASC"
	TodoOrderByInputIDDesc        TodoOrderByInput = "id_DESC"
	TodoOrderByInputCreatedAtAsc  TodoOrderByInput = "createdAt_ASC"
	TodoOrderByInputCreatedAtDesc TodoOrderByInput = "createdAt_DESC"
	TodoOrderByInputTextAsc       TodoOrderByInput = "text_ASC"
	TodoOrderByInputTextDesc      TodoOrderByInput = "text_DESC"
)

type MutationType string

const (
	MutationTypeCreated MutationType = "CREATED"
	MutationTypeUpdated MutationType = "UPDATED"
	MutationTypeDeleted MutationType = "DELETED"
)

type TodoWhereUniqueInput struct {
	ID   *string `json:"id,omitempty"`
	Text *string `json:"text,omitempty"`
}

type TodoWhereInput struct {
	ID                *string          `json:"id,omitempty"`
	IDNot             *string          `json:"id_not,omitempty"`
	IDIn              []string         `json:"id_in,omitempty"`
	IDNotIn           []string         `json:"id_not_in,omitempty"`
	IDLt              *string          `json:"id_lt,omitempty"`
	IDLte             *string          `json:"id_lte,omitempty"`
	IDGt              *string          `json:"id_gt,omitempty"`
	IDGte             *string          `json:"id_gte,omitempty"`
	IDContains        *string          `json:"id_contains,omitempty"`
	IDNotContains     *string          `json:"id_not_contains,omitempty"`
	IDStartsWith      *string          `json:"id_starts_with,omitempty"`
	IDNotStartsWith   *string          `json:"id_not_starts_with,omitempty"`
	IDEndsWith        *string          `json:"id_ends_with,omitempty"`
	IDNotEndsWith     *string          `json:"id_not_ends_with,omitempty"`
	CreatedAt         *string          `json:"createdAt,omitempty"`
	CreatedAtNot      *string          `json:"createdAt_not,omitempty"`
	CreatedAtIn       []string         `json:"createdAt_in,omitempty"`
	CreatedAtNotIn    []string         `json:"createdAt_not_in,omitempty"`
	CreatedAtLt       *string          `json:"createdAt_lt,omitempty"`
	CreatedAtLte      *string          `json:"createdAt_lte,omitempty"`
	CreatedAtGt       *string          `json:"createdAt_gt,omitempty"`
	CreatedAtGte      *string          `json:"createdAt_gte,omitempty"`
	Text              *string          `json:"text,omitempty"`
	TextNot           *string          `json:"text_not,omitempty"`
	TextIn            []string         `json:"text_in,omitempty"`
	TextNotIn         []string         `json:"text_not_in,omitempty"`
	TextLt            *string          `json:"text_lt,omitempty"`
	TextLte           *string          `json:"text_lte,omitempty"`
	TextGt            *string          `json:"text_gt,omitempty"`
	TextGte           *string          `json:"text_gte,omitempty"`
	TextContains      *string          `json:"text_contains,omitempty"`
	TextNotContains   *string          `json:"text_not_contains,omitempty"`
	TextStartsWith    *string          `json:"text_starts_with,omitempty"`
	TextNotStartsWith *string          `json:"text_not_starts_with,omitempty"`
	TextEndsWith      *string          `json:"text_ends_with,omitempty"`
	TextNotEndsWith   *string          `json:"text_not_ends_with,omitempty"`
	And               []TodoWhereInput `json:"AND,omitempty"`
	Or                []TodoWhereInput `json:"OR,omitempty"`
	Not               []TodoWhereInput `json:"NOT,omitempty"`
}

type TodoCreateInput struct {
	ID   *string `json:"id,omitempty"`
	Text string  `json:"text"`
}

type TodoUpdateInput struct {
	Text *string `json:"text,omitempty"`
}

type TodoUpdateManyMutationInput struct {
	Text *string `json:"text,omitempty"`
}

type TodoSubscriptionWhereInput struct {
	MutationIn                 []MutationType               `json:"mutation_in,omitempty"`
	UpdatedFieldsContains      *string                      `json:"updatedFields_contains,omitempty"`
	UpdatedFieldsContainsEvery []string                     `json:"updatedFields_contains_every,omitempty"`
	UpdatedFieldsContainsSome  []string                     `json:"updatedFields_contains_some,omitempty"`
	Node                       *TodoWhereInput              `json:"node,omitempty"`
	And                        []TodoSubscriptionWhereInput `json:"AND,omitempty"`
	Or                         []TodoSubscriptionWhereInput `json:"OR,omitempty"`
	Not                        []TodoSubscriptionWhereInput `json:"NOT,omitempty"`
}

type TodoExec struct {
	exec *prisma.Exec
}

func (instance TodoExec) Exec(ctx context.Context) (*Todo, error) {
	var v Todo
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance TodoExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type TodoExecArray struct {
	exec *prisma.Exec
}

func (instance TodoExecArray) Exec(ctx context.Context) ([]Todo, error) {
	var v []Todo
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type Todo struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	Text      string `json:"text"`
}

type TodoConnectionExec struct {
	exec *prisma.Exec
}

func (instance *TodoConnectionExec) PageInfo() *PageInfoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "PageInfo"},
		"pageInfo",
		[]string{"hasNextPage", "hasPreviousPage", "startCursor", "endCursor"})

	return &PageInfoExec{ret}
}

func (instance *TodoConnectionExec) Edges() *TodoEdgeExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "TodoEdge"},
		"edges",
		[]string{"cursor"})

	return &TodoEdgeExec{ret}
}

func (instance *TodoConnectionExec) Aggregate(ctx context.Context) (Aggregate, error) {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "AggregateTodo"},
		"aggregate",
		[]string{"count"})

	var v Aggregate
	_, err := ret.Exec(ctx, &v)
	return v, err
}

func (instance TodoConnectionExec) Exec(ctx context.Context) (*TodoConnection, error) {
	var v TodoConnection
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance TodoConnectionExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type TodoConnectionExecArray struct {
	exec *prisma.Exec
}

func (instance TodoConnectionExecArray) Exec(ctx context.Context) ([]TodoConnection, error) {
	var v []TodoConnection
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type TodoConnection struct {
}

type PageInfoExec struct {
	exec *prisma.Exec
}

func (instance PageInfoExec) Exec(ctx context.Context) (*PageInfo, error) {
	var v PageInfo
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance PageInfoExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type PageInfoExecArray struct {
	exec *prisma.Exec
}

func (instance PageInfoExecArray) Exec(ctx context.Context) ([]PageInfo, error) {
	var v []PageInfo
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
}

type TodoEdgeExec struct {
	exec *prisma.Exec
}

func (instance *TodoEdgeExec) Node() *TodoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "Todo"},
		"node",
		[]string{"id", "createdAt", "text"})

	return &TodoExec{ret}
}

func (instance TodoEdgeExec) Exec(ctx context.Context) (*TodoEdge, error) {
	var v TodoEdge
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance TodoEdgeExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type TodoEdgeExecArray struct {
	exec *prisma.Exec
}

func (instance TodoEdgeExecArray) Exec(ctx context.Context) ([]TodoEdge, error) {
	var v []TodoEdge
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type TodoEdge struct {
	Cursor string `json:"cursor"`
}

type TodoSubscriptionPayloadExec struct {
	exec *prisma.Exec
}

func (instance *TodoSubscriptionPayloadExec) Node() *TodoExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "Todo"},
		"node",
		[]string{"id", "createdAt", "text"})

	return &TodoExec{ret}
}

func (instance *TodoSubscriptionPayloadExec) PreviousValues() *TodoPreviousValuesExec {
	ret := instance.exec.Client.GetOne(
		instance.exec,
		nil,
		[2]string{"", "TodoPreviousValues"},
		"previousValues",
		[]string{"id", "createdAt", "text"})

	return &TodoPreviousValuesExec{ret}
}

func (instance TodoSubscriptionPayloadExec) Exec(ctx context.Context) (*TodoSubscriptionPayload, error) {
	var v TodoSubscriptionPayload
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance TodoSubscriptionPayloadExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type TodoSubscriptionPayloadExecArray struct {
	exec *prisma.Exec
}

func (instance TodoSubscriptionPayloadExecArray) Exec(ctx context.Context) ([]TodoSubscriptionPayload, error) {
	var v []TodoSubscriptionPayload
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type TodoSubscriptionPayload struct {
	Mutation      MutationType `json:"mutation"`
	UpdatedFields []string     `json:"updatedFields,omitempty"`
}

type TodoPreviousValuesExec struct {
	exec *prisma.Exec
}

func (instance TodoPreviousValuesExec) Exec(ctx context.Context) (*TodoPreviousValues, error) {
	var v TodoPreviousValues
	ok, err := instance.exec.Exec(ctx, &v)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoResult
	}
	return &v, nil
}

func (instance TodoPreviousValuesExec) Exists(ctx context.Context) (bool, error) {
	return instance.exec.Exists(ctx)
}

type TodoPreviousValuesExecArray struct {
	exec *prisma.Exec
}

func (instance TodoPreviousValuesExecArray) Exec(ctx context.Context) ([]TodoPreviousValues, error) {
	var v []TodoPreviousValues
	err := instance.exec.ExecArray(ctx, &v)
	return v, err
}

type TodoPreviousValues struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	Text      string `json:"text"`
}
