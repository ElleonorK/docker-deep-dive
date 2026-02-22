# Docker Exercise Style Guide

## Document Structure

### Header Pattern
```markdown
# Exercise XX: [Topic Name]

[One-line description of what they'll learn]

## Objectives

[Bullet list starting with "By the end of this exercise, you'll be comfortable with:"]
- [Specific skill 1]
- [Specific skill 2]
- [Specific skill 3]
```

### Task Structure
```markdown
### Task N: [Action-Oriented Title]

[Context paragraph explaining the scenario or setup]

**Your challenge:** [Clear goal statement]

**Success criteria:** [Observable outcome they can verify]

**Think:** [Reflective question about why/when/how]
```

## Language & Tone

### Voice
- Direct, conversational, second person ("you")
- Encouraging but not patronizing
- Assumes intelligence, not prior knowledge
- Uses contractions naturally ("you'll", "that's", "what's")

### Instruction Style
- Start with action verbs: "Run", "Create", "Build", "Try", "Verify"
- Use "Your goal:" or "Your challenge:" for objectives
- Use "Success criteria:" for verification
- Use "Think:" for reflection prompts
- Use "Hint:" sparingly, only when needed

### Formatting Conventions
- Use `code formatting` for commands, filenames, URLs, ports, environment variables
- Use **bold** for section labels: "Your challenge:", "Success criteria:", "Think:", "Hint:"
- Use bullet points (*) not dashes (-) for lists
- Use checkboxes (- [ ]) for verification checklists

## Task Design Principles

### Discovery-Based Learning
- Don't give direct answers
- Provide enough context to figure it out
- Let them experiment and fail
- Guide with questions, not solutions

Example:
```markdown
What's annoying about modifying the Dockerfile for each version? There's a better way coming in Task 3.
```

### Observable Outcomes
Every task must have a clear way to know you succeeded:
- "You see JSON with app info when you visit the URL"
- "Both containers should run simultaneously without conflicts"
- "Check the logs - there should be no connection errors"

### Progressive Complexity
- Start simple, add layers
- Each task builds on previous understanding
- Introduce one new concept at a time
- Revisit concepts in new contexts

Example progression:
1. Run a container (basic)
2. Run with environment variable (add config)
3. Make it configurable at build time (add build args)

### Reflection Questions
Use "Think:" prompts to encourage deeper understanding:
- "Why would you want to have a default command for an image?"
- "How does this architecture pattern improve security?"
- "Does the build stage size matter, or only the final stage?"

## Task Types

### Type 1: Experimentation
"Try X. What happens? Now try Y. What's different?"

Example:
```markdown
Try running the `danielkraic/asciiquarium` image. What happens? Can you see the aquarium?

Now try running it in a different way. Can you see the fish swimming now?
```

### Type 2: State Transitions
"Move something through different states and observe"

Example:
```markdown
Transition it through these states:
- Start it
- Check status (running)
- Pause it
- Check status (paused)
```

### Type 3: Build & Verify
"Create something, test it works, verify specific behavior"

Example:
```markdown
Create a `Dockerfile` in this directory that runs the simple-web app.

Your goal: build an image, run a container from it, and access the app at http://localhost:8080

**Success criteria:** You see JSON with app info when you visit the URL.
```

### Type 4: Security Challenge
"Try to break it, then fix it"

Example:
```markdown
Try to modify the application code from inside the container.
[Shows attack commands]
What do you see? Were you able to modify the running application?

**Your challenge:** Secure your image so that an attacker who compromises the app can't modify the application files.
```

### Type 5: Multi-Part Scenarios
"Set up a realistic architecture with multiple components"

Example:
```markdown
Create a scenario with three containers:
- frontend (use your simple-web image)
- backend (your api image)
- database (postgres)

**Requirements:**
- frontend can reach backend
- backend can reach database
- frontend CANNOT reach database
```

## Section Templates

### Prerequisites Section
```markdown
## Prerequisites

You should have completed exercise XX and have a `[artifact]` [ready/built]. If not, go back and [action] first.
```

### Resources Section
```markdown
## Resources

- [Link text](URL)
- [Link text](URL)
```

### Verification Section
```markdown
## Verification Checklist

You've completed this exercise when:
- [ ] [Observable outcome 1]
- [ ] [Observable outcome 2]
- [ ] [Observable outcome 3]
```

### Cleanup Section
```markdown
## Cleanup

[Specific commands or instructions to clean up]
```

## Content Patterns

### Introducing New Concepts
1. Show the problem first
2. Let them experience the pain point
3. Then introduce the solution
4. Have them implement it
5. Reflect on why it matters

### Hints
- Use sparingly
- Don't give the answer
- Point to documentation or a concept
- Ask a guiding question

Good hint:
```markdown
**Hint:** The app needs the version as an environment variable at runtime. But you want to set it at build time. What do you do?
```

Bad hint:
```markdown
**Hint:** Use ARG in your Dockerfile
```

### Code Examples
- Provide complete, runnable commands when showing setup
- Don't provide solutions to the actual challenges
- Use realistic values (not foo/bar)
- Include comments only when necessary

### Questions
Use questions to:
- Prompt observation: "What do you see?"
- Encourage comparison: "What's different about these?"
- Drive reflection: "Why is that a security risk?"
- Guide discovery: "What files does Node.js actually need to run your app?"

## Anti-Patterns to Avoid

❌ Don't say "simply" or "just" - it's condescending
❌ Don't give step-by-step solutions - let them figure it out
❌ Don't use passive voice - "The container should be started" → "Start the container"
❌ Don't apologize or hedge - "You might want to try" → "Try"
❌ Don't use academic language - "utilize" → "use"
❌ Don't explain everything upfront - let discovery happen

## Quality Checklist

Before finalizing an exercise, verify:
- [ ] Every task has an observable success criteria
- [ ] Instructions are action-oriented (start with verbs)
- [ ] No direct solutions are given
- [ ] Reflection questions encourage deeper thinking
- [ ] Tasks build progressively in complexity
- [ ] Code examples are complete and runnable
- [ ] Formatting is consistent (bold labels, code formatting, bullets)
- [ ] Verification checklist covers all major objectives
- [ ] Resources are relevant and linked
- [ ] Cleanup instructions are clear
