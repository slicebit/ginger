package db

import (
	"github.com/aacanakin/qb"
)

// New returns a db context given driver & dsn
func New(driver string, dsn string) (*Context, error) {
	engine, err := qb.New(driver, dsn)
	if err != nil {
		return nil, err
	}

	metadata := qb.MetaData()
	return &Context{
		Engine:   engine,
		MetaData: metadata,
	}, nil
}

// Context is the db context representation
type Context struct {
	Engine   *qb.Engine
	MetaData *qb.MetaDataElem
}

// DropAll drops all tables registered to metadata
func (c *Context) DropAll() error {
	return c.MetaData.DropAll(c.Engine)
}

// CreateAll creates all tables registered to metadata
func (c *Context) CreateAll() error {
	return c.MetaData.CreateAll(c.Engine)
}

// Migrate drops & add tables registered to metadata given reset option
func (c *Context) Migrate(reset bool) error {
	var err error
	if reset {
		err = c.DropAll()
	}

	err = c.CreateAll()
	return err
}

// RegisterTable registers table to metadata
func (c *Context) RegisterTable(table qb.TableElem) {
	c.MetaData.AddTable(table)
}
