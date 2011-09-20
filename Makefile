all:
	make -C google-api clean install
	make -C google-api-go-generator clean install
	google-api-go-generator/google-api-go-gen -cache=false -install -api=*
