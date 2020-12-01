package model



import(
	"time"
	"math/rand"
	"encoding/hex"
	
	"github.com/jinzhu/gorm"
)


type BaseModel struct{
	ID        string  `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time 

}


func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	b := make([]byte, 4) //equals 8 charachters
    rand.Read(b) 
    s := hex.EncodeToString(b)
	
	return scope.SetColumn("ID", s)
   }