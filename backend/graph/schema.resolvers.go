package graph

import (
    "context"
    "fmt"
    "graph_note/graph/model"
)

// グローバル変数としてモックデータを定義
var users = []*model.User{
    {ID: "1", Name: "ユーザー1"},
    {ID: "2", Name: "ユーザー2"},
}

var todos = []*model.Todo{
    {ID: "1", Text: "GraphQLを学ぶ", Done: false, User: users[0]},
    {ID: "2", Text: "N+1問題を理解する", Done: false, User: users[0]},
    {ID: "3", Text: "DataLoaderを実装する", Done: false, User: users[1]},
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
    // ユーザーIDをチェック
    var user *model.User
    for _, u := range users {
        if u.ID == input.UserID {
            user = u
            break
        }
    }
    
    if user == nil {
        return nil, fmt.Errorf("user not found")
    }
    
    // 新しいTodoを作成
    todo := &model.Todo{
        ID:   fmt.Sprintf("%d", len(todos)+1),
        Text: input.Text,
        Done: false,
        User: user,
    }
    
    todos = append(todos, todo)
    return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
    // ここでN+1問題が発生します（各Todoに対して別々のクエリが発生するケース）
    // 実際のケースではデータベースからデータを取得します
    fmt.Println("Todos query executed") // クエリログ
    
    return todos, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }