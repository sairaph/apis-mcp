---
title: ReaderRequest
page_id: schema-readerrequest-e101b35b
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# ReaderRequest

```yaml
{"type": "object", "properties": {"url": {"type": "string", "description": "The URL to retrieve"}, "timeout": {"type": "integer", "description": "Request timeout in seconds. Default is 20", "default": 20}, "no_cache": {"type": "boolean", "description": "Whether to disable caching (true/false). Default is false", "default": false}, "return_format": {"type": "string", "description": "Return format (e.g., markdown, text). Default is markdown", "default": "markdown"}, "retain_images": {"type": "boolean", "description": "Whether to retain images (true/false). Default is true", "default": true}, "no_gfm": {"type": "boolean", "description": "Whether to disable GitHub Flavored Markdown (true/false). Default is false", "default": false}, "keep_img_data_url": {"type": "boolean", "description": "Whether to keep image data URLs (true/false). Default is false", "default": false}, "with_images_summary": {"type": "boolean", "description": "Whether to include image summary (true/false). Default is false", "default": false}, "with_links_summary": {"type": "boolean", "description": "Whether to include links summary (true/false). Default is false", "default": false}}, "required": ["url"]}
```
