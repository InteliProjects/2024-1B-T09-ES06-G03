package types

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type CategoryStore interface {
	GetCategories() ([]Category, error)
	GetCategoryByID(id int) (*Category, error)
	CreateCategory(category *Category) error
	UpdateCategory(id int, category *Category) error
	DeleteCategory(id int) error
}

type SubcategoryStore interface {
	GetSubcategories() ([]Subcategory, error)
	GetSubcategoryByID(id int) (*Subcategory, error)
	GetSubcategoriesByCategory(categoryID int) ([]Subcategory, error)
	CreateSubcategory(subcategory *Subcategory) error
	UpdateSubCategory(id int, subcategory *Subcategory) error
	DeleteSubcategory(id int) error
}

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Company     string `json:"company"`
	Instagram   string `json:"instagram"`
	Linkedin    string `json:"linkedin"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Subcategory struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
}

type RegisterUserPayload struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=3,max=130"`
	Company     string `json:"company" validate:"required"`
	Instagram   string `json:"instagram"`
	Linkedin    string `json:"linkedin"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CategoryPayload struct {
	Name string `json:"name" validate:"required"`
}

type SubcategoryPayload struct {
	Name       string `json:"name" validate:"required"`
	CategoryID int    `json:"category_id" validate:"required"`
}
