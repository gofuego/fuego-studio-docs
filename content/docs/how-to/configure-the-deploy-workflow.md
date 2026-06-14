---
title: Configure the deploy workflow
layout: doc
summary: A complete GitHub Actions workflow that builds your site at the studio mount path and publishes it to the serve branch.
nav_section: "How-to Guides"
nav_weight: 2
tags: [how-to, deployment, ci]
---

Studio serves what your CI publishes to a serve branch. Here's a complete workflow
for a Fuego site.

## The workflow

```yaml
name: Deploy
on:
  push:
    branches: [main]
permissions:
  contents: write
jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: "1.25"
      - name: Build for fuego-studio
        run: go run . build --base-url /<owner>/<repo>
      - name: Publish to fuego-pages
        uses: peaceiris/actions-gh-pages@v4
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          publish_dir: build
          publish_branch: fuego-pages
```

## What matters

- **`--base-url /<owner>/<repo>`** — studio mounts each site under its
  `owner/repo` path, so the build must use that prefix or links and assets break.
- **`fuego-pages`** — studio prefers this branch. Publishing here keeps your studio
  build separate from any public GitHub Pages build (see
  [Set up dual hosting](/docs/how-to/set-up-dual-hosting/)).
- **No studio secrets needed** — the workflow only uses the built-in
  `secrets.GITHUB_TOKEN`. Studio's own credentials live on the studio host, not in
  your repo.

## After it runs

The push fires a webhook to studio, which re-syncs and serves the new build. If you
had a page open, it reloads itself.
