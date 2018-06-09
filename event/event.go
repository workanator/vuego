package event

type Event struct {
	Category Category    `json:"category"`
	Target   string      `json:"target"`
	Name     string      `json:"name"`
	Data     interface{} `json:"data,omitempty"`
}

func (e *Event) Conforms(target, name string) bool {
	return e.Target == target && e.Name == name
}

// Event categories
const (
	CategorySystem Category = "system"
	CategoryModel  Category = "model"
	CategoryDom    Category = "dom"
	CategoryUser   Category = "user"
)

type Category string

// Test if the event category is system
func (ec Category) IsSystem() bool {
	return ec == CategorySystem
}

// Test if the event category is model
func (ec Category) IsModel() bool {
	return ec == CategoryModel
}

// Test if the event category is DOM
func (ec Category) IsDom() bool {
	return ec == CategoryDom
}

// Test if the event category is user
func (ec Category) IsUser() bool {
	return ec == CategoryUser
}

// Test if the event category is custom
func (ec Category) IsCustom() bool {
	return ec != CategorySystem && ec != CategoryModel && ec != CategoryDom && ec != CategoryUser
}
