package repository

import (
	"api-go/src/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		"VALUES($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {

	query, err := pr.connection.Prepare("SELECT * FROM product WHERE od = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product

	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &produto, nil
}

func (pr *ProductRepository) UpdateProduct(product model.Product) error {
	query, err := pr.connection.Prepare("UPDATE product SET product_name = $1, price = $2 WHERE id = $3")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer query.Close()

	_, err = query.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {
	query, err := pr.connection.Prepare("DELETE FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer query.Close()

	_, err = query.Exec(id)
	if err != nil {
		//fmt.Println(err)
		return err
	}

	return nil
}
