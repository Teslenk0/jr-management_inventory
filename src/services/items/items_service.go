package items

//Interface with methods
type itemsServiceInterface interface {
}

//Struct
type itemsService struct {
}

//Implementing the interface
var (
	UsersService itemsServiceInterface = &itemsService{}
)
