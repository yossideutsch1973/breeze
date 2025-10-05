"""
Core Breeze module - Python wrapper for the Breeze Go binary.

This module provides a Pythonic interface to the Breeze Go library,
enabling local LLM interactions via Ollama with zero configuration.
"""

import subprocess
import sys
import os
import json
from typing import Optional, List, Callable
from pathlib import Path


class BreezeError(Exception):
    """Exception raised for errors in Breeze operations."""
    pass


def _find_breeze_binary() -> str:
    """
    Find the Breeze binary in the package directory or PATH.
    
    Returns:
        str: Path to the Breeze binary
        
    Raises:
        BreezeError: If binary cannot be found
    """
    # Check in package directory first
    package_dir = Path(__file__).parent.parent.parent
    binary_name = "breeze.exe" if sys.platform == "win32" else "breeze"
    binary_path = package_dir / binary_name
    
    if binary_path.exists() and binary_path.is_file():
        return str(binary_path)
    
    # Check in PATH
    try:
        result = subprocess.run(
            ["which" if sys.platform != "win32" else "where", "breeze"],
            capture_output=True,
            text=True,
            check=True
        )
        return result.stdout.strip().split('\n')[0]
    except (subprocess.CalledProcessError, FileNotFoundError):
        pass
    
    raise BreezeError(
        "Breeze binary not found. Please ensure Go is installed and run: "
        "go build ./cmd/breeze"
    )


def _run_breeze_command(args: List[str], input_text: Optional[str] = None) -> str:
    """
    Run a Breeze command and return the output.
    
    Args:
        args: Command line arguments for Breeze
        input_text: Optional input text to send to stdin
        
    Returns:
        str: Command output
        
    Raises:
        BreezeError: If command fails
    """
    try:
        binary = _find_breeze_binary()
        result = subprocess.run(
            [binary] + args,
            capture_output=True,
            text=True,
            input=input_text,
            check=True
        )
        return result.stdout.strip()
    except subprocess.CalledProcessError as e:
        error_msg = e.stderr if e.stderr else str(e)
        raise BreezeError(f"Breeze command failed: {error_msg}")
    except FileNotFoundError as e:
        raise BreezeError(f"Failed to execute Breeze: {e}")


def ai(prompt: str, model: Optional[str] = None, temperature: Optional[float] = None,
       concise: bool = False, docs: Optional[List[str]] = None) -> str:
    """
    Generate a response for a single prompt using AI.
    
    Args:
        prompt: The prompt to send to the AI
        model: Optional model name (e.g., "codellama", "llama2")
        temperature: Optional temperature setting (0.0 to 1.0)
        concise: If True, request concise responses
        docs: Optional list of document file paths to process
        
    Returns:
        str: AI-generated response
        
    Raises:
        BreezeError: If the operation fails
        
    Example:
        >>> response = breeze.ai("Explain recursion")
        >>> print(response)
    """
    args = [prompt]
    
    # Note: The Go CLI doesn't support all options via command line
    # For full feature support, users should use the Go library directly
    # This wrapper provides basic functionality
    
    return _run_breeze_command(args)


def chat(prompt: str) -> str:
    """
    Send a message in a conversational context (maintains history).
    
    Args:
        prompt: The message to send
        
    Returns:
        str: AI response
        
    Raises:
        BreezeError: If the operation fails
        
    Example:
        >>> breeze.chat("Hello!")
        >>> breeze.chat("Tell me more")
    """
    return _run_breeze_command(["chat", prompt])


def code(prompt: str) -> str:
    """
    Generate code based on a prompt (optimized for code generation).
    
    Args:
        prompt: Description of the code to generate
        
    Returns:
        str: Generated code
        
    Raises:
        BreezeError: If the operation fails
        
    Example:
        >>> code = breeze.code("Write a factorial function in Python")
        >>> print(code)
    """
    return _run_breeze_command(["code", prompt])


def clear() -> None:
    """
    Clear the conversation history.
    
    Example:
        >>> breeze.clear()
    """
    _run_breeze_command(["clear"])


def stream(prompt: str, callback: Optional[Callable[[str], None]] = None) -> str:
    """
    Stream AI responses token by token.
    
    Note: The Python wrapper currently doesn't support true streaming.
    This method returns the complete response. For true streaming,
    use the Go library directly.
    
    Args:
        prompt: The prompt to send to the AI
        callback: Optional callback function (not currently used)
        
    Returns:
        str: Complete AI response
        
    Example:
        >>> response = breeze.stream("Write a story")
        >>> print(response)
    """
    # Note: True streaming requires more complex subprocess handling
    # For now, return complete response
    return ai(prompt)


def batch(prompts: List[str]) -> List[str]:
    """
    Process multiple prompts in batch.
    
    Args:
        prompts: List of prompts to process
        
    Returns:
        List[str]: List of responses corresponding to each prompt
        
    Example:
        >>> results = breeze.batch(["Explain AI", "Explain ML"])
        >>> for result in results:
        ...     print(result)
    """
    return [ai(prompt) for prompt in prompts]
