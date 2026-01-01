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