Low-Level Local LLM Chatbot Guide: Building from Scratch with Raw Completions
=============================================================================

Introduction and Philosophy
---------------------------

This guide teaches you to build a terminal-based chatbot **entirely locally**, without external services or high-level abstractions like LangChain. We stick to raw LLM mechanics: **text in, text out** (tokens, really). No "chat" endpoints, no multi-message histories passed as structured inputs---just a single string prompt and basic hyperparameters (temperature, top_p, max_tokens, stop sequences).

Why? Because understanding LLMs at this level reveals their true nature: sophisticated text parsers and predictors. Anything more is abstraction hiding details. By building memory, RAG, and tools manually, you'll see the "sharp edges" and why treating LLMs as untrusted interfaces is critical for security.

Key principles:

-   LLMs have **no inherent state** → Memory is manual prompt engineering.
-   Embeddings and generation use the **same model** for consistency.
-   Tools/agentic behavior: Parse text output yourself---no built-in tool calling.
-   Security: Prompts aren't secret; users can extract them. Tools are the real attack surface.

Everything runs on your machine using llama-cpp-python for inference on quantized GGUF models.

Step 1: Environment Setup
-------------------------

### Install Dependencies

Bash

```
pip install llama-cpp-python numpy
```

-   For GPU (CUDA example): CMAKE_ARGS="-DGGML_CUDA=on" pip install llama-cpp-python
-   Optional for faster search: pip install faiss-cpu

### Download a Model

Use quantized GGUF models (efficient for local run):

-   Starter: Llama 3.1 8B Instruct Q4_K_M (~5GB) → [https://huggingface.co/bartowski/Meta-Llama-3.1-8B-Instruct-GGUF](https://huggingface.co/bartowski/Meta-Llama-3.1-8B-Instruct-GGUF?referrer=grok.com)
-   Better: Qwen2.5 7B or Gemma2 9B/27B GGUF variants. Download the .gguf file to a models/ folder.

Step 2: Raw Completion Wrapper
------------------------------

Load the model and create a simple completion function.

Python

```
from llama_cpp import Llama

# Adjust n_ctx for larger context (e.g., 8192+), n_gpu_layers=-1 for full GPU offload
llm = Llama(
    model_path="models/Meta-Llama-3.1-8B-Instruct-Q4_K_M.gguf",
    n_ctx=8192,
    n_threads=8,
    n_gpu_layers=-1,  # Use all layers on GPU if available
    verbose=False
)

def complete(
    prompt: str,
    max_tokens: int = 512,
    temperature: float = 0.7,
    top_p: float = 0.95,
    stop: list = None
) -> str:
    if stop is None:
        stop = ["\nUser:", "</s>", "<|eot_id|>"]  # Model-specific stops
    output = llm(
        prompt,
        max_tokens=max_tokens,
        temperature=temperature,
        top_p=top_p,
        stop=stop,
        echo=False
    )
    return output["choices"][0]["text"].strip()
```

**How it works**: This calls the raw completion endpoint internally. Single prompt string + hypers only. Test it:

Python

```
print(complete("The capital of France is"))
```

Step 3: Adding Conversation Memory
----------------------------------

LLMs are stateless---memory is built by prepending history to the prompt.

Python

```
history = []  # List of "User: ..." and "Assistant: ..." strings

def build_prompt(user_input: str) -> str:
    global history
    history.append(f"User: {user_input}")
    return "\n".join(history) + "\nAssistant:"

def chat_with_memory(user_input: str) -> str:
    global history
    prompt = build_prompt(user_input)
    response = complete(prompt, stop=["\nUser:", "</s>", "<|eot_id|>"])
    history.append(f"Assistant: {response}")

    # Prevent context overflow: Keep last ~6000 tokens (rough estimate)
    if len("\n".join(history)) > 6000:
        history = history[-10:]  # Keep recent turns

    return response
```

**Terminal Loop**:

Python

```
print("Local Chatbot Ready (type 'quit' to exit)")
while True:
    user = input("\nYou: ")
    if user.lower() == "quit":
        break
    print("\nBot:", chat_with_memory(user))
```

**How memory works**: Full history is in every prompt. Truncation prevents overflow. This is the "raw" way---no vector stores or summaries yet.

Step 4: Building RAG (Retrieval-Augmented Generation)
-----------------------------------------------------

Use the **same LLM** for embeddings. Store chunks in a vector DB (simple NumPy or FAISS).

Python

```
import numpy as np
from typing import List

documents: List[str] = []
embeddings: List[np.ndarray] = []

def get_embedding(text: str) -> np.ndarray:
    emb = llm.embed(text)  # Built-in embedding support
    return np.array(emb)

def add_document(text: str, chunk_size: int = 512, overlap: int = 100):
    chunks = [text[i:i+chunk_size] for i in range(0, len(text), chunk_size - overlap)]
    for chunk in chunks:
        documents.append(chunk)
        embeddings.append(get_embedding(chunk))

def cosine_similarity(a: np.ndarray, b: np.ndarray) -> float:
    return np.dot(a, b) / (np.linalg.norm(a) * np.linalg.norm(b))

def retrieve(query: str, k: int = 4) -> List[str]:
    query_emb = get_embedding(query)
    scores = [cosine_similarity(query_emb, emb) for emb in embeddings]
    top_indices = np.argsort(scores)[-k:][::-1]
    return [documents[i] for i in top_indices]
```

**Load Knowledge** (example):

Python

```
with open("knowledge.txt", "r", encoding="utf-8") as f:
    add_document(f.read())
```

**Integrate into Chat**:

Python

```
def chat_with_rag(user_input: str) -> str:
    relevant_chunks = retrieve(user_input)
    context = "\n\n".join(relevant_chunks)
    rag_prompt = f"""Use only the following context to answer the question.
Context:
{context}

Question: {user_input}
Answer:"""
    return complete(rag_prompt, max_tokens=768)
```

**How RAG works here**: Semantic search via embeddings from the same model. No external embedder. Retrieval injects context into a single prompt. Combine with memory for full chatbot.

For scale: Use FAISS instead of list for O(1) search.

Step 5: Building a Custom Agentic System with Tools
---------------------------------------------------

No native tool calling---LLM outputs structured text (e.g., JSON). You parse and execute.

**Define Tools**:

Python

```
import json
import datetime

def get_current_time() -> str:
    return datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")

def get_weather(location: str) -> str:
    # Simulated---replace with real local API if desired
    return f"Weather in {location}: 72°F, sunny (simulated data)"

tools = {
    "get_current_time": get_current_time,
    "get_weather": get_weather
}

tool_instructions = """
You have access to tools. If needed, output ONLY a JSON object:
{"tool": "tool_name", "args": {"param": "value"}}

Available:
- get_current_time() → returns current time
- get_weather(location: str) → returns simulated weather

Otherwise, respond normally."""
```

**Agent Loop** (ReAct-style):

Python

```
def agent_chat(user_input: str) -> str:
    global history
    history.append(f"User: {user_input}")

    while True:
        prompt = "\n".join(history) + f"\n{tool_instructions}\nAssistant:"
        response = complete(prompt, max_tokens=256, stop=["\nUser:", "}"])

        # Attempt tool call
        if "{" in response and "}" in response:
            try:
                json_str = response[response.find("{"):response.rfind("}")+1]
                call = json.loads(json_str)
                tool_name = call["tool"]
                args = call.get("args", {})

                if tool_name in tools:
                    result = tools[tool_name](**args)
                    history.append(f"Assistant: {response}")
                    history.append(f"Tool Result: {result}")
                    continue  # Feed result back for reasoning
            except json.JSONDecodeError:
                pass

        # No valid tool or final answer
        history.append(f"Assistant: {response}")
        return response
```

**How tools work**: LLM "decides" via prompted JSON. You safely parse/validate/execute. Never eval raw output. Add sandboxing for real tools (e.g., subprocess limits).

Security Considerations
-----------------------

-   **Prompts leak easily**: Users can jailbreak to reveal system instructions.
-   **LLM ≠ security control**: Sanitize all inputs/outputs like browser data.
-   **Tool interfaces are the risk**: Validate args strictly; rate-limit; log.
-   Model extraction: Economically infeasible (as you noted).
-   Red-team your build: Try prompt injection to abuse tools.

Full Integration & Extensions
-----------------------------

Combine all: Use memory + RAG context in agent prompt. Add better stopping, streaming, or evaluation.

This raw approach exposes LLM realities. Build, break, iterate---you'll understand deeply.