import json
import datetime
from llama_cpp import Llama
import numpy as np
import os


'''
Few simple tools to feed weather and date and time back to LLM. 

'''


# This is just to get rid of annoying pop ups llama gives in the terminal
import sys, os
sys.stderr = open(os.devnull, 'w')

# Load models
print("Loading models...")
llm = Llama(
    model_path="/home/brett/studying/Notes/LLMs/StarterProjects/models/Meta-Llama-3.1-8B-Instruct-Q5_K_M.gguf",
    n_ctx=8192,
    n_threads=8,
    n_gpu_layers=-1,
    verbose=False,
    embedding=True
)
embed_llm = Llama(
    model_path="/home/brett/studying/Notes/LLMs/StarterProjects/models/nomic-embed-text-v1.5.f32.gguf",
    n_ctx=8192,
    n_threads=8,
    n_gpu_layers=-1,
    verbose=False,
    embedding=True
)



# Suppress llama.cpp warnings
os.environ["LLAMA_CPP_LOG_LEVEL"] = "ERROR"


# Load RAG data
print("Loading knowledge base...")
chunks = np.load("chunks.npy", allow_pickle=True).tolist()
embeddings = np.load("embeddings.npy")  # (N, 768)
print(f"Loaded {len(chunks)} chunks")

# Tools
def get_current_time() -> str:
    return datetime.datetime.now().strftime("%Y/%m/%d, %H:%M:%S")

def get_weather(location: str) -> str:
    return f"Weather in {location}: 72°F, sunny (simulated)"

tools = {
    "get_current_time": get_current_time,
    "get_weather": get_weather,
}

tool_instructions = """
You have access to tools. When you need real-time or external information (like current time or weather), you MUST use a tool.

To use a tool, respond with ONLY this exact JSON format — no extra text, no explanation:

{"tool": "tool_name", "args": {"param": "value"}}

Available tools:
- get_current_time(): returns current date and time
- get_weather(location: str): returns current weather for a city

Examples:
User: What time is it?
→ {"tool": "get_current_time", "args": {}}

User: What's the weather in Paris?
→ {"tool": "get_weather", "args": {"location": "Paris"}}

Only respond normally if no tool is required.
"""

history = []
system_prompt = (
    "You are a helpful, concise assistant. Use provided context when relevant. "
    "Answer directly. Never mention sources."
)

def format_message(role: str, content: str) -> str:
    return f"<|start_header_id|>{role}<|end_header_id|>\n\n{content}<|eot_id|>"

def embed_query(query: str) -> np.ndarray:
    text = "search_query: " + query
    result = embed_llm.embed(text)
    vec = np.array(result, dtype=np.float32).flatten()
    vec /= np.linalg.norm(vec) + 1e-8
    return vec

def retrieve(query: str, k: int = 3) -> str:
    if not query.strip():
        return ""
    query_vec = embed_query(query)
    scores = np.dot(embeddings, query_vec)
    top_k = np.argsort(scores)[-k:][::-1]
    context_parts = []
    for i in top_k:
        if scores[i] > 0.3:  # threshold
            context_parts.append(chunks[i])
    if not context_parts:
        return ""
    return "\n\n".join(context_parts)

def try_parse_tool(text: str):
    text = text.strip()
    # Look for JSON block
    start = text.find("{") # finds position of first {
    if start == -1:
        return None
    end = text.rfind("}") + 1 # finds last }
    if end == 0:
        return None
    try:
        json_str = text[start:end] # extracts {"tool": ...}
        return json.loads(json_str) # → {'tool': 'get_weather', 'args': {...}}
    except:
        return None

def chat_turn(user_input: str, rag_context: str = "") -> str:
    global history

    full_prompt = format_message("system", system_prompt)

    # Add RAG context as system message if available
    if rag_context:
        full_prompt += format_message("system", f"Use this context if relevant:\n{rag_context}")

    # Add conversation history
    for msg in history:
        full_prompt += msg

    # Add current user message
    full_prompt += format_message("user", user_input)
    full_prompt += "<|start_header_id|>assistant<|end_header_id|>\n\n"

    output = llm(
        full_prompt,
        max_tokens=512,
        temperature=0.7,
        top_p=0.95,
        stop=["<|eot_id|>", "<|end_of_text|>"],
        echo=False
    )

    response = output["choices"][0]["text"].strip()

    # Save to history
    history.append(format_message("user", user_input))
    history.append(format_message("assistant", response))

    if len(history) > 40:
        history = history[-40:]

    return response

def agent_chat(user_input: str) -> str:
    rag_context = retrieve(user_input)

    # Strong prompt: force tool use when needed
    enhanced_input = user_input + "\n\n" + tool_instructions

    response = chat_turn(enhanced_input, rag_context)

    # Tool loop with better JSON detection
    for _ in range(5):
        tool_call = try_parse_tool(response)

        if not tool_call or "tool" not in tool_call:
            return response  # Final answer

        tool_name = tool_call["tool"]
        args = tool_call.get("args", {})

        if tool_name not in tools:
            result = f"Error: Unknown tool {tool_name}"
        else:
            try:
                func = tools[tool_name]
                result = str(func(**args) if isinstance(args, dict) else func(args))
            
            except Exception as e:
                result = f"Error: {e}"

        # Feed result and continue
        response = chat_turn(f"Observation: {result}")

    return response

'''
The Conditional Expression (Ternary)
Pythonfunc(**args) if isinstance(args, dict) else func(args)
            
This is a conditional expression (Python’s version of a ternary operator). It means:
    "If args is a dictionary, call the function with keyword unpacking (**args).
Otherwise, call it with positional arguments (args)."
Why Both Options?
    Because the JSON from the LLM could come in two different formats:
    Case A: Keyword arguments (most common and correct)
        JSON{"tool": "get_weather", "args": {"location": "cincinnati"}}
        → args is a dict → use **args → equivalent to get_weather(location="cincinnati")
    Case B: Positional arguments (less common, but possible)
        JSON{"tool": "some_tool", "args": ["value1", "value2"]}
        → args is a list → use func(args) → equivalent to some_tool("value1", "value2")


Example 1: get_weather (your actual tool)
Pythonargs = {"location": "cincinnati"}  # dict
isinstance(args, dict) → True
→ func(**args) → get_weather(location="cincinnati")
→ returns "Weather in cincinnati: 72°F, sunny (simulated)"
→ str(...) → same string

Example 2: Hypothetical tool that takes two positional args
Pythondef add_numbers(a, b):
    return a + b

# LLM outputs:
{"tool": "add_numbers", "args": [5, 7]}

→ args = [5, 7] → not a dict
→ func(args) → add_numbers(5, 7) → 12
→ str(12) → "12"
'''
# REPL
print("\nLocal RAG + Tools + Memory Chatbot Ready!\nType 'exit' to quit.\n")
while True:
    user = input("You: ").strip()
    if user.lower() in ["exit", "quit", "bye"]:
        print("Goodbye!")
        break
    if user:
        print("Bot:", agent_chat(user))
        print()