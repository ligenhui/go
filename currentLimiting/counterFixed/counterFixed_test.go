package counterFixed

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestCounterFixed_Check(t *testing.T) {
	/*url1 := "a/b/c"
	url2 := "abc"
	cf := GetCfInstance()
	cf.Add(url1, time.Second*10, 100)  //设置10秒内最大放行100个请求进来
	cf.Add(url2, time.Second*60, 1000) //设置60秒内最大放行1000个请求进来
	http.HandleFunc("/a/b/c", func(w http.ResponseWriter, r *http.Request) {
		if cf.Check(url1) { //验证请求是否超出限制
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte("success"))
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("fail"))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	})
	http.HandleFunc("/abc", func(w http.ResponseWriter, r *http.Request) {
		if cf.Check(url2) { //验证请求是否超出限制
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte("success"))
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("fail"))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	})
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err.Error())
	}*/

	url1 := "a/b/c"
	cf := GetCfInstance()
	cf.Add(url1, time.Second*10, 100) //设置10秒内最大放行100个请求进来
	var count uint32 = 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !cf.Check(url1) {
				atomic.AddUint32(&count, 1)
			}
		}()
	}
	wg.Wait()
	if count != 900 {
		t.Errorf("count %v", count)
	}
}
