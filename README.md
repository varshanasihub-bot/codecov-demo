# Codecov Demo

A complete, ready-to-run demo project configured for **[Codecov](https://about.codecov.io/)** using Go and GitHub Actions.

---

## 📁 Repository Structure

```
├── .github/
│   └── workflows/
│       └── codecov.yml      # GitHub Actions workflow for tests & Codecov upload
├── cal.go                   # Calculator functions (Add, Sub, Mul, Div)
├── cal_test.go              # Unit tests with ~83.3% initial coverage
├── codecov.yml              # Codecov configuration (thresholds, PR comments)
├── Makefile                 # Shortcuts for running tests and coverage
├── go.mod                   # Go module definition
└── .gitignore               # Ignores coverage reports and binaries
```

---

## 🚀 Local Testing & Coverage

Run tests and generate coverage reports locally:

```bash
# Run unit tests
make test

# Run tests and output statement coverage report
make coverage

# Generate and view interactive HTML coverage in your browser
make html
open coverage.html
```

> **Note:** The current tests deliberately leave the division-by-zero branch in `Div()` uncovered (~83.3% coverage). This allows you to see uncovered lines highlighted in Codecov!

---

## ⚙️ How to Connect with Codecov

### 1. Initialize Git and Push to GitHub

```bash
git init
git add .
git commit -m "Initial commit for Codecov demo"
git branch -M main
git remote add origin https://github.com/<your-username>/<your-repo-name>.git
git push -u origin main
```

### 2. Set Up Codecov

1. Log into **[Codecov.io](https://about.codecov.io/)** with your GitHub account.
2. Under **Repositories**, click **Configure** or add your repository.
3. Copy your repository's **Upload Token** (`CODECOV_TOKEN`).

### 3. Add Token to GitHub Secrets

1. In your GitHub repository, navigate to **Settings** > **Secrets and variables** > **Actions**.
2. Click **New repository secret**.
3. Name: `CODECOV_TOKEN`
4. Value: Paste the token from Codecov.
5. Click **Add secret**.

---

## 📊 Viewing Coverage Results

- Every time you push to `main` or create a pull request, the GitHub Action automatically runs tests, generates `coverage.txt`, and uploads it to Codecov.
- Codecov comments on your pull requests with coverage diffs and adds status checks!
