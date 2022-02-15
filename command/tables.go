package command

import (
	"fmt"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"io"
	"lyky-go/global"
	"os"
	"strings"
)

// Table 表
type Table struct {
	TableName    string `gorm:"column:TABLE_NAME"`    // 表名
	TableComment string `gorm:"column:TABLE_COMMENT"` // 表名注释
}

// Field 表字段
type Field struct {
	Field      string `gorm:"column:Field"`      // 字段
	Type       string `gorm:"column:Type"`       // 类型
	Collation  string `gorm:"column:Collation"`  // 编码
	Null       string `gorm:"column:Null"`       // 是否为空
	Key        string `gorm:"column:Key"`        // 主键
	Default    string `gorm:"column:Default"`    // 默认值
	Extra      string `gorm:"column:Extra"`      // 额外类型
	Privileges string `gorm:"column:Privileges"` // 权限
	Comment    string `gorm:"column:Comment"`    // 注释
}

// GenerateAllModel 生成所有model的结构体
func GenerateAllModel(dbName string) {
	tables := GetAllTables(dbName)
	for _, val := range tables {
		err := GenerateSingleModel(dbName, val.TableName, val)
		if err != nil {
			continue
		}
	}
}

// GenerateSingleModel 生成单个model的结构体
func GenerateSingleModel(dbName string, tbName string, table Table) error {
	var fields []Field
	fields = GetFieldsByTable(tbName)
	content := ParseFieldsByTable(tbName, table.TableComment, fields)
	camelTbName := generator.CamelCase(tbName)

	// 生成model文件
	fileName := "./models/" + tbName + ".go"
	var f *os.File
	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		fmt.Println(tbName + " 已存在.")
		return err
	} else {
		f, err = os.Create(fileName)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	if _, err := io.WriteString(f, content); err != nil {
		fmt.Println(camelTbName + " 生成失败.")
		return err
	} else {
		fmt.Println(camelTbName + " 已生成.")
	}
	return nil
}

// GetAllTables 获取所有表信息和字段信息
func GetAllTables(dbName string) []Table {
	var tables []Table
	err := global.DB.Raw(`
			SELECT
				TABLE_NAME,      -- 表名
				TABLE_COMMENT    -- 表名注释
				FROM
				INFORMATION_SCHEMA.TABLES
				WHERE TABLE_SCHEMA = "` + dbName + `" -- 数据库名
		`).Scan(&tables).Error
	if err != nil {
		return []Table{}
	}
	return tables
}

// GetSingleTable 获取单个表信息和字段信息
func GetSingleTable(dbName string, tbName string) Table {
	var table Table
	err := global.DB.Raw(`
			SELECT
				TABLE_NAME,      -- 表名
				TABLE_COMMENT    -- 表名注释
				FROM
				INFORMATION_SCHEMA.TABLES
				WHERE TABLE_SCHEMA = "` + dbName + `" AND TABLE_NAME = "` + tbName + `"
		`).Find(&table).Error
	if err != nil {
		return Table{}
	}
	return table
}

// ParseFieldsByTable 转换数据表字段的类型
func ParseFieldsByTable(tbName string, tbComment string, fields []Field) string {
	content := "package models\n\n"
	if len(tbComment) > 0 {
		content += "// " + tbComment + "\n"
	}
	camelTbName := generator.CamelCase(tbName)
	content += "type " + camelTbName + " struct {\n"
	for _, val := range fields {
		// 生成字段
		columnField := generator.CamelCase(val.Field)
		columnJson := "`gorm:\""
		if val.Key == "PRI" {
			columnJson += "primaryKey;"
		}
		if val.Extra == "auto_increment" {
			columnJson += "autoIncrement;"
		}
		columnJson += "column:" + val.Field + ";type:" + val.Type + ";"
		if val.Default != "" {
			columnJson += "default '" + val.Default + "';"
		}
		if val.Null == "NO" {
			columnJson += "NOT NULL;"
		} else {
			columnJson += "NULL;"
		}
		if val.Comment != "" {
			columnJson += "comment:" + val.Comment
		}
		columnJson += "\" json:\"" + val.Field + "\"`"
		columnType := ParseFieldTypeByTable(val.Type)
		columnComment := ""
		if len(val.Comment) > 0 {
			columnComment = "// " + val.Comment
		}
		content += "    " + columnField + " " + columnType + " " + columnJson + " " + columnComment + "\n"
	}
	content += "}"
	return content
}

// GetFieldsByTable 根据表获取表字段
func GetFieldsByTable(tbName string) []Field {
	var field []Field
	err := global.DB.Raw(`
		SHOW FULL COLUMNS FROM ` + tbName + `
	`).Find(&field).Error
	if err != nil {
		return []Field{}
	}
	return field
}

// ParseFieldTypeByTable 转义数据库字段类型到model
func ParseFieldTypeByTable(fieldType string) string {
	typeArr := strings.Split(fieldType, "(")
	switch typeArr[0] {
	case "int", "integer", "mediumint", "bit", "year", "smallint", "int unsigned", "mediumint unsigned", "smallint unsigned":
		return "int"
	case "tinyint":
		return "int8"
	case "tinyint unsigned":
		return "uint8"
	case "bigint":
		return "int64"
	case "bigint unsigned":
		return "uint64"
	case "decimal", "double", "float", "real", "numeric":
		return "float32"
	case "timestamp", "datetime", "time":
		return "time.Time"
	default:
		return "string"
	}
}
