---
title: mq_event-source
page_id: schema-mq-event-source-7860ec0c
path: schemas
description: Source configuration for the subscription
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_event-source

Source configuration for the subscription

```yaml
{"description": "Source configuration for the subscription", "type": "object", "oneOf": [{"$ref": "#/components/schemas/mq_event-source-images"}, {"$ref": "#/components/schemas/mq_event-source-kv"}, {"$ref": "#/components/schemas/mq_event-source-r2"}, {"$ref": "#/components/schemas/mq_event-source-super-slurper"}, {"$ref": "#/components/schemas/mq_event-source-vectorize"}, {"$ref": "#/components/schemas/mq_event-source-workers-ai-model"}, {"$ref": "#/components/schemas/mq_event-source-workers-builds-worker"}, {"$ref": "#/components/schemas/mq_event-source-workflows-workflow"}]}
```
