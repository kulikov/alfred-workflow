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

	prompt := `Translate the following text. If the text is in ` + lang1 + `, translate to ` + lang2 + `. If the text is in ` + lang2 + `, translate to ` + lang1 + `. Output ONLY the translation, nothing else:

` + text

	chatgptPath := "chatgpt"
	if usr, err := user.Current(); err == nil {
		chatgptPath = filepath.Join(usr.HomeDir, ".dotfiles", "bin", "chatgpt")
	}

	cmd := exec.Command(chatgptPath, "-n", "--model", model, "--completions-path", "/v1/chat/completions", prompt)

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
