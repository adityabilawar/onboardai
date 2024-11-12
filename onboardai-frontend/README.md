# OnboardAI Frontend

A modern web application for automating the onboarding process for new developers using AI.

## Features

- Authentication (Login/Signup)
- Role-based dashboards (Manager/Engineer)
- AI-powered onboarding checklist generation
- Progress tracking for engineers
- Integration capabilities with Notion, Jira, GitHub, and GitLab

## Project Setup

```sh
# Install dependencies
npm install

# Start development server
npm run dev

# Build for production
npm run build
```

## Environment Setup

Create a `.env` file in the root directory:

```
VITE_API_URL=http://localhost:8080
```

## Backend Requirements

Make sure the backend server is running at `http://localhost:8080`. See the backend README for setup instructions.
