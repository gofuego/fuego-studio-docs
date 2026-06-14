---
title: Why gateway-first
layout: doc
summary: Studio hosts a site it never builds. Being the origin for the served bytes is what makes auth, editing, and reload simple instead of hard.
nav_section: "Explanation"
nav_weight: 1
tags: [explanation, architecture]
---

fuego-studio could have been a build service, or a proxy in front of GitHub Pages.
It is neither. It is a **gateway**: it serves a site's already-built output behind a
permission check, and never builds the site itself.

## The choice

Your repo's normal CI builds the site and pushes the static output to a serve
branch. Studio fetches that branch and serves it, gated by GitHub permission. The
build stays where it belongs — in your workflow — and studio is just the
authenticating front door.

## Why it's the origin, not a proxy

Because studio serves the actual bytes (rather than redirecting you to another host),
the site and the gate share **one origin**. That single fact removes a stack of
problems:

- **Auth is simple.** No cross-origin cookies, no CORS negotiation, no bouncing
  between a login host and a content host. A request either clears the permission
  check or it doesn't.
- **The editor ships with the page.** Studio injects it into the same HTML it's
  already serving — there's no separate editor app to host, authenticate, or keep in
  sync.
- **Reload is possible.** Studio holds the served bytes in memory, so when a new
  build lands it can tell your open page to update.

## What you give up

Studio doesn't build, preview, or transform your site. If your build is broken, your
site is broken — studio just hosts whatever you published. That's the trade: less
magic, far less surface area, and an auth-and-edit layer that's easy to reason about.

This is the same separation-of-concerns the Fuego engine follows: the engine builds
and emits a manifest of facts; studio reads those facts and adds hosting, gating, and
editing on top.
