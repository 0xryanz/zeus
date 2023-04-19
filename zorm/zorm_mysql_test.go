package zorm

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func OpenMysqlDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("mysql", "root:@/test")
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	return engine
}

func TestMysqlNewEngine(t *testing.T) {
	engine := OpenMysqlDB(t)
	defer engine.Close()
}
