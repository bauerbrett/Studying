from llama_cpp import Llama

llm = Llama(
    model_path="/home/brett/studying/Notes/LLMs/StarterProjects/models/Meta-Llama-3.1-8B-Instruct-Q5_K_M.gguf",
    n_ctx=8192,
    n_threads=8,
    n_gpu_layers=-1,
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
    system_prompt = "You are a helpful, concise assistant. Answer questions directly with only the essential information, no extra lists or examples unless asked."

    formatted_prompt = f"<|start_header_id|>system<|end_header_id|>\n{system_prompt}<|eot_id|><|start_header_id|>user<|end_header_id|>\n{prompt}<|eot_id|><|start_header_id|>assistant<|end_header_id|>\n"
    print(formatted_prompt)
    output = llm(
        formatted_prompt,
        max_tokens=max_tokens,
        temperature=temperature,
        top_p=top_p,
        stop=stop,
        echo=False
    )
    return output

print(complete("Answer briefly: What is the capital of France?"))