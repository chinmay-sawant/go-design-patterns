# Go Design Patterns Viewer

A minimal, dark-themed React application to visualize Go design patterns.

## Features

- **File Tree Navigation**: Browse the directory structure of your Go design patterns.
- **Syntax Highlighting**: Beautiful syntax highlighting for Go code.
- **Resizable Sidebar**: Draggable sidebar width for optimal viewing.
- **Dark Mode**: Premium dark theme with no gradients, focused on content readability.

## Setup

1. Install dependencies:

   ```bash
   npm install
   ```

2. Run development server:

   ```bash
   npm run dev
   ```

   This will automatically scan `../design_patterns` and generate the data structure.

3. Build for production:
   ```bash
   npm run build
   ```
   This generates a `docs` folder in the project root, ready for GitHub Pages deployment.

## Architecture

- **Frontend**: Vite + React
- **Styling**: Vanilla CSS with CSS Variables (Tailwind for utility classes only where explicitly needed)
- **Data Source**: A pre-build script (`scripts/generate-tree.cjs`) scans the local directory and bundles content into `src/data.json`. This ensures the site works identically locally and when deployed to GitHub Pages.

## Deployment

Simply commit the `docs` folder to your repository and enable GitHub Pages from the `/docs` folder in repository settings.
