# AI Language Teacher Application

This repository is for an AI-powered language learning chatbot and tool. The goal is to build an intelligent assistant to help practice and learn new languages.

## Key Features

- Chat interface for conversing with AI bot
- NLP powered by BERT models and intent classification
- Speech recognition and synthesis for audio conversation
- Personalized learning based on knowledge gaps
- Vocabulary practice and quizzes
- Translation and pronunciation help
- Admin dashboard for managing content

## Tech Stack

- Go backend with Gin framework
- PostgreSQL database
- Docker for environment and deployment
- React frontend
- Google Cloud Speech & Text-to-Speech APIs
- BERT NLP models
- TailwindCSS styling
- GitHub Actions CI/CD
- Fly.io hosting

## Getting Started

1. Install Go 1.18+
2. Clone the repo
3. Initialize the DB schema
4. Start the API with `go run .`
5. Connect to http://localhost:8080


## Testing 

```bash
go test -cover ./... # test coverage for recursively

go test -coverprofile=c.out # create a test coverage file
go tool cover -func=c.out # print nice coverage report to console
```


Contributions welcome!

## License

This project uses the MIT license - see LICENSE.md for details
