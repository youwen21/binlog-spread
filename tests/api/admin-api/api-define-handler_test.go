package admin_api

import (
	"binlog_spread/app/dto"
	"binlog_spread/app/models"
	lib2 "binlog_spread/lib"
	"binlog_spread/router"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var r *gin.Engine

func init() {

}

func setup() {
	dir, _ := os.Getwd()
	rootDir, err := lib2.GetRootDir(dir)
	err = godotenv.Load(rootDir + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	print("setup define handler")
	r = router.SetupRouter()
}

func teardown() {
	fmt.Println("After all tests")
}

//https://juejin.cn/post/6862992090985562125
func TestMain(m *testing.M)  {
	setup()
	fmt.Println("Test begins....")
	code := m.Run() // 如果不加这句，只会执行Main
	teardown()
	os.Exit(code)
}


func Test_apiDefineHandler_Add(t *testing.T) {
	// 方法与路由
	//method := "POST"
	//url := "/admin-api/v1/api_define"
}

func Test_apiDefineHandler_Delete(t *testing.T) {
	// 方法与路由
	//method := "DELETE"
	//url := "/admin-api/v1/api_define"
}

func RandString(len int) string {
	rd := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := rd.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

type formReader models.ApiDefine

func (f formReader) Read(p []byte) (n int, err error) {
	return len(p), nil
}

func Test_apiDefineHandler_Edit(t *testing.T) {
	// 方法与路由
	method := "POST"
	url := "/admin-api/v1/api_define/51"
	form := models.ApiDefine{
		ApiName: RandString(10),
	}

	//var bio io.Reader
	//bio = frd

	// http api gin的单元测试
	req, _ := http.NewRequest(method, url, formReader(form))
	w2 := httptest.NewRecorder()
	req.Header.Add("AdminAuthorization", getToken())
	r.ServeHTTP(w2, req)

	// 输出调用接口， 返回内容
	print(method, ":", url, "  -  ", w2.Body.String(), "\n")
	assert.Equal(t, 200, w2.Code)
}

func Test_apiDefineHandler_GetInfo(t *testing.T) {
	// 方法与路由
	method := "GET"
	url := "/admin-api/v1/api_define/" + "51"

	// http api gin的单元测试
	req, _ := http.NewRequest(method, url, nil)
	w2 := httptest.NewRecorder()
	req.Header.Add("AdminAuthorization", getToken())
	r.ServeHTTP(w2, req)

	// 输出调用接口， 返回内容
	print(method, ":", url, "  -  ", w2.Body.String(), "\n")
	assert.Equal(t, 200, w2.Code)

	type ApiResp struct {
		Code int              `json:"code"`
		Msg  string           `json:"msg"`
		Data models.ApiDefine `json:"data"`
	}
	apiResp := ApiResp{}
	if err := json.Unmarshal(w2.Body.Bytes(), &apiResp); err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, 51, apiResp.Data.ApiDefineId, "返回内容有错误")
	//lib.PrettyPrint(apiResp)
}

func Test_apiDefineHandler_Query1(t *testing.T) {
	method := "GET"
	url := "/admin-api/v1/api_define"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, nil)
	r.ServeHTTP(w, req)
	print(method, ":", url, "  -  ", w.Body.String(), "\n")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"code\":1,\"data\":\"\",\"msg\":\"缺失认证token\"}", w.Body.String())
}

func Test_apiDefineHandler_Query2(t *testing.T) {
	// 方法与路由
	method := "GET"
	url := "/admin-api/v1/api_define"

	// http api gin的单元测试
	req, _ := http.NewRequest(method, url, nil)
	w2 := httptest.NewRecorder()
	req.Header.Add("AdminAuthorization", getToken())
	r.ServeHTTP(w2, req)

	// 输出调用接口， 返回内容
	print(method, ":", url, "  -  ", w2.Body.String(), "\n")
	assert.Equal(t, 200, w2.Code)

	// 以下开始返回内容断言
	type ApiResp struct {
		Code int                 `json:"code"`
		Msg  string              `json:"msg"`
		Data dto.ApiDefineResult `json:"data"`
	}
	apiResp := ApiResp{}
	_ = json.Unmarshal(w2.Body.Bytes(), &apiResp)
	//lib.PrettyPrint(apiResp)
	assert.True(t, apiResp.Data.Count > 0, "返回数据为空")
}

func getToken() string {
	return os.Getenv("UNIT_ADMIN_TOKEN")
}