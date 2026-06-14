---
title: Server configuration
layout: doc
summary: The environment variables a fuego-studio instance reads, and where each value comes from.
nav_section: "Reference"
nav_weight: 2
tags: [reference, configuration, ops]
---

A studio instance is configured entirely through the environment (copy
`.env.example` to `.env`, which is gitignored — never commit it).

| Variable | Required | Where to get it |
|----------|----------|-----------------|
| `APP_ID` | yes | GitHub App settings → "App ID" |
| `CLIENT_ID` | yes | App settings → "Client ID" |
| `CLIENT_SECRET` | yes | App settings → "Generate a client secret" |
| `WEBHOOK_SECRET` | yes | the value you set as the App's webhook secret |
| `PRIVATE_KEY_PATH` | yes | path to the App's downloaded `.pem` (keep it gitignored) |
| `PORT` | no | listen port; default `3000` |
| `EDITING` | no | `off` makes the whole instance read-only; default on |

## Notes

- These are **runtime secrets of the studio host**, not GitHub Actions secrets.
  Repos that you host need nothing studio-related in their own CI — only the
  built-in `GITHUB_TOKEN`.
- The private key signs the App JWT; the client secret backs user OAuth; the webhook
  secret verifies push deliveries. Keep all three out of version control.
- Sessions persist to a gitignored `.sessions.json` (mode `0600`) so signed-in users
  survive a restart.

## Local webhooks

To receive webhooks on a local instance, forward them with [smee](https://smee.io):

```bash
smee -u https://smee.io/<channel> --target http://127.0.0.1:3000/webhook
```
