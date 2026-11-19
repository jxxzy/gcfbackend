package gcfbackend

import (
	"context"
	"fmt"
	"os"
	"strings"
	pasproj "github.com/HRMonitorr/PasetoprojectBackend"
	"github.com/whatsauth/watoken"
	"testing"
)

//func TestGCHandlerFunc(t *testing.T) {
//	data := GCHandlerFunc("string", "GIS", "geogis")
//
//	fmt.Printf("%+v", data)
//}

func TestGeneratePaseto(t *testing.T) {
	//privateKey, publicKey := watoken.GenerateKey()
	//fmt.Println(privateKey)
	//fmt.Println(publicKey)
	hasil, err := watoken.Encode("rofiasoy", "PrivateKey")
	fmt.Println(hasil, err)
}

// integrationDB requires an explicit opt-in and a disposable test database.
func integrationDB(t *testing.T) string {
	t.Helper()
	if os.Getenv("GCFBACKEND_RUN_INTEGRATION") != "1" {
		t.Skip("MongoDB integration test disabled")
	}
	db := os.Getenv("GCFBACKEND_TEST_DB")
	if !strings.HasPrefix(db, "test_") || len(db) <= len("test_") {
		t.Fatal("GCFBACKEND_TEST_DB must be a disposable database beginning with test_")
	}
	if os.Getenv("GCFBACKEND_TEST_MONGO") == "" {
		t.Fatal("GCFBACKEND_TEST_MONGO is required")
	}
	return db
}

func TestUpdateData(t *testing.T) {
	db := integrationDB(t)
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "cihuy",
		Volume: "1",
	}
	up := UpdateNameGeo("GCFBACKEND_TEST_MONGO", db, context.Background(), data)
	fmt.Println(up)
}

func TestDeleteDataGeo(t *testing.T) {
	db := integrationDB(t)
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "cihuy",
		Volume: "1",
	}
	up := DeleteDataGeo("GCFBACKEND_TEST_MONGO", db, context.Background(), data)
	fmt.Println(up)
}

func TestInsertUser(t *testing.T) {
	db := integrationDB(t)
	conn := GetConnectionMongo("GCFBACKEND_TEST_MONGO", db)
	pass, _ := pasproj.HashPass("rofiganteng")
	data := RegisterStruct{
		Username: "rofiasoy",
		Password: pass,
	}
	ins := InsertUserdata(conn, data.Username, data.Password)
	fmt.Println(ins)
}

func TestIsexist(t *testing.T) {
	data := IsExist("token", "PublicKey")
	fmt.Println(data)
}
