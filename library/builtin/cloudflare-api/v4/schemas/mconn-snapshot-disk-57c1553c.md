---
title: mconn_snapshot_disk
page_id: schema-mconn-snapshot-disk-57c1553c
path: schemas
description: Snapshot Disk
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_snapshot_disk

Snapshot Disk

```yaml
{"description": "Snapshot Disk", "type": "object", "properties": {"discards": {"description": "Discards completed successfully", "type": "number"}, "discards_merged": {"description": "Discards merged", "type": "number"}, "flushes": {"description": "Flushes completed successfully", "type": "number"}, "in_progress": {"description": "I/Os currently in progress", "type": "number"}, "major": {"description": "Device major number", "type": "number"}, "merged": {"description": "Reads merged", "type": "number"}, "minor": {"description": "Device minor number", "type": "number"}, "name": {"description": "Device name", "type": "string"}, "reads": {"description": "Reads completed successfully", "type": "number"}, "sectors_discarded": {"description": "Sectors discarded", "type": "number"}, "sectors_read": {"description": "Sectors read successfully", "type": "number"}, "sectors_written": {"description": "Sectors written successfully", "type": "number"}, "time_discarding_ms": {"description": "Time spent discarding (milliseconds)", "type": "number"}, "time_flushing_ms": {"description": "Time spent flushing (milliseconds)", "type": "number"}, "time_in_progress_ms": {"description": "Time spent doing I/Os (milliseconds)", "type": "number"}, "time_reading_ms": {"description": "Time spent reading (milliseconds)", "type": "number"}, "time_writing_ms": {"description": "Time spent writing (milliseconds)", "type": "number"}, "weighted_time_in_progress_ms": {"description": "Weighted time spent doing I/Os (milliseconds)", "type": "number"}, "writes": {"description": "Writes completed", "type": "number"}, "writes_merged": {"description": "Writes merged", "type": "number"}}, "required": ["in_progress", "major", "merged", "minor", "name", "reads", "sectors_read", "sectors_written", "time_in_progress_ms", "time_reading_ms", "time_writing_ms", "weighted_time_in_progress_ms", "writes", "writes_merged"]}
```
