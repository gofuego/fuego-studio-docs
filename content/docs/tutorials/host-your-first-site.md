---
title: Host your first site
layout: doc
summary: Install the app on a repo, point its build at the studio mount path, and watch the gated site come up.
nav_section: "Tutorials"
nav_weight: 1
order: 1
tags: [tutorial, hosting, getting-started]
---

By the end of this tutorial you'll have a site hosted behind GitHub-permission
gating at `studio/<owner>/<repo>/`.

You need a repo whose built static site is published to a branch (a Fuego site, or
any project that deploys to `gh-pages`/`fuego-pages`).

## 1. Install the app

Install the **fuego-studio** GitHub App on the repo (Settings → GitHub Apps, or the
app's install page). Installing grants studio read access so it can fetch the built
site to serve it.

## 2. Build at the studio mount path

Studio serves each site under a path prefix, `/<owner>/<repo>`. Your build must set
its `base_url` to that prefix so links and assets resolve. In your deploy workflow:

```bash
go run . build --base-url /<owner>/<repo>
```

Publish the result to the **`fuego-pages`** branch (studio prefers it). See
[Configure the deploy workflow](/docs/how-to/configure-the-deploy-workflow/) for a
complete workflow file.

## 3. Visit it

Open `studio/<owner>/<repo>/`:

- If the repo is **public**, it serves to anyone.
- If it's **private**, studio sends you through GitHub sign-in and checks you have
  read access before serving a byte.

Your accessible sites are also listed on the dashboard at `/`.

That's hosting. Next: [edit a page in place](/docs/tutorials/edit-your-first-page/).
