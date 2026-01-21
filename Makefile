.PHONY: generate lint clean deps

generate:
	buf generate

lint:
	buf lint

clean:
	rm -rf api/ docs/

deps:
	buf dep update
