# Contributing to Safa Life API 🤝

First off, thank you for considering contributing to Safa Life API! It's people like you that make Safa Life such a great tool for the Muslim community.

Following these guidelines helps to communicate that you respect the time of the developers managing and developing this open source project. In return, they should reciprocate that respect in addressing your issue, assessing changes, and helping you finalize your pull requests.

## 🌟 What We're Looking For

Safa Life API is an open source project and we love to receive contributions from our community — you! There are many ways to contribute, from writing tutorials or blog posts, improving the documentation, submitting bug reports and feature requests or writing code which can be incorporated into Safa Life API itself.

### Types of Contributions We Welcome

- 🐛 **Bug Reports**: Help us identify and fix issues
- 💡 **Feature Requests**: Suggest new Islamic lifestyle features
- 📚 **Documentation**: Improve our guides and API documentation
- 🧪 **Testing**: Write tests to improve code coverage
- 🔧 **Code Contributions**: Implement new features or fix bugs
- 🌍 **Translations**: Help make the API accessible in different languages
- 🎨 **UI/UX Improvements**: Enhance user experience
- 📖 **Islamic Content**: Contribute authentic Islamic content and data

## 🚀 Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:
- Go 1.21 or higher
- PostgreSQL 14+
- Redis 6+
- Docker & Docker Compose
- Git

### Setting Up Your Development Environment

1. **Fork the repository**
   ```bash
   # Fork the repo on GitHub, then clone your fork
   git clone https://github.com/YOUR_USERNAME/safa-life-api.git
   cd safa-life-api
   ```

2. **Add the upstream remote**
   ```bash
   git remote add upstream https://github.com/achmdndy/safa-life-api.git
   ```

3. **Set up the development environment**
   ```bash
   # Start dependencies
   docker-compose up -d postgres redis jaeger prometheus
   
   # Install Go dependencies
   go mod download
   
   # Copy and configure environment
   cp configs/config.yaml configs/config.local.yaml
   # Edit configs/config.local.yaml with your local settings
   ```

4. **Run the application**
   ```bash
   go run src/cmd/server/main.go start
   ```

5. **Verify the setup**
   ```bash
   curl http://localhost:8080/health
   ```

## 📋 Development Guidelines

### Code Style and Standards

We follow Go best practices and maintain high code quality standards:

#### Go Code Standards
- Follow [Effective Go](https://golang.org/doc/effective_go.html) guidelines
- Use [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Run `gofmt` and `golint` before submitting
- Maintain test coverage above 80%
- Use meaningful variable and function names
- Write clear, concise comments for public APIs

#### Architecture Principles
- Follow **Clean Architecture** principles
- Maintain clear separation between layers:
  - **Domain**: Business logic and entities
  - **Application**: Use cases and interfaces
  - **Infrastructure**: External concerns (database, monitoring)
  - **Presentation**: HTTP handlers and middleware
- Use dependency injection
- Write testable code with proper interfaces

#### Islamic Content Guidelines
- Ensure all Islamic content is authentic and properly sourced
- Include references for Quranic verses and Hadiths
- Respect different schools of Islamic thought (madhabs)
- Use appropriate Arabic transliterations
- Maintain sensitivity to Islamic values and principles

### Commit Message Convention

We use [Conventional Commits](https://www.conventionalcommits.org/) for clear and consistent commit messages:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

**Types:**
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing tests or correcting existing tests
- `chore`: Changes to the build process or auxiliary tools

**Examples:**
```bash
feat(auth): add JWT authentication middleware
fix(health): resolve database connection timeout issue
docs(api): update endpoint documentation
test(handlers): add unit tests for health handler
```

### Branch Naming Convention

Use descriptive branch names that indicate the type of work:

```
<type>/<short-description>
```

**Examples:**
- `feature/prayer-times-api`
- `bugfix/database-connection-timeout`
- `docs/api-documentation-update`
- `refactor/clean-architecture-implementation`

## 🧪 Testing

We maintain high test coverage to ensure code quality and reliability.

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with detailed coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific package tests
go test ./src/presentation/handlers/health

# Run integration tests
go test -tags=integration ./...
```

### Writing Tests

- Write unit tests for all public functions
- Use table-driven tests for multiple test cases
- Mock external dependencies using interfaces
- Write integration tests for critical workflows
- Test both success and error scenarios

**Example Test Structure:**
```go
func TestHealthHandler_GetHealth(t *testing.T) {
    tests := []struct {
        name           string
        mockSetup      func(*mocks.HealthUseCase)
        expectedStatus int
        expectedBody   string
    }{
        {
            name: "successful health check",
            mockSetup: func(m *mocks.HealthUseCase) {
                m.On("CheckHealth", mock.Anything).Return(&domain.Health{
                    Status: "healthy",
                }, nil)
            },
            expectedStatus: 200,
            expectedBody:   `{"status":"healthy"}`,
        },
        // Add more test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation...
        })
    }
}
```

## 📝 Pull Request Process

### Before Submitting a Pull Request

1. **Update your fork**
   ```bash
   git fetch upstream
   git checkout main
   git merge upstream/main
   ```

2. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **Make your changes**
   - Write clean, well-documented code
   - Add tests for new functionality
   - Update documentation if needed
   - Ensure all tests pass

4. **Commit your changes**
   ```bash
   git add .
   git commit -m "feat: add your feature description"
   ```

5. **Push to your fork**
   ```bash
   git push origin feature/your-feature-name
   ```

### Pull Request Checklist

Before submitting your pull request, please ensure:

- [ ] Code follows Go best practices and project conventions
- [ ] All tests pass (`go test ./...`)
- [ ] Code coverage is maintained or improved
- [ ] Documentation is updated (if applicable)
- [ ] Commit messages follow conventional commit format
- [ ] Islamic content is authentic and properly sourced (if applicable)
- [ ] No sensitive information (passwords, keys) is committed
- [ ] The PR description clearly explains the changes

### Pull Request Template

When creating a pull request, please use this template:

```markdown
## Description
Brief description of the changes made.

## Type of Change
- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update

## Testing
- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] Manual testing completed

## Islamic Content (if applicable)
- [ ] Content is authentic and properly sourced
- [ ] References are included for Quranic verses/Hadiths
- [ ] Content respects different Islamic perspectives

## Checklist
- [ ] Code follows project style guidelines
- [ ] Self-review of code completed
- [ ] Documentation updated
- [ ] Tests added and passing
- [ ] No breaking changes (or clearly documented)
```

## 🐛 Bug Reports

When filing an issue, make sure to answer these questions:

1. **What version of Go are you using?** (`go version`)
2. **What operating system and processor architecture are you using?**
3. **What did you do?** (Provide steps to reproduce)
4. **What did you expect to see?**
5. **What did you see instead?**

### Bug Report Template

```markdown
## Bug Description
A clear and concise description of what the bug is.

## Steps to Reproduce
1. Go to '...'
2. Click on '....'
3. Scroll down to '....'
4. See error

## Expected Behavior
A clear and concise description of what you expected to happen.

## Actual Behavior
A clear and concise description of what actually happened.

## Environment
- OS: [e.g. macOS, Linux, Windows]
- Go Version: [e.g. 1.21.0]
- Application Version: [e.g. 1.0.0]

## Additional Context
Add any other context about the problem here, including logs, screenshots, etc.
```

## 💡 Feature Requests

We welcome feature requests! Please provide:

1. **Clear description** of the feature
2. **Use case** - why is this feature needed?
3. **Islamic context** - how does this align with Islamic values?
4. **Implementation ideas** (if you have any)

### Feature Request Template

```markdown
## Feature Description
A clear and concise description of the feature you'd like to see.

## Islamic Context
How does this feature serve the Muslim community or align with Islamic values?

## Use Case
Describe the problem this feature would solve or the value it would add.

## Proposed Solution
A clear and concise description of what you want to happen.

## Alternative Solutions
A clear and concise description of any alternative solutions you've considered.

## Additional Context
Add any other context, mockups, or examples about the feature request here.
```

## 🌍 Internationalization

We welcome contributions to make Safa Life API accessible in different languages:

- Arabic (العربية) - Primary Islamic language
- English - International communication
- Indonesian (Bahasa Indonesia)
- Urdu (اردو)
- Turkish (Türkçe)
- Malay (Bahasa Melayu)

### Translation Guidelines

- Maintain Islamic terminology accuracy
- Use appropriate cultural context
- Ensure proper Arabic transliteration
- Test with native speakers when possible

## 📚 Documentation

Good documentation is crucial for project success. You can contribute by:

- Improving existing documentation
- Adding code examples
- Creating tutorials
- Translating documentation
- Adding API documentation
- Writing blog posts about using the API

## 🤝 Community

### Getting Help

- 📧 Email: achmdndy@gmail.com
- 🐛 Issues: [GitHub Issues](https://github.com/achmdndy/safa-life-api/issues)
- 📖 Documentation: [Project Wiki](https://github.com/achmdndy/safa-life-api/wiki)

### Code of Conduct

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md). By participating in this project you agree to abide by its terms.

## 🎯 Recognition

Contributors will be recognized in:
- README.md contributors section
- Release notes for significant contributions
- Special recognition for Islamic content contributions
- Community highlights for exceptional contributions

## 📄 License

By contributing to Safa Life API, you agree that your contributions will be licensed under the same license as the project (MIT License).

---

**Thank you for contributing to Safa Life API! May Allah reward your efforts in serving the Muslim community. Barakallahu feek! 🤲**

## 🔗 Quick Links

- [Code of Conduct](CODE_OF_CONDUCT.md)
- [License](LICENSE)
- [README](README.md)
- [Monitoring Guide](MONITORING.md)
- [GitHub Issues](https://github.com/achmdndy/safa-life-api/issues)
- [GitHub Discussions](https://github.com/achmdndy/safa-life-api/discussions)