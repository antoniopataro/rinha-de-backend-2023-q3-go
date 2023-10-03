package people

type Person struct {
	Birthdate string   `json:"nascimento"`
	ID        string   `json:"id"`
	Name      string   `json:"nome"`
	Nickname  string   `json:"apelido"`
	Stack     []string `json:"stack"`
}
