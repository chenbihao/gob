package model

import (
	"fmt"
	"github.com/chenbihao/gob/framework/cobra"
	"github.com/chenbihao/gob/framework/contract"
	"github.com/chenbihao/gob/framework/provider/orm"
)

// modelTestCommand 检测数据库连接是否正常
var modelTestCommand = &cobra.Command{
	Use:          "test",
	Short:        "测试数据库",
	SilenceUsage: true,
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()

		gormService := container.MustMake(contract.ORMKey).(contract.ORM)
		db, err := gormService.GetDB(orm.WithConfigPath(database))
		if err != nil {
			fmt.Println("数据库连接：" + database + "失败，请检查配置")
			return err
		}

		// 获取所有表
		dbTables, err := db.Migrator().GetTables()
		if err != nil {
			fmt.Println("数据库连接：" + database + "失败，请检查配置")
			return err
		}
		fmt.Println("数据库连接：" + database + "成功")

		// 一共存在多少张表
		fmt.Printf("一共存在%d张表\n", len(dbTables))
		for _, table := range dbTables {
			fmt.Println(table)
		}
		return nil
	},
}
