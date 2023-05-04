package repository

import (
	"context"
	"fmt"
	"golang_db/entity"
	"golang_db/helper"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsert(t *testing.T) {
	commentRepository := NewCommentRepository(helper.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email: "test@localhost.com",
		Comment: "Test Comment",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil{
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T){
	commentRepository := NewCommentRepository(helper.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindById(ctx, 1)
	if err != nil{
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAll(t *testing.T){
	commentRepository := NewCommentRepository(helper.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindAll(ctx)
	if err != nil{
		panic(err)
	}

	fmt.Println(result)
}

func TestUpdate(t *testing.T){
	commentRepository := NewCommentRepository(helper.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Id: 1,
		Email: "update_email@localhost.com",
		Comment: "Test Update",
	}
	result, err := commentRepository.Update(ctx, comment)
	if err != nil{
		panic(err)
	}

	fmt.Println(result)
}

func TestDelete(t *testing.T){
	commentRepository := NewCommentRepository(helper.GetConnection())

	ctx := context.Background()
	result, err := commentRepository.Delete(ctx, 1)
	if err != nil{
		panic(err)
	}

	fmt.Println(result)
}