package handlers

import (
	"context"
	"fmt"
	"github.com/evenfound/even-go/core"
	"github.com/evenfound/even-go/server/api"
	"gx/ipfs/QmVmDhyTTUcQXFD1rRQ64fGLMSAoaQvNH3hwuaCFAPq2hy/errors"
	"io"
	"io/ioutil"
	"os"
)

type FilesHandler struct {
}

func (handler *FilesHandler) GetFileByHash(ctx context.Context, req *api.FileRequest) (*api.FileResponse, error) {

	var (
		fname  = req.Cid
		output = req.Output
	)

	reader, err := core.GetFileByHash(fname)

	if err != nil {
		return &api.FileResponse{}, err
	}

	f, err := os.OpenFile(output, os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		return &api.FileResponse{}, errors.New(fmt.Sprintf("Error  : %v", err))
	}

	_, err = io.Copy(f, reader)

	if err != nil {
		return &api.FileResponse{}, errors.New(fmt.Sprintf("Error  : %v", err))
	}

	return &api.FileResponse{}, nil
}

func (handler *FilesHandler) Create(_ context.Context, req *api.FileCreateRequest) (*api.FileResponse, error) {

	path := req.GetFname()
	source := req.GetSource()

	f, err := ioutil.ReadFile(source)

	if err != nil {
		return &api.FileResponse{}, err
	}

	err = core.CreateNewFile(path, f)

	if err != nil {
		return &api.FileResponse{}, err
	}

	return &api.FileResponse{}, nil
}

func (handler *FilesHandler) Mkdir(ctx context.Context, req *api.FileMkdirRequest) (*api.FileDataResponse, error) {

	var dirName = req.GetDirName()

	err := core.CreateDirectory(dirName)

	if err != nil {
		return &api.FileDataResponse{}, err
	}

	return &api.FileDataResponse{}, nil
}

func (handler *FilesHandler) Stat(ctx context.Context, req *api.FileStateRequest) (*api.FileStatResponse, error) {

	return core.FileStat(req.GetPath())
}
