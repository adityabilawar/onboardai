# onboardai

OnboardAI is a platform that automates the onboarding process for new developers using AI. It generates personalized onboarding checklists and coding challenges based on your codebase documentation.

## Features

- AI-powered onboarding checklist generation
- Role-based access (Manager/Engineer)
- Progress tracking for new hires
- Integration capabilities with Notion, Jira, GitHub, and GitLab
- Real-time progress updates

## Project Structure

- `/onboardai-frontend`: Vue.js frontend application
- `/onboardai-backend`: Go backend service

## Getting Started

1. Start the backend:
```sh
cd onboardai-backend
go mod download
go run main.go
```

2. Start the frontend:
```sh
cd onboardai-frontend
npm install
npm run dev
```

See individual README files in each directory for detailed setup instructions.