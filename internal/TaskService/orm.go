package TaskService

type RequestBody struct {
	ID      uint   `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Task    string `json:"task"`
	Is_done bool   `json:"is-done"`
}
