---
title: "How to Compile the App"
description: "Instructions for compiling Pitch-by-Pitch Pro on Windows, macOS, and Linux"
icon: "build"
toc: true
draft: false
weight: 10
---

# How to Compile the App

This guide walks you through compiling **Pitch-by-Pitch Pro** on different operating systems.

---

## Windows


1. **Install Go**

   Download and install Go for Windows from the [official website](https://go.dev/dl/).

2. **Open Command Prompt or PowerShell**

3. **Clone the repository**
```shell
git clone https://github.com/AndrewDamico/Pitch-by-Pitch-Pro.git
cd Pitch-by-Pitch-Pro
```

4. **Build the application**


```bash
go build -o generator.exe ./app/scenariogenerator
```

5. **Run the executable**

```shell
./generator.exe
```

---

## macOS

1. **Install Go**

   Use Homebrew:

```bash
brew install go
```

2. **Open Terminal**

3. **Clone the repository**

```bash
git clone https://github.com/AndrewDamico/Pitch-by-Pitch-Pro.git
cd Pitch-by-Pitch-Pro
```

4. **Build the application**

```bash
go build -o pitch-by-pitch-pro ./app/scenariogenerator
```

5. **Run the executable**

```bash
./pitch-by-pitch-pro
```

---

## Linux

1. **Install Go**

   Use your distro’s package manager, e.g., for Ubuntu:

```shell
sudo apt update
sudo apt install golang-go
```

2. **Open Terminal**

3. **Clone the repository**

```bash
git clone https://github.com/AndrewDamico/Pitch-by-Pitch-Pro.git
cd Pitch-by-Pitch-Pro
```

4. **Build the application**
```bash
go build -o pitch-by-pitch-pro ./app/scenariogenerator
```

5. **Run the executable**

```bash
./pitch-by-pitch-pro
```

---

## Notes

<aside class="success">
Make sure your `GOPATH` and `GOROOT` environment variables are configured correctly.
</aside>

- For any issues, please check the [GitHub Issues](https://github.com/AndrewDamico/Pitch-by-Pitch-Pro/issues) page.

---

If you need additional help, contact the development team or consult the [full documentation](./docs).
