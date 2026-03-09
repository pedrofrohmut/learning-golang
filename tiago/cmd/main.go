package main

func main() {
	var conf = Config {
		addr: ":8080",
		db: dbConfig {},
	}

	var app = Application {
		config: config,
	}
}
