package core

import (
	"github.com/evenfound/even-go/server/api"

	"errors"
	"strings"

	"github.com/evenfound/even-go/pb"

	"github.com/ipfs/interface-go-ipfs-core"
	//"github.com/ipfs/go-ipfs/core/coreapi/interface"
	//dag "github.com/ipfs/go-merkledag"
	"github.com/ipfs/go-mfs"
	//ft "github.com/ipfs/go-unixfs"
)

func GetFileByHash(hash string) (iface.Reader, error) {

	// TODO: commented mode refactoring

	//fmt.Println("Processing hash " + hash)
	//
	//path, err := iface.ParsePath(hash)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//cApi, _ := coreapi.NewCoreAPI(Node.IpfsNode)
	//
	//fs := cApi.Unixfs()
	//
	//ctx := context.Background()
	//
	//reader, err := fs.Cat(ctx, path)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//return reader, nil
	return nil, errors.New("GetFileByHash not implemented")
}

func CreateNewFile(name string, data []byte) error {
	// TODO: commented mode refactoring

	//cnts := strings.Split(name, "/")
	//
	//if len(cnts) > 2 {
	//	return errors.New("Invalid file path")
	//}
	//
	//var (
	//	directory = Node.IpfsNode.FilesRoot.GetValue().(*mfs.Directory)
	//
	//	fname string
	//)
	//
	//if len(cnts) > 1 {
	//
	//	dirName := cnts[0]
	//
	//	dirLookup, err := mfs.Lookup(Node.IpfsNode.FilesRoot, dirName)
	//
	//	if err != nil {
	//		return err
	//	}
	//	var ok bool
	//
	//	directory, ok = dirLookup.(*mfs.Directory)
	//
	//	if !ok {
	//		return errors.New("Directory not found")
	//	}
	//
	//	fname = cnts[1]
	//
	//} else {
	//	fname = cnts[0]
	//}
	//
	//nrd := dag.NodeWithData(ft.FilePBData(nil, 0))
	//
	//prfx := dag.V0CidPrefix()
	//
	//nrd.SetPrefix(&prfx)
	//
	//err := directory.AddChild(fname, nrd)
	//
	//if err != nil {
	//	return err
	//}
	//
	//fNode, err := directory.Child(fname)
	//
	//if err != nil {
	//	return err
	//}
	//
	//file, ok := fNode.(*mfs.File)
	//
	//if !ok {
	//	return errors.New("File not found or has not been created")
	//}
	//
	//wfd, err := file.Open(mfs.OpenReadWrite, true)
	//
	//if err != nil {
	//	return err
	//}
	//
	//defer wfd.Close()
	//
	//wfd.Seek(0, 0)
	//
	//wfd.Truncate(0)
	//
	//wfd.Write(data)
	//
	//return nil
	return errors.New("CreateNewFile not implemented")
}

func CreateDirectory(name string) error {
	return mfs.Mkdir(Node.IpfsNode.FilesRoot, name, mfs.MkdirOpts{
		Flush: true,
	})
}

func FileStat(path string) (*api.FileStatResponse, error) {

	cnts := strings.Split(path, "/")

	if len(cnts) > 2 {
		return &api.FileStatResponse{}, errors.New("Invalid path")
	}

	dirName := cnts[0]

	dirLookup, err := mfs.Lookup(Node.IpfsNode.FilesRoot, dirName)

	if err != nil {
		return &api.FileStatResponse{}, err
	}

	dirNode, err := dirLookup.GetNode()

	if err != nil {
		return &api.FileStatResponse{}, err
	}

	state, err := dirNode.Stat()

	if err != nil {
		return &api.FileStatResponse{}, err
	}

	numLinks := int64(state.NumLinks)
	dataSize := int64(state.DataSize)
	linkSize := int64(state.LinksSize)
	cumulativeSize := int64(state.CumulativeSize)

	if len(cnts) == 1 {
		return &api.FileStatResponse{
			Cid:            dirNode.Cid().String(),
			NumLinks:       numLinks,
			CumulativeSize: cumulativeSize,
			LinksSize:      linkSize,
			DataSize:       dataSize,
			Type:           api.FileStatResponse_Directory,
		}, nil
	}

	fname := cnts[1]

	directory, ok := dirLookup.(*mfs.Directory)

	if !ok {
		return &api.FileStatResponse{}, errors.New("Invalid path")
	}

	file, err := directory.Child(fname)

	if err != nil {
		return &api.FileStatResponse{}, err
	}

	fileNode, err := file.GetNode()

	if err != nil {
		return &api.FileStatResponse{}, err
	}

	fDesc, ok := file.(*mfs.File)

	fDescNode, _ := fDesc.GetNode()

	state, err = fDescNode.Stat()

	if err != nil {
		return &api.FileStatResponse{}, err
	}

	numLinks = int64(state.NumLinks)
	dataSize = int64(state.DataSize)
	linkSize = int64(state.LinksSize)
	cumulativeSize = int64(state.CumulativeSize)

	return &api.FileStatResponse{
		Cid:            fileNode.Cid().KeyString(),
		NumLinks:       numLinks,
		CumulativeSize: cumulativeSize,
		LinksSize:      linkSize,
		DataSize:       dataSize,
		Type:           api.FileStatResponse_File,
	}, nil

}

func ProcessMessage(msg *pb.Message) error {

	// TODO: commented mode refactoring

	// First get file data from hash

	//data := msg.GetPayload().GetValue()
	//
	//fmt.Println(data)
	//
	//fHash := string(data[2:])
	//
	//fmt.Printf("Received hash %v \n", fHash)
	//
	//path, err := coreapi.ParsePath(fHash)
	//
	//if err != nil {
	//	return err
	//}
	//
	//cApi := coreapi.NewCoreAPI(Node.IpfsNode)
	//
	//fs := cApi.Unixfs()
	//
	//ctx := context.Background()
	//
	//reader, err := fs.Cat(ctx, path)
	//
	//if err != nil {
	//	return err
	//}
	//
	//defer reader.Close()
	//
	//f, err := os.OpenFile(fHash, os.O_CREATE|os.O_RDWR, 0644)
	//
	//if err != nil {
	//	return err
	//}
	//
	//defer f.Close()
	//
	//io.Copy(f, reader)
	//
	//return nil

	return errors.New("ProcessMessage not imlemented")
}
