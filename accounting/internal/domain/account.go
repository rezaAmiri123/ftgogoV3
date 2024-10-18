package domain
import "github.com/stackus/errors"
type Account struct {
    ID string
    Name string
    Enabled bool
} 

func RegisterAccount(id, name string)(*Account, error){
    if id == ""{

    }
}