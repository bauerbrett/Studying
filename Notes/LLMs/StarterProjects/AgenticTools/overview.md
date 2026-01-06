```
# Local LLM Agent Tool Flow Guide

This document explains how **tool calling** works in your custom, low-level, fully local LLM agent built with `llama-cpp-python`. No frameworks --- just raw text parsing and prompt engineering.

## Overview

Your agent can:
- Answer questions using its own knowledge
- Retrieve relevant context from a local RAG knowledge base
- Use **tools** (like getting current time or weather) when needed
- Remember conversation history across turns

The **tool system** is built entirely from scratch using text parsing --- no structured tool-calling endpoints.

## How Tool Calling Works (Step-by-Step Flow)

### 1. User Asks a Question
Example:
`what is the weather in cincinnati`

### 2. Agent Builds the Prompt
The prompt sent to the LLM includes:
- System prompt
- Conversation history
- RAG context (if any)
- **Tool instructions** appended to the user message

```text
[system prompt]
[previous conversation]

User: what is the weather in cincinnati

You have access to tools. When you need real-time info like weather or time, you MUST use a tool.

To use a tool, respond with ONLY valid JSON:
{"tool": "get_weather", "args": {"location": "cincinnati"}}

Available tools:
- get_current_time()
- get_weather(location: str)
```

### 3\. LLM Decides to Use a Tool

The model outputs **only** JSON:

JSON

```
{"tool": "get_weather", "args": {"location": "cincinnati"}}
```

(No extra text --- critical for reliable parsing)

### 4\. Your Code Parses the JSON

Python

```
tool_call = try_parse_tool(response)
```

-   Finds the first { and last }
-   Extracts the JSON string
-   Parses it with json.loads()

If successful → tool_call = {'tool': 'get_weather', 'args': {'location': 'cincinnati'}}

### 5\. Tool Is Executed

Python

```
func = tools["get_weather"]
result = func(**args)  # → "Weather in cincinnati: 72°F, sunny (simulated)"
```

### 6\. Tool Result Is Fed Back to the Model

A new prompt is built:

text

```
[previous history including the JSON tool call]

User: Tool result: Weather in cincinnati: 72°F, sunny (simulated)

Assistant:
```

The model now generates the **final natural answer**:

text

```
The weather in Cincinnati is 72°F and sunny.
```

### 7\. Final Answer Returned to User

No JSON → loop ends → this response is shown.

Full Flow Diagram
-----------------

text

```
User Input
   ↓
agent_chat()
   ↓
Add tool instructions to prompt
   ↓
LLM generates response
   ↓
Is it JSON tool call? → try_parse_tool()
   ├─ No  → Return response (final answer)
   └─ Yes → Execute tool
             ↓
         Feed result back as "Tool result: ..."
             ↓
         LLM generates next response
             ↓
         Loop back to parsing
```

Why This Works So Well
----------------------

-   **Full control**: You see and manage every step
-   **Robust**: Handles malformed output gracefully
-   **Transparent**: All tool calls are visible in history
-   **Flexible**: Easy to add new tools
-   **Local & private**: No external APIs

What's Saved in History
-----------------------

Every step is preserved:

-   Original user question + tool instructions
-   Raw JSON tool call (as assistant message)
-   Tool result message
-   Final natural language answer

This helps the model learn the tool-use pattern over time.

Example Full History After One Tool Call
----------------------------------------

text

```
User: what is the weather in cincinnati

[tool instructions]

Assistant: {"tool": "get_weather", "args": {"location": "cincinnati"}}

User: Tool result: Weather in cincinnati: 72°F, sunny (simulated)

Assistant: The weather in Cincinnati is 72°F and sunny.
```

Perfect for multi-turn reasoning.