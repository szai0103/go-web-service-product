package data

import (
	"context"
	"database/sql"
	"errors"
	"github.com/schwarz/inventoryservice/product/data/mysql"
	"github.com/schwarz/inventoryservice/product/model"
	"log"
	"time"
)

func GetProductList() ([]model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	results, err := mysql.DbConnection.QueryContext(ctx, `SELECT 
	productID, 
	productName, 
	productPrice 
	FROM products`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	products := make([]model.Product, 0)
	for results.Next() {
		var product model.Product
		results.Scan(&product.ProductID,
			&product.ProductName,
			&product.ProductPrice)

		products = append(products, product)
	}
	return products, nil
}

func GetTopThreeProducts() ([]model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	results, err := mysql.DbConnection.QueryContext(ctx, `SELECT 
	productID,
	productName, 
	productPrice 
	FROM products ORDER BY productID DESC LIMIT 3
	`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	products := make([]model.Product, 0)
	for results.Next() {
		var product model.Product
		results.Scan(&product.ProductID,
			&product.ProductName,
			&product.ProductPrice)

		products = append(products, product)
	}
	return products, nil
}

func GetProduct(productID int) (*model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := mysql.DbConnection.QueryRowContext(ctx, `SELECT 
	productID, 
	productName, 
	productPrice 
	FROM products 
	WHERE productID = ?`, productID)
	product := &model.Product{}
	err := row.Scan(&product.ProductID,
		&product.ProductName,
		&product.ProductPrice,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return product, nil
}

func RemoveProduct(productID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := mysql.DbConnection.ExecContext(ctx,
		`DELETE FROM products where productId = ?`, productID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func UpdateProduct(product model.Product) error {
	// if the product id is set, update, otherwise add
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if product.ProductID == nil || *product.ProductID == 0 {
		return errors.New("product has invalid ID")
	}
	_, err := mysql.DbConnection.ExecContext(ctx, `UPDATE products SET 
		productName=?,
		productPrice=?
		WHERE productId=?`,
		product.ProductName,
		product.ProductPrice,
		product.ProductID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func CreateProduct(product model.Product) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := mysql.DbConnection.ExecContext(ctx, `INSERT INTO products  
	(
	productName,
	productPrice
	) VALUES (?, ?)`,
		product.ProductName,
		product.ProductPrice)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return int(insertID), nil
}
