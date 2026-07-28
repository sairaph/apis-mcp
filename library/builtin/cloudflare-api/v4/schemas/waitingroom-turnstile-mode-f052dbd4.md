---
title: waitingroom_turnstile_mode
page_id: schema-waitingroom-turnstile-mode-f052dbd4
path: schemas
description: |-
    Which Turnstile widget type to use for detecting bot traffic. See
    [the Turnstile documentation](https://developers.cloudflare.com/turnstile/concepts/widget/#widget-types)
    for the definitions of these widget types. Set to `off` to disable the
    Turnstile integration entirely. Setting this to anything other than
    `off` or `invisible` requires Advanced Waiting Room.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waitingroom_turnstile_mode

Which Turnstile widget type to use for detecting bot traffic. See
[the Turnstile documentation](https://developers.cloudflare.com/turnstile/concepts/widget/#widget-types)
for the definitions of these widget types. Set to `off` to disable the
Turnstile integration entirely. Setting this to anything other than
`off` or `invisible` requires Advanced Waiting Room.

```yaml
{"description": "Which Turnstile widget type to use for detecting bot traffic. See\n[the Turnstile documentation](https://developers.cloudflare.com/turnstile/concepts/widget/#widget-types)\nfor the definitions of these widget types. Set to `off` to disable the\nTurnstile integration entirely. Setting this to anything other than\n`off` or `invisible` requires Advanced Waiting Room.\n", "type": "string", "default": "invisible", "enum": ["off", "invisible", "visible_non_interactive", "visible_managed"], "x-auditable": true}
```
