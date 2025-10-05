# Publishing Breeze to PyPI

This guide explains how to publish the `breeze-ai` package to PyPI so users can install it with `pip install breeze-ai`.

## Prerequisites

1. **PyPI Account**: Create accounts on:
   - [PyPI](https://pypi.org/account/register/) (production)
   - [TestPyPI](https://test.pypi.org/account/register/) (testing)

2. **API Tokens**: Generate API tokens for both accounts
   - PyPI: https://pypi.org/manage/account/token/
   - TestPyPI: https://test.pypi.org/manage/account/token/

3. **Required Tools**:
   ```bash
   pip install --upgrade build twine
   ```

## Building the Package

1. **Clean previous builds**:
   ```bash
   rm -rf dist/ build/ *.egg-info
   ```

2. **Build the Go binary**:
   ```bash
   go build ./cmd/breeze
   ```

3. **Build the Python package**:
   ```bash
   python -m build
   ```
   
   This creates:
   - `dist/breeze_ai-2.0.0-py3-none-any.whl` (wheel)
   - `dist/breeze-ai-2.0.0.tar.gz` (source distribution)

## Testing the Package

### Test on TestPyPI (recommended first)

1. **Upload to TestPyPI**:
   ```bash
   python -m twine upload --repository testpypi dist/*
   ```
   
   When prompted, use your TestPyPI username and API token.

2. **Test installation**:
   ```bash
   # Create a test virtual environment
   python -m venv test_env
   source test_env/bin/activate
   
   # Install from TestPyPI
   pip install --index-url https://test.pypi.org/simple/ breeze-ai
   
   # Test it
   python -c "import breeze; print(breeze.__version__)"
   
   # Cleanup
   deactivate
   rm -rf test_env
   ```

### Test locally before uploading

```bash
# Install locally from the built package
pip install dist/breeze_ai-2.0.0-py3-none-any.whl

# Or test in editable mode
pip install -e .
```

## Publishing to PyPI

Once you've tested on TestPyPI:

1. **Upload to PyPI**:
   ```bash
   python -m twine upload dist/*
   ```
   
   When prompted, use your PyPI username and API token.

2. **Verify the upload**:
   - Visit https://pypi.org/project/breeze-ai/
   - Check that the description renders correctly

3. **Test installation**:
   ```bash
   # Create a fresh environment
   python -m venv test_install
   source test_install/bin/activate
   
   # Install from PyPI
   pip install breeze-ai
   
   # Test
   python -c "import breeze; print(breeze.__version__)"
   breeze --version
   
   # Cleanup
   deactivate
   rm -rf test_install
   ```

## Using GitHub Actions (Automated Publishing)

Create `.github/workflows/publish-python.yml`:

```yaml
name: Publish Python Package

on:
  release:
    types: [published]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    
    - name: Build Go binary
      run: go build ./cmd/breeze
    
    - name: Set up Python
      uses: actions/setup-python@v4
      with:
        python-version: '3.x'
    
    - name: Install dependencies
      run: |
        python -m pip install --upgrade pip
        pip install build twine
    
    - name: Build package
      run: python -m build
    
    - name: Publish to PyPI
      env:
        TWINE_USERNAME: __token__
        TWINE_PASSWORD: ${{ secrets.PYPI_API_TOKEN }}
      run: twine upload dist/*
```

Then add your PyPI API token as a secret:
1. Go to repository Settings → Secrets → Actions
2. Add `PYPI_API_TOKEN` with your PyPI API token

## Version Management

To release a new version:

1. Update version in:
   - `setup.py` (line with `version=`)
   - `pyproject.toml` (line with `version =`)
   - `python/breeze/__init__.py` (line with `__version__ =`)

2. Create a git tag:
   ```bash
   git tag v2.0.1
   git push origin v2.0.1
   ```

3. Create a GitHub release (if using GitHub Actions)

## Troubleshooting

### "Package already exists"
- You can't re-upload the same version
- Increment the version number

### "Invalid distribution"
- Check that `README.md` exists and is valid markdown
- Verify `setup.py` and `pyproject.toml` syntax

### "Binary not found" after installation
- Users need to build the Go binary: `go build ./cmd/breeze`
- Consider adding pre-built binaries for common platforms
- Update `setup.py` `PostInstallCommand` to handle this better

### Package size too large
- PyPI has a 100MB limit per file
- Exclude unnecessary files in `MANIFEST.in`
- Don't include large binaries in the package

## Notes

- The package name on PyPI is `breeze-ai` (not just `breeze`, which may be taken)
- Users import it as `import breeze` (the hyphen becomes an underscore internally)
- The package includes the Go binary, so users need Go installed to build it
- Consider providing pre-compiled binaries for better user experience

## References

- [PyPI Publishing Guide](https://packaging.python.org/tutorials/packaging-projects/)
- [Twine Documentation](https://twine.readthedocs.io/)
- [Python Packaging Guide](https://packaging.python.org/)
