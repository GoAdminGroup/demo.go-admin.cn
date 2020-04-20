all: check remove-no-use combine compile

check:
	ifeq (, $(shell which adm))
		$(error "adm not found, consider doing: go install github.com/GoAdminGroup/go-admin/adm")
	endif

remove-no-use:
	find ./ -name ".DS_Store" -depth -exec rm {} \;

combine:
	adm combine js --src=./assets/src/js/combine/ --dist=./assets/login/dist/all.min.js
	adm combine css --src=./assets/src/css/ --dist=./assets/login/dist/all.min.css

compile:
	adm compile asset --src=./assets/login/dist/ --dist=./ --package=login

.PHONY: all check remove-no-use combine compile