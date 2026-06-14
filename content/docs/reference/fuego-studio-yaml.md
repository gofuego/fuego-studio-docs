---
title: .fuego-studio.yaml
layout: doc
summary: "The optional per-repo editing-policy file: its three keys, their defaults, and how they resolve."
nav_section: "Reference"
nav_weight: 3
tags: [reference, editing, policy]
---

`.fuego-studio.yaml` is an **optional** file on a repo's **default branch** that
overrides studio's editing policy for that repo. Omit it and studio uses smart
defaults.

```yaml
editing: true               # may this repo be edited at all?
edit_branch: fuego-studio-edits   # which branch edits land on
edit_branch_deploys: false  # does that branch deploy (instant publish)?
```

## Keys

| Key | Type | Default | Meaning |
|-----|------|---------|---------|
| `editing` | bool | smart | `false` makes the repo read-only in studio. |
| `edit_branch` | string | smart | Branch edits are committed to. |
| `edit_branch_deploys` | bool | smart | Whether `edit_branch` deploys; selects instant-publish vs the pull-request flow. |

## How resolution works

The effective policy is decided in three layers, highest priority first:

1. **Server `EDITING=off`** → read-only, regardless of this file.
2. **This file** → any key you set wins over the default.
3. **Smart default** (any key you omit):
   - **`editing`** → on.
   - **`edit_branch`** → the default branch if it's unprotected; otherwise an auto
     `fuego-studio-edits` branch.
   - **`edit_branch_deploys`** → true when the edit branch *is* the deploying default
     branch, false otherwise.

> This file is studio's concern — workflow policy, not a build fact. The Fuego
> engine and its `site-manifest.json` know nothing about it.
