---
title: r2_enable_sippy_gcs
page_id: schema-r2-enable-sippy-gcs-44d3600d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_enable_sippy_gcs

```yaml
{"properties": {"destination": {"description": "R2 bucket to copy objects to.", "type": "object", "properties": {"accessKeyId": {"description": "ID of a Cloudflare API token.\nThis is the value labelled \"Access Key ID\" when creating an API.\ntoken from the [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).\n\nSippy will use this token when writing objects to R2, so it is\nbest to scope this token to the bucket you're enabling Sippy for.\n", "type": "string"}, "provider": {"type": "string", "enum": ["r2"]}, "secretAccessKey": {"description": "Value of a Cloudflare API token.\nThis is the value labelled \"Secret Access Key\" when creating an API.\ntoken from the [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).\n\nSippy will use this token when writing objects to R2, so it is\nbest to scope this token to the bucket you're enabling Sippy for.\n", "type": "string", "x-sensitive": true}}}, "source": {"description": "GCS bucket to copy objects from.", "type": "object", "properties": {"bucket": {"description": "Name of the GCS bucket.", "type": "string", "x-auditable": true}, "clientEmail": {"description": "Client email of an IAM credential (ideally scoped to a single GCS bucket).", "type": "string"}, "privateKey": {"description": "Private Key of an IAM credential (ideally scoped to a single GCS bucket).", "type": "string", "x-sensitive": true}, "provider": {"type": "string", "enum": ["gcs"]}}}}}
```
