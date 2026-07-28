---
title: dlp_RegexValidationQuery
page_id: schema-dlp-regexvalidationquery-f081bc58
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_RegexValidationQuery

```yaml
{"type": "object", "properties": {"max_match_bytes": {"description": "Maximum number of bytes that the regular expression can match.\n\nIf this is `null` then there is no limit on the length. Patterns can use\n`*` and `+`. Otherwise repeats should use a range `{m,n}` to restrict\npatterns to the length. If this field is missing, then a default length\nlimit is used.\n\nNote that the length is specified in bytes. Since regular expressions\nuse UTF-8 the pattern `.` can match up to 4 bytes. Hence `.{1,256}`\nhas a maximum length of 1024 bytes.", "type": "integer", "format": "int32", "minimum": 0, "nullable": true}, "regex": {"type": "string"}}, "required": ["regex"]}
```
