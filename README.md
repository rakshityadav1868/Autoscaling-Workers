# TaskForge

An AI-powered asynchronous job processing platform. Applications submit long-running tasks — text summarization, email generation, translations, or other AI workflows — without blocking the user. Background workers consume jobs, call LLMs where needed, update status and results in the database, and expose everything through a simple API.

## Architecture

```
Client → API (submit job) → SQLite (store metadata) → Queue → Workers → LLM
                                                  ↓
                                           GET /jobs/:id
```

- **API**: Accepts job submissions and returns status
- **SQLite**: Stores job metadata, status, and results
- **Queue**: Holds pending jobs for workers (in-memory, moving to Redis)
- **Workers**: Consume jobs, execute logic or call LLMs, update results

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| POST | `/submit` | Submit a new job. Body: `{"type": "...", "payload": "..."}` |
| GET | `/jobs/:id` | Get job status and result by ID |
| GET | `/jobs` | List all jobs |

## Tech Stack

- **Language**: Go
- **Database**: SQLite (migrating to Redis for queue)
- **Planned**: Redis-backed queue, multi-LLM worker pool, dashboard UI

## Getting Started

```bash
git clone https://github.com/rakshityadav1868/TaskForge.git
cd TaskForge/autoworkers
go run cmd/server/main.go
```

The server starts and accepts job submissions. Workers pick up jobs automatically.

## Example

```bash
# Submit a summarization job
curl -X POST http://localhost:8080/submit \
  -H "Content-Type: application/json" \
  -d '{"type": "summarize", "payload": "Long text to summarize..."}'

# Check job status
curl http://localhost:8080/jobs/job-1
```

## Roadmap

- Redis-backed queue for persistence and scale
- Pluggable LLM workers (OpenAI, Anthropic, Gemini)
- Job retries and dead-letter queue
- Web dashboard for job monitoring
- Worker autoscaling based on queue depth