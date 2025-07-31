def fibonacci(n):
    if n <= 0:
        return 0
    elif n == 1:
        return 1
    return fibonacci(n - 1) + fibonacci(n - 2)

# Example usage:
print(fibonacci(6))  # Output: 8

# 0, 1, 2, 3, 5, 8, 

