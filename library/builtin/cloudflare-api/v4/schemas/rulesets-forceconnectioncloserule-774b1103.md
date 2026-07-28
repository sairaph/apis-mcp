---
title: rulesets_ForceConnectionCloseRule
page_id: schema-rulesets-forceconnectioncloserule-774b1103
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_ForceConnectionCloseRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["force_connection_close"]}, "description": {"example": "Close ongoing HTTP connections. This action does not block a request, but it forces the client to reconnect. For HTTP/2 and HTTP/3 connections, the connection will be closed even if it breaks other requests running on the same connection."}}, "title": "Force Connection Close Rule"}]}
```
