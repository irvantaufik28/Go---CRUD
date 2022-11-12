package repository

import (
	"CRUD/helper"
	"CRUD/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into customer(name) values (?)"
	result, err := tx.ExceContext(ctx, SQL, category.Name)
	helper.PanicifError(err)

	id, err := result.LastInsertId()
	helper.PanicifError(err)

	category.Id = int(id)
	return category

}
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where id = ?"
	result, err := tx.ExceContext(ctx, SQLcategory.Name, category.id)
	helper.PanicifError(err)

	return category
}
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	SQL := "delete category set name = ? where id = ?"
	result, err := tx.ExceContext(ctx, SQLcategory.Name, category.id)
	helper.PanicifError(err)

	return category
}
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "select id, name from category where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicifError(err)

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicifError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}

}
func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "seletc id, name form category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicifError(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicifError(err)
		categories = append(categories, category)
	}

	return categories
}
