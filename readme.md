# Dynamic HTML Story Renderer

A Go project that dynamically renders a **"Choose Your Adventure"-style story** using HTML templates. The story is defined in JSON format, and users can navigate through different arcs using links displayed on the web interface.

## Features

- **Dynamic Story Rendering**: Each story arc is rendered based on JSON data.
- **HTML Templates**: Uses Go's `html/template` package for dynamic web content generation.
- **User Interaction**: Provides clickable links to navigate through story arcs.
- **Flexible Story Input**: Stories are defined in JSON format, making them easy to extend or modify.
- **Lightweight Web Server**: A simple Go HTTP server powers the application.

---

## Getting Started

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/ah-naf/cyoa.git
   cd cyoa
   ```

2. **Add Your Story**:

   - Modify or replace the `story.json` file to include your own story arcs.
   - Ensure each arc has:
     - `title`: Title of the story arc.
     - `story`: Array of paragraphs describing the arc.
     - `options`: Array of options to navigate to other arcs.

3. **Run the Application**:

   ```bash
   go run main.go
   ```

4. **Access the Application**:
   - Open your browser and navigate to: [http://localhost:8080](http://localhost:8080)

---

## JSON Story Format

Here’s an example structure for the `story.json` file:

```json
{
  "intro": {
    "title": "The Mysterious Envelope",
    "story": [
      "You find a black envelope under your door...",
      "Inside is a cryptic note warning you of danger."
    ],
    "options": [
      { "text": "Call the police", "arc": "call-police" },
      { "text": "Investigate on your own", "arc": "solo-investigation" }
    ]
  },
  "call-police": {
    "title": "Calling the Police",
    "story": [
      "The police arrive but find nothing unusual...",
      "You wonder if this was the right choice."
    ],
    "options": []
  }
}
```
