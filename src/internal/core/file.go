package core

import (
	"archive/tar"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/klauspost/compress/zstd"
)

type FileEntry struct {
	Path    string
	ModTime time.Time
	Size    int64
	Hash    string // 可选，以后用 SHA256
}

type Snapshot struct {
	Time  time.Time
	Files []FileEntry
}

func CreateFullBackup(srcDir, dstDir string, log func(string)) error {
	// 生成文件名
	timestamp := time.Now().Format("20060102_150405")
	backupFile := filepath.Join(dstDir, "backup_"+timestamp+".tar.zst")
	metaFile := filepath.Join(dstDir, "backup_"+timestamp+".json")

	// 创建压缩包
	outFile, err := os.Create(backupFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	zstdWriter, _ := zstd.NewWriter(outFile, zstd.WithEncoderLevel(zstd.SpeedDefault))
	defer zstdWriter.Close()

	tarWriter := tar.NewWriter(zstdWriter)
	defer tarWriter.Close()

	var snapshot Snapshot
	snapshot.Time = time.Now()

	// 遍历文件并写入 tar
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == srcDir {
			return nil
		}

		relPath, _ := filepath.Rel(srcDir, path)
		relPath = filepath.ToSlash(relPath)

		// 记录到快照
		entry := FileEntry{Path: relPath, ModTime: info.ModTime(), Size: info.Size()}
		// 跳过目录在快照中的记录（可调整）
		if !info.IsDir() {
			snapshot.Files = append(snapshot.Files, entry)
		}

		// 创建 tar 头
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		header.Name = relPath
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// 写入文件内容
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		io.Copy(tarWriter, file)

		log("已压缩: " + relPath)
		return nil
	})
	if err != nil {
		return err
	}

	// 保存快照 JSON
	metaData, _ := json.MarshalIndent(snapshot, "", "  ")
	os.WriteFile(metaFile, metaData, 0644)

	log("全量备份完成：" + backupFile)
	return nil
}
