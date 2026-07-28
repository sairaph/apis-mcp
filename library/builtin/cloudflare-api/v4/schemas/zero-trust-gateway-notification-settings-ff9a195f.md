---
title: zero-trust-gateway_notification_settings
page_id: schema-zero-trust-gateway-notification-settings-ff9a195f
path: schemas
description: Configure the message the user's device shows during an antivirus scan.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_notification_settings

Configure the message the user's device shows during an antivirus scan.

```yaml
{"description": "Configure the message the user's device shows during an antivirus scan.", "type": "object", "properties": {"enabled": {"description": "Specify whether to enable notifications.", "type": "boolean", "x-auditable": true}, "include_context": {"description": "Specify whether to include context information as query parameters.", "type": "boolean", "x-auditable": true}, "msg": {"description": "Specify the message to show in the notification.", "type": "string", "x-auditable": true}, "support_url": {"description": "Specify a URL that directs users to more information. If unset, the notification opens a block page.", "type": "string", "x-auditable": true}}, "nullable": true, "x-stainless-terraform-configurability": "computed_optional"}
```
