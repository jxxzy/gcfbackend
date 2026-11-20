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

// mongoTestDatabase requires explicit opt-in and a disposable database name.
func mongoTestDatabase(t *testing.T) string {
	t.Helper()
	if os.Getenv("RUN_MONGO_INTEGRATION_TESTS") != "1" {
		t.Skip("Set RUN_MONGO_INTEGRATION_TESTS=1 to run MongoDB integration tests")
	}
	if os.Getenv("MONGO") == "" {
		t.Fatal("MONGO must point to a disposable test MongoDB instance")
	}
	db := os.Getenv("MONGO_TEST_DB")
	if !strings.HasPrefix(db, "test_") || len(db) <= len("test_") {
		t.Fatal("MONGO_TEST_DB must be a disposable database starting with test_")
	}
	return db
}

func TestGeneratePaseto(t *testing.T) {
	//privateKey, publicKey := watoken.GenerateKey()
	//fmt.Println(privateKey)
	//fmt.Println(publicKey)
	hasil, err := watoken.Encode("rofiasoy", "PrivateKey")
	fmt.Println(hasil, err)
}

func TestUpdateData(t *testing.T) {
	db := mongoTestDatabase(t)
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "cihuy",
		Volume: "1",
	}
	up := UpdateNameGeo("MONGO", db, context.Background(), data)
	fmt.Println(up)
}

func TestDeleteDataGeo(t *testing.T) {
	db := mongoTestDatabase(t)
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "cihuy",
		Volume: "1",
	}
	up := DeleteDataGeo("MONGO", db, context.Background(), data)
	fmt.Println(up)
}

func TestInsertUser(t *testing.T) {
	db := mongoTestDatabase(t)
	conn := GetConnectionMongo("MONGO", db)
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
