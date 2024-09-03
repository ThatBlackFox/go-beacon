package main

func main() {
	config := read_config()
	dbc := DBConstructor{DBurl: config.DBurl, ConfigPath: "creds.json", AuthOverride: config.AuthOverride}
	utils := dbc.make()

}
