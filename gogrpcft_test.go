package gogrpcft

import (
    "testing"
    "log"
    "context"
    "net"
    "os"
    "io/ioutil"
    "bytes"
    "path/filepath"

    "google.golang.org/grpc"
    "google.golang.org/grpc/test/bufconn"
)

var lis *bufconn.Listener

func init() {
    lis = bufconn.Listen(1024 * 1024)
    s := grpc.NewServer()
    RegisterFilesTransferServer(s)
    RegisterBytesTransferServer(s)
    go func() {
        if err := s.Serve(lis); err != nil {
            log.Fatalf("failed to listen: %v", err)
        }
    }()
}

func dialer(context.Context, string) (net.Conn, error) {
    return lis.Dial()
}

func TestUploadFile(t *testing.T) {

    // Create the files transfer client

    conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(dialer), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("gRPC connect failed: %v", err)
    }
    defer conn.Close()
    client := CreateTransferClient(conn)

    t.Run("successful download", func(t *testing.T) {

        // Create a temp file and upload dest

        content := make([]byte, 2048 * 1024 + 1024)

        tmpDir, err := ioutil.TempDir("", "test")
        if err != nil {
            t.Fatalf("failed to create temp remote directory: %v", err)
        }
        defer os.RemoveAll(tmpDir)

        srcPath := filepath.Join(tmpDir, "src_tmpfile")
        if err := ioutil.WriteFile(srcPath, content, 0666); err != nil {
            t.Fatalf("failed to create temp src file: %v", err)
        }

        uploadPath := filepath.Join(tmpDir, "upload_tempfile")

        // Perform upload

        if err := UploadFile(client, context.Background(), srcPath, uploadPath); err != nil {
            t.Fatalf("client failed: %v", err)
        }

        // Compare uploaded file to source

        srcf, err := ioutil.ReadFile(srcPath)
        if err != nil{
            t.Fatalf("failed to read src file: %v", err)
        }
        uploadf, err := ioutil.ReadFile(uploadPath)
        if err != nil{
            t.Fatalf("failed to read uploaded file: %v", err)
        }
        if !bytes.Equal(srcf, uploadf) {
            t.Fatalf("mismatch between remote and downloaded files")
        }
    })
}

func TestDownloadFile(t *testing.T) {

    // Create the files transfer client

    conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(dialer), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("gRPC connect failed: %v", err)
    }
    defer conn.Close()
    client := CreateTransferClient(conn)

    t.Run("successful download", func(t *testing.T) {

        // Create a temp file and download dest

        content := make([]byte, 2048 * 1024 + 1024)

        tmpDir, err := ioutil.TempDir("", "test")
        if err != nil {
            t.Fatalf("failed to create temp remote directory: %v", err)
        }
        defer os.RemoveAll(tmpDir)

        remotePath := filepath.Join(tmpDir, "remote_tmpfile")
        if err := ioutil.WriteFile(remotePath, content, 0666); err != nil {
            t.Fatalf("failed to create temp remote file: %v", err)
        }

        dstPath := filepath.Join(tmpDir, "dst_tempfile")

        // Perform download

        if err := DownloadFile(client, context.Background(), remotePath, dstPath); err != nil {
            t.Fatalf("client failed: %v", err)
        }

        // Compare downloaded file to source

        remotef, err := ioutil.ReadFile(remotePath)
        if err != nil{
            t.Fatalf("failed to read remote file: %v", err)
        }
        downloadedf, err := ioutil.ReadFile(dstPath)
        if err != nil{
            t.Fatalf("failed to read downloaded file: %v", err)
        }
        if !bytes.Equal(remotef, downloadedf) {
            t.Fatalf("mismatch between remote and downloaded files")
        }
    })
}

func TestUploadBytes(t *testing.T) {

    // Create the files transfer client

    conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(dialer), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("gRPC connect failed: %v", err)
    }
    defer conn.Close()
    client := CreateTransferClient(conn)

    t.Run("successful download", func(t *testing.T) {

        // Create a temp file and upload dest

        content := make([]byte, 2048 * 1024 + 1024)

        tmpDir, err := ioutil.TempDir("", "test")
        if err != nil {
            t.Fatalf("failed to create temp remote directory: %v", err)
        }
        defer os.RemoveAll(tmpDir)

        srcPath := filepath.Join(tmpDir, "src_tmpfile")
        if err := ioutil.WriteFile(srcPath, content, 0666); err != nil {
            t.Fatalf("failed to create temp src file: %v", err)
        }

        uploadPath := filepath.Join(tmpDir, "upload_tempfile")

        // Perform upload

        if err := UploadFile(client, context.Background(), srcPath, uploadPath); err != nil {
            t.Fatalf("client failed: %v", err)
        }

        // Compare uploaded file to source

        srcf, err := ioutil.ReadFile(srcPath)
        if err != nil{
            t.Fatalf("failed to read src file: %v", err)
        }
        uploadf, err := ioutil.ReadFile(uploadPath)
        if err != nil{
            t.Fatalf("failed to read uploaded file: %v", err)
        }
        if !bytes.Equal(srcf, uploadf) {
            t.Fatalf("mismatch between remote and downloaded files")
        }
    })
}

func TestDownloadBytes(t *testing.T) {

    // Create the bytes transfer client

    conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(dialer), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("gRPC connect failed: %v", err)
    }
    defer conn.Close()
    client := CreateTransferClient(conn)

    t.Run("successful download", func(t *testing.T) {

        // Create receiver/streamer

        content := make([]byte, 2048 * 1024 + 1024)

        receiver := &simpleBufferReceiver{}

        streamer := &simpleBufferStreamer{buffer: content}
        SetBytesTransferServerStreamer(

        // Perform download

        if err := DownloadBytes(client, context.Background(), "<unused>", "<unused>", receiver); err != nil {
            t.Fatalf("client failed: %v", err)
        }

        // Compare downloaded bytes to source

        if !bytes.Equal(content, receiver.buffer) {
            t.Fatalf("mismatch between remote and downloaded files")
        }
    })
}


type simpleBufferStreamer struct {
    buffer []byte
    index int
}

func (sbs *simpleBufferStreamer) Init(msg string) error {
    sbs.index = 0
}

func (sbs *simpleBufferStreamer) HasNext() bool {
    return sbs.index < len(sbs.buffer)
}

func (sbs *simpleBufferStreamer) GetNext() ([]byte, error) {
    if sbs.index + 1234 - 1 < len(sbs.buffer) {
        data := sbs.buffer[sbs.index:(sbs.index + 1234)]
        sbs.index += 1234
        return data, nil
    }
    reminder := len(sbs.buffer) - sbs.index
    data := sbs.buffer[sbs.index:(sbs.index + reminder)]
    sbs.index += reminder
    return data, nil
}

func (sbs *simpleBufferStreamer) Finalize() error {
    return nil
}


type simpleBufferReceiver struct {
    buffer []byte
}

func (sbr *simpleBufferReceiver) Init(msg string) error {
    sbr.buffer = make([]byte, 0)
    return nil
}

func (sbr *simpleBufferReceiver) Push(data []byte) error {
    sbr.buffer = append(sbr.buffer, data)
    return nil
}

func (sbr *simpleBufferReceiver) Finalize() error {
    return nil
}

