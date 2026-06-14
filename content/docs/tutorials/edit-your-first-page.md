---
title: Edit your first page
layout: doc
summary: Open a Fuego page you have write access to, edit its source, commit, and watch it republish and reload itself.
nav_section: "Tutorials"
nav_weight: 2
order: 2
tags: [tutorial, editing]
---

This picks up from [hosting your first site](../host-your-first-site/). You'll edit
a page's source and see it go live.

You need **write access** to the repo, and the site must be a **Fuego** site (studio
reads its `site-manifest.json` to find each page's source).

## 1. Find the edit button

Open any content page on your hosted site. If you have write access, an **Edit
source** button appears. (No button? You may lack write access, or editing may be
turned off — see [Control editing](../../how-to/control-editing/).)

## 2. Edit the source

The button opens a **raw source editor** — the page's Markdown (or whatever DSL the
site uses), exactly as it is in the repo. There's no live preview, by design:
editing is about *what the page says*. Make your change.

## 3. Commit

Click **Commit**. Studio writes the file back to GitHub:

- The commit is **authored by you** (co-authored by `fuego-studio[bot]`).
- It's conditional on the version you loaded — if someone else changed the file
  first, you get a conflict instead of a silent overwrite.

## 4. Watch it publish

If your edit branch deploys (the common case), your normal CI rebuilds the site and
studio re-syncs. A toast walks you through *Committing → deploying → Published*, and
the page **reloads itself** with your change — scroll preserved.

If the site lives in a protected repo, your commit becomes a **pull request** instead
— see [Edit in a protected repo](../../how-to/edit-in-a-protected-repo/).

That's the whole loop. To understand what just happened, read
[The publish & reload loop](../../explanation/publish-and-reload/).
