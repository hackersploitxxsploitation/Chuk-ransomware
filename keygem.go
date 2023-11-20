package main

import "fmt"
import "golang.org/x/crypto/curve25519"
import "crypto/rand"
import "io"
import "os"
import "log"

//Usamos  a  seguinte formula algebrica
// keygen para ransomware ,gere seu par de chave publica privada  e  criptografe seu conteudo com a chave publica
func main(){
var chave_publica [32]byte
var privada [32]byte
	
	

	buffer := make([]byte, 32)
	io.ReadFull(rand.Reader, buffer)//gereamos um buffer 32 bytes aleatorio vamos testar
	copy(privada[:], buffer)
	curve25519.ScalarBaseMult(&chave_publica, &privada)//publicKey=k*G ponto gerador
	fmt.Println("seu par de chaves publica é privada",chave_publica,privada)//aqui printamos no console as chaves
	
f, err := os.Create("testFile.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    fmt.Println(f.Name())
	os.WriteFile("testFile.txt",int(chave_publica),0664)



}


//Iremos  criar nosso keygen em go  para gerar nossa par de chaves de privada publica usamos a  funçao curve25519 para gerar nossa chave publica e privada
//nosso ransomware irar usar curvas elpitcas do tipo curve2599 curvas NIST nao serao usadas por fatores de segurança
