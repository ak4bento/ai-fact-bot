# 🤖 Golang AI Facts Bot 🧠
> Auto commit daily Golang facts using AI (OpenAI or OpenRouter).

![Auto Commit](https://github.com/ak4bento/golang-ai-facts-bot/actions/workflows/auto-commit.yml/badge.svg)
![Made with Go](https://img.shields.io/badge/Made%20with-Go-blue?logo=go)
![License](https://img.shields.io/github/license/ak4bento/golang-ai-facts-bot)
![Template](https://img.shields.io/badge/template-repo-brightgreen)

---

## 🧠 Tentang Proyek

Repo ini otomatis menambahkan 1 fakta unik tentang **bahasa pemrograman Go** ke file `thoughts.md` setiap hari menggunakan AI — tanpa harus kamu buka laptop!

Didukung dua API:
- ✅ **OpenAI** (dengan `OPENAI_API_KEY`)
- ✅ **OpenRouter** (misal: DeepSeek R1 Zero)

---

## 🚀 Cara Pakai (GitHub Template)

1. Klik tombol `Use this template` (di atas)
2. Fork repo ini ke akun kamu
3. Buka **Settings > Secrets > Actions**
4. Tambahkan salah satu secret berikut:

### ✅ Pakai OpenAI
### ✅ Pakai OpenRouter (contoh: DeepSeek R1 Zero)


---

## 🛠️ Cara Kerja

- Script Go (`scripts/generate.go`) akan memanggil ChatGPT/DeepSeek untuk membuat 1 fakta Golang
- Fakta disimpan ke `thoughts.md`
- Hasilnya di-commit dan di-push otomatis tiap hari

---

## ⏰ Schedule

- Otomatis jalan setiap hari jam **09:00 WIB**
- Bisa dijalankan manual lewat tab `Actions`

---

## 📄 Output Contoh

### 🤖 AI Fact of the Day
<!-- AI-FACT-START -->
AI can predict earthquakes by analyzing patterns in seismic data.
<!-- AI-FACT-END -->

---

## ⚙️ Teknologi

- [Go](https://golang.org/)
- [OpenAI API](https://platform.openai.com/)
- [OpenRouter](https://openrouter.ai/)
- [GitHub Actions](https://github.com/features/actions)

---

## 📜 Lisensi

[MIT](LICENSE)
