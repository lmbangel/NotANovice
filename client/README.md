# Daily Discipleship Quiz - Frontend Client

A vanilla JavaScript, HTML, and Tailwind CSS frontend for the Daily Discipleship Quiz application.

## Features

- 📱 Mobile-first responsive design
- 🎨 Beautiful UI inspired by the provided design references
- ⚡ Vanilla JavaScript (no frameworks)
- 🎯 Single question display format
- 📡 REST API integration with backend
- ✅ Interactive radio button answers
- 🎉 Results feedback

## Getting Started

### Prerequisites

- Node.js (for the development server)
- Backend API running on `http://localhost:8000`

### Installation & Running

1. Navigate to the client directory:
   ```bash
   cd NotANovice/client
   ```

2. Start the development server:
   ```bash
   npm start
   ```

3. Open your browser and go to:
   ```
   http://localhost:3000
   ```

### Backend Integration

The client connects to the backend API at `http://localhost:8000/v1/questions`. Make sure your Go backend is running before testing the client.

## Project Structure

```
client/
├── index.html          # Main HTML file
├── app.js             # JavaScript application logic
├── server.js          # Simple Node.js development server
├── package.json       # NPM configuration
├── README.md          # This file
└── docs/              # Design references
    ├── App Design.png
    └── App Design 2.png
```

## API Integration

Currently integrates with:
- `GET /v1/questions` - Fetches available questions (displays first one)

## Design

The interface follows the design references with:
- Coral/orange accent colors (#FF7F7F, #FF9A7F)
- Teal for success states (#4ECDC4, #26D0CE)
- Clean white cards with rounded corners
- Mobile-first responsive layout
- Modern typography and spacing