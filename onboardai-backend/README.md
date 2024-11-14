# OnboardAI Backend

Go backend service for the OnboardAI platform.

## Requirements

- Go 1.23 or higher
- OpenAI API key

## Setup

1. Install dependencies:
```sh
go mod download
```

2. Create a `.env` file with your OpenAI API key:
```
OPENAI_API_KEY=your_api_key_here
```

3. Run the server:
```sh
go run main.go
```

The server will start on `localhost:8080`.

## API Endpoints

- `POST /api/signup`: Create new user account
- `POST /api/login`: Authenticate user
- `POST /api/generate-onboarding`: Generate onboarding checklist using GPT-4
- `GET /api/onboarding/:email`: Get existing onboarding data for an engineer
