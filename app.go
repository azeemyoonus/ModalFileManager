package main

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	goruntime "runtime"
	"strconv"
	"strings"
	"time"

	clip "github.com/atotto/clipboard"
	"github.com/go-git/go-git/v5"
	github "github.com/google/go-github/v49/github"
	cp "github.com/otiai10/copy"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App application struct and other structs
type App struct {
	ctx           context.Context
	err           string
	lastRightDir  string
	lastLeftDir   string
	Commands      []string
	Timer         *time.Timer
	Stopped       bool
	LenLeftFiles  int
	LenRightFiles int
	LeftHash      string
	RightHash     string
}

type FileParts struct {
	Dir       string
	Name      string
	Extension string
}

type FileInfo struct {
	Dir       string
	Name      string
	Extension string
	IsDir     bool
	Size      int64
	Modtime   string
	Index     int
	Mode      fs.FileMode
	Link      bool
}

type GitHubRepos struct {
	Name        string           `json:"name"`
	URL         string           `json:"url"`
	Stars       int              `json:"stars"`
	Owner       string           `json:"owner"`
	ID          int64            `json:"id"`
	Description string           `json:"description"`
	UpdatedAt   github.Timestamp `json:"updatedat"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (b *App) startup(ctx context.Context) {
	b.ctx = ctx

	//
	// We need to start the backend and setup the signaling.
	//
	go backend(b, ctx)
}

// domReady is called after the front-end dom has been loaded
func (b *App) domReady(ctx context.Context) {
	//
	// Start watching the directories.
	//
	b.lastLeftDir = ""
	b.lastRightDir = ""
	b.Stopped = false
	go b.StartWatcher()
}

// shutdown is called at application termination
func (b *App) shutdown(ctx context.Context) {
	//
	// Stop watching directories.
	//
	b.lastLeftDir = ""
	b.lastRightDir = ""
	b.Stopped = true
	b.StopWatcher()
}

// These functions are for watching the current directories in the file manager.
func (b *App) SetRightDirWatch(path string) {
	b.lastRightDir = path
}

func (b *App) SetLeftDirWatch(path string) {
	b.lastLeftDir = path
}

func (b *App) CloseRightWatch() {
	b.lastRightDir = ""
}

func (b *App) CloseLeftWatch() {
	b.lastLeftDir = ""
}

func (b *App) StartWatcher() {
	for !b.Stopped {
		//
		// Create the timer.
		//
		b.Timer = time.NewTimer(time.Millisecond * 30000)

		//
		// Do the Job. check for changes in the current directories.
		//
		if b.lastLeftDir != "" {
			leftFiles := b.ReadDir(b.lastLeftDir)
			LenLeftFiles := len(leftFiles)
			if LenLeftFiles != b.LenLeftFiles {
				//
				// The number of files have changed. Reload the directory.
				//
				b.LenLeftFiles = LenLeftFiles
				rt.EventsEmit(b.ctx, "leftSideChange", "")
			} else {
				//
				// See if a file name changed. This takes longer.
				//
				fileNames := ""
				for i := 0; i < LenLeftFiles; i++ {
					fileNames += leftFiles[i].Name + strconv.FormatInt(leftFiles[i].Size, 10)
				}
				if b.LeftHash != fileNames {
					b.LeftHash = fileNames
					rt.EventsEmit(b.ctx, "leftSideChange", "")
				}
			}
		}
		if b.lastRightDir != "" {
			rightFiles := b.ReadDir(b.lastRightDir)
			LenRightFiles := len(rightFiles)
			if LenRightFiles != b.LenRightFiles {
				//
				// The number of files have changed. Reload the directory.
				//
				b.LenRightFiles = LenRightFiles
				rt.EventsEmit(b.ctx, "rightSideChange", "")
			} else {
				//
				// See if a file name changed. This takes longer.
				//
				fileNames := ""
				for i := 0; i < LenRightFiles; i++ {
					fileNames += rightFiles[i].Name + strconv.FormatInt(rightFiles[i].Size, 10)
				}
				if b.RightHash != fileNames {
					b.RightHash = fileNames
					rt.EventsEmit(b.ctx, "rightSideChange", "")
				}
			}
		}

		//
		// Wait for the timer to finish.
		//
		<-b.Timer.C
	}
}

func (b *App) StopWatcher() {
	b.lastLeftDir = ""
	b.lastRightDir = ""
	b.Timer.Stop()
}

func (b *App) GetCommandLineCommands() []string {
	return b.Commands
}

func (b *App) ReadFile(path string) string {
	b.err = ""
	contents, err := os.ReadFile(path)
	if err != nil {
		b.err = err.Error()
	}
	return string(contents[:])
}

func (b *App) GetHomeDir() string {
	b.err = ""
	hdir, err := os.UserHomeDir()
	if err != nil {
		b.err = err.Error()
	}
	return hdir
}

func (b *App) WriteFile(path string, data string) {
	err := os.WriteFile(path, []byte(data), 0666)
	if err != nil {
		b.err = err.Error()
	}
}

func (b *App) FileExists(path string) bool {
	b.err = ""
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func (b *App) DirExists(path string) bool {
	b.err = ""
	dstat, err := os.Stat(path)
	if err != nil {
		b.err = err.Error()
		return false
	}
	return dstat.IsDir()
}

func (b *App) SplitFile(path string) FileParts {
	b.err = ""
	var parts FileParts
	parts.Dir, parts.Name = filepath.Split(path)
	parts.Extension = filepath.Ext(path)
	return parts
}

func (b *App) ReadDir(path string) []FileInfo {
	b.err = ""
	var result []FileInfo
	result = make([]FileInfo, 0)
	files, err := os.ReadDir(path)
	if err != nil {
		b.err = err.Error()
	} else {
		for index, file := range files {
			var fileInfo FileInfo
			info, err := file.Info()
			if err != nil {
				b.err = err.Error()
			}
			fileInfo.Name = file.Name()
			fileInfo.Size = info.Size()
			fileInfo.IsDir = file.IsDir()
			fileInfo.Modtime = info.ModTime().Format(time.ANSIC)
			fileInfo.Dir = path
			fileInfo.Extension = filepath.Ext(file.Name())
			fileInfo.Index = index
			fileInfo.Mode = info.Mode().Perm()
			fileInfo.Link = false

			//
			// Determine if it is a symlink and if so if it's a directory. Currently,
			// if before it wasn't a directory but is now, then it is treated as a
			// symlink directory. Nothing is in place to detect a file symlink.
			//
			ninfo, err := os.Stat(filepath.Join(path, fileInfo.Name))
			if err != nil {
				b.err = err.Error()
			} else {
				dir := ninfo.Mode().IsDir()
				if dir && !fileInfo.IsDir {
					fileInfo.IsDir = true
					fileInfo.Link = true
				}
				if ninfo.Mode()&os.ModeSymlink == 'l' {
					fileInfo.Link = true
				}
			}
			//
			// Add it to the rest.
			//
			result = append(result, fileInfo)
		}
	}
	return result
}

func (b *App) MakeDir(path string) {
	b.err = ""
	err := os.MkdirAll(path, 0755)
	if err != nil {
		b.err = err.Error()
	}
}

func (b *App) MakeFile(path string) {
	b.err = ""
	b.WriteFile(path, "")
}

func (b *App) MoveEntries(from string, to string) {
	b.err = ""
	err := os.Rename(from, to)
	if err != nil {
		b.err = err.Error()
	}
}

func (b *App) RenameEntry(from string, to string) {
	b.err = ""
	err := os.Rename(from, to)
	if err != nil {
		b.err = err.Error()
	}
}

func (b *App) GetError() string {
	return b.err
}

func (b *App) CopyEntries(from string, to string) {
	b.err = ""
	info, err := os.Stat(from)
	if os.IsNotExist(err) {
		b.err = err.Error()
		return
	}
	if info.IsDir() {
		//
		// It's a directory! Do a deap copy.
		//
		err := cp.Copy(from, to)
		if err != nil {
			b.err = err.Error()
			return
		}
	} else {
		//
		// It's a file. Just copy it.
		//
		source, err := os.Open(from)
		if err != nil {
			b.err = err.Error()
			return
		}
		defer source.Close()

		destination, err := os.Create(to)
		if err != nil {
			b.err = err.Error()
			return
		}
		defer destination.Close()
		_, err = io.Copy(destination, source)
		if err != nil {
			b.err = err.Error()
		}
	}
}

func (b *App) DeleteEntries(path string) {
	b.err = ""
	err := os.RemoveAll(path)
	if err != nil {
		b.err = err.Error()
	}
}

func (b *App) RunCommandLine(cmd string, args []string, env []string, dir string) string {
	b.err = ""
	cmdline := exec.Command(cmd)
	cmdline.Args = args
	cmdline.Env = env
	cmdline.Dir = dir
	result, err := cmdline.CombinedOutput()
	if err != nil {
		b.err = err.Error()
	}

	return string(result[:])
}

func (b *App) GetClip() string {
	result, err := clip.ReadAll()
	if err != nil {
		b.err = err.Error()
	}
	return result
}

func (b *App) SetClip(msg string) {
	err := clip.WriteAll(msg)
	if err != nil {
		b.err = err.Error()
	}
}

func (b *App) GetEnvironment() []string {
	return os.Environ()
}

func (b *App) AppendPath(dir string, name string) string {
	return filepath.Join(dir, name)
}

func (b *App) Quit() {
	rt.Quit(b.ctx)
}

func (b *App) GetOSName() string {
	os := goruntime.GOOS
	result := ""
	switch os {
	case "windows":
		result = "windows"
	case "darwin":
		result = "macos"
	case "linux":
		result = "linux"
	default:
		result = os
	}
	return result
}

func (b *App) GetGitHubThemes() []GitHubRepos {
	var result []GitHubRepos
	client := github.NewClient(nil)
	topics, _, err := client.Search.Repositories(context.Background(), "in:topic modalfilemanager in:topic theme", nil)
	if err == nil {
		total := *topics.Total
		result = make([]GitHubRepos, total)
		for i := 0; i < total; i++ {
			result[i].ID = *topics.Repositories[i].ID
			result[i].Name = *topics.Repositories[i].Name
			result[i].Owner = *topics.Repositories[i].Owner.Login
			result[i].URL = *topics.Repositories[i].CloneURL
			result[i].Stars = *topics.Repositories[i].StargazersCount
			result[i].Description = *topics.Repositories[i].Description
			result[i].UpdatedAt = *topics.Repositories[i].UpdatedAt
		}
	}
	return result
}

func (b *App) GetGitHubScripts() []GitHubRepos {
	var result []GitHubRepos
	client := github.NewClient(nil)
	topics, _, err := client.Search.Repositories(context.Background(), "in:topic modalfilemanager in:topic V2 in:topic extension", nil)
	if err == nil {
		total := *topics.Total
		result = make([]GitHubRepos, total)
		for i := 0; i < total; i++ {
			result[i].ID = *topics.Repositories[i].ID
			result[i].Name = *topics.Repositories[i].Name
			result[i].Owner = *topics.Repositories[i].Owner.Login
			result[i].URL = *topics.Repositories[i].CloneURL
			result[i].Stars = *topics.Repositories[i].StargazersCount
			result[i].Description = *topics.Repositories[i].Description
			result[i].UpdatedAt = *topics.Repositories[i].UpdatedAt
		}
	}
	return result
}

func (b *App) SearchMatchingDirectories(path string, pat string, max int) []string {
	result := make([]string, max)
	count := 0
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err == nil && info.IsDir() {
			//
			// check for a directory that matches the pattern.
			//
			if strings.Contains(path, pat) {
				result[count] = path
				count++
				if count >= max {
					return filepath.SkipAll
				}
			}
		}
		return nil
	})
	if err != nil {
		b.err = err.Error()
	}

	//
	// Return the results.
	//
	return result
}

func (b *App) CloneGitHub(url string, dir string) {
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		b.err = err.Error()
	}
}
