"""Tests for Breeze Python package."""

import pytest
import subprocess
from breeze import breeze


def test_breeze_error():
    """Test BreezeError exception."""
    error = breeze.BreezeError("Test error")
    assert str(error) == "Test error"
    assert isinstance(error, Exception)


def test_find_breeze_binary_raises_when_not_found(monkeypatch):
    """Test that _find_breeze_binary raises when binary not found."""
    # Mock Path.exists to return False
    from pathlib import Path
    original_exists = Path.exists
    
    def mock_exists(self):
        # Return False for breeze binary checks
        if "breeze" in str(self):
            return False
        return original_exists(self)
    
    monkeypatch.setattr(Path, "exists", mock_exists)
    
    # Mock subprocess.run to raise FileNotFoundError
    def mock_run(*args, **kwargs):
        raise subprocess.CalledProcessError(1, "which")
    
    import subprocess as sp
    monkeypatch.setattr(sp, "run", mock_run)
    
    with pytest.raises(breeze.BreezeError) as exc_info:
        breeze._find_breeze_binary()
    
    assert "Breeze binary not found" in str(exc_info.value)


# Note: Integration tests that actually call the binary would require
# Ollama to be running. These should be marked as integration tests
# and skipped in CI if Ollama is not available.

@pytest.mark.skip(reason="Requires Breeze binary and Ollama")
def test_ai_basic():
    """Test basic AI functionality (integration test)."""
    response = breeze.ai("Hello")
    assert isinstance(response, str)
    assert len(response) > 0


@pytest.mark.skip(reason="Requires Breeze binary and Ollama")
def test_chat_basic():
    """Test basic chat functionality (integration test)."""
    response = breeze.chat("Hello")
    assert isinstance(response, str)
    assert len(response) > 0


@pytest.mark.skip(reason="Requires Breeze binary and Ollama")
def test_code_basic():
    """Test basic code functionality (integration test)."""
    response = breeze.code("factorial function")
    assert isinstance(response, str)
    assert len(response) > 0


@pytest.mark.skip(reason="Requires Breeze binary and Ollama")
def test_clear():
    """Test clear functionality (integration test)."""
    breeze.clear()  # Should not raise


@pytest.mark.skip(reason="Requires Breeze binary and Ollama")
def test_batch():
    """Test batch functionality (integration test)."""
    results = breeze.batch(["Hello", "World"])
    assert isinstance(results, list)
    assert len(results) == 2
