---
title: intel-sinkholes_sinkhole_create_params
page_id: schema-intel-sinkholes-sinkhole-create-params-32c3883e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel-sinkholes_sinkhole_create_params

```yaml
{"type": "object", "properties": {"name": {"description": "The name of the sinkhole.", "type": "string"}, "r2_bucket": {"description": "The name of the R2 bucket to store results. Required if you want to store large request bodies in R2.", "type": "string"}, "r2_id": {"description": "The id of the R2 instance. Required if you want to store large request bodies in R2.", "type": "string"}, "r2_secret": {"description": "The secret key for the R2 API token. Required if you want to store large request bodies in R2.", "type": "string", "writeOnly": true}}, "required": ["name"]}
```
