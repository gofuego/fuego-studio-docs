---
title: Install the GitHub App
layout: doc
summary: Install fuego-studio on the repos you want hosted, and check the install with the debug endpoints.
nav_section: "How-to Guides"
nav_weight: 1
tags: [how-to, github, setup]
---

## Install on a repo

1. Go to the **fuego-studio** GitHub App's page and choose **Install** (or
   **Configure** for an existing installation).
2. Select the account/org and either **All repositories** or specific repos. With
   "All repositories", new repos are covered automatically.
3. Installing grants studio the access it needs to fetch the built site (and, for
   editing, to read branch protection and commit on the user's behalf).

## Confirm it worked

A repo shows up on the dashboard once it has a **serve branch**
(`gh-pages`/`fuego-pages`) — that's how studio recognises a hosted site. So after
installing, make sure your [deploy workflow](../configure-the-deploy-workflow/) has
run at least once.

## Operator smoke tests

If you run the instance, two endpoints confirm credentials and installations:

```bash
curl <host>/healthz                 # -> ok
curl <host>/debug/installations     # lists installations + mints a token
```

A successful `/debug/installations` means the App ID, private key, and webhook
wiring are correct.
