type Config struct {
	env: "local" // local, dev, prod
	storage_path: "./storage/storage.db"
	http_server:
		address: "localhost:8082"
		timeout: 4s # read query & response timeout
		idle_timeout: 60s # client-server connection lifespan
}
