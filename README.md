# 🤖 AI Fact Bot

![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/ak4bento/ai-fact-bot/ai-fact.yml?branch=main)
![GitHub license](https://img.shields.io/github/license/ak4bento/ai-fact-bot)
![Last Commit](https://img.shields.io/github/last-commit/ak4bento/ai-fact-bot)

**AI Fact Bot** is a GitHub Action that automatically generates daily AI-related facts using large language models (LLMs) like **OpenAI GPT** or **DeepSeek R1 Zero** via [OpenRouter](https://openrouter.ai). The fact is directly written to the `README.md` file and committed automatically to your repository.

## 🚀 Features

- Uses OpenAI or OpenRouter API
- Can be triggered on a schedule (`cron`) or manually
- Automatically commits fact updates
- Reusable in **any repo**

---

## 🧰 Requirements

1. **Go 1.22+** (automatically handled by workflow)
2. Add **2 GitHub Secrets** to your repo:
   - `AI_API_KEY` → Your API key from OpenAI or OpenRouter
   - `GH_PAT` → A Personal Access Token with `repo` permission (for pushing commits)

---

## 🔧 How to Use

Add the following to `.github/workflows/ai-fact.yml` in your repo:

```yaml
name: Daily AI Fact

on:
  workflow_dispatch:  # Trigger manually via GitHub UI
  schedule:
    - cron: '0 0 * * *'  # Every day at 00:00 UTC

jobs:
  update-fact:
    runs-on: ubuntu-latest
    steps:
      - uses: ak4bento/ai-fact-bot@v1
        with:
          api_key: ${{ secrets.AI_API_KEY }}
          api_base_url: "https://openrouter.ai/api/v1"

      - name: Commit & Push updated README
        env:
          GH_PAT: ${{ secrets.GH_PAT }}
        run: |
          git config --global user.name "${{ github.actor }}"
          git config --global user.email "${{ github.actor }}@users.noreply.github.com"
          git add README.md || true
          git diff --cached --quiet || (
            git commit -m "🤖 update AI fact" &&
            git push https://${{ github.actor }}:${GH_PAT}@github.com/${{ github.repository }}.git HEAD:main
          )
```

---

### 🧠 Custom Prompt (optional)

You can optionally set a secret named `FACT_PROMPT` to customize the topic or style of the generated content.

#### Example:

```env
FACT_PROMPT="Tell me a mind-blowing AI fact in a single sentence."
```

If `FACT_PROMPT` is not set, the default prompt will be used:

> "Give me one interesting fact about AI in 1 sentence."

This allows you to repurpose the bot for other use cases, such as:

- 🧠 Tech facts  
- 🤖 Random developer tips  
- 💡 Programming quotes  
- 🔮 Fun facts or trivia  

Customize it to fit your project's theme!

---

## 📝 Output

The AI-generated fact will be automatically written to the `README.md`. Example:

```md
## 🔍 AI Fact of the Day

AI can improve software development efficiency by up to 40% through automated testing and code generation.
```

---

## 📄 Example Output 

### 🤖 AI Fact of the Day
<!-- AI-FACT-START -->
AI can analyze vast amounts of data to detect patterns and make predictions faster and more accurately than humans.
<!-- AI-FACT-END -->

---

## 🔒 Security

- Your API key is never printed or logged.
- Always store `AI_API_KEY` and `GH_PAT` in GitHub Secrets.

---

## 💡 Tips

- Fork this repo and customize the fact content or output.
- You can also use it to generate quotes, fun facts, or any kind of daily snippet!

---

## 👨‍💻 Contributing

Pull requests and feedback are welcome!  
Feel free to open an issue or submit a PR if you want to improve this action.

---

## 📄 License

MIT License © [ak4bento](https://github.com/ak4bento)

