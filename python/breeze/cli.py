#!/usr/bin/env python3
"""
Breeze CLI - Command-line interface for Breeze Python package.

This provides a command-line interface to the Breeze Python wrapper,
which itself wraps the Breeze Go binary.
"""

import sys
import argparse
from . import breeze
from . import __version__


def main():
    """Main CLI entry point."""
    parser = argparse.ArgumentParser(
        description="Breeze - Ultra-simple local LLM interactions"
    )
    parser.add_argument(
        "command",
        nargs="?",
        help="Command: chat, code, clear, or a direct prompt"
    )
    parser.add_argument(
        "prompt",
        nargs="*",
        help="Prompt text or additional arguments"
    )
    parser.add_argument(
        "--version",
        action="version",
        version=f"Breeze {__version__}"
    )
    
    args = parser.parse_args()
    
    try:
        if not args.command:
            parser.print_help()
            sys.exit(0)
        
        # Handle different commands
        if args.command == "chat":
            if not args.prompt:
                print("Error: chat requires a prompt")
                sys.exit(1)
            prompt = " ".join(args.prompt)
            response = breeze.chat(prompt)
            print(response)
        
        elif args.command == "code":
            if not args.prompt:
                print("Error: code requires a prompt")
                sys.exit(1)
            prompt = " ".join(args.prompt)
            response = breeze.code(prompt)
            print(response)
        
        elif args.command == "clear":
            breeze.clear()
            print("Conversation cleared.")
        
        else:
            # Treat as a direct AI prompt
            if args.prompt:
                prompt = args.command + " " + " ".join(args.prompt)
            else:
                prompt = args.command
            response = breeze.ai(prompt)
            print(response)
    
    except breeze.BreezeError as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)
    except KeyboardInterrupt:
        print("\nInterrupted", file=sys.stderr)
        sys.exit(130)


if __name__ == "__main__":
    main()
