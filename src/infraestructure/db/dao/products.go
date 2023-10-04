package infraestructure

import (
	"database/sql"
	"fmt"
	db "golang-template-clean-architecture/src/infraestructure/db/adapter"
	entities "golang-template-clean-architecture/src/modules/products/domain/entities"
)

type PostgreProductDao struct {
	sql *sql.DB
}

func NewPostgreProductDao(connection *db.DBConnection) *PostgreProductDao{
	return &PostgreProductDao{ sql: connection.DB}
}
func (dao *PostgreProductDao) Create(p entities.Product) (err error) {
	_, err = dao.sql.Exec("INSERT INTO products(name,price)VALUES($1,$2)", p.Name, p.Price)
	if err != nil{
		return
	}
	return
}

func (dao *PostgreProductDao) FindAll(page, limit int) ([]entities.Product, error) {
	var products []entities.Product
	fmt.Println("LLega aqui", page, limit)
	data, err := dao.sql.Query("SELECT * FROM products LIMIT $2 OFFSET $1", page, limit)
	if err != nil{
		return nil, err
	}
	for data.Next(){
		var p entities.Product
		err := data.Scan(&p.Id, &p.Name, &p.Price, &p.Brand, &p.CreatedAt)
		if err != nil{
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}