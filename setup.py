#!/usr/bin/env python3
"""Setup script for breeze Python package."""

from setuptools import setup, find_packages
from setuptools.command.install import install
import subprocess
import os
import sys
import platform


class PostInstallCommand(install):
    """Post-installation for building the Go binary."""
    
    def run(self):
        install.run(self)
        # Build the Go binary after installation
        try:
            print("Building Breeze Go binary...")
            subprocess.check_call(["go", "build", "./cmd/breeze"], 
                                cwd=os.path.dirname(os.path.abspath(__file__)))
            print("Breeze binary built successfully!")
        except subprocess.CalledProcessError as e:
            print(f"Warning: Failed to build Go binary: {e}")
            print("You may need to run 'go build ./cmd/breeze' manually")
        except FileNotFoundError:
            print("Warning: Go compiler not found. Please install Go from https://golang.org/")
            print("Then run 'go build ./cmd/breeze' in the package directory")


with open("README.md", "r", encoding="utf-8") as fh:
    long_description = fh.read()


setup(
    name="breeze-ai",
    version="2.0.0",
    author="Yossi Deutsch",
    author_email="",
    description="Ultra-simple local LLM interactions via Ollama",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/yossideutsch1973/breeze",
    packages=find_packages(where="python"),
    package_dir={"": "python"},
    classifiers=[
        "Development Status :: 4 - Beta",
        "Intended Audience :: Developers",
        "Topic :: Software Development :: Libraries :: Python Modules",
        "Topic :: Scientific/Engineering :: Artificial Intelligence",
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Programming Language :: Python :: 3.11",
        "Programming Language :: Python :: 3.12",
        "Operating System :: OS Independent",
    ],
    python_requires=">=3.7",
    install_requires=[],
    extras_require={
        "dev": ["pytest>=6.0", "black", "flake8"],
    },
    entry_points={
        "console_scripts": [
            "breeze=breeze.cli:main",
        ],
    },
    cmdclass={
        "install": PostInstallCommand,
    },
    include_package_data=True,
    package_data={
        "breeze": ["../breeze", "../breeze.exe"],
    },
)
