package rpc

import (
	"context"
	"fmt"
	"github.com/evenfound/even-go/node/cmd/evenctl/config"
	"github.com/evenfound/even-go/node/server/api"

	"google.golang.org/grpc"
)

type FileCMD struct {
	ctx  context.Context
	stub api.FileServiceClient
}

func NewFileCMD() (*FileCMD, error) {

	ctx := context.Background()

	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	stub := api.NewFileServiceClient(conn)

	return &FileCMD{
		ctx:  ctx,
		stub: stub,
	}, nil

}

func (cmd *FileCMD) GetFileByHash(filename, output string) error {

	_, err := cmd.stub.GetFileByHash(cmd.ctx, &api.FileRequest{
		Cid:    filename,
		Output: output,
	})

	if err == nil {
		fmt.Println("Success")
	} else {
		fmt.Printf("Error : %v\n", err)
	}

	return nil

}

func (cmd *FileCMD) Mkdir(dirName string) error {

	req := &api.FileMkdirRequest{
		DirName: dirName,
	}

	res, err := cmd.stub.Mkdir(cmd.ctx, req)

	if err != nil {
		return err
	}

	fmt.Printf("Created directory with name %v  hash %v\n", dirName, res.Cid)

	return nil

}


func (cmd *FileCMD) Create(source, destination string) error {

	req := &api.FileCreateRequest{
		Fname:  destination,
		Source: source,
	}

	_, err := cmd.stub.Create(cmd.ctx, req)

	fmt.Println(err)

	return nil
}

func (cmd *FileCMD) Stat(path string) error {

	fmt.Println(path)

	req := &api.FileStateRequest{
		Path: path,
	}

	res, err := cmd.stub.Stat(cmd.ctx, req)

	if err != nil {
		return err
	}

	fmt.Printf("Hash %v\n", res.GetCid())
	fmt.Printf("BlockSize %v\n", res.GetBlockSize())
	fmt.Printf("CumulativeSize %v\n", res.GetCumulativeSize())
	fmt.Printf("DataSize %v\n", res.GetDataSize())
	fmt.Printf("NumLinks %v\n", res.GetNumLinks())
	fmt.Printf("LinksSize %v\n", res.GetLinksSize())
	fmt.Printf("Type %v\n", res.GetType().String())

	return nil
}
