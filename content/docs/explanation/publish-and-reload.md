---
title: The publish & reload loop
layout: doc
summary: What actually happens between clicking Commit and the page updating itself — and why there's no live preview.
nav_section: "Explanation"
nav_weight: 3
tags: [explanation, editing, reload]
---

When you commit an edit and the page updates on its own, a specific chain of events
runs. Understanding it explains both what you see and what you *don't*.

## The chain

1. **Commit** — studio writes your file back to GitHub on the resolved edit branch,
   authored by you, conditional on the version you loaded.
2. **CI rebuilds** — your normal deploy workflow runs and pushes the new build to the
   serve branch.
3. **Webhook** — GitHub tells studio about the push. Studio re-fetches the new build
   (by the push's exact commit, so it never reads a stale cache) and swaps it in.
4. **Reload** — studio's re-sync advances a generation counter; your open page is
   watching it and reloads when it ticks, preserving your scroll position.

The toast you see — *Committing → deploying → Published* — is narrating this chain.

## Why "by the exact commit" matters

GitHub's branch archives are CDN-cached and can briefly lag a fresh push. Early on,
studio re-fetched by branch name and sometimes pulled the *old* bytes — so a page
would say "published" but show stale content even after a hard refresh. Studio now
fetches by the push's immutable commit SHA, which is always current. It's an
invisible detail that makes the whole loop trustworthy.

## Why there's no live preview

You might expect a side-by-side preview. There isn't one, on purpose. Editing in
studio is aimed at **content** — fixing what a page *says* — for developers who
already work in the repo. "See it" means the real build ran and the real page
updated, which is the same thing your readers will see. Skipping a preview keeps the
editor a thin, honest layer over your actual publish pipeline rather than a
second renderer that could disagree with it.

## On a protected repo, the loop pauses

If edits go to a non-deploying branch (the [pull-request flow](/docs/how-to/edit-in-a-protected-repo/)),
there's no instant reload — feedback comes after the PR merges and deploys. That's
the deliberate cost of review.
