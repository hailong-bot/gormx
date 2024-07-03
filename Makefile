GOPROXY=GOPROXY=https://goproxy.cn,direct

mod_tidy:
	@$(GOPROXY) go mod tidy
