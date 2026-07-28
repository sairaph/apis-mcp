---
title: zones_content_converter
page_id: schema-zones-content-converter-c39fa689
path: schemas
description: |-
    When enabled and the client sends an Accept header requesting text/markdown,
    Cloudflare will convert HTML responses to Markdown format using the toMarkdown() service.
    Refer to the [developer documentation](https://developers.cloudflare.com/workers-ai/features/markdown-conversion/) for more information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_content_converter

When enabled and the client sends an Accept header requesting text/markdown,
Cloudflare will convert HTML responses to Markdown format using the toMarkdown() service.
Refer to the [developer documentation](https://developers.cloudflare.com/workers-ai/features/markdown-conversion/) for more information.

```yaml
{"description": "When enabled and the client sends an Accept header requesting text/markdown,\nCloudflare will convert HTML responses to Markdown format using the toMarkdown() service.\nRefer to the [developer documentation](https://developers.cloudflare.com/workers-ai/features/markdown-conversion/) for more information.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "content_converter", "enum": ["content_converter"]}, "value": {"$ref": "#/components/schemas/zones_content_converter_value"}}}], "title": "Content Converter"}
```
