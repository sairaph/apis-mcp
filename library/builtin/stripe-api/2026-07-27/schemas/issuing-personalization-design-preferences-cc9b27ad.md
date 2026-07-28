---
title: issuing_personalization_design_preferences
page_id: schema-issuing-personalization-design-preferences-cc9b27ad
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_personalization_design_preferences

```yaml
{"title": "IssuingPersonalizationDesignPreferences", "required": ["is_default"], "type": "object", "properties": {"is_default": {"type": "boolean", "description": "Whether we use this personalization design to create cards when one isn't specified. A connected account uses the Connect platform's default design if no personalization design is set as the default design."}, "is_platform_default": {"type": "boolean", "description": "Whether this personalization design is used to create cards when one is not specified and a default for this connected account does not exist.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
