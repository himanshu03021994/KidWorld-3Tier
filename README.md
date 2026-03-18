# 🌈 Kid World - Interactive 3-Tier Educational App

![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white)
![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Render](https://img.shields.io/badge/Render-%46E3B7.svg?style=for-the-badge&logo=render&logoColor=white)
![GitHub Pages](https://img.shields.io/badge/github%20pages-121013?style=for-the-badge&logo=github&logoColor=white)

**Kid World** is a zero-latency, gamified educational web application designed for early childhood learning. Built with a modern, interactive UI, it features regional language support, an AI voice engine, and a cloud-powered backend for seamless content delivery.

🚀 **Play the Game Live:** [Kid World - Live Demo](https://himanshu03021994.github.io/KidWorld-3Tier/)

---

## ✨ Key Features

* **🎮 Advanced Gamified Learning:** Features a dynamic 6-button multiple-choice layout with sequential learning (A to Z) to keep kids engaged.
* **🗣️ AI Voice Engine:** Native Web Speech API integration talks to kids ("Find A", "अ कहाँ है?") and supports regional accents (English, Hindi, Punjabi, Gujarati, Marathi) without external audio files.
* **⚡ Zero-Latency Gameplay:** Implements background "Pre-fetching" of cloud data directly on the dashboard, ensuring the actual game loads instantly with zero lag.
* **🌍 Regional Language Support:** Dynamically serves distinct datasets (Alphabets, Emojis, Words) based on the user's selected State.
* **❤️ Health & Reward System:** Includes a 3-lives (hearts) system to encourage accuracy, plus a real-time Coin counter (🪙) for every correct answer.
* **🪄 Magic Reveal UI:** Smooth CSS transitions and visual feedback. Correct answers trigger a custom pop-up card that transforms the target letter into a visual object (e.g., `A` ➡️ `🍎 A for Apple`).

---

## 🏗️ 3-Tier Architecture

This application strictly follows a modern, scalable 3-Tier Cloud Architecture:

1. **Presentation Tier (Frontend):** * Built with pure HTML, CSS, and Vanilla JavaScript.
   * State management is handled completely client-side via browser `localStorage`.
   * Hosted globally and continuously deployed via **GitHub Pages**.
2. **Logic Tier (Backend Server):** * RESTful API built in **Go (Golang)**.
   * Handles dynamic routing (`/setup` and `/get-alphabet`) and CORS policies.
   * Hosted and continuously deployed on **Render Cloud**.
3. **Data Tier (Data Structures):** * Structured JSON maps inside the Go environment serving rich datasets (Letters, Words, Emojis) mapped to specific geographic regions.

---

## 🛠️ Local Development Setup

To run this project on your local machine, follow these steps:

### 1. Run the Backend (Go)
```bash
# Navigate to the backend directory
cd backend

# Run the Go server locally
go run main.go
# Server will start on http://localhost:8080
