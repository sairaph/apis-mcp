---
title: waitingroom_turnstile_action
page_id: schema-waitingroom-turnstile-action-fea2bc9b
path: schemas
description: |-
    Which action to take when a bot is detected using Turnstile. `log` will
    have no impact on queueing behavior, simply keeping track of how many
    bots are detected in Waiting Room Analytics. `infinite_queue` will send
    bots to a false queueing state, where they will never reach your
    origin. `infinite_queue` requires Advanced Waiting Room.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waitingroom_turnstile_action

Which action to take when a bot is detected using Turnstile. `log` will
have no impact on queueing behavior, simply keeping track of how many
bots are detected in Waiting Room Analytics. `infinite_queue` will send
bots to a false queueing state, where they will never reach your
origin. `infinite_queue` requires Advanced Waiting Room.

```yaml
{"description": "Which action to take when a bot is detected using Turnstile. `log` will\nhave no impact on queueing behavior, simply keeping track of how many\nbots are detected in Waiting Room Analytics. `infinite_queue` will send\nbots to a false queueing state, where they will never reach your\norigin. `infinite_queue` requires Advanced Waiting Room.\n", "type": "string", "default": "log", "enum": ["log", "infinite_queue"], "x-auditable": true}
```
