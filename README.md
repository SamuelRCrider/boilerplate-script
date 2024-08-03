<div align="center">

# 🐣 Hatch Your Next Project With Chiks. 🐣

[![npm version](https://img.shields.io/npm/v/chiks.svg?style=flat-square)](https://www.npmjs.com/package/chiks)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Downloads](https://img.shields.io/npm/dt/chiks.svg?style=flat-square)](https://www.npmjs.com/package/chiks)

#### Chiks is your one-stop solution for hatching modern web projects with ease!

[Installation](#-installation) •
[Features](#-features) •
[Usage](#-usage) •
[Updating](#-updating) •
[Uninstalling](#-uninstalling) •
[Contributing](#-contributing)

</div>

## 🚀 Installation

Make sure you have the following prerequisites installed:

- Docker
- PostgreSQL

Then, install Chiks globally:

```bash
npm install -g chiks
```

## ✨ Features

### 🛠 Project Creation

- Full Stack
  - Next.js
  - Vite + Express
- Frontend Only
  - Next.js
  - Vite

### 🐳 Docker Integration

Automatically spins up a Docker container

### 🗄 Database Setup

Sets up a PostgreSQL database (Prisma ORM) and runs your first migration

### 🎨 Customization Options

- Choose your preferred auth integration
- Opt for a database, or not
- Select your desired Docker port
- Opt for Tailwind CSS integration
  - Additional UI framework options:
    - Shadcn UI
    - DaisyUI

## 📝 Usage

1. Open your terminal
2. Navigate to your projects directory
3. Run the following command:

```bash
chiks
```

4. Follow the prompts and watch your project hatch!

## 🔄 Updating

Keep Chiks up-to-date with a simple command:

```bash
npm update -g chiks
```

## 🗑 Uninstalling

If you need to uninstall Chiks:

1. Uninstall the package:

```bash
npm uninstall -g chiks
```

2. Remove the binary:

```bash
rm $(which chiks)
```

NOTE: Due to npm v7+, uninstall scripts won't work, so the package's preuninstall script cannot remove the binary from your system. That's why you need to manually remove the binary as shown above.

## 🤝 Contributing

#### Please contribute!

We have a contributing guide! Check it out [here](https://github.com/SamuelRCrider/chiks/blob/main/CONTRIBUTING.md). This is especially important if you want to add a new stack to Chiks.

## 📄 License

This project is licensed under the MIT License - basically, do whatever you want as long as you give us credit.

<h3 align="center">
  Made with ❤️ by the Chiks team
</h3>

<div align="center">

[Report a bug](https://github.com/SamuelRCrider/chiks/issues/new?assignees=&labels=bug&template=bug_report.md&title=) •
[Request a feature](https://github.com/SamuelRCrider/chiks/issues/new?assignees=&labels=enhancement&template=feature_request.md&title=) •
[Contact us](https://github.com/SamuelRCrider/chiks/issues/new?assignees=&labels=question&template=question.md&title=)

</div>
