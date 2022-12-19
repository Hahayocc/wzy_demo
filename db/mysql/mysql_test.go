package mysql

import (
	"fmt"
	"sync"
	"testing"
)

// go test -v -test.run TestInitMysql
func TestInitMysql(t *testing.T) {
	InitMysql()
}

func TestCreateLockTable(t *testing.T) {
	InitMysql()
	var wg sync.WaitGroup
	count := 10
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(index int) {
			defer wg.Done()
			resp, err := CreateLockTable(fmt.Sprintf("测试%v", index))
			t.Logf("resp: %v, err: %v", resp, err)
		}(i)
	}
	wg.Wait()
	t.Log("success")
}

func TestUpdateCounts(t *testing.T) {
	InitMysql()
	var wg sync.WaitGroup
	count := 10
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(index int) {
			defer wg.Done()
			err := UpdateCounts("", fmt.Sprintf("%v", index))
			t.Logf("err: %v", err)
		}(i)
	}
	wg.Wait()
	t.Log("success")
}
