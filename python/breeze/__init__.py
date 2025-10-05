"""
Breeze - Ultra-simple local LLM interactions via Ollama.

A Python wrapper for the Breeze Go library that provides zero-configuration
AI queries, chat, code generation, streaming, and batch processing.

Example usage:
    >>> import breeze
    >>> response = breeze.ai("Explain quantum physics")
    >>> print(response)
    
    >>> # Chat with context
    >>> breeze.chat("Hello AI!")
    >>> breeze.chat("Tell me more about AI")
    
    >>> # Generate code
    >>> code = breeze.code("Write a factorial function in Python")
    >>> print(code)
"""

__version__ = "2.0.0"
__author__ = "Yossi Deutsch"

from .breeze import (
    ai,
    chat,
    code,
    clear,
    stream,
    batch,
    BreezeError,
)

__all__ = [
    "ai",
    "chat",
    "code",
    "clear",
    "stream",
    "batch",
    "BreezeError",
]
