package workflows

import (
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func Translate(langPair, text, keyFile, model string) []Item {
	langs := strings.Split(langPair, "-")
	if len(langs) != 2 {
		return []Item{{
			Title:    "Invalid language pair",
			Subtitle: "Use format: ru-en",
			Arg:      "",
			Icon:     Icon{"default", "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns"},
		}}
	}

	lang1, lang2 := langs[0], langs[1]

	prompt := `Translate the text below. If it's in ` + lang1 + `, translate to ` + lang2 + `. If it's in ` + lang2 + `, translate to ` + lang1 + `.

Style rules:
- Write the way a regular person texts a coworker or friend, not like an essay or article.
- Use simple, everyday words. Avoid formal, academic, or marketing-style vocabulary.
- Keep it short. Don't add words that weren't in the original. Don't "improve" or expand the message.
- Use common contractions (it's, don't, I'm, can't, we'll).
- Don't start sentences with "Certainly", "Of course", "Indeed", "Furthermore", "Moreover", "Additionally". Don't end with summaries or polite filler.
- It's fine to leave a slightly imperfect, casual phrasing if that's how a non-native speaker would actually say it. Don't over-polish.
- Match the original tone: casual stays casual, blunt stays blunt, a question stays a question.
- Keep punctuation light. No em-dashes (—). Prefer commas, periods, or just a new sentence.
- Preserve technical terms, code, names, URLs, numbers, and emojis exactly as written.
- This is a Slack message to coworkers. Lowercase first letter is fine if the original is lowercase. No period at the end of a one-sentence message.
- If the original mixes English words or product terms into Russian (e.g. "деплой", "пулл реквест", "воркспейс"), keep them as the plain English term (deploy, PR, workspace), don't rephrase them.
- Write like a software engineer in a dev team chat. Use the short forms the team actually uses: PR, QA, prod, env, config, repo, deps, msg, tx, btw, imo, afaik, wip, lgtm, pls, thx, np, asap. Don't force them in, but pick them over the long form when a dev would.
- Prefer "gonna", "wanna", "kinda", "yeah", "nope", "ok" where the original is that casual.
- Plain ASCII only for punctuation: straight apostrophes and quotes ('), never curly ones.
- Keep the original line breaks and list structure as is.

Output ONLY the translation. No quotes, no notes, no explanations.

Text:
` + text

	chatgptPath := "chatgpt"
	if usr, err := user.Current(); err == nil {
		chatgptPath = filepath.Join(usr.HomeDir, ".dotfiles", "bin", "chatgpt")
	}

	cmd := exec.Command(chatgptPath, "-q", "--omit-history", "--effort", "low", "--model", model, prompt)

	if keyFile != "" {
		keyData, err := os.ReadFile(keyFile)
		if err != nil {
			return []Item{{
				Title:    "Cannot read key file",
				Subtitle: err.Error(),
				Arg:      "",
				Icon:     Icon{"default", "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns"},
			}}
		}
		cmd.Env = append(os.Environ(), "OPENAI_API_KEY="+strings.TrimSpace(string(keyData)))
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		errMsg := strings.TrimSpace(string(output))
		if errMsg == "" {
			errMsg = err.Error()
		}
		return []Item{{
			Title:    "Translation error",
			Subtitle: errMsg,
			Arg:      "",
			Icon:     Icon{"default", "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns"},
		}}
	}

	translation := strings.TrimSpace(string(output))

	return []Item{{
		Title: translation,
		Arg:   translation,
		Icon:  Icon{"default", "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/SidebariCloud.icns"},
	}}
}
