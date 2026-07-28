---
title: pages_env_vars
page_id: schema-pages-env-vars-0d3e693e
path: schemas
description: Environment variables used for builds and Pages Functions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# pages_env_vars

Environment variables used for builds and Pages Functions.

```yaml
{"description": "Environment variables used for builds and Pages Functions.", "type": "object", "additionalProperties": {"discriminator": {"mapping": {"plain_text": "#/components/schemas/pages_plain_text_env_var", "secret_text": "#/components/schemas/pages_secret_text_env_var"}, "propertyName": "type"}, "nullable": true, "oneOf": [{"$ref": "#/components/schemas/pages_plain_text_env_var"}, {"$ref": "#/components/schemas/pages_secret_text_env_var"}], "type": "object"}, "nullable": true}
```
