package repositories

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/inact25/PickMyFood-BackEnd/masters/apis/models"
)

type ProductRepoImpl struct {
	db *sql.DB
}

func (s *ProductRepoImpl) GetProducts() ([]*models.ProductModels, error) {
	var products []*models.ProductModels
	query := "SELECT * FROM tb_product"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := models.ProductModels{}
		err := rows.Scan(&product.ProductID, &product.StoreID, &product.ProductName, &product.ProductCategoryID, &product.ProductStock, &product.ProductStatus)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)

	}

	return products, nil
}

func (s *ProductRepoImpl) GetProductByID(ID string) (*models.ProductModels, error) {
	results := s.db.QueryRow("SELECT * FROM tb_store WHERE store_id = ?", ID)

	var d models.ProductModels
	err := results.Scan(&d.ProductID, &d.StoreID, &d.ProductName, &d.ProductCategoryID, &d.ProductStock, &d.ProductStatus)
	if err != nil {
		return nil, errors.New("ID Not Found")
	}

	return &d, nil
}

func (s *ProductRepoImpl) PostProduct(d models.ProductModels) (*models.ProductModels, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmnt, _ := tx.Prepare(`INSERT INTO tb_product(product_id, store_id, product_name, product_category_id, product_stock, product_status) VALUES(?, ?, ?, ?, ?, ?)`)
	defer stmnt.Close()

	result, err := stmnt.Exec(d.ProductID, d.StoreID, d.ProductName, d.ProductCategoryID, d.ProductStock, d.ProductStatus)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	lastInsertID, _ := result.LastInsertId()
	tx.Commit()
	return s.GetProductByID(strconv.Itoa(int(lastInsertID)))
}

func (s *ProductRepoImpl) UpdateProduct(ID string, data models.ProductModels) (*models.ProductModels, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(`UPDATE tb_product SET store_id=?, product_name=?, product_category_id=?, product_stock=?, product_status=? WHERE product_id=?`,
		data.StoreID, data.ProductName, data.ProductCategoryID, data.ProductStock, data.ProductStatus, ID)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return s.GetProductByID(ID)
}

func (s *ProductRepoImpl) DeleteProduct(ID string) (*models.ProductModels, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec("DELETE FROM tb_product WHERE product_id = ?", ID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return s.GetProductByID(ID)

}

func (s *ProductRepoImpl) GetProductsPrice() ([]*models.ProductPrice, error) {
	var productsPrice []*models.ProductPrice
	query := "SELECT * FROM tb_product_price"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		productPrice := models.ProductPrice{}
		err := rows.Scan(&productPrice.ProductPriceID, &productPrice.ProductID, &productPrice.Price, &productPrice.DateModified)

		if err != nil {
			return nil, err
		}

		productsPrice = append(productsPrice, &productPrice)

	}

	return productsPrice, nil
}

func (s *ProductRepoImpl) GetProductPriceByID(ID string) (*models.ProductPrice, error) {
	results := s.db.QueryRow("SELECT * FROM tb_product_price WHERE product_price_id = ?", ID)

	var d models.ProductPrice
	err := results.Scan(&d.ProductPriceID, &d.ProductID, &d.Price, &d.DateModified)
	if err != nil {
		return nil, errors.New("ID Not Found")
	}

	return &d, nil
}

func (s *ProductRepoImpl) PostProductPrice(d models.ProductPrice) (*models.ProductPrice, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmnt, _ := tx.Prepare(`INSERT INTO tb_product_price(product_price_id, product_id, price, date_modified) VALUES (?, ?, ?, ?)`)
	defer stmnt.Close()

	result, err := stmnt.Exec(d.ProductPriceID, d.ProductID, d.Price, d.DateModified)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	lastInsertID, _ := result.LastInsertId()
	tx.Commit()
	return s.GetProductPriceByID(strconv.Itoa(int(lastInsertID)))
}

func (s *ProductRepoImpl) UpdateProductPrice(ID string, data models.ProductPrice) (*models.ProductPrice, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(`UPDATE tb_product_price SET product_id=?, price=?, date_modified=? WHERE product_price_id=?`,
		data.ProductID, data.Price, data.DateModified, ID)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return s.GetProductPriceByID(ID)
}

func (s *ProductRepoImpl) DeleteProductPrice(ID string) (*models.ProductPrice, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec("DELETE FROM tb_product_price WHERE product_price_id = ?", ID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return s.GetProductPriceByID(ID)

}

func (s *ProductRepoImpl) GetProductsCategory() ([]*models.ProductCategory, error) {
	var productsCategory []*models.ProductCategory
	query := "SELECT * FROM tb_product_category"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		productCategory := models.ProductCategory{}
		err := rows.Scan(&productCategory.ProductCategoryID, &productCategory.ProductCategoryName)

		if err != nil {
			return nil, err
		}

		productsCategory = append(productsCategory, &productCategory)

	}

	return productsCategory, nil
}

func (s *ProductRepoImpl) GetProductCategoryByID(ID string) (*models.ProductCategory, error) {
	results := s.db.QueryRow("SELECT * FROM tb_product_category WHERE product_category_id = ?", ID)

	var d models.ProductCategory
	err := results.Scan(&d.ProductCategoryID, &d.ProductCategoryName)
	if err != nil {
		return nil, errors.New("ID Not Found")
	}

	return &d, nil
}

func (s *ProductRepoImpl) PostProductCategory(d models.ProductCategory) (*models.ProductCategory, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmnt, _ := tx.Prepare(`INSERT INTO tb_product_category(product_category_id, product_category_name) VALUES (?, ?)`)
	defer stmnt.Close()

	result, err := stmnt.Exec(d.ProductCategoryID, d.ProductCategoryName)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	lastInsertID, _ := result.LastInsertId()
	tx.Commit()
	return s.GetProductCategoryByID(strconv.Itoa(int(lastInsertID)))
}

func (s *ProductRepoImpl) UpdateProductCategory(ID string, data models.ProductCategory) (*models.ProductCategory, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(`UPDATE tb_product_category SET product_category_name=? WHERE product_category_id=?`,
		data.ProductCategoryName, ID)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return s.GetProductCategoryByID(ID)
}

func (s *ProductRepoImpl) DeleteProductCategory(ID string) (*models.ProductCategory, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec("DELETE FROM tb_product_category WHERE product_category_id = ?", ID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return s.GetProductCategoryByID(ID)

}

func InitProductRepoImpl(db *sql.DB) ProductsRepo {
	return &ProductRepoImpl{db}

}
