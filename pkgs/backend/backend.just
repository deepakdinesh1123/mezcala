backend-api-docs:
    npx @redocly/cli@1.25.11 preview-docs -p 8048 pkgs/backend/spec/api.yml

generate-backend-api:
    cd pkgs/backend && go generate ./spec

lint-backend-api:
	npx @redocly/cli@1.25.11 lint --config pkgs/backend/spec/redocly.yml pkgs/backend/spec/api.yml

compile-backend:
	cd pkgs/backend && go build -o ../../mezb