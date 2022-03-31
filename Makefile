check:
	@yq -n '.someNew="content"' > newfile.yml
	@cat newfile.yaml
