---
title: Control who can edit
layout: doc
summary: Turn editing off entirely, or override the per-repo policy with a .fuego-studio.yaml. Most repos need neither.
nav_section: "How-to Guides"
nav_weight: 4
tags: [how-to, editing, policy]
---

Editing is resolved per repo from three layers. Most setups need **zero config** —
reach for these only to override the default.

## Turn editing off for the whole instance

Set the server flag when running studio:

```bash
EDITING=off go run ./cmd/fuego-studio
```

The instance becomes a **read-only host** — it still serves and gates sites, but no
editor appears anywhere.

## Override per repo with `.fuego-studio.yaml`

Add `.fuego-studio.yaml` to the repo's **default branch**:

```yaml
editing: false              # disable editing for just this repo
edit_branch: docs-edits     # commit edits to this branch instead of the default
edit_branch_deploys: true   # this branch deploys → instant publish (else: PR flow)
```

All three keys are optional:

- **`editing: false`** — read-only for this repo even though the instance allows editing.
- **`edit_branch`** — land edits on a dedicated branch instead of the default.
- **`edit_branch_deploys`** — tell studio whether that branch deploys. Set it `true`
  if your workflow deploys the edit branch (so studio uses the instant-publish loop);
  otherwise studio uses the [pull-request flow](/docs/how-to/edit-in-a-protected-repo/).

## The default when you configure nothing

- **Unprotected default branch** → edit it directly, instant publish.
- **Protected default branch** → commit to an auto `fuego-studio-edits` branch via a
  pull request. Safe by default, no config required.

In all cases a user still needs **GitHub write access** to see an editor —
`canEdit` is *policy-allows-it* **and** *you-have-write*.
