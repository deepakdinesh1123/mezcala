set working-directory := "pkgs/mezui"

dev:
    pnpm run dev

install:
    pnpm i

add package:
    pnpm add {{package}}

remove package:
    pnpm remove {{package}}

audit:
    pnpm audit

add-shadcn component:
    pnpm dlx shadcn@latest add {{component}}

pretty:
    npx prettier --write ./

generate-frontend:
    openapi-generator-cli generate -g typescript-axios -i ../backend/spec/api.yml -o /pkgs/mezui/oas-client
