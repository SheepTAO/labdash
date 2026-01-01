package main

import (
	"encoding/json"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"
)

// --- Docs Handlers ---

type FileNode struct {
	Name     string      `json:"name"`
	Path     string      `json:"path"` // Relative path
	Type     string      `json:"type"` // "file" or "dir"
	Owner    string      `json:"owner"`
	ModTime  string      `json:"modTime"`
	Children []*FileNode `json:"children,omitempty"`
}

func getFileMetadata(path string) (string, string) {
	info, err := os.Stat(path)
	if err != nil {
		return "--", "--"
	}

	// ModTime
	modTime := info.ModTime().Format("2006-01-02 15:04")

	// Owner
	owner := "unknown"
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		uid := strconv.Itoa(int(stat.Uid))
		u, err := user.LookupId(uid)
		if err == nil {
			owner = u.Username
		} else {
			owner = uid
		}
	}
	return owner, modTime
}

// hasMarkdownFiles checks if a directory (recursively) contains any .md files
func hasMarkdownFiles(basePath, relPath string) bool {
	fullPath := filepath.Join(basePath, relPath)
	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return false
	}

	for _, entry := range entries {
		name := entry.Name()
		// Skip hidden files
		if strings.HasPrefix(name, ".") {
			continue
		}

		childRelPath := filepath.Join(relPath, name)

		if entry.IsDir() {
			// Recursively check subdirectories
			if hasMarkdownFiles(basePath, childRelPath) {
				return true
			}
		} else if strings.HasSuffix(strings.ToLower(name), ".md") {
			// Found a markdown file
			return true
		}
	}

	return false
}

func walkDocs(basePath, relPath string, currentDepth, maxDepth int) []*FileNode {
	if currentDepth > maxDepth {
		return []*FileNode{}
	}

	fullPath := filepath.Join(basePath, relPath)
	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return []*FileNode{}
	}

	var nodes []*FileNode

	for _, entry := range entries {
		name := entry.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}

		childRelPath := filepath.Join(relPath, name)

		if entry.IsDir() {
			// If we are at maxDepth, we cannot show the content of this directory (which would be at maxDepth+1).
			// So we ignore the directory itself to avoid showing an empty folder.
			if currentDepth >= maxDepth {
				continue
			}

			// Check if this directory contains any markdown files (recursively)
			if !hasMarkdownFiles(basePath, childRelPath) {
				// Skip directories without markdown files
				continue
			}

			children := walkDocs(basePath, childRelPath, currentDepth+1, maxDepth)
			nodes = append(nodes, &FileNode{
				Name:     name,
				Path:     childRelPath,
				Type:     "dir",
				Children: children,
			})
		} else if strings.HasSuffix(strings.ToLower(name), ".md") {
			owner, modTime := getFileMetadata(filepath.Join(basePath, childRelPath))
			nodes = append(nodes, &FileNode{
				Name:    name,
				Path:    childRelPath,
				Type:    "file",
				Owner:   owner,
				ModTime: modTime,
			})
		}
	}

	// Sort: Default doc first, then Files, then Directories
	sort.Slice(nodes, func(i, j int) bool {
		// 1. Default document always comes first
		if globalConfig.DefaultDoc != "" {
			if nodes[i].Path == globalConfig.DefaultDoc {
				return true
			}
			if nodes[j].Path == globalConfig.DefaultDoc {
				return false
			}
		}
		// 2. Type priority: File > Dir
		if nodes[i].Type != nodes[j].Type {
			return nodes[i].Type == "file"
		}
		// 3. Name order (Case insensitive)
		return strings.ToLower(nodes[i].Name) < strings.ToLower(nodes[j].Name)
	})

	return nodes
}

func handleDocsTree(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	children := walkDocs(globalConfig.DocsPath, "", 1, globalConfig.DocsDepth)
	root := &FileNode{
		Name:     "root",
		Path:     "",
		Type:     "dir",
		Children: children,
	}

	json.NewEncoder(w).Encode(root)
}

func handleDocsContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	relPath := r.URL.Query().Get("path")
	if relPath == "" {
		http.Error(w, "Path is required", http.StatusBadRequest)
		return
	}

	// Security: Prevent directory traversal
	if strings.Contains(relPath, "..") {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	fullPath := filepath.Join(globalConfig.DocsPath, relPath)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Determine content type
	lower := strings.ToLower(relPath)
	if !strings.HasSuffix(lower, ".md") {
		http.Error(w, "Forbidden: Only Markdown files are allowed", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "text/markdown")
	w.Write(content)
}
