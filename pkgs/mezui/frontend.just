
[working-directory: 'pkgs/mezui']
pwd-c:
    pwd

[working-directory: 'pkgs/mezui']
dev:
    pnpm run dev

[working-directory: 'pkgs/mezui']
build:
    pnpm run build

[working-directory: 'pkgs/mezui']
install:
    pnpm i

[working-directory: 'pkgs/mezui']
add package:
    pnpm add {{package}}

[working-directory: 'pkgs/mezui']
remove package:
    pnpm remove {{package}}

[working-directory: 'pkgs/mezui']
audit:
    pnpm audit

[working-directory: 'pkgs/mezui']
add-shadcn component:
    pnpm dlx shadcn@latest add {{component}}

[working-directory: 'pkgs/mezui']
pretty:
    npx prettier --write ./

[working-directory: 'pkgs/mezui']
generate-frontend-dev:
    openapi-generator-cli generate -g typescript-axios -i ../backend/spec/api.yml -o ./src/oas-client --openapi-normalizer SET_TAGS_FOR_ALL_OPERATIONS=Default
