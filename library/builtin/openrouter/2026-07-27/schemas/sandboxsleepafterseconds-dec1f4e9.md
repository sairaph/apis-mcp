---
title: SandboxSleepAfterSeconds
page_id: schema-sandboxsleepafterseconds-dec1f4e9
path: schemas
description: 'How long (in seconds) the container stays warm after its last command before sleeping, freeing its capacity slot. Idle-based: each command renews the timer. Defaults to 900 (15 minutes); capped at 2592000 (30 days).'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# SandboxSleepAfterSeconds

How long (in seconds) the container stays warm after its last command before sleeping, freeing its capacity slot. Idle-based: each command renews the timer. Defaults to 900 (15 minutes); capped at 2592000 (30 days).

```yaml
{"description": "How long (in seconds) the container stays warm after its last command before sleeping, freeing its capacity slot. Idle-based: each command renews the timer. Defaults to 900 (15 minutes); capped at 2592000 (30 days).", "example": 900, "type": "integer"}
```
