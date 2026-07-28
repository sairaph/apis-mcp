---
title: mq_consumer-response
page_id: schema-mq-consumer-response-a5f4c0c8
path: schemas
description: Response body representing a consumer
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_consumer-response

Response body representing a consumer

```yaml
{"description": "Response body representing a consumer", "type": "object", "discriminator": {"mapping": {"http_pull": "#/components/schemas/mq_http-consumer-response", "worker": "#/components/schemas/mq_worker-consumer-response"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/mq_worker-consumer-response"}, {"$ref": "#/components/schemas/mq_http-consumer-response"}]}
```
