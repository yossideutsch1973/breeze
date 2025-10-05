#!/usr/bin/env python3
"""
Demo script showing Python API usage for Breeze.

This demonstrates the basic functionality of the breeze Python package.
Note: Requires Ollama to be running and breeze binary to be built.
"""

import breeze

def main():
    print("=" * 60)
    print("Breeze Python Package Demo")
    print("=" * 60)
    print()
    
    # Example 1: Simple AI query
    print("Example 1: Simple AI Query")
    print("-" * 60)
    try:
        response = breeze.ai("What is 2+2? Give a very brief answer.")
        print(f"Prompt: What is 2+2?")
        print(f"Response: {response}")
    except breeze.BreezeError as e:
        print(f"Error: {e}")
        print("\nMake sure:")
        print("1. Ollama is installed and running")
        print("2. Go binary is built: go build ./cmd/breeze")
        return
    
    print()
    
    # Example 2: Chat (maintains context)
    print("Example 2: Chat with Context")
    print("-" * 60)
    try:
        # First message
        response1 = breeze.chat("My name is Alice. Remember this.")
        print(f"User: My name is Alice. Remember this.")
        print(f"AI: {response1[:100]}..." if len(response1) > 100 else f"AI: {response1}")
        
        # Second message (should remember name)
        response2 = breeze.chat("What is my name?")
        print(f"\nUser: What is my name?")
        print(f"AI: {response2}")
    except breeze.BreezeError as e:
        print(f"Error: {e}")
    
    print()
    
    # Example 3: Code generation
    print("Example 3: Code Generation")
    print("-" * 60)
    try:
        code = breeze.code("Write a Python function to calculate fibonacci")
        print(f"Prompt: Write a Python function to calculate fibonacci")
        print(f"Generated code:\n{code}")
    except breeze.BreezeError as e:
        print(f"Error: {e}")
    
    print()
    
    # Example 4: Batch processing
    print("Example 4: Batch Processing")
    print("-" * 60)
    try:
        prompts = ["What is Python?", "What is Go?"]
        results = breeze.batch(prompts)
        for i, (prompt, result) in enumerate(zip(prompts, results), 1):
            print(f"\n{i}. Prompt: {prompt}")
            print(f"   Response: {result[:80]}..." if len(result) > 80 else f"   Response: {result}")
    except breeze.BreezeError as e:
        print(f"Error: {e}")
    
    print()
    
    # Clear conversation
    print("Clearing conversation history...")
    breeze.clear()
    print("Done!")
    
    print()
    print("=" * 60)
    print("Demo complete!")
    print("=" * 60)


if __name__ == "__main__":
    main()
