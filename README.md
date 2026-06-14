# fuego-studio-docs

Public, user-facing documentation for [fuego-studio](https://github.com/gofuego/fuego-studio) —
the authenticating host and in-place editor for [Fuego](https://github.com/gofuego/fuego) sites.

**[Read the docs](https://gofuego.github.io/fuego-studio-docs/)**

It's itself a Fuego site (dogfooding), using the shared
[public docs theme pack](https://github.com/gofuego/fuego-doctheme) and following the
[Diátaxis](https://diataxis.fr/) structure: tutorials, how-to guides, reference, and
explanation.

## Develop

```bash
go run . serve      # dev server with live rebuild
go run . build      # one-shot build into ./build
```

## Deploy

Dual-hosted (see `.github/workflows/deploy.yml`): built twice and published to two
branches —

- `gh-pages` (base URL `/fuego-studio-docs`) → public **GitHub Pages**.
- `fuego-pages` (base URL `/gofuego/fuego-studio-docs`) → **fuego-studio**, where the
  docs are editable in place.
