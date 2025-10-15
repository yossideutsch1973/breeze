# Security Policy

## Supported Versions

We actively support the following versions of Breeze with security updates:

| Version | Supported          |
| ------- | ------------------ |
| 2.0.x   | :white_check_mark: |
| < 2.0   | :x:                |

## Reporting a Vulnerability

We take the security of Breeze seriously. If you discover a security vulnerability, please follow these steps:

### How to Report

1. **Do NOT open a public GitHub issue** for security vulnerabilities
2. Email security concerns to: [Create an email or use GitHub Security Advisories]
3. Alternatively, use [GitHub's private vulnerability reporting](https://github.com/yossideutsch1973/breeze/security/advisories/new)

### What to Include

Please include the following information in your report:

- **Description** of the vulnerability
- **Steps to reproduce** the issue
- **Potential impact** of the vulnerability
- **Suggested fix** (if you have one)
- **Affected versions**
- Your **contact information** for follow-up

### What to Expect

- **Acknowledgment**: We will acknowledge receipt of your report within 48 hours
- **Updates**: We will provide regular updates on our progress
- **Timeline**: We aim to release a fix within 7-14 days for critical issues
- **Credit**: We will credit you in the security advisory (unless you prefer to remain anonymous)

## Security Best Practices

When using Breeze, please follow these security best practices:

### 1. Protect Sensitive Data
- Never pass sensitive information (passwords, API keys, personal data) in prompts
- Be cautious when processing documents that may contain sensitive information
- Review AI-generated code for security vulnerabilities before deployment

### 2. Keep Dependencies Updated
- Regularly update Go and Ollama to the latest versions
- Monitor security advisories for dependencies
- Run `go mod tidy` and `go get -u` periodically

### 3. Secure Local Development
- Run Ollama in a secure, isolated environment
- Use firewall rules to restrict Ollama network access
- Be aware that LLM responses may include unvetted code or commands

### 4. Validate AI Outputs
- Always review and validate AI-generated code before execution
- Do not blindly execute commands suggested by the AI
- Implement proper input validation in production applications

### 5. Model Security
- Only use trusted models from the official Ollama registry
- Be aware that models may have biases or produce unexpected outputs
- Review model documentation and licenses before use

## Known Security Considerations

### Local LLM Risks
- **Data Privacy**: All data stays local, but be mindful of what you input
- **Code Execution**: The library does not execute AI-generated code automatically
- **Network Access**: Ollama runs locally but requires internet for initial model downloads

### Dependencies
- **Ollama**: Ensure you download from the official source (https://ollama.ai)
- **Go Libraries**: We use minimal dependencies, primarily Go stdlib

## Security Updates

Security updates will be announced via:
- GitHub Security Advisories
- Repository releases with `[SECURITY]` tag
- README.md updates for critical issues

## Compliance

Breeze is designed for:
- **Local-first**: All processing happens on your machine
- **Privacy**: No data is sent to external servers (except for model downloads)
- **Open Source**: MIT License, fully auditable code

## Questions?

If you have questions about security that don't constitute a vulnerability report, please:
- Open a GitHub Discussion
- Check existing security-related issues
- Contact the maintainers through GitHub

Thank you for helping keep Breeze and its users secure!
