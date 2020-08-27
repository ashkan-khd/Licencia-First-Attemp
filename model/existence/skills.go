package existence

type Field struct {
	ID     string   `json:"id" binding:"min=6,max=10" pg:"id,pk"`
	Name   string   `json:"name" pg:"name"`
	Skills []string `json:"skills" pg:"skills"`
}
