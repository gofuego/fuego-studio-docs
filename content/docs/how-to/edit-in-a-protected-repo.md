---
title: Edit a site in a protected repo
layout: doc
summary: When the default branch is protected, edits become pull requests instead of direct commits. Here's the flow you'll see.
nav_section: "How-to Guides"
nav_weight: 5
tags: [how-to, editing, git]
---

If your docs live in a repo whose default branch is **protected** — like a code repo
where `main` requires review — studio won't (and can't) commit straight to `main`.
Instead it routes edits through a **pull request**. You don't configure anything;
this is the smart default.

## What happens when you commit

1. Studio ensures an edit branch (`fuego-studio-edits` by default) exists, creating
   it from the default branch.
2. It merges the default branch into the edit branch so you start from current
   content. If that merge **conflicts**, studio tells you rather than guessing.
3. It commits your change to the edit branch (authored by you).

## What you see

Instead of "Published", the editor shows:

> **Committed to `fuego-studio-edits` — Open a pull request →**

The link is a **prefilled compare page** (`base...edit`, expanded, with a title
filled in). A pending-edits prompt stays visible, counting how many commits are
waiting, until the branch is merged.

## Finishing up

Open the PR, get it reviewed, and merge it like any other. Your normal deploy runs
on merge and the site updates then — there's no instant reload on this path, by
design (the edit branch doesn't deploy).

## Want instant publish on a protected repo?

Point editing at a branch your workflow *does* deploy, and tell studio so, via
[`.fuego-studio.yaml`](/docs/how-to/control-editing/):

```yaml
edit_branch: docs-live
edit_branch_deploys: true
```
