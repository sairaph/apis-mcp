---
title: builds_WorkerResponse
page_id: schema-builds-workerresponse-72bb6c1d
path: schemas
description: Worker build configuration including git repository linkage and production settings
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_WorkerResponse

Worker build configuration including git repository linkage and production settings

```yaml
{"description": "Worker build configuration including git repository linkage and production settings", "type": "object", "properties": {"git_repository": {"$ref": "#/components/schemas/builds_WorkerGitRepository"}, "production_settings": {"$ref": "#/components/schemas/builds_WorkerBuildSettings"}, "script_tag": {"$ref": "#/components/schemas/builds_external_script_id"}}}
```
