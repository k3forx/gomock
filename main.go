package main

type ApiClient interface {
	Request(string) (string, error)
}
type DataRegister struct {
	Client ApiClient
}

func (d *DataRegister) Register(data string) (string, error) {
	result, err := d.Client.Request(data)
	if err != nil {
		return "", err
	}
	return result, nil
}

func main() {
}
