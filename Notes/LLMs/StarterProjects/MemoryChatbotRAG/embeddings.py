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