# lavt-frontend

the frontend uses `bun` as a runtime with `vite` to build the vanilla html/css/js

when the app is built, there are two cmds that get run in this frontend folder (see README.md):

- frontend:install
- frontend:build

these commands take care of building the frontend to the `/dist` directory. which is then passed along to the `embed.FS` call in the main go package.

everything else is handled by the go application.

### tech stack

- vanilla html/css/js
- [tailwind](https://tailwindcss.com/)
- [htmx](https://htmx.org/)
