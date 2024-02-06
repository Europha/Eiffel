## About locker

Locker is an **independent** project focused on **educational purposes**. Its main goal is to be **fast and easy to compile and use**. The code was written in Golang, using libraries to map and identify other storage partitions and network folders automatically mounted on the machine. The system uses **goroutines** and threads to facilitate **time management** in situations with **multiple files**, in addition to using modern CBC encryption, known as **ChaCha20**.



https://github.com/Redshifteye/Eiffel/assets/109265897/baddf75e-e076-4e93-9d90-a3ad91c3760e

![aa](https://github.com/Redshifteye/Eiffel/assets/109265897/1b79bbe1-8390-4aa7-aa3f-bd3da51e6720)



### Disclaimer:

Locker was created for educational purposes only and should not be used in production environments. The author is not responsible for any damage or data loss that may occur as a result of using this project. It is important for the user to understand the risks and responsibilities before using Locker.

## Sobre o Locker

O Locker é um projeto **independente** focado em **propósitos educacionais**. Sua principal meta é ser **rápido e fácil de compilar e usar**. O código foi escrito em Golang, utilizando bibliotecas para mapear e identificar outras partições de armazenamento e pastas de rede automaticamente montadas na máquina. O sistema utiliza **goroutines** e threads para facilitar o gerenciamento de tempo em situações com **múltiplos arquivos**, além de contar com criptografia CBC moderna, conhecida como **ChaCha20**.

### Disclaimer:

O Locker foi criado apenas para fins didáticos e não deve ser utilizado em ambientes de produção. O autor não se responsabiliza por qualquer dano ou perda de dados que possa ocorrer como resultado do uso deste projeto. É importante que o usuário compreenda os riscos e responsabilidades antes de utilizar o Locker.

```
installation (requires golang 1.21.1 or above):

git clone https://github.com/redshifteye/eiffel

go install

go build a.go #For linux dists GOOS=windows GOARCH=amd64 go build -o bin/app-amd64.exe app.go
```
