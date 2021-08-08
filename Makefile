install: main.go utils/
	@sh utils/install.sh

clean: /usr/local/bin/wanty
	@rm /usr/local/bin/wanty
