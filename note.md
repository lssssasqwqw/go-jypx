data, _ := ioutil.ReadAll(c.Request.Body)
var c_json map[string]interface{}
err := json.Unmarshal([]byte(string(data)), &c_json)
if err != nil {
	panic(err)
}