package standard

type DatabaseRepository interface {
	Select(selector interface{}, output interface{}) error
	Update(selector interface{}, updater interface{}) error
	Insert(payload interface{}) error
	// Delete() error
}
