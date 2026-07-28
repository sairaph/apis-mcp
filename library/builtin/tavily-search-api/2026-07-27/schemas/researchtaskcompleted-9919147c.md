---
title: ResearchTaskCompleted
page_id: schema-researchtaskcompleted-9919147c
path: schemas
source: https://docs.tavily.com/documentation/api-reference/openapi.json
source_type: openapi
imported_from: https://docs.tavily.com/documentation/api-reference/openapi.json
---

# ResearchTaskCompleted

```yaml
{"title": "Completed", "type": "object", "properties": {"request_id": {"type": "string", "description": "The unique identifier of the research task.", "example": "123e4567-e89b-12d3-a456-426614174111"}, "created_at": {"type": "string", "description": "Timestamp when the research task was created.", "example": "2025-01-15T10:30:00Z"}, "status": {"type": "string", "description": "The current status of the research task.", "enum": ["completed"]}, "content": {"oneOf": [{"type": "string"}, {"type": "object"}], "description": "The research report content. Can be a string or a structured object if output_schema was provided."}, "sources": {"type": "array", "description": "List of sources used in the research.", "items": {"type": "object", "properties": {"title": {"type": "string", "description": "Title or name of the source.", "example": "Latest AI Developments"}, "url": {"type": "string", "format": "uri", "description": "URL of the source.", "example": "https://example.com/ai-news"}, "favicon": {"type": "string", "format": "uri", "description": "URL to the source's favicon.", "example": "https://example.com/favicon.ico"}}}}, "response_time": {"type": "integer", "description": "Time in seconds it took to complete the request.", "example": 1.23}}, "required": ["request_id", "created_at", "status", "content", "sources", "response_time"], "example": {"request_id": "123e4567-e89b-12d3-a456-426614174111", "created_at": "2025-01-15T10:30:00Z", "status": "completed", "content": "Research Report: Latest Developments in AI\n\n## Executive Summary\n\nArtificial Intelligence has seen significant advancements in recent months, with major breakthroughs in large language models, multimodal AI systems, and real-world applications...", "sources": [{"title": "Latest AI Developments", "url": "https://example.com/ai-news", "favicon": "https://example.com/favicon.ico"}, {"title": "AI Research Breakthroughs", "url": "https://example.com/ai-research", "favicon": "https://example.com/favicon.ico"}], "response_time": 1.23}}
```
