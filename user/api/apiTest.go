package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MiteshSharma/project/setting"
	"github.com/MiteshSharma/project/wrapper"

	"github.com/MiteshSharma/project/model"

	"github.com/MiteshSharma/project/core/eventdispatcher"
	"github.com/MiteshSharma/project/core/repository/docker"

	"github.com/MiteshSharma/project/core/bi"
	"github.com/MiteshSharma/project/core/bus"
	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/core/metrics"
	"github.com/MiteshSharma/project/user/repository"
	"github.com/gorilla/mux"
)

type APITest struct {
	API         *UserAPI
	MySQLDocker *docker.MysqlDocker
}

var apiTest *APITest

func SetupApiTest() *APITest {
	config := setting.GetConfigFromFile("test")
	logger := logger.NewTestLogger(config.LoggerConfig)
	metrics := metrics.NewTestMetrics()
	biEventHandler := bi.NewBiTestEventHandler()
	bus := bus.NewTestBus()
	router := mux.NewRouter()
	eventDispatcher := eventdispatcher.NewEventDispatcher(logger, bus, 10, 2)

	mysqlDocker := &docker.MysqlDocker{
		ContainerName: "mysql-api-container",
	}
	fmt.Println("Starting docker mysql container")
	mysqlDocker.StartMysqlDocker()
	fmt.Println("Started docker mysql container")
	fmt.Println("waiting for 10 sec before start")
	// Wait for docker mysql server to start
	time.Sleep(10 * time.Second)
	fmt.Println("waiting complete")

	storage := repository.NewPersistentRepository(logger, config, metrics)
	external := repository.NewExternalRepository(logger, config, metrics, bus)

	serverParam := model.NewServerParam(logger, metrics, bus, config, eventDispatcher, biEventHandler)

	api := NewUserAPI(router, storage, external, serverParam)
	apiTest = &APITest{
		API:         api,
		MySQLDocker: mysqlDocker,
	}
	return apiTest
}

func GetApiTest() *APITest {
	return apiTest
}

func (at *APITest) CleanUpApiTest() {
	at.MySQLDocker.Stop()
	time.Sleep(10 * time.Second)
}

func (at *APITest) CheckValidTestUser(t *testing.T, expectedUser *model.User, receivedUser *model.User) {
	t.Helper()

	if expectedUser.Email != receivedUser.Email {
		t.Errorf("handler returned wrong email: got %v want %v",
			receivedUser.Email, expectedUser.Email)
	}

	if expectedUser.FirstName != receivedUser.FirstName {
		t.Errorf("handler returned wrong first name: got %v want %v",
			receivedUser.FirstName, expectedUser.FirstName)
	}

	if expectedUser.LastName != receivedUser.LastName {
		t.Errorf("handler returned wrong last name: got %v want %v",
			receivedUser.LastName, expectedUser.LastName)
	}
}

func (at *APITest) CreateUserAuthFromTestAPI(t *testing.T, api *UserAPI, user *model.User) *model.UserAuth {
	t.Helper()

	jsonUser, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := wrapper.RequestHandler(apiTest.API.createUser)
	handler.ServeHTTP(res, req)

	CheckCreatedStatus(t, res.Code)

	t.Logf("Create user response %s", res.Body.String())

	return model.UserAuthFromString(res.Body.String())
}
