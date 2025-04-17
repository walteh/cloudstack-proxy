package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/rs/zerolog"
	"gitlab.com/tozd/go/errors"
	"gopkg.in/yaml.v3"
)

// Version will be set during build
var Version = "dev"

type Buf3pdDep struct {
	Type   string   `yaml:"type"`
	Repo   string   `yaml:"repo"`
	Path   string   `yaml:"path"`
	Ref    string   `yaml:"ref"`
	Filter []string `yaml:"filter"`
}

// Buf3pdConfig represents the configuration structure in buf.yaml
type Buf3pdConfig struct {
	Path string      `yaml:"path"`
	Deps []Buf3pdDep `yaml:"deps"`
}

// BufYaml represents the buf.yaml file structure
type BufYaml struct {
	Version  string                 `yaml:"version"`
	Modules  []map[string]string    `yaml:"modules"`
	Breaking map[string]interface{} `yaml:"breaking"`
	Deps     []string               `yaml:"deps"`
	Lint     map[string]interface{} `yaml:"lint"`
	Buf3pd   Buf3pdConfig           `yaml:"buf3pd"`
}

// LockDep represents a dependency entry in the lock file
type LockDep struct {
	Repo     string          `yaml:"repo"`
	Path     string          `yaml:"path"`
	Ref      string          `yaml:"ref"`
	Digest   string          `yaml:"digest"`
	Prefix   string          `yaml:"prefix"`
	Metadata LockDepMetadata `yaml:"metadata"`
}

type LockDepMetadata struct {
	Commit string `yaml:"commit"`
	Type   string `yaml:"type"`
}

// LockFile represents the structure of the buf3pd.lock file
type LockFile struct {
	Version        string     `yaml:"version"`
	Deps           []*LockDep `yaml:"deps"`
	commitMetadata string
}

func main() {
	ctx := context.Background()
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	ctx = logger.WithContext(ctx)

	log := zerolog.Ctx(ctx)
	log.Info().Str("version", Version).Msg("starting buf3pd")

	// Parse command line flags
	var (
		bufYamlPath = flag.String("config", "buf.yaml", "Path to buf.yaml file")
		workDir     = flag.String("workdir", ".", "Working directory")
	)
	flag.Parse()

	// Ensure workDir is absolute
	absWorkDir, err := filepath.Abs(*workDir)
	if err != nil {
		log.Fatal().Err(errors.Errorf("resolving absolute path for workdir: %w", err)).Msg("failed to start")
	}

	cfg, err := readBuf3pdConfig(ctx, filepath.Join(absWorkDir, *bufYamlPath))
	if err != nil {
		log.Fatal().Err(errors.Errorf("reading buf3pd config: %w", err)).Msg("failed to read buf3pd config")
	}

	// Process dependencies
	lockFile, err := readLockFile(filepath.Join(absWorkDir, "buf3pd.lock"))
	if err != nil {
		log.Fatal().Err(errors.Errorf("reading lock file: %w", err)).Msg("failed to read lock file")
	}

	// Create the output directory if it doesn't exist
	outputPath := filepath.Join(absWorkDir, cfg.Path)
	if err := os.MkdirAll(outputPath, 0755); err != nil {
		log.Fatal().Err(errors.Errorf("creating output directory: %w", err)).Msg("failed to create output directory")
	}

	depFilesToUpdate := []*DepFiles{}

	for _, dep := range cfg.Deps {
		if dep.Type != "git" {
			log.Warn().Str("type", dep.Type).Msg("unsupported dependency type, skipping")
			continue
		}

		storedLockDep := lockFile.EntryFor(dep)

		var ok bool
		var tryLoc *DepFiles

		// if storedLockDep != nil {
		tryLoc, ok, err = NewDepFilesFromLocal(ctx, cfg, dep)
		if err != nil {
			log.Fatal().Err(errors.Errorf("processing git dependency: %w", err)).Msg("failed to process git dependency")
		}
		// }

		var depFiles *DepFiles
		var lockDep *LockDep

		var skipRemote = false

		if ok {
			// Local dependency found
			realLockDep, err := tryLoc.LockEntry()
			if err != nil {
				log.Fatal().Err(errors.Errorf("creating lock entry: %w", err)).Msg("failed to create lock entry")
			}

			if storedLockDep != nil && storedLockDep.Compare(realLockDep) {
				log.Info().Str("repo", dep.Repo).Str("path", dep.Path).Str("ref", dep.Ref).Msg("dependency already processed")
				skipRemote = true
				realLockDep.Metadata = storedLockDep.Metadata
			} else {
				log.Warn().Any("storedLockDep", storedLockDep).Any("realLockDep", realLockDep).Msg("dependency already processed, but with different commit")
			}

			log.Info().Str("repo", dep.Repo).Str("path", dep.Path).Str("ref", dep.Ref).Msg("using local dependency")
			depFiles = tryLoc
			lockDep = realLockDep
		}

		if !skipRemote {
			// No local dependency, fetch from remote
			log.Info().Str("repo", dep.Repo).Str("path", dep.Path).Str("ref", dep.Ref).Msg("processing git dependency from remote")

			remoteDepFiles, err := NewDepFilesFromRemote(ctx, dep)
			if err != nil {
				log.Fatal().Err(errors.Errorf("processing git dependency: %w", err)).Msg("failed to process git dependency")
			}

			remoteLockDep, err := remoteDepFiles.LockEntry()
			if err != nil {
				log.Fatal().Err(errors.Errorf("creating lock entry: %w", err)).Msg("failed to create lock entry")
			}

			depFiles = remoteDepFiles
			lockDep = remoteLockDep
		}

		if storedLockDep != nil {
			*storedLockDep = *lockDep
		} else {
			lockFile.Deps = append(lockFile.Deps, lockDep)
		}

		depFilesToUpdate = append(depFilesToUpdate, depFiles)

		log.Info().Str("repo", dep.Repo).Str("prefix", lockDep.Prefix).Msg("successfully processed dependency")
	}

	for _, depFiles := range depFilesToUpdate {
		err := depFiles.WriteToDir(filepath.Join(outputPath, filepath.Base(depFiles.DepInfo.Repo)))
		if err != nil {
			log.Fatal().Err(errors.Errorf("writing dependency files: %w", err)).Msg("failed to write dependency files")
		}
	}

	// Write the lock file
	lockFileContent, err := yaml.Marshal(lockFile)
	if err != nil {
		log.Fatal().Err(errors.Errorf("marshaling lock file: %w", err)).Msg("failed to create lock file")
	}

	lockFilePath := filepath.Join(absWorkDir, "buf3pd.lock")
	if err := os.WriteFile(lockFilePath, []byte("# Generated by buf3pd. DO NOT EDIT.\n"+string(lockFileContent)), 0644); err != nil {
		log.Fatal().Err(errors.Errorf("writing lock file: %w", err)).Msg("failed to write lock file")
	}

	log.Info().Str("path", lockFilePath).Msg("created lock file")
	log.Info().Msg("buf3pd completed successfully")
}

func readLockFile(path string) (*LockFile, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &LockFile{
				Version: "v2",
				Deps:    []*LockDep{},
			}, nil
		}
		return nil, errors.Errorf("reading lock file: %w", err)
	}

	var lockFile LockFile
	if err := yaml.Unmarshal(content, &lockFile); err != nil {
		return nil, errors.Errorf("unmarshalling lock file: %w", err)
	}

	return &lockFile, nil
}

func readBuf3pdConfig(ctx context.Context, path string) (*Buf3pdConfig, error) {

	log := zerolog.Ctx(ctx)

	log.Info().Str("path", path).Msg("reading buf3pd config")
	// Parse buf.yaml
	bufYamlContent, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Errorf("reading buf.yaml: %w", err)
	}

	// Split by "---" to find the sections
	sections := strings.Split(string(bufYamlContent), "---")
	if len(sections) < 2 {
		return nil, errors.New("buf.yaml does not contain the expected sections separated by '---'")
	}

	if len(sections) < 2 {
		return nil, errors.New("buf.yaml does not contain the expected sections separated by '---'")
	}

	content := sections[1]

	var buf3pdConfig struct {
		Buf3pd *Buf3pdConfig `yaml:"buf3pd"`
	}
	if err := yaml.Unmarshal([]byte(content), &buf3pdConfig); err != nil {
		return nil, errors.Errorf("unmarshalling buf3pd config: %w", err)
	}

	if buf3pdConfig.Buf3pd == nil || len(buf3pdConfig.Buf3pd.Deps) == 0 {
		return nil, errors.New("buf3pd config does not contain any dependencies")
	}

	return buf3pdConfig.Buf3pd, nil
}

type DepFiles struct {
	DepInfo        Buf3pdDep `yaml:"dep"`
	Files          []*File   `yaml:"files"`
	commitMetadata string
}

func NewDepFilesFromLocal(ctx context.Context, cfg *Buf3pdConfig, dep Buf3pdDep) (*DepFiles, bool, error) {

	pth := filepath.Join(cfg.Path, filepath.Base(dep.Repo))

	zerolog.Ctx(ctx).Info().Str("path", pth).Msg("processing local dependency")

	// Check if directory exists
	if _, err := os.Stat(pth); os.IsNotExist(err) {
		zerolog.Ctx(ctx).Warn().Str("path", pth).Msg("directory does not exist, skipping local dependency")
		return nil, false, nil
	}

	// Find all proto files in the directory
	protoFiles, err := doublestar.Glob(os.DirFS(pth), "**/*.proto")
	if err != nil {
		return nil, false, errors.Errorf("walking directory: %w", err)
	}

	if len(protoFiles) == 0 {
		zerolog.Ctx(ctx).Warn().Str("path", pth).Msg("no proto files found, skipping local dependency")
		return nil, false, nil
	}

	depFiles := &DepFiles{
		DepInfo: dep,
		Files:   []*File{},
	}

	for _, file := range protoFiles {
		fullPath := filepath.Join(pth, file)
		content, err := os.ReadFile(fullPath)
		if err != nil {
			return nil, false, errors.Errorf("reading file: %w", err)
		}
		depFiles.Files = append(depFiles.Files, &File{
			Path:    file,
			Content: content,
		})
	}

	return depFiles, true, nil
}

func NewDepFilesFromRemote(ctx context.Context, dep Buf3pdDep) (*DepFiles, error) {
	tempDir, err := os.MkdirTemp("", "buf3pd-git-")
	if err != nil {
		return nil, errors.Errorf("creating temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Clone the repository
	cmd := exec.Command("git", "clone", "--depth", "1", "https://"+dep.Repo, tempDir)
	if output, err := cmd.CombinedOutput(); err != nil {
		return nil, errors.Errorf("git clone: %w: %s", err, string(output))
	}

	// fetch the tags
	cmd = exec.Command("git", "fetch", "origin", "--tags")
	cmd.Dir = tempDir
	if output, err := cmd.CombinedOutput(); err != nil {
		return nil, errors.Errorf("git fetch: %w: %s", err, string(output))
	}

	// Checkout the specified reference
	cmd = exec.Command("git", "checkout", dep.Ref)
	cmd.Dir = tempDir
	if output, err := cmd.CombinedOutput(); err != nil {
		return nil, errors.Errorf("git checkout: %w: %s", err, string(output))
	}

	// get the commit hash
	cmd = exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = tempDir
	commitHash, err := cmd.Output()
	if err != nil {
		return nil, errors.Errorf("git rev-parse: %w", err)
	}
	commit := strings.TrimSpace(string(commitHash))

	depFiles := &DepFiles{
		commitMetadata: commit,
		DepInfo:        dep,
		Files:          []*File{},
	}

	if err := depFiles.AddAllNestedProtoFiles(ctx, filepath.Join(tempDir, dep.Path), dep.Filter...); err != nil {
		return nil, errors.Errorf("adding all nested proto files: %w", err)
	}

	if len(depFiles.Files) == 0 {
		return nil, errors.New("no proto files found")
	}

	return depFiles, nil
}

func (d *DepFiles) SortedFiles() []*File {

	return d.Files
}

// func longestCommonPrefix(strs []string) string {
// 	var longestPrefix string = ""
// 	var endPrefix = false

// 	if len(strs) > 0 {
// 		sort.Strings(strs)
// 		first := string(strs[0])
// 		last := string(strs[len(strs)-1])

// 		for i := 0; i < len(first); i++ {
// 			if !endPrefix && string(last[i]) == string(first[i]) {
// 				longestPrefix += string(last[i])
// 			} else {
// 				endPrefix = true
// 			}
// 		}
// 	}
// 	return longestPrefix
// }

// func (d *DepFiles) MostSpecificSharedPath() string {
// 	files := make([]string, len(d.Files))
// 	for i, file := range d.Files {
// 		files[i] = file.Path
// 	}
// 	lcp := longestCommonPrefix(files)
// 	if strings.HasSuffix(lcp, "/") {
// 		return strings.TrimSuffix(lcp, "/")
// 	}

// 	return filepath.Dir(lcp)
// }

func (d *DepFiles) AddAllNestedProtoFiles(ctx context.Context, path string, filters ...string) error {
	files, err := doublestar.Glob(os.DirFS(path), "**/*.proto")
	if err != nil {
		return errors.Errorf("getting files: %w", err)
	}

	if len(filters) > 0 {
		files = slices.DeleteFunc(files, func(file string) bool {
			for _, filter := range filters {
				ok, err := doublestar.PathMatch(filter, file)
				if err != nil || !ok {
					return true
				}
			}
			return false
		})
	}

	if len(files) == 0 {
		return errors.New("no proto files found in: " + path)
	}

	for _, file := range files {
		zerolog.Ctx(ctx).Info().Str("file", file).Str("path", path).Msg("adding file")

		if err := d.AddFile(path, file); err != nil {
			return errors.Errorf("adding file: %w", err)
		}
	}

	slices.SortFunc(d.Files, func(a, b *File) int {
		return strings.Compare(a.Path, b.Path)
	})

	return nil
}

func (d *DepFiles) AddFile(path string, file string) error {
	content, err := os.ReadFile(filepath.Join(path, file))
	if err != nil {
		return errors.Errorf("reading file: %w", err)
	}

	d.Files = append(d.Files, &File{
		Path:    file,
		Content: content,
	})

	return nil
}

const zeroHash = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

func (d *DepFiles) CalculateDigest() (string, error) {
	hash := sha256.New()

	// ensure the files are sorted
	slices.SortFunc(d.Files, func(a, b *File) int {
		return strings.Compare(a.Path, b.Path)
	})

	for _, file := range d.Files {
		_, err := hash.Write([]byte(file.Path))
		if err != nil {
			return "", errors.Errorf("writing to hash: %w", err)
		}
		_, err = hash.Write(file.Content)
		if err != nil {
			return "", errors.Errorf("writing to hash: %w", err)
		}
	}

	out := hex.EncodeToString(hash.Sum(nil))

	if out == zeroHash {
		return "", errors.New("zero hash")
	}

	return out, nil
}

func (d *DepFiles) WriteToDir(relPath string) error {

	for _, file := range d.Files {
		outfilePath := filepath.Join(relPath, file.Path)
		if err := os.MkdirAll(filepath.Dir(outfilePath), 0755); err != nil {
			return errors.Errorf("creating output directory: %w", err)
		}
		if err := os.WriteFile(outfilePath, file.Content, 0644); err != nil {
			return errors.Errorf("writing file: %w", err)
		}
	}

	// // write a commit.txt file
	// if err := os.WriteFile(filepath.Join(relPath, "commit.txt"), []byte(d.Commit), 0644); err != nil {
	// 	return errors.Errorf("writing commit file: %w", err)
	// }

	return nil
}

func (d *DepFiles) LockEntry() (*LockDep, error) {
	digest, err := d.CalculateDigest()
	if err != nil {
		return nil, errors.Errorf("calculating digest: %w", err)
	}

	return &LockDep{
		Metadata: LockDepMetadata{
			Type:   d.DepInfo.Type,
			Commit: d.commitMetadata,
		},
		Repo: d.DepInfo.Repo,
		Path: d.DepInfo.Path,
		Ref:  d.DepInfo.Ref,
		// Commit: d.Commit,
		Digest: digest,
		// Prefix: d.MostSpecificSharedPath(),
	}, nil
}

// compare two lock entries
func (l *LockDep) Compare(other *LockDep) bool {
	return l.Repo == other.Repo && l.Path == other.Path && l.Ref == other.Ref && l.Digest == other.Digest && l.Prefix == other.Prefix
}

func (me *LockFile) EntryFor(dep Buf3pdDep) *LockDep {
	for _, lockDep := range me.Deps {
		if lockDep.Repo == dep.Repo && lockDep.Path == dep.Path && lockDep.Ref == dep.Ref {
			return lockDep
		}
	}
	return nil
}

type File struct {
	Path    string `json:"path"`
	Content []byte `json:"content"`
}
