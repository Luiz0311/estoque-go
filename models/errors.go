package models

import "errors"

var (
	ErrDB               = errors.New("erro na base de dados")
	ErrRead             = errors.New("erro ao ler produto")
	ErrProductsNotFound = errors.New("nenhum produto encontrado")
	ErrProductNotFound  = errors.New("produto não encontrado")
	ErrSearchProduct    = errors.New("erro ao buscar o produto")
	ErrNoPrice          = errors.New("atributo \"price\" tem que ser >= 0")
	ErrNoName           = errors.New("falta o atributo \"name\"")
	ErrNoType           = errors.New("falta o atributo \"type\"")
	ErrCreateProduct    = errors.New("erro ao criar produto")
	ErrDelete           = errors.New("erro ao deletar produto")
	ErrNoData           = errors.New("nenhum campo enviado")
)
