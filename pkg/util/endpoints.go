package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/neoxelox/kit"
	"github.com/rs/xid"

	"service/pkg/config"
)

type HealthEndpoints struct {
	config   config.Config
	observer *kit.Observer
	database *kit.Database
	cache    *kit.Cache
}

func NewHealthEndpoints(observer *kit.Observer, database *kit.Database, cache *kit.Cache, config config.Config) *HealthEndpoints {
	return &HealthEndpoints{
		config:   config,
		observer: observer,
		database: database,
		cache:    cache,
	}
}

type HealthEndpointsGetHealthResponseItem struct {
	Error   *error `json:"error"`
	Latency int64  `json:"latency"`
}

type HealthEndpointsGetHealthResponse struct {
	Database HealthEndpointsGetHealthResponseItem `json:"database"`
	Cache    HealthEndpointsGetHealthResponseItem `json:"cache"`
}

func (self *HealthEndpoints) GetServerHealth(ctx echo.Context) error {
	response := HealthEndpointsGetHealthResponse{}

	start := time.Now()
	errD := self.database.Health(ctx.Request().Context())
	response.Database.Latency = time.Since(start).Milliseconds()
	if errD != nil {
		response.Database.Error = &errD
	} else {
		response.Database.Error = nil
	}

	start = time.Now()
	errC := self.cache.Health(ctx.Request().Context())
	response.Cache.Latency = time.Since(start).Milliseconds()
	if errC != nil {
		response.Cache.Error = &errC
	} else {
		response.Cache.Error = nil
	}

	if errD != nil || errC != nil {
		return ctx.JSON(http.StatusServiceUnavailable, &response)
	}

	return ctx.JSON(http.StatusOK, &response)
}

func (self *HealthEndpoints) GetWorkerHealth(res http.ResponseWriter, req *http.Request) {
	response := HealthEndpointsGetHealthResponse{}

	start := time.Now()
	errD := self.database.Health(req.Context())
	response.Database.Latency = time.Since(start).Milliseconds()
	if errD != nil {
		response.Database.Error = &errD
	} else {
		response.Database.Error = nil
	}

	start = time.Now()
	errC := self.cache.Health(req.Context())
	response.Cache.Latency = time.Since(start).Milliseconds()
	if errC != nil {
		response.Cache.Error = &errC
	} else {
		response.Cache.Error = nil
	}

	res.Header().Set("Content-Type", "application/json")
	if errD != nil || errC != nil {
		res.WriteHeader(http.StatusServiceUnavailable)
	} else {
		res.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(res).Encode(&response) // nolint:errcheck,errchkjson
}

type FileEndpoints struct {
	config   config.Config
	observer *kit.Observer
}

func NewFileEndpoints(observer *kit.Observer, config config.Config) *FileEndpoints {
	return &FileEndpoints{
		config:   config,
		observer: observer,
	}
}

type FileEndpointsGetFileRequest struct {
	Name string `param:"name"`
}

func (self *FileEndpoints) GetFile(ctx echo.Context) error {
	request := FileEndpointsGetFileRequest{}

	err := ctx.Bind(&request)
	if err != nil {
		return kit.HTTPErrInvalidRequest.Cause(err)
	}

	return ctx.File(fmt.Sprintf("%s/%s", self.config.Service.FilesPath, request.Name))
}

type FileEndpointsPostFileResponse struct {
	URI string `json:"uri"`
}

func (self *FileEndpoints) PostFile(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return kit.HTTPErrInvalidRequest.Cause(err)
	}

	src, err := file.Open()
	if err != nil {
		return kit.HTTPErrInvalidRequest.Cause(err)
	}
	defer src.Close()

	fileName := fmt.Sprintf("%s%s", xid.New().String(), filepath.Ext(file.Filename))
	filePath := fmt.Sprintf("%s/%s", self.config.Service.FilesPath, fileName)

	dst, err := os.Create(filePath)
	if err != nil {
		return kit.HTTPErrServerGeneric.Cause(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return kit.HTTPErrServerGeneric.Cause(err)
	}

	response := FileEndpointsPostFileResponse{}
	response.URI = fmt.Sprintf("/file/%s", fileName)

	return ctx.JSON(http.StatusOK, &response)
}
