package orm

import (
	"context"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/jianfengye/collection"
	"gorm.io/gorm"
)

func (app *GormService) GetTables(ctx context.Context, db *gorm.DB) ([]string, error) {
	return db.Migrator().GetTables()
}

func (app *GormService) HasTable(ctx context.Context, db *gorm.DB, table string) (bool, error) {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return false, err
	}
	tableColl := collection.NewStrCollection(tables)
	isContain := tableColl.Contains(table)
	return isContain, nil
}

func (app *GormService) GetTableColumns(ctx context.Context, db *gorm.DB, table string) ([]contract.TableColumn, error) {
	columnTypes, err := db.Migrator().ColumnTypes(table)
	if err != nil {
		return nil, err
	}

	var columns []contract.TableColumn
	for _, ct := range columnTypes {
		name := ct.Name()
		columnType, _ := ct.ColumnType()
		nullable, _ := ct.Nullable()
		primaryKey, _ := ct.PrimaryKey()
		defaultValue, _ := ct.DefaultValue()
		comment, _ := ct.Comment()
		c := contract.TableColumn{
			Field:   name,
			Type:    columnType,
			Null:    nullable,
			Key:     primaryKey,
			Default: defaultValue,
			Comment: comment,
		}
		columns = append(columns, c)
	}

	return columns, nil
}
