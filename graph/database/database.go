package database

import(

	"bufio"
    "fmt"
    "os"
    "strings"
    "syscall"
	"golang.org/x/crypto/ssh/terminal"
	
    _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	
	"github.com/baadjis/transferservice/graph/model"


)
var db *gorm.DB

func getCredentials() (string, string, error) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter databse Username: ")
    username, err := reader.ReadString('\n')
    if err != nil {
        return "", "", err
    }

    fmt.Print("Enter database Password: ")
    bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
    if err != nil {
        return "", "", err
    }

    password := string(bytePassword)
    return strings.TrimSpace(username), strings.TrimSpace(password), nil
}


func initDB() {
	var err error
	username,password ,_:=getCredentials()
    dataSourceName := username+":"+password+"@tcp(localhost:3306)/?parseTime=True"
    db, err := gorm.Open("mysql", dataSourceName)

    if err != nil {
        fmt.Println(err)
        panic("failed to connect database")
    }
    fmt.Println("succesfuly connected to mysql")
    db.LogMode(true)

    // Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("DROP DATABASE transfer_db")
	db.Exec("CREATE DATABASE transfer_db")
    
	db.Exec("USE transfer_db")
	

    // Migration to create tables for transfer
    db.AutoMigrate(&model.Transaction{}, &model.TransactionDetails{},&model.Customer{})	
  }