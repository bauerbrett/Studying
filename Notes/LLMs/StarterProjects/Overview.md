Low-Level Local LLM Chatbot Guide: Building from Scratch 
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
from llama_cpp import Llama

llm = Llama(
    model_path="/home/brett/studying/Notes/LLMs/StarterProjects/models/Meta-Llama-3.1-8B-Instruct-Q5_K_M.gguf",
    n_ctx=8192,
    n_threads=8,
    n_gpu_layers=-1,
    verbose=False
)

history = []  # Stores formatted message blocks: user and assistant turns
system_prompt = "You are a helpful, concise assistant. Answer directly with only essential information."

def format_message(role: str, content: str) -> str:
    return f"<|start_header_id|>{role}<|end_header_id|>\n\n{content}<|eot_id|>"

def chat_turn(user_input: str) -> str:
    global history
    
    # Build full prompt in correct order
    full_prompt = format_message("system", system_prompt)
    
    # Add all previous conversation history
    for msg in history:
        full_prompt += msg
    
    # Add current user message
    full_prompt += format_message("user", user_input)
    
    # Start assistant response
    full_prompt += "<|start_header_id|>assistant<|end_header_id|>\n\n"

    # Generate
    output = llm(
        full_prompt,
        max_tokens=512,
        temperature=0.7,
        top_p=0.95,
        stop=["<|eot_id|>", "<|end_of_text|>"],
        echo=False
    )
    
    response = output["choices"][0]["text"].strip()
    
    # Update history
    history.append(format_message("user", user_input))
    history.append(format_message("assistant", response))
    
    # Optional: limit history to prevent context overflow
    if len(history) > 40:  # ~20 full turns
        history = history[-40:]
    
    return response

```

**Terminal Loop**:

Python

```
# Simple REPL
print("Chatbot ready! Type 'exit' to quit.\n")
while True:
    user = input("You: ").strip()
    if user.lower() in ["exit", "quit"]:
        break
    if user:
        print("Bot:", chat_turn(user))
        print()  # blank line for readability
```

**How memory works**: Full history is in every prompt. Truncation prevents overflow. This is the "raw" way---no vector stores or summaries yet.

Step 4: Building RAG (Retrieval-Augmented Generation)
-----------------------------------------------------

Use the **same LLM** for embeddings. Store chunks in a vector DB (simple NumPy or FAISS).

Python

```
import numpy as np
from llama_cpp import Llama

llm = Llama(
    model_path="/home/brett/studying/Notes/LLMs/StarterProjects/models/nomic-embed-text-v1.5.f32.gguf",
    n_threads=8,
    n_gpu_layers=-1,
    embedding=True,
    verbose=False
)

# Load text and chunk
with open("facts.txt", "r", encoding="utf-8") as f:
    full_text = f.read()

def create_chunks(text, size=100, overlap=20):
    words = text.split()
    chunks = []
    for i in range(0, len(words), size - overlap):
        chunk = " ".join(words[i:i+size])
        if chunk.strip():
            chunks.append(chunk)
    return chunks

chunks = create_chunks(full_text)
print(f"Created {len(chunks)} chunks")

embeddings = []

for chunk in chunks:
    text = "search_document: " + chunk
    vec = np.array(llm.embed(text), dtype=np.float32).flatten()
    vec /= np.linalg.norm(vec)
    embeddings.append(vec)

embeddings = np.vstack(embeddings)

# Save chunks + embeddings
np.save("chunks.npy", np.array(chunks, dtype=object))
np.save("embeddings.npy", embeddings)

print("Embeddings saved.")
print("Chunks shape:", np.load("chunks.npy", allow_pickle=True).shape)
print("Embeddings shape:", embeddings.shape)
```

**Integrate into Chat**:

Python

```
from llama_cpp import Llama
import numpy as np

# This is just to get rid of annoying pop ups llama gives in the terminal
import sys, os
sys.stderr = open(os.devnull, 'w')

llm = Llama(
    model_path="/home/brett/studying/Notes/LLMs/StarterProjects/models/Meta-Llama-3.1-8B-Instruct-Q5_K_M.gguf",
    n_ctx=8192,
    n_threads=8,
    n_gpu_layers=-1,
    verbose=False,
    embedding=True
)
llm_embed = Llama(
    model_path="/home/brett/studying/Notes/LLMs/StarterProjects/models/nomic-embed-text-v1.5.f32.gguf",
    n_ctx=8192,
    n_threads=8,
    n_gpu_layers=-1,
    verbose=False,
    embedding=True
)

history = []  # Stores formatted message blocks: user and assistant turns
system_prompt = "You are a helpful, concise assistant. You will be given context for some questions" \
", use the context to answer. If the context does not apply use your own judgment. Do not put " \
"in the response where you got the answer from, just answer directly with only essential information."

print("Loading documentation into chatbot....")
chunks = np.load("chunks.npy", allow_pickle=True)
embeddings = np.load("embeddings.npy")  # shape: (N, 4096)
print("Documents loaded")


def format_message(role: str, content: str) -> str:
    return f"<|start_header_id|>{role}<|end_header_id|>\n\n{content}<|eot_id|>"

def embed_query(query: str) -> np.ndarray:
    # Add the model-specific prefix
    text = "search_query: " + query

    # Call embed
    result = llm_embed.embed(text)

    # If returned as list of lists, flatten and convert to float32
    if isinstance(result, dict):
        vec = np.array(result["data"][0]["embedding"], dtype=np.float32)
    else:
        vec = np.array(result, dtype=np.float32).flatten()

    # L2 normalize
    vec /= np.linalg.norm(vec)
    return vec

def retrieve(query, k=2, min_score=0.2):
    query_vec = embed_query(query)
    scores = embeddings @ query_vec
    top_idx = np.argsort(scores)[::-1]
    results = []
    for idx in top_idx[:k]:
        if scores[idx] < min_score:
            continue
        results.append(f"score: {float(scores[idx])} text: {chunks[idx]}")
    return results

def chat_turn(user_input: str) -> str:
    global history
    
    # Build full prompt in correct order
    full_prompt = format_message("system", system_prompt)
    
    # Add all previous conversation history
    for msg in history:
        full_prompt += msg
    
    rag_result = retrieve(user_input)

    # Add current user message
    full_prompt += format_message("user", user_input + "RAG context: " + " ".join(rag_result))
    
    # Start assistant response
    full_prompt += "<|start_header_id|>assistant<|end_header_id|>\n\n"

    # Generate
    output = llm(
        full_prompt,
        max_tokens=512,
        temperature=0.7,
        top_p=0.95,
        stop=["<|eot_id|>", "<|end_of_text|>"],
        echo=False
    )
    
    response = output["choices"][0]["text"].strip()
    
    # Update history
    history.append(format_message("user", user_input))
    history.append(format_message("assistant", response))
    
    # Optional: limit history to prevent context overflow
    if len(history) > 40:  # ~20 full turns
        history = history[-40:]
    
    return response


# Simple REPL
print("Chatbot ready! Type 'exit' to quit.\n")
while True:
    user = input("You: ").strip()
    if user.lower() in ["exit", "quit"]:
        break
    if user:
        print("Bot:", chat_turn(user))
        print()  # blank line for readability
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
-   Model extraction: Economically infeasible.


Full Integration & Extensions
-----------------------------

Combine all: Use memory + RAG context in agent prompt. Add better stopping, streaming, or evaluation.

This raw approach exposes LLM realities. Build, break, iterate---you'll understand deeply.