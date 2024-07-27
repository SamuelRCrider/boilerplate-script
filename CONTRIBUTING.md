# Contributing to Chiks 🐣

First off, thank you for considering contributing to Chiks! It's people like you that make Chiks such a great tool.

## 👀 Where to start

We love your input! We want to make contributing to this project as easy and transparent as possible, whether it's:

- 🐛 Reporting a bug
- 💡 Discussing the current state of the code
- 🛠️ Submitting a fix
- 🚀 Proposing new features

## 🛠️ Our Development Environment

- VSCode
- Golang

## 🔀 Pull Requests

Pull requests are the best way to propose changes to the codebase. We actively welcome your pull requests:

1. Fork the repo and create your branch from `main`.
2. **If you're adding a new stack:**
   - Put things where they belong - full stack scripts go in the `full_stack` folder, and frontend only scripts go in the `frontend_only` folder.
   - Note that frontend only scripts must include (Frontend Only) in the name of the script.
   - If you're adding a full stack script:
     - create two folders, one for the frontend framework and one for the backend framework
     - include optionality in your scripts (e.g. database y/n, UI framework y/n, auth integration y/n)
     - make sure to link them into the full_stack.go switch case and add them to the `stacks` array in the main.go file
   - If you're adding a frontend only script:
     - create a folder for the frontend framework
     - include optionality in your script (e.g. UI framework y/n)
     - make sure to link it into the frontend_only.go switch case and add it to the `stacks` array in the main.go file
3. If you've added code that should be tested, add tests.
4. Ensure the test suite passes.
5. Issue that pull request!

## 📜 License

Any contributions you make will be under the MIT Software License. In short, when you submit code changes, your submissions are understood to be under the same [MIT License](http://choosealicense.com/licenses/mit/) that covers the project.

## 🐛 Report bugs using Github's [issues](https://github.com/samuelrcrider/chiks/issues)

We use GitHub issues to track public bugs. Report a bug by [opening a new issue](https://github.com/samuelrcrider/chiks/issues/new); it's that easy!

## 📝 Write bug reports with detail, background, and sample code

**Great Bug Reports** tend to have:

- A quick summary and/or background
- Steps to reproduce
  - Be specific!
  - Give sample code if you can.
- What you expected would happen
- What actually happens
- Notes (possibly including why you think this might be happening, or stuff you tried that didn't work)

People _love_ thorough bug reports.

## 💅 Use a Consistent Coding Style

- You can try running `go fmt` for style unification

## 🤝 Code of Conduct

Please note we have a [code of conduct](CODE_OF_CONDUCT.md), please follow it in all your interactions with the project.

## 📚 References

This document was adapted from the open-source contribution guidelines for [Facebook's Draft](https://github.com/facebook/draft-js/blob/master/CONTRIBUTING.md).

---

Happy coding! 🎉
