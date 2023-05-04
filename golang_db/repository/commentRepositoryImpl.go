package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_db/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository{
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error){
	query := "INSERT INTO COMMENT(EMAIL, COMMENT) VALUES(?, ?)"
	result, err := repository.DB.ExecContext(ctx, query, comment.Email, comment.Comment)
	if err != nil{
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil{
		return comment, err
	}
	comment.Id = id
	return comment, nil
}
func (repository *commentRepositoryImpl) Update(ctx context.Context, comment entity.Comment) (string, error){
	query := "UPDATE COMMENT SET EMAIL = ? , COMMENT = ? WHERE COMMENT_ID = ?"
	result, err := repository.DB.ExecContext(ctx, query, comment.Email, comment.Comment, comment.Id)
	if err != nil{
		return "Failed to Update Comment", err
	}
	id, err := result.LastInsertId()
	if err != nil{
		return "Failed to Update Comment", err
	}
	comment.Id = id
	return "Comment Updated", nil
}
func (repository *commentRepositoryImpl) Delete(ctx context.Context, id int64) (string, error){
	query := "DELETE FROM COMMENT WHERE COMMENT_ID = ?"
	_, err := repository.DB.ExecContext(ctx, query, id)
	if err != nil{
		return "Failed to Delete Comment", err
	}
	return "Comment Deleted", nil
}
func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int64) (entity.Comment, error){
	query := "SELECT COMMENT_ID, EMAIL, COMMENT FROM COMMENT WHERE COMMENT_ID = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, query, id)
	comment := entity.Comment{}
	if err != nil{
		return comment, err
	}
	defer rows.Close()
	if rows.Next(){
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	}else{
		return comment, errors.New("Comment with Id: " + strconv.Itoa(int(id)) + " Not Found")
	}
}
func (repository *commentRepositoryImpl) FindAll(ctx context.Context)([]entity.Comment, error){
	query := "SELECT COMMENT_ID, EMAIL, COMMENT FROM COMMENT"
	rows, err := repository.DB.QueryContext(ctx, query)
	
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment

	for rows.Next(){
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
