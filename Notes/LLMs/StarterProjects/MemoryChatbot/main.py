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

# Simple REPL
print("Chatbot ready! Type 'exit' to quit.\n")
while True:
    user = input("You: ").strip()
    if user.lower() in ["exit", "quit"]:
        break
    if user:
        print("Bot:", chat_turn(user))
        print()  # blank line for readability