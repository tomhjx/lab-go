package mysqlhd

import (
	"testing"
	"time"
)

func TestGORMHandleQuery(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/lab?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:123456@tcp(mysql8.lab.local:3306)/lab?charset=utf8mb4&parseTime=True&loc=Local"
	h, err := NewGORMHandle(dsn)
	if err != nil {
		t.Fatal(err)
	}

	for i := 1; i < 10; i++ {
		go func(i int) {
			r, err := h.Query("select sleep(30),id,title,summary,update_time from app_sample where id =?", i)
			if err != nil {
				t.Log(err)
				return
			}
			t.Logf("query res:%s", r)
		}(i)
	}

	time.Sleep(time.Minute)
}
