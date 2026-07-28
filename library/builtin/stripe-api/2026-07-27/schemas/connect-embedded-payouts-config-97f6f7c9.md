---
title: connect_embedded_payouts_config
page_id: schema-connect-embedded-payouts-config-97f6f7c9
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# connect_embedded_payouts_config

```yaml
{"title": "ConnectEmbeddedPayoutsConfig", "required": ["enabled", "features"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "Whether the embedded component is enabled."}, "features": {"$ref": "#/components/schemas/connect_embedded_payouts_features"}}, "description": "", "x-expandableFields": ["features"]}
```
