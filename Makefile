all:
	make -C google-api install
	make -C google-api-go-generator install
	google-api-go-generator/google-api-go-gen -cache -install -api=*
