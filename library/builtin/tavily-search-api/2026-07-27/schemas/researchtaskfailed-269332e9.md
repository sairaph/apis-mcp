---
title: ResearchTaskFailed
page_id: schema-researchtaskfailed-269332e9
path: schemas
source: https://docs.tavily.com/documentation/api-reference/openapi.json
source_type: openapi
imported_from: https://docs.tavily.com/documentation/api-reference/openapi.json
---

# ResearchTaskFailed

```yaml
{"title": "Failed", "type": "object", "properties": {"request_id": {"type": "string", "description": "The unique identifier of the research task.", "example": "123e4567-e89b-12d3-a456-426614174111"}, "status": {"type": "string", "description": "The current status of the research task.", "enum": ["failed"]}, "response_time": {"type": "integer", "description": "Time in seconds it took to complete the request.", "example": 1.23}}, "required": ["request_id", "status", "response_time"], "example": {"request_id": "123e4567-e89b-12d3-a456-426614174111", "status": "failed"}}
```
