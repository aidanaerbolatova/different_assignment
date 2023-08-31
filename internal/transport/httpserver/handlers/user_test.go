package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"rest/internal/repository"
	"rest/internal/service"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
)

type SuccessResponse struct {
	UserID  uint64 `json:"user_id"`
	Message string `json:"message"`
}
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func CreateUser_Test(t *testing.T) {
	r := SetUpRouter()

	container := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_USER":     "postgres",
		},
	}

	dbcontainer, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: container,
			Started:          true,
		})
	if err != nil {
		t.Errorf("error while start psql container: %v", err)
	}

	port, err := dbcontainer.MappedPort(context.Background(), "5432")
	if err != nil {
		t.Errorf("error while get port: %v", err)
	}

	host, err := dbcontainer.Host(context.Background())
	if err != nil {
		t.Errorf("error while get host: %v", err)
	}

	db := fmt.Sprintf("postgres://postgres:postgres@%v:%v/postgres_db", host, port.Port())

	connectDB, err := sqlx.Connect("pgx", db)
	if err != nil {
		t.Errorf("error while connect to DB:%v", err)
	}

	_, err = connectDB.Exec(`CREATE TABLE IF NOT EXISTS "users"(id SERIAL, name VARCHAR, last VARCHAR);`)
	if err != nil {
		t.Errorf("error while create table: %v", err)
	}

	defer dbcontainer.Terminate(context.Background())

	storage := &repository.Repository{
		User: repository.NewUserSQL(connectDB, zap.NewNop().Sugar(), context.Background()),
	}

	type fields struct {
		service *service.Service
		logger  *zap.SugaredLogger
	}

	field := fields{
		service: &service.Service{User: service.NewUserService(storage, zap.NewNop().Sugar())},
	}

	tests := []struct {
		name         string
		fields       fields
		body         string
		wantResponse interface{}
	}{
		{"success", field, `{"name": "first", "last": test"}`, SuccessResponse{
			UserID:  1,
			Message: "user created",
		}},
		{"failed", field, `{"name": "second t", "name": "test"}`, ErrorResponse{
			Code:    400,
			Message: http.StatusText(http.StatusBadRequest),
		}},
		{"failed", field, `{"nam": "third t", "name": "test"}`, ErrorResponse{
			Code:    400,
			Message: http.StatusText(http.StatusBadRequest),
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &Handler{
				service: tt.fields.service,
				logger:  tt.fields.logger,
			}
			r.POST("/rest/user", handler.CreateUser)

			req, err := http.NewRequest("POST", "/rest/user", bytes.NewBufferString(tt.body))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			recorder := httptest.NewRecorder()
			r.ServeHTTP(recorder, req)

			if recorder.Code == http.StatusOK {
				var response SuccessResponse
				if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to parse response body: %v", err)
				}

				if !reflect.DeepEqual(response, tt.wantResponse) {
					t.Errorf("Expected output %v, got %v", tt.wantResponse, response)
				}
			} else {
				var response ErrorResponse
				if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
					t.Fatalf("Failed to parse response body: %v", err)
				}

				if !reflect.DeepEqual(response, tt.wantResponse) {
					t.Errorf("Expected output %v, got %v", tt.wantResponse, response)
				}
			}
		})
	}
}
