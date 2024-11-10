.PHONY: gen genproto genopenapi

gen: genproto genopenapi

genproto:
	@./scripts/genproto.sh

genopenapi:
	@./scripts/genopenapi.sh
