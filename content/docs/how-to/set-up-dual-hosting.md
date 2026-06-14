---
title: Serve on Pages and Studio at once
layout: doc
summary: Public GitHub Pages and studio use different base URLs. Build twice, publish to two branches, and each host serves the right one.
nav_section: "How-to Guides"
nav_weight: 3
tags: [how-to, deployment, hosting]
---

A site can be **both** a public GitHub Pages site *and* an in-place-editable studio
site. The catch: the two hosts mount at different paths, so they need different
`base_url` builds.

| Host | Mount path | `base_url` | Branch |
|------|-----------|------------|--------|
| GitHub Pages (project site) | `/<repo>` | `/<repo>` | `gh-pages` |
| fuego-studio | `/<owner>/<repo>` | `/<owner>/<repo>` | `fuego-pages` |

## Build twice

```yaml
      - name: Build for fuego-studio
        run: go run . build --base-url /<owner>/<repo>
      - name: Publish studio build to fuego-pages
        uses: peaceiris/actions-gh-pages@v4
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          publish_dir: build
          publish_branch: fuego-pages

      - name: Build for GitHub Pages
        run: go run . build --base-url /<repo>
      - name: Publish public build to gh-pages
        uses: peaceiris/actions-gh-pages@v4
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          publish_dir: build
          publish_branch: gh-pages
```

## Why studio prefers `fuego-pages`

When resolving a site's source, studio checks `fuego-pages` **first**, then the
Pages source, then `gh-pages`. So the studio-base build on `fuego-pages` is what
studio serves, while GitHub Pages independently serves the `/<repo>`-base build on
`gh-pages`. One source repo, two correct hosts.

> Private repos can't use public GitHub Pages — for those, build once to
> `fuego-pages` and serve only through studio.
