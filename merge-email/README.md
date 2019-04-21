go run main.go moby-moby-gitdata.dump kubernetes-kubernetes-gitdata.dump

ENVS

	if os.Getenv("API_USER") != "" {
		API_USER = os.Getenv("API_USER")
	}
	if os.Getenv("FROM") != "" {
		FROM = os.Getenv("FROM")
	}
	if os.Getenv("FROM_USER") != "" {
		FROM_USER = os.Getenv("FROM_USER")
	}
	if os.Getenv("KEY") != "" {
		KEY = os.Getenv("KEY")
	}
