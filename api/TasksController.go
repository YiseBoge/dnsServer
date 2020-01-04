package api

type API int

func (api *API) Lookup(name string, address *string) error {
	print("abebe")
	*address = name

	client := ParentClient()
	var result string
	err := client.Call("API.Lookup", "abebe", &result)
	if err != nil {
		return err
	}
	return nil
}
