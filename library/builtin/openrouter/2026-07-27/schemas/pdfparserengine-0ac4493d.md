---
title: PDFParserEngine
page_id: schema-pdfparserengine-0ac4493d
path: schemas
description: The engine to use for parsing PDF files. "pdf-text" is deprecated and automatically redirected to "cloudflare-ai".
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# PDFParserEngine

The engine to use for parsing PDF files. "pdf-text" is deprecated and automatically redirected to "cloudflare-ai".

```yaml
{"anyOf": [{"enum": ["mistral-ocr", "native", "cloudflare-ai"], "type": "string", "x-speakeasy-unknown-values": "allow"}, {"enum": ["pdf-text"], "type": "string"}], "description": "The engine to use for parsing PDF files. \"pdf-text\" is deprecated and automatically redirected to \"cloudflare-ai\".", "example": "cloudflare-ai"}
```
