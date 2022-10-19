package todo

type TodoList struct {
	Id          int    `json:"-" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
}

type TodoItem struct {
	Id          int    `json:"-" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
	Importance  int    `json:"importance" db:"importance"`
}

type ListItems struct {
	Id     int
	ListId int
	ItemId int
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

// struct for list updating (if some field not in request -> get nil value)
type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// struct for items updating (if some field not in request -> get nil value)
type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
	Importance  *int    `json:"importance"`
}
