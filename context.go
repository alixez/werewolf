package werewolf

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/alixez/werewolf/utils"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
)

type (
	Context struct {
		echo.Context
		services    map[string]ServiceInterface
		apiResponse *APIResponse
		dbHelper    map[string]interface{}
		Config      *Env
	}

	File struct {
		Filename     string
		Path         string
		AbstructPath string
		Host         string
		Extension    string
	}
)

func (this *Context) SaveFilesToStorage(fields string, subpath string) ([]*File, error) {
	var fileList []*File
	form, err := this.MultipartForm()
	if err != nil {
		return nil, err
	}

	files := form.File[fields]

	for _, file := range files {
		if fileModel, err := this.executeUploadedFile(file, subpath); err == nil {
			fileList = append(fileList, fileModel)
		} else {
			return nil, err
		}

	}

	return fileList, nil
}

func (this *Context) SaveFileToStorage(fields string, subpath string) (*File, error) {
	var fileModel *File

	file, err := this.FormFile(fields)
	if err != nil {
		return nil, err
	}
	fileModel, err = this.executeUploadedFile(file, subpath)
	if err != nil {
		return nil, err
	}
	return fileModel, nil
}

func (this *Context) executeUploadedFile(file *multipart.FileHeader, subpath string) (*File, error) {
	config := this.Config
	storage := config.GetConfig("storage").(map[string]string)
	rootPath := storage["root"]
	tumbnailPath := filepath.Join(rootPath, storage["tumbnail"])
	orignailPath := filepath.Join(rootPath, storage["orignail"])
	mimeType := file.Header["Content-Type"][0]
	filename := uuid.NewV1().String() + "." + strings.Split(mimeType, "/")[1]
	dstPath := filepath.Join(orignailPath, subpath)

	if !utils.IsDirExist(rootPath) {
		os.Mkdir(rootPath, 0666)
	}
	if !utils.IsDirExist(tumbnailPath) {
		os.Mkdir(tumbnailPath, 0666)
	}
	if !utils.IsDirExist(orignailPath) {
		os.Mkdir(orignailPath, 0666)
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()
	if !utils.IsDirExist(dstPath) {
		os.MkdirAll(dstPath, 0666)
	}
	dst, err := os.Create(filepath.Join(dstPath, filename))
	if err != nil {
		return nil, err
	}

	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return nil, err
	}

	absPath, err := filepath.Abs(filepath.Join(dstPath, filename))
	if err != nil {
		return nil, err
	}

	fileModel := &File{
		Filename:     filename,
		Path:         filepath.Join(dstPath, filename),
		AbstructPath: absPath,
		Host:         storage["host"],
		Extension:    strings.Split(mimeType, "/")[1],
	}

	return fileModel, nil
}

func (this *Context) AddDBHelper(name string, value interface{}) {
	this.dbHelper[name] = value
}

func (this *Context) GetDB(name string) interface{} {
	return this.dbHelper[name]
}

func (this *Context) SetServices(services map[string]ServiceInterface) {

	this.services = services
}

func (this *Context) GetService(name string) ServiceInterface {
	service := this.services[name]
	service.Init(this)
	return service
}
