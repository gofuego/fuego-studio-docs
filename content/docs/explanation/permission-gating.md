---
title: How permission gating works
layout: doc
summary: Studio doesn't invent access rules — it mirrors GitHub's. Your repo permission decides what you can see, host, and edit.
nav_section: "Explanation"
nav_weight: 2
tags: [explanation, auth, security]
---

The most important thing to understand about studio's security model: it **has no
access model of its own**. Everything maps to your GitHub repo permissions.

## Reading a site

When you request a site:

- **Public repo** → served to anyone. (The site is public on GitHub Pages anyway;
  gating it would be theatre.)
- **Private repo** → you must have **repo read**. Not signed in? Studio sends you
  through GitHub OAuth. Signed in but no access? Denied.

This is what makes "GitHub Pages for private sites" real: a private docs site is
visible to exactly the people who can already see the repo, enforced by GitHub, not
by a password studio invented.

## The dashboard

The dashboard shows the sites *you* can access — built by walking the app's
installations and keeping repos that both host a site and that your token can see.
You never see a site you couldn't already reach on GitHub.

## Editing

An edit button appears only when **both** are true:

> the repo's policy allows editing **and** you have GitHub **write** access.

So a read-only viewer of a private site sees the content but no editor; a writer sees
the editor. Commits are made with **your** token, so they're attributed to you and
respect your permissions — studio can never commit something you couldn't commit
yourself.

## Why mirror GitHub instead of adding roles

A separate permission system would be a second source of truth to keep in sync, and a
new thing to get wrong. By deferring entirely to GitHub, studio's access control is
exactly as correct as GitHub's — and there's nothing extra to administer.
