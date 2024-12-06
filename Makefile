#comment
check:
	@yq --version
	@echo '{"foo": "bar"}' | yq e '.foo' -
