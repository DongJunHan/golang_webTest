package myapp
import(
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"encoding/json"
	"net/http/httptest"

)
func TestIndexPathHandler(t *testing.T){
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET","/",nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res,req)
	assert.Equal(http.StatusOK,res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello world",string(data))
}

func TestIndexPathHandler_WithoutName(t *testing.T){
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET","/bar",nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res,req)
	assert.Equal(http.StatusOK,res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World!",string(data))
}

func TestIndexPathHandler_WithName(t *testing.T){
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET","/bar?name=dongjun",nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res,req)
	assert.Equal(http.StatusOK,res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello dongjun!",string(data))
}

func TestFooHandler_WithoutJson(t *testing.T){
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET","/foo",nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res,req)

	assert.Equal(http.StatusBadRequest,res.Code)


}
func TestFooHandler_WithJson(t *testing.T){
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST","/foo",
	strings.NewReader(`{"first_name":"dongjun","last_name":"han","email":"gamedokdok@naver.com"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res,req)

	assert.Equal(http.StatusCreated,res.Code)
	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("dongjun",user.FirstName)
	assert.Equal("han",user.LastName)

}
