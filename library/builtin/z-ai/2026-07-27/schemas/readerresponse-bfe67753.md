---
title: ReaderResponse
page_id: schema-readerresponse-bfe67753
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# ReaderResponse

```yaml
{"type": "object", "properties": {"id": {"type": "string", "description": "Task ID"}, "created": {"type": "integer", "format": "int64", "description": "Request creation time as a Unix timestamp in seconds"}, "request_id": {"type": "string", "description": "Client-provided unique identifier to distinguish requests. If not provided, the platform will generate one."}, "model": {"type": "string", "description": "Model code"}, "reader_result": {"type": "object", "description": "Web reading result", "properties": {"content": {"type": "string", "description": "Main content parsed from the page (body, images, links, etc.)"}, "description": {"type": "string", "description": "Brief description of the page"}, "title": {"type": "string", "description": "Page title"}, "url": {"type": "string", "description": "Original page URL"}, "external": {"type": "object", "description": "External resources referenced by the page", "properties": {"stylesheet": {"type": "object", "description": "Collection of external stylesheets", "additionalProperties": {"type": "object", "properties": {"type": {"type": "string", "description": "Stylesheet MIME type, typically `text/css`"}}}}}}, "metadata": {"type": "object", "description": "Page metadata", "properties": {"keywords": {"type": "string", "description": "Page keywords"}, "viewport": {"type": "string", "description": "Viewport settings"}, "description": {"type": "string", "description": "Meta description"}, "format-detection": {"type": "string", "description": "Format detection settings, e.g., `telephone=no`"}}}}}}}
```
