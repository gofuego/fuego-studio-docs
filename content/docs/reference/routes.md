---
title: Routes
layout: doc
summary: Every HTTP route fuego-studio exposes and what it does.
nav_section: "Reference"
nav_weight: 1
tags: [reference, routes]
---

| Route | Purpose |
|-------|---------|
| `GET /` | Dashboard of sites you can access (or a sign-in screen). |
| `GET /<owner>/<repo>/…` | Serve the gated site; inject the editor for write-users. |
| `GET /auth/login` | Start GitHub OAuth login. |
| `GET /auth/callback` | OAuth callback. |
| `GET /auth/logout` | Sign out. |
| `GET /_studio/file` | Editor API — load a page's source (write-gated). |
| `POST /_studio/file` | Editor API — commit a page's source (write-gated). |
| `GET /_studio/status` | Publish status — generation counter, or PR `ahead`/`pr_url`. |
| `GET /_studio/events` | Live-reload SSE stream (write-gated). |
| `POST /webhook` | GitHub push deliveries → re-sync + reload broadcast. |
| `GET /healthz` | Liveness check (`ok`). |
| `GET /debug/installations` | List app installations + mint a token (operator check). |

## Notes

- Site and asset routes under `/<owner>/<repo>/` are **permission-gated**: public
  repos serve to anyone; private repos require repo read.
- `/_studio/*` editor and event routes are **write-gated** — only users with repo
  write reach them.
- `/webhook` verifies an HMAC-SHA256 signature; unsigned or mismatched deliveries
  are rejected.
